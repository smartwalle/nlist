package smap

import (
	"testing"
	"fmt"
)

func TestSyncMap(t *testing.T) {
	var m = NewSyncMap()
	m.Set("a", "bbb")

	fmt.Println(m)
}

func TestIter(t *testing.T) {
	fmt.Println("=====TestIter=====")
	var m = NewSyncMap()

	m.Set("a", "aaa")
	m.Set("b", "bbb")

	for v := range m.Iter() {
		fmt.Println(v.Key, v.Value)
	}
}