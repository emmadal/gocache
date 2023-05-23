package cache2

import (
	"errors"
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	NoExpiration      time.Duration = -1
	DefaultExpiration time.Duration = 0
)

type Item[V any] struct {
	object     V
	expiration int64
}

type cache[K ~string, V any] struct {
	mu         sync.RWMutex
	items      map[K]*Item[V]
	done       chan struct{}
	expTime    time.Duration
	cleanupInt time.Duration
}

type Cache[K ~string, V any] struct {
	*cache[K, V]
}

func newCache[K ~string, V any](expTime, cleanupInt time.Duration, item map[K]*Item[V]) *cache[K, V] {
	c := &cache[K, V]{
		mu:         sync.RWMutex{},
		items:      item,
		expTime:    expTime,
		cleanupInt: cleanupInt,
		done:       make(chan struct{}),
	}
	return c
}

func New[K ~string, V any](expTime, cleanupTime time.Duration) *Cache[K, V] {
	items := make(map[K]*Item[V])
	c := newCache(expTime, cleanupTime, items)

	if cleanupTime > 0 {
		go c.cleanup()
		runtime.SetFinalizer(c, stopCleanup[K, V])
	}

	return &Cache[K, V]{c}
}

func (c *Cache[K, V]) Set(key K, val V, d time.Duration) error {
	item, err := c.Get(key)
	if item != nil && err == nil {
		return fmt.Errorf("já existe um item com a chave '%v'. Use o método Atualizar", key)
	}
	c.add(key, val, d)

	return nil
}

func (c *Cache[K, V]) add(key K, val V, d time.Duration) error {
	var exp int64

	if d == DefaultExpiration {
		d = c.expTime
	}
	if d > 0 {
		exp = time.Now().Add(d).UnixNano()
	} else if d < 0 {
		exp = int64(NoExpiration)
	}

	item, err := c.Get(key)
	if item != nil && err != nil {
		return fmt.Errorf("item com chave '%v' já existe", key)
	}

	switch any(val).(type) {
	case string:
		if len(any(val).(string)) == 0 {
			return fmt.Errorf("valor do tipo string não pode estar vazio")
		}
	}

	c.mu.Lock()
	c.items[key] = &Item[V]{
		object:     val,
		expiration: exp,
	}
	c.mu.Unlock()

	return nil
}
func (c *Cache[K, V]) Get(key K) (*Item[V], error) {
	c.mu.RLock()
	if item, ok := c.items[key]; ok {
		if item.expiration > 0 {
			now := time.Now().UnixNano()
			if now > item.expiration {
				c.mu.RUnlock()
				return nil, fmt.Errorf("item com chave '%v' expirou", key)
			}
		}
		c.mu.RUnlock()
		return item, nil
	}
	c.mu.RUnlock()
	return nil, fmt.Errorf("item com chave '%v' não encontrado", key)
}

func (it *Item[V]) Val() V {
	var v V
	if it != nil {
		return it.object
	}
	return v
}

func (c *Cache[K, V]) IsExpired(key K) bool {
	item, err := c.Get(key)
	if item != nil && err != nil {
		if item.expiration > time.Now().UnixNano() {
			return true
		}
	}
	return false
}

func (c *cache[K, V]) cleanup() {
	tick := time.NewTicker(c.cleanupInt)

	for {
		select {
		case <-tick.C:
			c.DeleteExpired()
		case <-c.done:
			tick.Stop()
			return
		}
	}
}

func (c *cache[K, V]) delete(key K) error {
	if _, ok := c.items[key]; ok {
		delete(c.items, key)

		return nil
	}

	return fmt.Errorf("item com chave '%v' não existe", key)
}

func (c *cache[K, V]) DeleteExpired() error {
	var err error

	now := time.Now().UnixNano()

	c.mu.Lock()
	for k, item := range c.items {
		if now > item.expiration && item.expiration != int64(NoExpiration) {
			if e := c.delete(k); e != nil {
				err = errors.Join(err, e)
			}
		}

	}
	c.mu.Unlock()

	return errors.Unwrap(err)
}

func stopCleanup[K ~string, V any](c *cache[K, V]) {
	c.done <- struct{}{}
}
