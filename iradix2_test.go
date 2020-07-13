package iradix

import (
	"fmt"
	"testing"
)

func TestRadix2(t *testing.T) {
	// Create a tree
	r := New()
	r, _, _ = r.Insert([]byte("Ni"), 1)
	r, _, _ = r.Insert([]byte("NiHao"), 2)
	r, _, _ = r.Insert([]byte("NiHaoMa"), 3)
	r, _, _ = r.Insert([]byte("NiHaoYa"), 4)
	r, _, _ = r.Insert([]byte("NiHaoYa"), 6)
	r, _, _ = r.Insert([]byte("NiHui"), 9)
	// r, _, _ = r.Insert([]byte("Jin"), 7)
	// r, _, _ = r.Insert([]byte("JinTian"), 8)

	s := formatNode(0, r.Root())
	fmt.Println(s)
}


