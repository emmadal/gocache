package gcache

import (
	"container/heap"
	"strconv"
	"sync"
	"time"
)

type Item struct {
	Value     interface{}
	ExpiresAt time.Time
	Index     int // required by heap.Interface
}

type Cache struct {
	mu    sync.Mutex
	ttl   time.Duration
	items map[string]*Item
	heap  expirationHeap
}

func New(ttl time.Duration) *Cache {
	cache := &Cache{
		ttl:   ttl,
		items: make(map[string]*Item),
	}
	go cache.cleanEvic()
	return cache
}

func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	//item, exists := c.items[key]
	item := &Item{
		Value:     value,
		ExpiresAt: time.Now().Add(c.ttl),
	}
	c.items[key] = item
	heap.Push(&c.heap, item)
	c.mu.Unlock()

	// if exists {
	// 	item.Value = value
	// 	item.ExpiresAt = time.Now().Add(c.ttl)
	// 	//heap.Fix(&c.heap, item.Index)
	// } else {
	// 	item = &Item{
	// 		Value:     value,
	// 		ExpiresAt: time.Now().Add(c.ttl),
	// 	}
	// 	c.items[key] = item
	// 	//heap.Push(&c.heap, item)
	// }
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	item, exists := c.items[key]
	c.mu.Unlock()
	if exists {
		if time.Now().After(item.ExpiresAt) {
			//delete(c.items, key)
			// if item.Index > 0 {
			// 	// heap.Remove(&c.heap, item.Index)
			// 	// c.Delete(key)
			// }
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
		//heap.Remove(&c.heap, item.Index)
		delete(c.items, key)
	}
	c.mu.Unlock()
}

func (c *Cache) cleanEvic() {
	for {
		time.Sleep(c.ttl)
		c.evict()
	}
}

func (c *Cache) evict() {
	c.mu.Lock()
	defer c.mu.Unlock()
	now := time.Now()
	for len(c.heap) > 0 && c.heap[0].ExpiresAt.Before(now) {
		item := heap.Pop(&c.heap).(*Item)

		switch item.Value.(type) {
		case string:
			delete(c.items, item.Value.(string))
		case int:
			delete(c.items, strconv.Itoa(item.Value.(int)))
		}
	}
}

type expirationHeap []*Item

func (h expirationHeap) Len() int {
	return len(h)
}

func (h expirationHeap) Less(i, j int) bool {
	return h[i].ExpiresAt.Before(h[j].ExpiresAt)
}

func (h expirationHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].Index = i
	h[j].Index = j
}

func (h *expirationHeap) Push(x interface{}) {
	n := len(*h)
	item := x.(*Item)
	item.Index = n
	*h = append(*h, item)
}

func (h *expirationHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	item.Index = -1 // for safety
	*h = old[0 : n-1]
	return item
}
