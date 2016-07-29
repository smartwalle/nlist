package set

import (
	"testing"
	"fmt"
)

func TestSet(t *testing.T) {
	var s = NewSet(1, 2, 3, 5, 6, 7, 7)

	if s.Len() != 6 {
		t.Error("集合的长度应该为 6")
	}

	if !s.Contains(1, 3) {
		t.Error("集合包含元素 1, 3")
	}

	if s.Contains(4) {
		t.Error("集合不包含元素 4")
	}

	s.Add(1, 4)

	if !s.Exists(4) {
		t.Error("集合包含元素 4")
	}

	fmt.Println(s.Values())
}

func TestIntersect(t *testing.T) {
	var s1 = NewSet(1, 2, 3)
	var s2 = NewSet(5, 6, 1, 3)

	// 1, 3
	fmt.Println(s1.Intersect(s2))
}

func TestUnion(t *testing.T) {
	var s1 = NewSet(1, 2, 3)
	var s2 = NewSet(5, 6, 1, 3)

	// 1, 2, 3, 5, 6
	fmt.Println(s1.Union(s2))
}

func TestDifference(t *testing.T) {
	var s1 = NewSet(1, 2, 3)
	var s2 = NewSet(5, 6, 1, 3)

	// 2
	fmt.Println(s1.Difference(s2))
}