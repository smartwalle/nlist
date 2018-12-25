package container

import (
	"fmt"
	"testing"
)

//func TestSliceContains(t *testing.T) {
//	var s1 = []interface{}{"a", "b", "e", "f", nil, "a"}
//	fmt.Println(Contains(s1, nil))
//	fmt.Println(Contains(s1, "aa"))
//}
//
//func TestMapContains(t *testing.T) {
//	var m1 = map[interface{}]string{}
//	m1["k1"] = "v1"
//	m1["k2"] = "v2"
//
//	fmt.Println(Contains(m1, "k2"))
//}
//
//func TestRemoveMapKey(t *testing.T)  {
//	var m1 = map[string]string{}
//	m1["k1"] = "v1"
//	m1["k2"] = "v2"
//
//	fmt.Println(m1)
//	Remove(m1, "k2")
//	fmt.Println(m1)
//}

func TestRemoveSliceKey(t *testing.T) {
	var s1 = []interface{}{"a", "b", "e", "f", nil, "a", "b"}

	fmt.Println(s1)
	RemoveAll(&s1, "a")
	RemoveAll(&s1, "b")
	//RemoveAll(nil, "b")
	fmt.Println(s1)
}

//func TestSliceIndex(t *testing.T) {
//	var s1 = []interface{}{"a", "b", "e", "f", nil, "a"}
//	fmt.Println(Index(s1, nil))
//	fmt.Println(Index(s1, "a"))
//}
