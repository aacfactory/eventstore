package tapedb

import (
	"encoding/binary"
	"fmt"
	"github.com/vmihailenco/msgpack/v5"
	"strconv"
	"time"
)

func encodeMetaToBytes(meta tapeMeta) (p []byte, err error) {
	b, encodeErr := msgpack.Marshal(meta)
	if encodeErr != nil {
		err = fmt.Errorf("encode meta failed, %v", encodeErr)
		return
	}
	bLen := len(b)
	p = make([]byte, 2+bLen)
	binary.LittleEndian.PutUint16(p[0:2], uint16(bLen))
	copy(p[2:], b)
	return
}

func decodeBytesToMeta(p []byte) (meta tapeMeta, err error) {
	if len(p) < 3 {
		err = fmt.Errorf("decode meta failed, invalid content")
		return
	}
	decodeErr := msgpack.Unmarshal(p[2:2+binary.LittleEndian.Uint16(p[0:2])], &meta)
	if decodeErr != nil {
		err = fmt.Errorf("decode meta failed, %v", decodeErr)
		return
	}
	return
}

type tapeMeta map[string]string

func (meta tapeMeta) setBlockCapacity(v uint64) {
	meta["bc"] = strconv.FormatUint(v, 10)
	return
}

func (meta tapeMeta) blockCapacity() (v uint64) {
	capacity, hasCapacity := meta["bc"]
	if !hasCapacity {
		panic(fmt.Errorf("there is no block capacity in meta"))
		return
	}
	var err error
	v, err = strconv.ParseUint(capacity, 10, 64)
	if err != nil {
		panic(fmt.Errorf("value of block capacity in meta is not uint64"))
		return
	}
	return
}

func (meta tapeMeta) setFlushTime(v time.Time) {
	meta["ft"] = v.Format(time.RFC3339)
	return
}

func (meta tapeMeta) latestFlushTime() (ft time.Time) {
	v, has := meta["ft"]
	if !has {
		return
	}
	ft0, parseErr := time.Parse(time.RFC3339, v)
	if parseErr != nil {
		return
	}
	ft = ft0
	return
}

func (meta tapeMeta) setOpening() {
	meta["op"] = "T"
	return
}

func (meta tapeMeta) isOpening() (ok bool) {
	v, has := meta["op"]
	if !has {
		return
	}
	ok = v == "T"
	return
}

func (meta tapeMeta) setClosed() {
	meta["op"] = "F"
	return
}
