package nlist_test

import (
	"github.com/smartwalle/nlist"
	"strconv"
	"testing"
)

func BenchmarkList_PushBack_Int(b *testing.B) {
	var l = nlist.New[int]()

	for i := 0; i < b.N; i++ {
		l.PushBack(i)
	}
}

func BenchmarkList_PushBack_String(b *testing.B) {
	var l = nlist.New[string]()

	for i := 0; i < b.N; i++ {
		l.PushBack(strconv.Itoa(i))
	}
}
