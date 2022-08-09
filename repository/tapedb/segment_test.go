package tapedb

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func TestSegment_Read(t *testing.T) {
	p := []byte(time.Now().String())
	seg := newSegment(32, p)
	r := seg.Read()
	fmt.Println(len(p) == len(r), bytes.Equal(p, r), string(r))
}
