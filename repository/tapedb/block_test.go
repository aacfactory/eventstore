package tapedb

import (
	"fmt"
	"testing"
)

func TestBlock_Read(t *testing.T) {
	b := newBlock(make([]byte, 16))
	fmt.Println(b.Write([]byte{1, 2, 3, 4}, 1, 1) == 4)
	fmt.Println(b.Read())
	fmt.Println(b.Write([]byte{1, 2, 3, 4, 5, 6, 7, 8}, 1, 1) == 8)
	fmt.Println(b.Read())
	fmt.Println(b.Write([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, 1) == 8)
	fmt.Println(b.Read())
}
