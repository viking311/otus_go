package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type cacheItem struct {
	key   Key
	value interface{}
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mx       sync.Mutex
}

func (lc *lruCache) Set(key Key, value interface{}) bool {
	lc.mx.Lock()
	defer lc.mx.Unlock()

	item, ok := lc.items[key]

	if ok {
		ci := item.Value.(cacheItem)
		ci.value = value
		item.Value = ci
		lc.queue.MoveToFront(item)
		return ok
	}

	if lc.queue.Len() == lc.capacity {
		back := lc.queue.Back()
		lc.queue.Remove(back)
		ci := back.Value.(cacheItem)
		backKey := ci.key
		delete(lc.items, backKey)
	}

	ci := cacheItem{
		key:   key,
		value: value,
	}
	lc.items[key] = lc.queue.PushFront(ci)

	return ok
}

func (lc *lruCache) Get(key Key) (interface{}, bool) {
	lc.mx.Lock()
	defer lc.mx.Unlock()

	if item, ok := lc.items[key]; ok {
		lc.queue.MoveToFront(item)
		ci := item.Value.(cacheItem)
		return ci.value, true
	}

	return nil, false
}

func (lc *lruCache) Clear() {
	lc.mx.Lock()
	defer lc.mx.Unlock()

	lc.items = make(map[Key]*ListItem, lc.capacity)
	lc.queue = NewList()
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
		mx:       sync.Mutex{},
	}
}
