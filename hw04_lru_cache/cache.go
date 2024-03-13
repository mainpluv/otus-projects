package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, v interface{}) bool {
	if _, exists := c.items[key]; exists {
		c.items[key].Value = v
		c.queue.MoveToFront(c.items[key])
		return true
	}
	newItem := c.queue.PushFront(key)
	c.items[key] = newItem
	c.items[key].Key = key
	newItem.Value = v

	if c.queue.Len() > c.capacity {
		oldest := c.queue.Back()
		delete(c.items, oldest.Key)
		c.queue.Remove(oldest)
	}
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if _, exists := c.items[key]; exists {
		c.queue.MoveToFront(c.items[key])
		return c.items[key].Value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}
