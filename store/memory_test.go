package store

import (
	"testing"
	"time"
)

type item struct {
	Name  string
	Value int
}

func TestMemoryStoreGetSetDel(t *testing.T) {
	m := NewMemoryStore(time.Second * 10)
	defer m.Close()

	it := &item{}
	if m.Get("foo", &it) != ErrNotFound {
		t.Error("key should not be found")
	}

	it = &item{Name: "Banana", Value: 76}
	if err := m.Create("fruit", it, 0); err != nil {
		t.Errorf("cannot set key: %s", err)
	}

	if m.Create("fruit", it, 0) != ErrExist {
		t.Error("should not be possible to overwrite key")
	}

	if err := m.Create("banana", it, 0); err != nil {
		t.Errorf("cannot save more than one object: %s", err)
	}

	banana := &item{}
	if err := m.Get("banana", &banana); err != nil {
		t.Errorf("cannot get banana: %s", err)
	}
	if banana.Name != it.Name || banana.Value != it.Value {
		t.Errorf("broken banana received: %v", banana)
	}

	if err := m.Del("does-not-exist"); err != nil {
		t.Errorf("cannot delete what does not exist: %s", err)
	}
	if err := m.Del("banana"); err != nil {
		t.Errorf("cannot delete banana: %s", err)
	}
	if err := m.Get("banana", &banana); err != ErrNotFound {
		t.Errorf("banana should be deleted: %s", err)
	}
}

func TestMemoryStoreExpiration(t *testing.T) {
	m := NewMemoryStore(time.Millisecond * 15)
	defer m.Close()

	orange := &item{Name: "orange", Value: 51}
	if err := m.Create("orange", orange, time.Millisecond*20); err != nil {
		t.Errorf("cannot save orange: %s", err)
	}
	apple := &item{Name: "apple", Value: 61}
	if err := m.Create("apple", apple, time.Millisecond*20); err != nil {
		t.Errorf("cannot save apple: %s", err)
	}

	if err := m.Get("orange", &orange); err != nil {
		t.Errorf("cannot get orange: %s", err)
	}
	if err := m.Get("apple", &apple); err != nil {
		t.Errorf("cannot get apple: %s", err)
	}

	time.Sleep(time.Millisecond * 20)

	if err := m.Get("orange", &orange); err != ErrNotFound {
		t.Errorf("orange has not expired: %s", err)
	}
	if err := m.Get("apple", &apple); err != ErrNotFound {
		t.Errorf("apple has not expired: %s", err)
	}
}
