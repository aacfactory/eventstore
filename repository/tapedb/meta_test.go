package tapedb

import (
	"fmt"
	"testing"
	"time"
)

func TestMeta(t *testing.T) {
	meta := tapeMeta{}
	meta.setBlockCapacity(64)
	meta.setOpening()
	meta.setFlushTime(time.Now())
	fmt.Println(meta.blockCapacity(), meta.isOpening(), meta.latestFlushTime())
	p, encodeErr := encodeMetaToBytes(meta)
	if encodeErr != nil {
		fmt.Println(encodeErr)
		return
	}
	meta1, decodeErr := decodeBytesToMeta(p)
	if decodeErr != nil {
		fmt.Println(decodeErr)
		return
	}
	fmt.Println(meta1.blockCapacity(), meta1.isOpening(), meta1.latestFlushTime())

}
