package tapedb

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func TestPage_Segment(t *testing.T) {
	buf := bytes.NewBufferString("")
	n := uint16(0)
	for i := 0; i < 4; i++ {
		v := []byte(time.Now().String())
		s := newSegment(32, v)
		n = n + s.blocks()
		buf.Write(s)
		fmt.Println(string(v))
	}
	p := &page{
		beg: 1,
		end: uint64(n),
		cap: 32,
		p:   buf.Bytes(),
	}
	for i := 0; i < 4; i++ {
		seg, remain := p.segment(uint64(i)*3 + 1)
		fmt.Println(remain, string(seg.Read()))
	}
}
