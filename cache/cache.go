package cache

import "time"

// Element is an element of a cache.
type Element struct {
	Value     interface{}
	Timestamp time.Time
}

// Cache is a cache that stores elements for a certain amount of time.
type Cache struct {
	Lifetime int
	Elements map[string]Element
}

// Set sets the value associated with the key.
func (c *Cache) Set(key string, value interface{}) {
	c.Elements[key] = Element{
		Value:     value,
		Timestamp: time.Now().Add(time.Duration(c.Lifetime) * time.Second),
	}
}

// Get returns the value associated with the key and a boolean indicating whether the key exists.
func (c *Cache) Get(key string) (interface{}, time.Time, bool) {
	element, exists := c.Elements[key]

	if !exists {
		return nil, time.Time{}, false
	}

	if time.Since(element.Timestamp) >= 0 {
		c.Delete(key)
		return nil, time.Time{}, false
	}

	return element.Value, element.Timestamp, true
}

// Delete deletes the value associated with the key.
func (c *Cache) Delete(key string) {
	delete(c.Elements, key)
}

// Update deletes all expired elements.
func (c *Cache) Update() {
	for key, element := range c.Elements {
		if time.Since(element.Timestamp) >= 0 {
			c.Delete(key)
		}
	}
}

// NewCache creates a new cache with the specified lifetime.
func NewCache(lifetime int) *Cache {
	return &Cache{
		Lifetime: lifetime,
		Elements: make(map[string]Element),
	}
}
