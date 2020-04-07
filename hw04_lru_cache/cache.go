package hw04_lru_cache //nolint:golint,stylecheck

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*listItem
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	cachedItem, ok := c.items[key]
	record := cacheItem{key: key, value: value}

	if ok {
		c.queue.MoveToFront(cachedItem)
		cachedItem.Value = record
	} else {
		c.queue.PushFront(record)
		c.items[key] = c.queue.Front()
		if c.queue.Len() > c.capacity {
			delete(c.items, c.queue.Back().Value.(cacheItem).key)
			c.queue.Remove(c.queue.Back())
		}
	}
	return ok
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	cachedItem, ok := c.items[key]
	if ok {
		c.queue.MoveToFront(cachedItem)
		return c.queue.Front().Value.(cacheItem).value, ok
	}
	return nil, ok
}

func (c *lruCache) Clear() {
	c.items = make(map[Key]*listItem)
	c.queue = NewList()
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	cache := &lruCache{capacity: capacity, queue: NewList(), items: make(map[Key]*listItem)}
	return cache
}
