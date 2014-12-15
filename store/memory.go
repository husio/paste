package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"
)

var (
	ErrTimeoutTooSmall = errors.New("expiration time too small")
)

type MemoryStore struct {
	sync.Mutex

	expTimeTick time.Duration
	items       map[string]*storeitem
	timeouts    map[int64][]string
	stop        chan bool
}

type storeitem struct {
	data      []byte
	validTill *time.Time
}

func NewMemoryStore(expTimeTick time.Duration) *MemoryStore {
	m := &MemoryStore{
		expTimeTick: expTimeTick,
		items:       make(map[string]*storeitem),
		stop:        make(chan bool, 1),
		timeouts:    make(map[int64][]string),
	}
	go m.cleanupExpired()
	return m
}

func (s *MemoryStore) Create(key string, obj interface{}, expireAfter time.Duration) error {
	if expireAfter != 0 && expireAfter < s.expTimeTick {
		return ErrTimeoutTooSmall
	}

	data, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	s.Lock()
	defer s.Unlock()
	if _, ok := s.items[key]; ok {
		return ErrExist
	}
	if expireAfter != 0 {
		if err := s.expireAt(key, expireAfter); err != nil {
			return err
		}
	}

	var exp *time.Time = nil
	if expireAfter != 0 {
		expTime := time.Now().Add(expireAfter)
		exp = &expTime
	}
	s.items[key] = &storeitem{data: data, validTill: exp}
	return nil
}

func (s *MemoryStore) Get(key string, obj interface{}) error {
	s.Lock()
	item, ok := s.items[key]
	s.Unlock()

	if !ok || (item.validTill != nil && item.validTill.Before(time.Now())) {
		return ErrNotFound
	}
	return json.Unmarshal(item.data, obj)
}

func (s *MemoryStore) Del(key string) error {
	s.Lock()
	delete(s.items, key)
	s.Unlock()
	return nil
}

func (s *MemoryStore) Close() error {
	s.stop <- true
	return nil
}

func (s *MemoryStore) cleanupExpired() {
	ticker := time.NewTicker(s.expTimeTick)
	defer ticker.Stop()

	tick := s.expTimeTick.Nanoseconds()
	for {
		select {
		case <-s.stop:
			return
		case now := <-ticker.C:
			t := now.UnixNano() / tick * tick
			s.Lock()
			if keys, ok := s.timeouts[t]; ok {
				for _, key := range keys {
					delete(s.items, key)
				}
				delete(s.timeouts, t)
			}
			s.Unlock()
		}
	}
}

func (s *MemoryStore) expireAt(key string, exp time.Duration) error {
	if exp <= s.expTimeTick {
		return fmt.Errorf("expiration period too short: %s", exp)
	}

	tick := s.expTimeTick.Nanoseconds()
	now := time.Now().UnixNano()
	t := (now + exp.Nanoseconds()) / tick * tick
	keys, ok := s.timeouts[t]
	if !ok {
		keys = make([]string, 0, 12)
	}
	s.timeouts[t] = append(keys, key)
	return nil
}
