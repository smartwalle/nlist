package smap

import (
	"sync"
	"fmt"
)

type SyncMap interface {
	//
	Set(key, value interface{})

	Remove(key interface{})

	RemoveAll()

	Exists(key interface{}) bool

	Contains(keys ...interface{}) bool

	Len() int

	Value(key interface{}) interface{}

	Keys() []interface{}

	Values() []interface{}
}

type syncMap struct {
	m     map[interface{}]interface{}
	rw    sync.RWMutex
	block bool
}

func NewMap() SyncMap {
	return newMap(false)
}

func NewSyncMap() SyncMap {
	return newMap(true)
}

func newMap(block bool) SyncMap {
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

	for k, _ := range this.m {
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

	for k, _ := range this.m {
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

func (this *syncMap) String() string {
	return fmt.Sprint(this.m)
}