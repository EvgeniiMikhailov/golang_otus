package hw04_lru_cache //nolint:golint,stylecheck

type Key string

type Cache interface {
	Set(key string, value interface{}) bool
	Get(key string) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    list
	items    map[string]listItem
}

func (c *lruCache) Set(key string, value interface{}) bool {
	cachedItem, ok := c.items[key]
	newCacheItem := cacheItem{key: key, value: value}
	if ok {
		c.queue.MoveToFront(&cachedItem)
		cachedItem.Value = newCacheItem
	} else {
		cachedItem = listItem{Value: newCacheItem}
		c.queue.PushFront(cachedItem)
		if c.queue.Len() > c.capacity {
			c.queue.Remove(c.queue.Back())
		}
	}
	return ok
}

func (c *lruCache) Get(key string) (interface{}, bool) {
	return 0, true
}

func (c *lruCache) Clear() {

}

type cacheItem struct {
	key   string
	value interface{}
}

func NewCache(capacity int) Cache {
	cache := &lruCache{}
	cache.capacity = capacity
	cache.items = make(map[string]listItem)
	return cache
}
