package internal

import (
	"os"
	"sync"
)

type Blocks struct {
	file                 os.File
	fixed                bool
	blockCapacity        uint64
	blockContentCapacity uint64
	counter              sync.WaitGroup
	lastBlockNo          *BlockNo
	closed               int64
	extra                *Blocks
}
