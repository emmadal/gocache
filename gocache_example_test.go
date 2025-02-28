// Package gocache provides a cache storage solution with expiration handling.
// It is designed for high-performance applications that require fast data retrieval
// with minimal contention.
//
// This file includes usage examples that are documented and available in GoDoc.

package gocache

import (
	"fmt"
	"time"
)

// This function is named ExampleNew()
// it with the Examples type.
func ExampleNew() {
	// Create a cache with a TTL of 10 seconds
	cache := New(10 * time.Second)

	// Print the default TTL of the cache
	fmt.Println("Default TTL:", cache.ttl)

	// Out put: Default TTL: 10s
}
