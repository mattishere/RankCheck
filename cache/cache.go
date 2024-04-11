package cache

import "time"

type Element struct {
	Value     interface{}
	Timestamp time.Time
}

type Cache struct {
	Lifetime int
	Elements map[string]Element
}

func (c *Cache) Set(key string, value interface{}) {
	c.Elements[key] = Element{
		Value:     value,
		Timestamp: time.Now().Add(time.Duration(c.Lifetime) * time.Second),
	}
}

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

func (c *Cache) Delete(key string) {
	delete(c.Elements, key)
}

func (c *Cache) Update() {
	for key, element := range c.Elements {
		if time.Since(element.Timestamp) >= 0 {
			c.Delete(key)
		}
	}
}

func NewCache(lifetime int) *Cache {
	return &Cache{
		Lifetime: lifetime,
		Elements: make(map[string]Element),
	}
}
