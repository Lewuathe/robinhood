package robinhood

import (
	"testing"
)

func TestLinear(t *testing.T) {
	var m HashMap
	m = NewLinear(10)

	keys := []string{
		"one",
		"two",
		"three",
	}

	for i, k := range keys {
		m.Put(k, i+1)
	}

	for i, k := range keys {
		if m.Get(k) != (i + 1) {
			t.Errorf("For %s expected=%d, but got %d", k, i+1, m.Get(k))
		}
	}

}
