package web

import (
	"testing"
)

func TestPubSub(t *testing.T) {
	ps := NewPubSub(8)

	ps.Publish("foo", "bar")

	cfirst := ps.Subscribe("first")
	ps.Publish("first", "message")
	msg, _ := (<-cfirst.C).(string)
	if msg != "message" {
		t.Fatalf("\"message\" expected, got \"%s\"", msg)
	}

	ps.Unsubscribe(cfirst)
	select {
	case val, ok := <-cfirst.C:
		if ok {
			t.Fatalf("message received, but channel should be closed: %v", val)
		}
	}

	cone := ps.Subscribe("test")
	ctwo := ps.Subscribe("test")

	ps.Publish("test", "foo")
	ps.Unsubscribe(cone)
	ps.Publish("test", "bar")
	ps.Unsubscribe(ctwo)

	assertConsume(t, cone.C, "foo")
	assertClosed(t, cone.C)

	assertConsume(t, ctwo.C, "foo")
	assertConsume(t, ctwo.C, "bar")
	assertClosed(t, ctwo.C)
}

func TestPubSubUnsubscribeAll(t *testing.T) {
	ps := NewPubSub(8)

	ps.UnsubscribeAll("does-not-exist")

	c1 := ps.Subscribe("x")
	c2 := ps.Subscribe("x")
	ps.UnsubscribeAll("x")

	assertClosed(t, c1.C)
	assertClosed(t, c2.C)
}

func assertConsume(t *testing.T, c <-chan interface{}, expected string) {
	select {
	case val, ok := <-c:
		if !ok {
			t.Fatalf("attempting to receive %s, but channel is closed", expected)
		}
		if s, ok := val.(string); !ok {
			t.Fatalf("expected to receive string, got %v", val)
		} else if s != expected {
			t.Fatalf("%s != %s", s, expected)
		}
	}
}

func assertClosed(t *testing.T, c <-chan interface{}) {
	select {
	case val, ok := <-c:
		if ok {
			t.Fatalf("expected closed channel, but %v received", val)
		}
	}
}
