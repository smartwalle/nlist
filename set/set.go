package set

import (
	"fmt"
	"sync"
)

type Set interface {
	// Add 向集合添加元素
	Add(values ...interface{})

	// Remove 从集合移除指定的元素
	Remove(values ...interface{})

	// RemoveAll 从集合移除所有的元素
	RemoveAll()

	// Exists 判断集合是否包含指定元素
	Exists(v interface{}) bool

	// Contains 判断集合是否包含指定的元素,包含所有的元素才会返回 true, 否则返回 false
	Contains(values ...interface{}) bool

	// Len 返回集合元素的长度
	Len() int

	// Values 返回集合所有元素组成的 Slice
	Values() []interface{}

	// Iter 返回集合的所有元素
	Iter() <-chan interface{}

	// Equal 判断和另一个集合是否相等
	Equal(s Set) bool

	// Clone 复制一个集合
	Clone() Set

	// Intersect 交集
	Intersect(s Set) Set

	// Union 并集
	Union(s Set) Set

	// Difference 差集
	Difference(s Set) Set
}

type set struct {
	m     map[interface{}]struct{}
	rw    sync.RWMutex
	block bool
}

func New(block bool, values ...interface{}) Set {
	var s = &set{}
	s.block = block
	s.m = make(map[interface{}]struct{})
	if len(values) > 0 {
		s.Add(values...)
	}
	return s
}

func (this *set) lock() {
	if this.block {
		this.rw.Lock()
	}
}

func (this *set) unlock() {
	if this.block {
		this.rw.Unlock()
	}
}

func (this *set) rLock() {
	if this.block {
		this.rw.RLock()
	}
}

func (this *set) rUnlock() {
	if this.block {
		this.rw.RUnlock()
	}
}

func (this *set) Add(values ...interface{}) {
	this.lock()
	defer this.unlock()

	for _, v := range values {
		this.m[v] = struct{}{}
	}
}

func (this *set) Remove(values ...interface{}) {
	this.lock()
	defer this.unlock()

	for _, v := range values {
		delete(this.m, v)
	}
}

func (this *set) RemoveAll() {
	this.lock()
	defer this.unlock()

	for k := range this.m {
		delete(this.m, k)
	}
}

func (this *set) Exists(v interface{}) bool {
	this.rLock()
	defer this.rUnlock()

	_, found := this.m[v]
	return found
}

func (this *set) Contains(values ...interface{}) bool {
	this.rLock()
	defer this.rUnlock()

	for _, v := range values {
		if _, found := this.m[v]; !found {
			return false
		}
	}
	return true
}

func (this *set) Len() int {
	this.rLock()
	defer this.rUnlock()

	return this.len()
}

func (this *set) len() int {
	return len(this.m)
}

func (this *set) Values() []interface{} {
	this.rLock()
	defer this.rUnlock()

	var vs = make([]interface{}, 0, this.len())
	for k := range this.m {
		vs = append(vs, k)
	}
	return vs
}

func (this *set) Iter() <-chan interface{} {
	var ch = make(chan interface{})

	go func(s *set) {
		if s.block {
			s.rLock()
		}

		for k := range this.m {
			ch <- k
		}

		close(ch)

		if s.block {
			s.rUnlock()
		}

	}(this)

	return ch
}

func (this *set) Equal(s Set) bool {
	this.rLock()
	defer this.rUnlock()

	if this.len() != s.Len() {
		return false
	}

	for k := range this.m {
		if !s.Exists(k) {
			return false
		}
	}

	return true
}

func (this *set) Clone() Set {
	this.rLock()
	defer this.rUnlock()

	var ns = New(this.block)
	for k := range this.m {
		ns.Add(k)
	}
	return ns
}

func (this *set) Intersect(s Set) Set {
	this.rLock()
	defer this.rUnlock()

	var ns = New(this.block)
	var vs = s.Values()
	for _, v := range vs {
		_, exists := this.m[v]
		if exists {
			ns.Add(v)
		}
	}
	return ns
}

func (this *set) Union(s Set) Set {
	this.rLock()
	defer this.rUnlock()

	var ns = New(this.block)
	ns.Add(this.Values()...)
	ns.Add(s.Values()...)
	return ns
}

func (this *set) Difference(s Set) Set {
	this.rLock()
	defer this.rUnlock()

	var ns = New(this.block)
	for k := range this.m {
		if !s.Contains(k) {
			ns.Add(k)
		}
	}
	return ns
}

func (this *set) String() string {
	return fmt.Sprint(this.Values()...)
}
