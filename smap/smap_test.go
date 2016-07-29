package smap

import (
	"fmt"
	"testing"
)

func TestSyncMap(t *testing.T) {
	var m = NewBlockMap()
	m.Set("a", "bbb")

	fmt.Println(m)
}

func TestIter(t *testing.T) {
	fmt.Println("=====TestIter=====")
	var m = NewBlockMap()

	m.Set("a", "aaa")
	m.Set("b", "bbb")

	for v := range m.Iter() {
		fmt.Println(v.Key, v.Value)
	}
}
