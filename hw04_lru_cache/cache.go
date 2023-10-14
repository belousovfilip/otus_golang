package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type queueItem struct {
	key   Key
	value interface{}
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

func (c *lruCache) Set(key Key, value interface{}) bool {
	item, exists := c.items[key]
	if exists {
		qItem := item.Value.(queueItem)
		qItem.value = value
		item.Value = qItem
		c.queue.MoveToFront(item)
		return true
	}
	if c.capacity == len(c.items) {
		last := c.queue.Back()
		qItem := last.Value.(queueItem)
		delete(c.items, qItem.key)
		c.queue.Remove(last)
	}
	c.items[key] = c.queue.PushFront(queueItem{
		key:   key,
		value: value,
	})
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	item, exists := c.items[key]
	if exists {
		c.queue.MoveToFront(item)
		qItem := item.Value.(queueItem)
		return qItem.value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}
