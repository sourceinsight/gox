package sync2

import (
	"fmt"
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	var d AtomicDuration
	if d.Get() != 0 {
		t.Errorf("atomicduration = %d; want %d", d, 0)
	}
	d.Set(time.Second)
	if fmt.Sprintf("%v", d) != "1s" {
		t.Errorf("atomicDuration display %s; want %s", d, "1s")
	}
}

func TestAtomicString(t *testing.T) {
	var s AtomicString
	if s.Get() != "" {
		t.Errorf("want empty, got %s", s.Get())
	}
	s.Set("a")
	if s.Get() != "a" {
		t.Errorf("want a, got %s", s.Get())
	}
	if s.CompareAndSwap("b", "c") {
		t.Errorf("want false, got true")
	}
	if s.Get() != "a" {
		t.Errorf("want a, got %s", s.Get())
	}
	if !s.CompareAndSwap("a", "c") {
		t.Errorf("want true, got false")
	}
	if s.Get() != "c" {
		t.Errorf("want c, got %s", s.Get())
	}
}
func TestAtomicBool(t *testing.T) {
	var b AtomicBool
	if b.Get() {
		t.Error("invalid value", b.Get())
	}
	b.Set(true)
	if !b.Get() {
		t.Error("invalid value", b.Get())
	}
}
