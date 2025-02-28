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

// This function is named ExampleCache_Set()
// it with the Examples type.
func ExampleCache_Set() {
	// Create a new cache with a default TTL of 10 seconds
	cache := New(10 * time.Second)

	// Store an item with the default expiration
	cache.Set("username", "john_doe", DefaultExpiration)

	// Store an item with a custom TTL of 5 minutes
	cache.Set("session_token", "xyz123", 5*time.Minute)

	// Retrieve and print the stored values
	user, found := cache.Get("username")
	if found {
		fmt.Println("Username:", user)
	} else {
		fmt.Println("Username not found")
	}

	token, found := cache.Get("session_token")
	if found {
		fmt.Println("Session Token:", token)
	} else {
		fmt.Println("Session Token not found")
	}

	// Out put: Username: john_doe
	// Session Token: xyz123
}

// This function is named ExampleCache_Get()
// it with the Examples type.
func ExampleCache_Get() {
	// Create a new cache with a default TTL of 2 seconds
	cache := New(2 * time.Second)

	// Store an item with a custom TTL of 1 second
	cache.Set("session_token", "xyz123", 1*time.Second)

	// Retrieve the item before expiration
	token, found := cache.Get("session_token")
	if found {
		fmt.Println("Before expiration:", token)
	} else {
		fmt.Println("Before expiration: Not found")
	}

	// Wait for the item to expire
	time.Sleep(2 * time.Second)

	// Try retrieving the item again after expiration
	token, found = cache.Get("session_token")
	if found {
		fmt.Println("After expiration:", token)
	} else {
		fmt.Println("After expiration: Not found")
	}

	// Out put:Before expiration: xyz123
	// After expiration: Not found
}

// This function is named ExampleCache_Delete()
// it with the Examples type.
func ExampleCache_Delete() {
	// Create a new cache with a default TTL of 10 seconds
	cache := New(10 * time.Second)

	// Store an item in the cache
	cache.Set("user_id", 42, DefaultExpiration)

	// Retrieve and print the stored value before deletion
	value, found := cache.Get("user_id")
	if found {
		fmt.Println("Before deletion:", value)
	} else {
		fmt.Println("Before deletion: Not found")
	}

	// Delete the item from the cache
	cache.Delete("user_id")

	// Try retrieving the item again after deletion
	value, found = cache.Get("user_id")
	if found {
		fmt.Println("After deletion:", value)
	} else {
		fmt.Println("After deletion: Not found")
	}

	// Out put:  Before deletion: 42
	// After deletion: Not found
}
