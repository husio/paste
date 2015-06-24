package paste

import "sync"

type memoryCache struct {
	sync.Mutex
	mem map[string]interface{}
}

func newMemoryCache() *memoryCache {
	return &memoryCache{
		mem: make(map[string]interface{}),
	}
}

func (c *memoryCache) Put(key string, value interface{}) {
	c.Lock()
	c.mem[key] = value
	c.Unlock()
}

func (c *memoryCache) Has(key string) bool {
	c.Lock()
	_, ok := c.mem[key]
	c.Unlock()
	return ok
}

func (c *memoryCache) Get(key string) (interface{}, bool) {
	c.Lock()
	value, ok := c.mem[key]
	c.Unlock()
	return value, ok
}

func (c *memoryCache) GetString(key string) (string, bool) {
	val, ok := c.Get(key)
	if !ok {
		return "", false
	}
	s, ok := val.(string)
	return s, ok
}

func (c *memoryCache) Delete(key string) {
	c.Lock()
	delete(c.mem, key)
	c.Unlock()
}

func (c *memoryCache) Reset() {
	c.Lock()
	c.mem = make(map[string]interface{})
	c.Unlock()
}
