package tapedb

import (
	"github.com/dgraph-io/ristretto"
	"os"
	"sync"
)

type DB struct {
	meta                 tapeMeta
	file                 *os.File
	blockCapacity        uint64
	blockContentCapacity uint64
	counter              sync.WaitGroup
	cache                *ristretto.Cache
	latestSeq            *sequence
	closed               int64
}
