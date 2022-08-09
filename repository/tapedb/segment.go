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
		b.ModSegment(i, blockSize)
		n = b.Write(p)
		p = p[n:]
	}
	return
}

type segment []byte

func (s segment) Read() (p []byte) {
	p = make([]byte, 0, len(s))
	blockSize := binary.LittleEndian.Uint16(s[6:8])
	blockCapacity := uint16(len(s)) / blockSize
	for i := uint16(1); i <= blockSize; i++ {
		length := binary.LittleEndian.Uint32(s[blockCapacity*(i-1) : blockCapacity*(i-1)+4])
		idx := binary.LittleEndian.Uint16(s[blockCapacity*(i-1)+4 : blockCapacity*(i-1)+6])
		if idx != i {
			panic(fmt.Errorf("tapedb: read invalid segment"))
		}
		p = append(p, s[uint32(blockCapacity*(i-1)):uint32(blockCapacity*i)][uint32(blockCapacity)-length:]...)
	}
	return
}
