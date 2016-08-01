package smap

import (
	"fmt"
	"sync"
)

type Map interface {
	// Set 添加一组键值对
	Set(key, value interface{})

	// Remove 移除指定的 key 及其 value
	Remove(key interface{})

	// RemoveAll 移除所有的 key 及 value
	RemoveAll()

	// Exists 判断指定的 key 是否存在
	Exists(key interface{}) bool

	// Contains 判断指定的 key 列表是否存在,只有当所有的 key 都存在的时候,才会返回 true
	Contains(keys ...interface{}) bool

	// Len 返回元素的个数
	Len() int

	// Value 获取指定 key 的 value
	Value(key interface{}) interface{}

	// Keys 返回由所有 key 组成的 Slice
	Keys() []interface{}

	// Values 返回由所有 value 组成的 Slice
	Values() []interface{}

	// Iter 返回所有 key 及 value
	Iter() <-chan mapValue
}

type mapValue struct {
	Key   interface{}
	Value interface{}
}

type syncMap struct {
	m     map[interface{}]interface{}
	rw    sync.RWMutex
	block bool
}

func New(block bool) Map {
	var sm = &syncMap{}
	sm.block = block
	sm.m = make(map[interface{}]interface{})
	return sm
}

func (this *syncMap) lock() {
	if this.block {
		this.rw.Lock()
	}
}

func (this *syncMap) unlock() {
	if this.block {
		this.rw.Unlock()
	}
}

func (this *syncMap) rLock() {
	if this.block {
		this.rw.RLock()
	}
}

func (this *syncMap) rUnlock() {
	if this.block {
		this.rw.RUnlock()
	}
}

func (this *syncMap) Set(key, value interface{}) {
	this.lock()
	defer this.unlock()

	this.m[key] = value
}

func (this *syncMap) Remove(key interface{}) {
	this.rLock()
	defer this.rUnlock()

	delete(this.m, key)
}

func (this *syncMap) RemoveAll() {
	this.lock()
	defer this.unlock()

	for k := range this.m {
		delete(this.m, k)
	}
}

func (this *syncMap) Exists(key interface{}) bool {
	this.rLock()
	defer this.rUnlock()

	_, found := this.m[key]
	return found
}

func (this *syncMap) Contains(keys ...interface{}) bool {
	this.rLock()
	defer this.rUnlock()

	for _, k := range keys {
		if _, found := this.m[k]; !found {
			return false
		}
	}
	return true
}

func (this *syncMap) Len() int {
	this.rLock()
	defer this.rUnlock()

	return this.len()
}

func (this *syncMap) len() int {
	return len(this.m)
}

func (this *syncMap) Value(key interface{}) interface{} {
	this.rLock()
	defer this.rUnlock()

	return this.m[key]
}

func (this *syncMap) Keys() []interface{} {
	this.rLock()
	defer this.rUnlock()
	var keys = make([]interface{}, 0, 0)

	for k := range this.m {
		keys = append(keys, k)
	}
	return keys
}

func (this *syncMap) Values() []interface{} {
	this.rLock()
	defer this.rUnlock()
	var values = make([]interface{}, 0, 0)

	for _, v := range this.m {
		values = append(values, v)
	}
	return values
}

func (this *syncMap) Iter() <-chan mapValue {
	var iv = make(chan mapValue)

	go func(m *syncMap) {
		if m.block {
			m.rLock()
		}

		for k, v := range this.m {
			iv <- mapValue{k, v}
		}

		close(iv)

		if m.block {
			m.rUnlock()
		}
	}(this)

	return iv
}

func (this *syncMap) String() string {
	return fmt.Sprint(this.m)
}
