package cache

import "sync"

type Cache struct {
	mu     *sync.Mutex
	values map[string]bool
}

func New() *Cache {
	return &Cache{
		mu:     &sync.Mutex{},
		values: make(map[string]bool),
	}
}

func (c *Cache) Set(name string, status bool) {
	c.mu.Lock()
	c.values[name] = status
	c.mu.Unlock()
}

func (c *Cache) GetAll() map[string]bool {
	valueCopy := make(map[string]bool)
	c.mu.Lock()
	for k, v := range c.values {
		valueCopy[k] = v
	}
	c.mu.Unlock()
	return valueCopy
}
