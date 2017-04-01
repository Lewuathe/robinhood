package robinhood

import "hash/fnv"

func FnvHash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

type HashMap interface {
	Size() uint32
	Get(string) interface{}
	Put(string, interface{})
	Erase(string)
	String() string
	LoadFactor() float32
	DibAverage() float32
}

type Entry struct {
	key   string
	value interface{}
	dib   uint32
}
