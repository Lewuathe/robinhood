package robinhood

import "fmt"

type Linear struct {
	values []Entry
	count  uint32
}

func NewLinear(size int) *Linear {
	return &Linear{
		values: make([]Entry, size, size),
	}
}

func (m Linear) Size() uint32 {
	return uint32(len(m.values))
}

func (m *Linear) Put(k string, v interface{}) {
	i := FnvHash(k) % m.Size()
	// Linear search of the target entry
	dib := uint32(0)
	for m.values[i].value != nil {
		i = (i + 1) % m.Size()
		dib += 1
	}
	m.values[i] = Entry{key: k, value: v, dib: dib}
	m.count += 1
}

func (m *Linear) Get(k string) interface{} {
	i := FnvHash(k) % m.Size()
	e := m.values[i]
	for e.key != k {
		i = (i + 1) % m.Size()
		e = m.values[i]
		if i == FnvHash(k)%m.Size() {
			// Not Found
			return nil
		}
	}
	return e.value
}

func (m *Linear) Erase(k string) {
	i := FnvHash(k) % m.Size()
	e := m.values[i]
	for e.key != k {
		i = (i + 1) % m.Size()
		e = m.values[i]
		if i == FnvHash(k)%m.Size() {
			return
		}
	}
	m.values[i] = Entry{key: "", value: nil}
	m.count -= 1
}

func (m *Linear) LoadFactor() float32 {
	return float32(m.count) / float32(m.Size())
}

func (m *Linear) DibAverage() float32 {
	sum := uint32(0)
	for _, v := range m.values {
		sum += v.dib
	}
	return float32(sum) / float32(m.count)
}

func (m *Linear) String() string {
	ret := ""
	for _, v := range m.values {
		ret = ret + fmt.Sprintf("[%s,%v] ", v.key, v.value)
	}
	return ret
}
