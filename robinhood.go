package robinhood

import (
	"fmt"
)

type RobinHood struct {
	values []Entry
	count  uint32
}

func NewRobinHood(size int) *RobinHood {
	return &RobinHood{
		values: make([]Entry, size, size),
	}
}

func (m RobinHood) Size() uint32 {
	return uint32(len(m.values))
}

func (m *RobinHood) Put(k string, v interface{}) {
	i := FnvHash(k) % m.Size()
	// DIB (distance to initial bucket) is zero initially
	e := Entry{key: k, value: v, dib: 0}
	for c := i; ; c = (c + 1) % m.Size() {
		if m.values[c].value == nil {
			// Able to insert here
			m.values[c] = e
			m.count += 1
			return
		} else {
			if m.values[c].dib < e.dib {
				// To be swapped
				tmp := e
				e = Entry{
					key:   m.values[c].key,
					value: m.values[c].value,
					dib:   m.values[c].dib,
				}
				m.values[c] = Entry{
					key:   tmp.key,
					value: tmp.value,
					dib:   tmp.dib,
				}

			}
		}
		e.dib += 1
	}
}

func (m *RobinHood) Get(k string) interface{} {
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

func (m *RobinHood) Erase(k string) {
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

func (m *RobinHood) LoadFactor() float32 {
	return float32(m.count) / float32(m.Size())
}

func (m *RobinHood) DibAverage() float32 {
	sum := uint32(0)
	for _, v := range m.values {
		sum += v.dib
	}
	return float32(sum) / float32(m.count)
}

func (m *RobinHood) String() string {
	ret := ""
	for _, v := range m.values {
		ret = ret + fmt.Sprintf("[%s,%v,%d] ", v.key, v.value, v.dib)
	}
	return ret
}
