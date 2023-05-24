package main

import (
	"log"
	"strconv"
	"testing"
	"time"

	freecache "github.com/coocood/freecache"
	gc1 "github.com/jeffotoni/gcache"
	gc2 "github.com/jeffotoni/gcache/v2"
	gocache "github.com/patrickmn/go-cache"
)

var c1 = gc1.New(10 * time.Second)
var c2 = gc2.New[string, int](time.Duration(time.Minute), 0)

var cache = gocache.New(10*time.Second, 1*time.Minute)

var fcacheSize = 100 * 1024 * 1024
var fcache = freecache.NewCache(fcacheSize)

func BenchmarkGcacheSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		key := strconv.Itoa(i)
		c1.Set(key, i, time.Duration(time.Minute))
	}
}

func BenchmarkGcacheSetGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		key := strconv.Itoa(i)
		c1.Set(key, i, time.Duration(time.Minute))
		i, ok := c1.Get(key)
		if !ok {
			log.Printf("Não encontrei: %v", i)
		}
	}
}

func BenchmarkGcacheSet2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		key := strconv.Itoa(i)
		c2.Set(key, i, time.Duration(time.Minute))
	}
}
func BenchmarkGcacheSetGet2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		key := strconv.Itoa(i)
		c2.Set(key, i, time.Duration(10*time.Second))
		i, ok := c2.Get(key)
		if ok != nil {
			log.Printf("Não encontrei: %v i=%v", ok, i)
		}
	}
}

func BenchmarkGo_cacheSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		key := strconv.Itoa(i)
		cache.Set(key, i, time.Duration(5*time.Second))
	}
}

func BenchmarkGo_cacheSetGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		key := strconv.Itoa(i)
		cache.Set(key, i, time.Duration(5*time.Second))
		i, ok := cache.Get(key)
		if !ok {
			log.Printf("Não encontrei: %v", i)
		}
	}
}

func BenchmarkFreeCacheSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		key := strconv.Itoa(i)
		fcache.Set([]byte(key), []byte(key), 3600)
	}
}

func BenchmarkFreeCacheSetGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		key := strconv.Itoa(i)
		fcache.Set([]byte(key), []byte(key), 3600)
		got, err := fcache.Get([]byte(key))
		if err != nil {
			log.Printf("\nError Get:%v %v", err, got)
		}
	}
}
