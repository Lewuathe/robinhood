package robinhood

import (
	"testing"
)

func TestRobinHood(t *testing.T) {
	var m HashMap
	m = NewRobinHood(5)

	keys := []string{
		"one",
		"two",
		"three",
	}

	for i, k := range keys {
		m.Put(k, i+1)
	}

	t.Log(m.String())

	for i, k := range keys {
		if m.Get(k) != (i + 1) {
			t.Errorf("For %s expected=%d, but got %d", k, i+1, m.Get(k))
		}
	}

	m.Erase("one")

	if m.Get("one") != nil {
		t.Errorf("Access existing item. %v", m.Get("one"))
	}
}
