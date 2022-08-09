package tapedb

import (
	"encoding/binary"
	"fmt"
)

func newSegment(blockCapacity uint16, p []byte) (s segment) {
	blockSize := calcBlockSize(blockCapacity, p)
	s = make([]byte, blockCapacity*blockSize)
	n := 0
	for i := uint16(1); i <= blockSize; i++ {
		b := newBlock(s[blockCapacity*(i-1) : blockCapacity*i])
		n = b.Write(p, i, blockSize)
		p = p[n:]
	}
	return
}

type segment []byte

func (s segment) Read() (p []byte) {
	p = make([]byte, 0, len(s))
	blockSize := s.blocks()
	blockCapacity := uint16(len(s)) / blockSize
	blockIdx := binary.LittleEndian.Uint16(s[4:6])
	for i := blockIdx; i <= blockSize; i++ {
		length := binary.LittleEndian.Uint32(s[blockCapacity*(i-1) : blockCapacity*(i-1)+4])
		idx := binary.LittleEndian.Uint16(s[blockCapacity*(i-1)+4 : blockCapacity*(i-1)+6])
		if idx != i {
			panic(fmt.Errorf("read invalid segment"))
		}
		p = append(p, s[uint32(blockCapacity*(i-1)):uint32(blockCapacity*i)][uint32(blockCapacity)-length:]...)
	}
	return
}

func (s segment) blocks() (n uint16) {
	n = binary.LittleEndian.Uint16(s[6:8])
	return
}

func mergeSegments(a segment, b segment) (c segment) {
	aLen := len(a)
	bLen := len(b)
	c = make([]byte, aLen+bLen)
	copy(c[0:aLen], a)
	copy(c[aLen:], b)
	return
}
