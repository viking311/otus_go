package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity    int
	queue       List
	items       map[Key]*ListItem
	reverseDict map[*ListItem]Key
	mx          sync.Mutex
}

func (lc *lruCache) Set(key Key, value interface{}) bool {
	lc.mx.Lock()
	defer lc.mx.Unlock()

	item, ok := lc.items[key]

	if ok {
		item.Value = value
		lc.queue.MoveToFront(item)
	} else {
		if lc.queue.Len() == lc.capacity {
			back := lc.queue.Back()
			lc.queue.Remove(back)
			backKey := lc.reverseDict[back]
			delete(lc.items, backKey)
			delete(lc.reverseDict, back)
		}
		lc.items[key] = lc.queue.PushFront(value)
		lc.reverseDict[lc.items[key]] = key
	}

	return ok
}

func (lc *lruCache) Get(key Key) (interface{}, bool) {
	lc.mx.Lock()
	defer lc.mx.Unlock()

	if item, ok := lc.items[key]; ok {
		lc.queue.MoveToFront(item)
		return item.Value, true
	}

	return nil, false
}

func (lc *lruCache) Clear() {
	lc.mx.Lock()
	defer lc.mx.Unlock()

	lc.reverseDict = make(map[*ListItem]Key, lc.capacity)
	lc.items = make(map[Key]*ListItem, lc.capacity)
	for lc.queue.Len() > 0 {
		lc.queue.Remove(lc.queue.Back())
	}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity:    capacity,
		queue:       NewList(),
		items:       make(map[Key]*ListItem, capacity),
		reverseDict: make(map[*ListItem]Key, capacity),
		mx:          sync.Mutex{},
	}
}
