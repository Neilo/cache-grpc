package cache

import (
	"getputer/internal/structure"
	"sync"
)

type Cache struct {
	sync.RWMutex
	items map[string]structure.User
}

func New() *Cache {
	items := make(map[string]structure.User)

	cache := Cache{
		items: items,
	}

	return &cache
}

func (c *Cache) Put(key string, value []byte) {
	c.Lock()
	defer c.Unlock()

	c.items[key] = structure.User{
		Value: value,
	}

}

func (c *Cache) Get(key string) string {
	c.RLock()
	defer c.RUnlock()

	item, found := c.items[key]

	if !found {
		return "Not found"
	}

	return string(item.Value)
}
