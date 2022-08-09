package tapedb

import (
	"encoding/binary"
	"math"
)

func calcBlockSize(capacity uint16, p []byte) (size uint16) {
	size = uint16(math.Ceil(float64(len(p)) / float64(capacity-8)))
	return
}

func newBlock(p []byte) (b block) {
	b = p
	return
}

type block []byte

func (b block) Write(p []byte, segmentIdx uint16, segmentSize uint16) (n int) {
	bLen := len(b) - 8
	pLen := len(p)
	if pLen-bLen < 0 {
		n = pLen
	} else {
		n = bLen
	}
	if segmentIdx == 0 {
		segmentIdx = 1
	}
	if segmentSize == 0 {
		segmentSize = 1
	}
	binary.LittleEndian.PutUint32(b[0:4], uint32(n))
	binary.LittleEndian.PutUint16(b[4:6], segmentIdx)
	binary.LittleEndian.PutUint16(b[6:8], segmentSize)
	copy(b[bLen+8-n:], p)
	return
}

func (b block) Read() (p []byte, segmentIdx uint16, segmentSize uint16, has bool) {
	length := binary.LittleEndian.Uint16(b[0:4])
	has = length > 0
	if !has {
		return
	}
	segmentIdx = binary.LittleEndian.Uint16(b[4:6])
	segmentSize = binary.LittleEndian.Uint16(b[6:8])
	p = b[uint16(len(b))-length:]
	return
}
