package bitree_test

import (
	"time"
	"math/rand"
	"testing"
	"mycodes/bitree"
)

var bt *bitree.BBiTree

func TestBBiTree(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	n := rand.Intn(100)
	for i:=0; i<n; i++ {
		bt, _ = bt.Insert(rand.Intn(1000))
	}
	bt.PreTraverse()
}
