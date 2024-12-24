package generator

import (
	"image"
	"sync"
)

// ConcurrentMap is a thread-safe map for storing image.Image objects.
type ConcurrentMap struct {
	mu   sync.RWMutex           // Mutex for synchronizing access to the map.
	data map[string]image.Image // The underlying map to store key-value pairs.
}

// NewConcurrentMap initializes and returns a new instance of ConcurrentMap.
func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		data: make(map[string]image.Image), // Create an empty map.
	}
}

// Set inserts or updates a key-value pair in the map.
func (c *ConcurrentMap) Set(key string, value image.Image) {
	c.mu.Lock()         // Acquire a write lock to ensure safe access.
	defer c.mu.Unlock() // Release the lock once the function returns.
	c.data[key] = value // Update the map.
}

// Get retrieves a value from the map by its key.
// If the key is not present, it fetches the image, stores it in the map asynchronously, and returns it.
func (c *ConcurrentMap) Get(key string) (image.Image, bool) {
	c.mu.RLock()         // Acquire a read lock for safe concurrent read.
	defer c.mu.RUnlock() // Release the lock once the function returns.

	value, ok := c.data[key] // Check if the key exists in the map.
	if ok {
		return value, true
	}

	// If the key is not found, fetch the image asynchronously.
	// Note: This is not strictly safe for concurrent updates on the same key.
	value = getImage(key) // Assumes getImage is a function to fetch or generate the image.

	// Use a goroutine to store the value in the map.
	go c.Set(key, value) // Non-blocking call to store the new value in the map.

	return value, true
}

// Delete removes a key-value pair from the map.
func (c *ConcurrentMap) Delete(key string) {
	c.mu.Lock()         // Acquire a write lock to ensure safe deletion.
	defer c.mu.Unlock() // Release the lock once the function returns.
	delete(c.data, key) // Delete the key from the map.
}
