package slist

import (
	"container/list"
	"sync"
)

type List interface {
	Len() int

	Front() *list.Element
	Back() *list.Element

	Remove(e *list.Element) interface{}
	PushFront(v interface{}) *list.Element
	PushBack(v interface{}) *list.Element

	InsertBefore(v interface{}, mark *list.Element) *list.Element
	InsertAfter(v interface{}, mark *list.Element) *list.Element

	MoveToFront(e *list.Element)
	MoveToBack(e *list.Element)
	MoveBefore(e, mark *list.Element)
	MoveAfter(e, mark *list.Element)

	PopFront() *list.Element
	PopBack() *list.Element
}

type syncList struct {
	list *list.List
	mu   sync.RWMutex
}

func New() *syncList {
	var sl = &syncList{}
	sl.list = list.New()
	return sl
}

func (this *syncList) Len() int {
	this.mu.RLock()
	defer this.mu.RUnlock()
	return this.list.Len()
}

func (this *syncList) Front() *list.Element {
	this.mu.RLock()
	defer this.mu.RUnlock()
	return this.list.Front()
}

func (this *syncList) Back() *list.Element {
	this.mu.RLock()
	defer this.mu.RUnlock()
	return this.list.Back()
}

func (this *syncList) Remove(e *list.Element) interface{} {
	this.mu.Lock()
	defer this.mu.Unlock()
	return this.list.Remove(e)
}

func (this *syncList) PushFront(v interface{}) *list.Element {
	this.mu.Lock()
	defer this.mu.Unlock()
	return this.list.PushFront(v)
}

func (this *syncList) PushBack(v interface{}) *list.Element {
	this.mu.Lock()
	defer this.mu.Unlock()
	return this.list.PushBack(v)
}

func (this *syncList) InsertBefore(v interface{}, mark *list.Element) *list.Element {
	this.mu.Lock()
	defer this.mu.Unlock()
	return this.list.InsertBefore(v, mark)
}

func (this *syncList) InsertAfter(v interface{}, mark *list.Element) *list.Element {
	this.mu.Lock()
	defer this.mu.Unlock()
	return this.list.InsertAfter(v, mark)
}

func (this *syncList) MoveToFront(e *list.Element) {
	this.mu.Lock()
	defer this.mu.Unlock()
	this.list.MoveToFront(e)
}

func (this *syncList) MoveToBack(e *list.Element) {
	this.mu.Lock()
	defer this.mu.Unlock()
	this.list.MoveToBack(e)
}

func (this *syncList) MoveBefore(e, mark *list.Element) {
	this.mu.Lock()
	defer this.mu.Unlock()
	this.list.MoveBefore(e, mark)
}

func (this *syncList) MoveAfter(e, mark *list.Element) {
	this.mu.Lock()
	defer this.mu.Unlock()
	this.list.MoveAfter(e, mark)
}

func (this *syncList) PopFront() *list.Element {
	this.mu.Lock()
	defer this.mu.Unlock()

	var item = this.list.Front()
	if item != nil {
		this.list.Remove(item)
	}
	return item
}

func (this *syncList) PopBack() *list.Element {
	this.mu.Lock()
	defer this.mu.Unlock()

	var item = this.list.Back()
	if item != nil {
		this.list.Remove(item)
	}
	return item
}
