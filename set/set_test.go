package set

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	var s = New(true, 1, 2, 3, 5, 6, 7, 7)

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
	var s1 = New(true, 1, 2, 3)
	var s2 = New(true, 5, 6, 1, 3)

	// 1, 3
	fmt.Println(s1.Intersect(s2))
}

func TestUnion(t *testing.T) {
	var s1 = New(true, 1, 2, 3)
	var s2 = New(true, 5, 6, 1, 3)

	// 1, 2, 3, 5, 6
	fmt.Println(s1.Union(s2))
}

func TestDifference(t *testing.T) {
	var s1 = New(true, 1, 2, 3)
	var s2 = New(true, 5, 6, 1, 3)

	// 2
	fmt.Println(s1.Difference(s2))
}

func TestIter(t *testing.T) {
	fmt.Println("=====TestIter=====")
	var s1 = New(true, 1, 2, 3, 4, 5)

	for v := range s1.Iter() {
		fmt.Println(v)
	}
}

func TestEqual(t *testing.T) {
	var s1 = New(true, 1, 2, 3)
	var s2 = New(true, 2, 1, 3)

	if !s1.Equal(s2) {
		t.Error("s1 与 s2 应该相等")
	}

	s2.Add("5")
	s1.Add(5)

	if s1.Equal(s2) {
		t.Error("s1 与 s2 不应该相等")
	}
}

func TestClone(t *testing.T) {
	var s1 = New(true, 1, 2, 3)
	var s2 = s1.Clone()

	if !s1.Equal(s2) {
		t.Error("s1 与 s2 应该相等")
	}

	s2.Add(4)

	if s1.Equal(s2) {
		t.Error("s1 与 s2 不应该相等")
	}
}
