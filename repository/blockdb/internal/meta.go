package internal

import (
	"fmt"
	"strconv"
)

type Meta map[string]string

func (meta Meta) setId(v string) {
	meta["id"] = v
	return
}

func (meta Meta) id() (v string) {
	id, hasId := meta["id"]
	if !hasId {
		panic(fmt.Errorf("blocks: there is no id in meta"))
		return
	}
	v = id
	return
}

func (meta Meta) setBlockCapacity(v uint64) {
	meta["blockCapacity"] = fmt.Sprintf("%d", v)
	return
}

func (meta Meta) blockCapacity() (v uint64) {
	capacity, hasCapacity := meta["blockCapacity"]
	if !hasCapacity {
		panic(fmt.Errorf("blocks: there is no blockCapacity in meta"))
		return
	}
	var err error
	v, err = strconv.ParseUint(capacity, 10, 64)
	if err != nil {
		panic(fmt.Errorf("blocks: blockSize in meta is not uint64"))
		return
	}
	return
}

func (meta Meta) setFixed(v bool) {
	if v {
		meta["fixed"] = "T"
	} else {
		meta["fixed"] = "F"
	}
	return
}

func (meta Meta) fixed() (v bool) {
	vv, has := meta["fixed"]
	if !has {
		return
	}
	v = vv == "T"
	return
}
