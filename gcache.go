package gcache

import (
	"sync"
	"time"
)

const (
	NoExpiration   time.Duration = -1
	Defaultexpires time.Duration = 0
)

type Item struct {
	value   interface{} // 8 bytes on 64-bit systems
	expires int64       // 8 bytes
}

type cache struct {
	mu    sync.RWMutex
	ttl   time.Duration
	items map[string]*Item
}

type Cache struct {
	*cache
}

func New(ttl time.Duration) *Cache {
	return &Cache{
		cache: &cache{
			ttl:   ttl,
			items: make(map[string]*Item),
		},
	}
	// go cache.cleanEvic()
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	var t int64
	if ttl == Defaultexpires {
		ttl = Defaultexpires
	} else if ttl > 0 {
		t = time.Now().Add(ttl).UnixNano()
	}

	c.mu.Lock()
	item := &Item{
		value:   value,
		expires: t,
	}
	c.items[key] = item
	c.mu.Unlock()
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	item, exists := c.items[key]
	c.mu.RUnlock()
	if exists {
		if item.expires > 0 && time.Now().UnixNano() > item.expires {
			return nil, false
		}
		return item.value, true
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
