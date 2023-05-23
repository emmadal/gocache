package gcache

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

type Item2[V any] struct {
	object     V
	expiration int64
}

type cache2[K ~string, V any] struct {
	mu         sync.RWMutex
	items      map[K]*Item2[V]
	done       chan struct{}
	expTime    time.Duration
	cleanupInt time.Duration
}

type Cache2[K ~string, V any] struct {
	*cache2[K, V]
}

func newCache2[K ~string, V any](expTime, cleanupInt time.Duration, item map[K]*Item2[V]) *cache2[K, V] {
	c := &cache2[K, V]{
		mu:         sync.RWMutex{},
		items:      item,
		expTime:    expTime,
		cleanupInt: cleanupInt,
		done:       make(chan struct{}),
	}
	return c
}

func New2[K ~string, V any](expTime, cleanupTime time.Duration) *Cache2[K, V] {
	items := make(map[K]*Item2[V])
	c2 := newCache2(expTime, cleanupTime, items)

	if cleanupTime > 0 {
		go c2.cleanup()
		runtime.SetFinalizer(c2, stopCleanup[K, V])
	}

	return &Cache2[K, V]{c2}
}

func (c *Cache2[K, V]) Set2(key K, val V, d time.Duration) error {
	item, err := c.Get2(key)
	if item != nil && err == nil {
		return fmt.Errorf("já existe um item com a chave '%v'. Use o método Atualizar", key)
	}
	c.add(key, val, d)

	return nil
}

func (c *Cache2[K, V]) add(key K, val V, d time.Duration) error {
	var exp int64

	if d == DefaultExpiration {
		d = c.expTime
	}
	if d > 0 {
		exp = time.Now().Add(d).UnixNano()
	} else if d < 0 {
		exp = int64(NoExpiration)
	}

	item, err := c.Get2(key)
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
	c.items[key] = &Item2[V]{
		object:     val,
		expiration: exp,
	}
	c.mu.Unlock()

	return nil
}
func (c *Cache2[K, V]) Get2(key K) (*Item2[V], error) {
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

func (it *Item2[V]) Val() V {
	var v V
	if it != nil {
		return it.object
	}
	return v
}

func (c *Cache2[K, V]) IsExpired(key K) bool {
	item, err := c.Get2(key)
	if item != nil && err != nil {
		if item.expiration > time.Now().UnixNano() {
			return true
		}
	}
	return false
}

func (c *cache2[K, V]) cleanup() {
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

func (c *cache2[K, V]) delete(key K) error {
	if _, ok := c.items[key]; ok {
		delete(c.items, key)

		return nil
	}

	return fmt.Errorf("item com chave '%v' não existe", key)
}

func (c *cache2[K, V]) DeleteExpired() error {
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

func stopCleanup[K ~string, V any](c *Cache2[K, V]) {
	c.done <- struct{}{}
}
