package gcache

import (
	"sync"
	"time"
)

const (
	NoExpiration   time.Duration = -1
	DefaultExpires time.Duration = 0
)

type Item struct {
	Value   interface{} // 8 bytes on 64-bit systems
	Expires int64       //  8 bytes
}

type Cache struct {
	mu    sync.RWMutex
	ttl   time.Duration
	items map[string]*Item
}

func New(ttl time.Duration) *Cache {
	cache := &Cache{
		ttl:   ttl,
		items: make(map[string]*Item),
	}
	// go cache.cleanEvic()
	return cache
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	var t int64
	if ttl == DefaultExpires {
		ttl = DefaultExpires
	} else if ttl > 0 {
		t = time.Now().Add(ttl).UnixNano()
	}

	c.mu.Lock()
	item := &Item{
		Value:   value,
		Expires: t,
	}
	c.items[key] = item
	c.mu.Unlock()
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	item, exists := c.items[key]
	c.mu.RUnlock()
	if exists {
		if item.Expires > 0 && time.Now().UnixNano() > item.Expires {
			return nil, false
		}
		return item.Value, true
	}
	return nil, false
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	_, exists := c.items[key]
	if exists {
		delete(c.items, key)
	}
	c.mu.Unlock()
}

func (c *Cache) cleanCache() {
	for {
		time.Sleep(c.ttl)
		c.clean()
	}
}

func (c *Cache) clean() {
	c.mu.Lock()
	defer c.mu.Unlock()
}
