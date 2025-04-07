# ⚡ gocache – High-Performance 

[![GoDoc](https://godoc.org/github.com/jeffotoni/gocache?status.svg)](https://godoc.org/github.com/jeffotoni/gocache) [![Go Report](https://goreportcard.com/badge/github.com/jeffotoni/gocache)](https://goreportcard.com/report/github.com/jeffotoni/gocache) [![License](https://img.shields.io/github/license/jeffotoni/gocache)](https://github.com/jeffotoni/gocache/blob/main/LICENSE) ![GitHub last commit](https://img.shields.io/github/last-commit/jeffotoni/gocache) ![GitHub contributors](https://img.shields.io/github/contributors/jeffotoni/gocache)
[![CircleCI](https://dl.circleci.com/status-badge/img/gh/jeffotoni/gocache/tree/main.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/jeffotoni/gocache/tree/main)
[![Coverage Status](https://coveralls.io/repos/github/jeffotoni/gocache/badge.svg)](https://coveralls.io/github/jeffotoni/gocache)
![GitHub stars](https://img.shields.io/github/forks/jeffotoni/gocache?style=social) 
![GitHub stars](https://img.shields.io/github/stars/jeffotoni/gocache)

---

## 🚀 About gocache

**GoCache** is a high-performance, **sharded** in-memory cache for Go applications, built for **speed and efficiency**. This implementation is **optimized for concurrent read/write operations**, leveraging **sharding, lock optimizations, and expiration management** to outperform other Go-based caches.

## 🔥 **Why use gocache?**  
✅ **Ultra-fast read & write operations**  
✅ **Sharded architecture for minimal lock contention**  
✅ **Automatic expiration & cleanup of stale data**  
✅ **Simple API with support for TTL (Time-To-Live)**  
✅ **Benchmarked & optimized for real-world workloads**

---

## 📦 Installation

```sh
$ go get github.com/jeffotoni/gocache
```
## 🔥 Quick Start

```go
package main

import (
	"fmt"
	"time"
	"github.com/jeffotoni/gocache"
)

func main() {
	// Create cache with a 10-minute TTL
	cache := gocache.New(10 * time.Minute)
	
	// Store items in cache
	cache.Set("key1", "Hello, gocache!", gocache.DefaultExpiration)
	cache.Set("key2", 12345, gocache.DefaultExpiration)
	
	var myPerson = struct {
		ID int
		Name   string
	}{
		ID: 564,
		Name:   "@jeffotoni",
	}
	cache.Set("key3", myPerson, gocache.DefaultExpiration)
	
	// Retrieve items
	val, found := cache.Get("key1")
	if found {
		fmt.Println("Found key1:", val)
	}

	val, found = cache.Get("key2")
	if found {
		fmt.Println("Found key2:", val)
	}
	
	val, found = cache.Get("key3")
	if found {
		fmt.Println("Found key3:", val)
	}
	
	// Deleting an item
	cache.Delete("key1")

	// Attempting to retrieve deleted item
	val, found = cache.Get("key1")
	if !found {
		fmt.Println("key1 not found in cache")
	}
}
```

## ⚡ Benchmark Results now with Risttreto, BigCache, Go-cache e Freecache

### 🚀 5-Second Benchmarks

| **Implementation**     | **Set Ops**     | **Set ns/op** | **Set/Get Ops** | **Set/Get ns/op** | **Observations**                                 |
|------------------------|------------------|----------------|------------------|--------------------|--------------------------------------------------|
| **gocache V1**         | 28,414,197       | 338.6 ns/op    | 22,687,808       | 294.9 ns/op        | Baseline version — decent speed, moderate allocs |
| **gocache V8**         | 26,022,742       | 364.5 ns/op    | 15,105,789       | 393.6 ns/op        | High memory cost, TTL enabled                    |
| **gocache V9**         | 44,026,141       | 265.4 ns/op    | 23,528,972       | 270.0 ns/op        | 🏆 **Fastest write throughput**                  |
| **gocache V10**        | 19,749,439       | 393.2 ns/op    | 16,217,510       | 495.9 ns/op        | ❌ Higher allocation and latency                  |
| **gocache V11 (Short)**| 39,719,458       | 264.2 ns/op    | 23,308,189       | 265.4 ns/op        | ⚡ Short TTL — very fast overall                  |
| **gocache V11 (Long)** | 22,334,095       | 348.8 ns/op    | 18,338,124       | 319.7 ns/op        | Balanced long TTL setup                          |
| **go-cache**           | 25,669,981       | 392.5 ns/op    | 20,485,022       | 306.0 ns/op        | Stable, but slower than newer gocache versions   |
| **freecache**          | 41,543,706       | 380.3 ns/op    | 14,433,577       | 425.2 ns/op        | 🚀 Fast writes, significantly slower reads        |
| **ristretto**          | 30,257,541       | 352.3 ns/op    | 10,055,701       | 547.8 ns/op        | 🧠 TinyLFU eviction, high allocation per op       |
| **bigcache**           | 30,260,250       | 320.6 ns/op    | 14,382,721       | 354.6 ns/op        | 🔥 Very consistent, low GC overhead               |

### 🚀 1-Second Benchmarks

```go
$ go test -bench=. -benchtime=1s
```

| **Implementation** | **Set Ops**    | **Set ns/op** | **Set/Get Ops** | **Set/Get ns/op** | **Observations**                      |
|--------------------|----------------|---------------|-----------------|-------------------|---------------------------------------|
| **gocache V7**     | 8,026,825      | 222.4 ns/op   | 4,978,083       | 244.3 ns/op       | 🏆 **Best write** (1s), fast reads    |
| **gocache V9**     | 9,295,434      | 215.9 ns/op   | 5,096,511       | 272.7 ns/op       | 🏆 **Fastest write** (lowest ns/op)   |
| **go-cache**       | 6,463,236      | 291.6 ns/op   | 4,698,109       | 290.7 ns/op       | Solid library, slower than V7/V9      |
| **freecache**      | 5,803,242      | 351.1 ns/op   | 2,183,834       | 469.7 ns/op       | 🚀 Decent writes, poor reads          |

### 🚀 3-Second Benchmarks

| **Implementation** | **Set Ops**     | **Set ns/op** | **Get Ops**     | **Get ns/op** | **Observations**                     |
|--------------------|-----------------|---------------|-----------------|---------------|--------------------------------------|
| **gocache V7**     | 27,229,544      | 252.4 ns/op   | 14,574,768      | 268.6 ns/op   | 🏆 **Best write** (3s)               |
| **gocache V9**     | 24,809,947      | 252.1 ns/op   | 13,225,228      | 275.7 ns/op   | 🏆 **Very fast write**, good read    |
| **go-cache**       | 15,594,752      | 375.4 ns/op   | 14,289,182      | 269.7 ns/op   | 🚀 Excellent reads, slower writes    |
| **freecache**      | 13,303,050      | 402.3 ns/op   | 8,903,779       | 421.4 ns/op   | ❌ Decent write, slow read           |

#### 🚀 You can find all the benchmarks here [benchmark-gocache](https://github.com/jeffotoni/benchmark-gocache)

# 🛠 API Reference

New Cache
```go
cache := gocache.New(10 * time.Minute) // Creates a cache with 10 min TTL
```

Set Value
```go
cache.Set("key", "value", gocache.DefaultExpiration)
```

Get Value
```go
val, found := cache.Get("key")
```

Delete Value
```go
cache.Delete("key")
```

## ⚖️ Trade-Offs & Performance

✅ V7 and V9 provide the fastest write performance.

✅ go-cache excel in retrieval speed.

❌ FreeCache struggles with read speed, despite having decent write speed.

## 🤝 Contributing

**🚀 Want to improve gocache? Follow these simple steps:**

 1️⃣ Fork this repo and add your own cache optimizations.

 2️⃣ Submit a Pull Request (PR) with your improvements.

 3️⃣ Open an issue if you have questions or ideas for enhancements.

**Your contributions are always welcome! 💡🔥**

## 📜 License

This project is **open-source** under the **MIT License**.

💡 Feel free to **fork, modify, and experiment** with these benchmarks in your own applications or libraries.  
🔬 The goal is to **help developers choose the best in-memory cache** for their needs.

