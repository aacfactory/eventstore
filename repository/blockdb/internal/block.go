package internal

import (
	"encoding/binary"
	"sync/atomic"
)

type BlockNo struct {
	value   uint64
	padding [7]uint64
}

func (bn *BlockNo) Value() (v uint64) {
	v = atomic.LoadUint64(&bn.value)
	return
}

func (bn *BlockNo) Next(n uint64) (no uint64) {
	no = atomic.AddUint64(&bn.value, n) - n + 1
	return
}

func NewBlock(capacity uint64, segmentNo uint16, segmentSize uint16) (b Block) {
	b = make([]byte, capacity)
	binary.LittleEndian.PutUint16(b[0:2], 0)
	binary.LittleEndian.PutUint16(b[2:4], segmentNo)
	binary.LittleEndian.PutUint16(b[4:6], segmentSize)
	binary.LittleEndian.PutUint16(b[6:8], 0)
	binary.LittleEndian.PutUint64(b[8:16], 0)
	return
}

type Block []byte

func (b Block) Write(p []byte) {
	binary.LittleEndian.PutUint16(b[6:8], uint16(len(p)))
	copy(b[16:], p)
	return
}

func (b Block) SetExtra(no uint64) {
	binary.LittleEndian.PutUint64(b[8:16], no)
	return
}

func (b Block) Commit() {
	binary.LittleEndian.PutUint16(b[0:2], 1)
	return
}

func (b Block) Read() (p []byte, segmentNo uint16, segmentSize uint16, extraNo uint64, has bool) {
	has = binary.LittleEndian.Uint16(b[0:2]) == 1
	if !has {
		return
	}
	segmentNo = binary.LittleEndian.Uint16(b[2:4])
	segmentSize = binary.LittleEndian.Uint16(b[4:6])
	extraNo = binary.LittleEndian.Uint64(b[8:16])
	p = b[16:binary.LittleEndian.Uint16(b[6:8])]
	return
}
