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

func NewBlock(segmentNo uint64, capacity int) (b Block) {
	b = make([]byte, capacity)
	binary.LittleEndian.PutUint64(b[0:8], 0)          // size
	binary.LittleEndian.PutUint64(b[8:16], segmentNo) // segmentNo
	binary.LittleEndian.PutUint64(b[16:24], 1)        // when gt 0 then spanNo; when eq 0 then link flag
	binary.LittleEndian.PutUint64(b[24:32], 1)        // span size or link no
	return
}

type Block []byte

func (b Block) Write(p []byte) {
	binary.LittleEndian.PutUint64(b[0:8], uint64(len(p)))
	copy(b[32:], p)
	return
}

func (b Block) SetLink(no uint64) {
	binary.LittleEndian.PutUint64(b[16:24], 0)
	binary.LittleEndian.PutUint64(b[24:32], no)
	return
}

func (b Block) SetSpan(no uint64, size uint64) {
	binary.LittleEndian.PutUint64(b[16:24], no)
	binary.LittleEndian.PutUint64(b[24:32], size)
	return
}

func (b Block) Read() (p []byte, has bool) {
	contentLen := binary.LittleEndian.Uint64(b[0:8])
	has = contentLen > 0
	if !has {
		return
	}
	p = b[32 : 32+contentLen]
	return
}

type Blocks []Block

func (bs Blocks) Bytes() (p []byte) {
	bc := cap(bs[0])
	p = make([]byte, len(bs)*bc)
	for i, b := range bs {
		copy(p[i*bc:(i+1)*bc], b)
	}
	return
}

func (bs Blocks) Content() (p []byte) {

	return
}
