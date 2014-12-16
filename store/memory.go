package store

import (
	"encoding/json"
	"errors"
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
	created   time.Time
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

	now := time.Now()
	var exp *time.Time = nil

	if expireAfter != 0 {
		expTime := now.Add(expireAfter)
		exp = &expTime

		if err := s.expireAt(key, expTime); err != nil {
			return err
		}
	}
	s.items[key] = &storeitem{data: data, validTill: exp, created: now}
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

	tick := int64(s.expTimeTick.Seconds())
	for {
		select {
		case <-s.stop:
			return
		case now := <-ticker.C:
			t := now.Unix() / tick * tick
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

func (s *MemoryStore) expireAt(key string, exp time.Time) error {
	tick := int64(s.expTimeTick.Seconds())
	t := exp.Unix() / tick * tick
	keys, ok := s.timeouts[t]
	if !ok {
		keys = make([]string, 0, 12)
	}
	s.timeouts[t] = append(keys, key)
	return nil
}
