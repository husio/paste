package web

import (
	"sync"
)

type PubSub struct {
	sync.RWMutex

	bufsize int
	topics  map[string]map[*Subscription]*Subscription
}

func NewPubSub(subBufSize int) *PubSub {
	ps := &PubSub{
		topics:  make(map[string]map[*Subscription]*Subscription),
		bufsize: subBufSize,
	}
	return ps
}

func (ps *PubSub) Subscribe(topic string) *Subscription {
	ps.Lock()
	subs, ok := ps.topics[topic]
	if !ok {
		subs = make(map[*Subscription]*Subscription)
		ps.topics[topic] = subs
	}
	c := make(chan interface{}, ps.bufsize)
	sub := &Subscription{C: c, c: c, topic: topic}
	subs[sub] = sub
	ps.Unlock()
	return sub
}

func (ps *PubSub) Publish(topic string, message interface{}) {
	ps.RLock()
	if subs, ok := ps.topics[topic]; ok {
		for _, s := range subs {
			s.c <- message
		}
	}
	ps.RUnlock()
}

func (ps *PubSub) Unsubscribe(sub *Subscription) {
	ps.Lock()
	if subs, ok := ps.topics[sub.topic]; ok {
		if _, ok := subs[sub]; ok {
			close(sub.c)
			delete(subs, sub)
		}
		if len(subs) == 0 {
			delete(ps.topics, sub.topic)
		}
	}
	ps.Unlock()
}

func (ps *PubSub) UnsubscribeAll(topic string) error {
	ps.Lock()
	if subs, ok := ps.topics[topic]; ok {
		for _, sub := range subs {
			close(sub.c)
		}
		delete(ps.topics, topic)
	}
	ps.Unlock()
	return nil
}

type Subscription struct {
	C     <-chan interface{}
	topic string
	c     chan interface{}
}
