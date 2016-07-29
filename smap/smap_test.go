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