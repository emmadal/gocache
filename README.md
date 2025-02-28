# ⚡ GoCache – High-Performance In-Memory Cache for Go

[![GoDoc](https://godoc.org/github.com/jeffotoni/gocache?status.svg)](https://godoc.org/github.com/jeffotoni/gocache)  
[![Go Report](https://goreportcard.com/badge/github.com/jeffotoni/gocache)](https://goreportcard.com/report/github.com/jeffotoni/gocache)  
[![License](https://img.shields.io/github/license/jeffotoni/gocache)](https://github.com/jeffotoni/gocache/blob/main/LICENSE)  
![GitHub last commit](https://img.shields.io/github/last-commit/jeffotoni/gocache)  
![GitHub contributors](https://img.shields.io/github/contributors/jeffotoni/gocache) 
![GitHub forks](https://img.shields.io/github/forks/jeffotoni/gocache?style=social) 
![GitHub stars](https://img.shields.io/github/stars/jeffotoni/gocache) 
---

## 🚀 About GoCache

**GoCache** is a high-performance, **sharded** in-memory cache for Go applications, built for **speed and efficiency**. This implementation is **optimized for concurrent read/write operations**, leveraging **sharding, lock optimizations, and expiration management** to outperform other Go-based caches.

## 🔥 **Why use GoCache?**  
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
	cache.Set("key1", "Hello, GoCache!", gocache.DefaultExpiration)
	cache.Set("key2", 12345, gocache.DefaultExpiration)

	// Retrieve items
	val, found := cache.Get("key1")
	if found {
		fmt.Println("Found key1:", val)
	}

	val, found = cache.Get("key2")
	if found {
		fmt.Println("Found key2:", val)
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

## ⚡ Benchmark Results

### 🚀 1-Second Benchmarks

$ go test -bench=. -benchtime=1s

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

## 🚀 Want to improve GoCache? Follow these simple steps:

 1️⃣ Fork this repo and add your own cache optimizations.

 2️⃣ Submit a Pull Request (PR) with your improvements.

 3️⃣ Open an issue if you have questions or ideas for enhancements.

**Your contributions are always welcome! 💡🔥**

## 📜 License

This project is **open-source** under the **MIT License**.

💡 Feel free to **fork, modify, and experiment** with these benchmarks in your own applications or libraries.  
🔬 The goal is to **help developers choose the best in-memory cache** for their needs.

