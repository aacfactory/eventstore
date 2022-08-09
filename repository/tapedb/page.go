package tapedb

import (
	"encoding/binary"
)

type page struct {
	beg uint64
	end uint64
	cap uint64
	p   []byte
}

func (p *page) segment(seq uint64) (seg segment, remainSeq uint64) {
	beg := seq - p.beg
	segmentIdx := binary.LittleEndian.Uint16(p.p[beg*p.cap+4 : beg*p.cap+6])
	segmentSize := binary.LittleEndian.Uint16(p.p[beg*p.cap+6 : beg*p.cap+8])
	end := seq + uint64(segmentSize-segmentIdx)
	if end > p.end {
		remainSeq = p.end + 1
		end = p.end
	}
	end = end - p.beg
	seg = p.p[beg*p.cap : (end+1)*p.cap]
	return
}
