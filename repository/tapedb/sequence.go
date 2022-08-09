package tapedb

import "sync/atomic"

type sequence struct {
	value   uint64
	padding [7]uint64
}

func (seq *sequence) Value() (v uint64) {
	v = atomic.LoadUint64(&seq.value)
	return
}

func (seq *sequence) Next(n uint64) (no uint64) {
	no = atomic.AddUint64(&seq.value, n) - n + 1
	return
}
