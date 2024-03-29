package cache

import (
	"errors"
	"sync"
	"time"
)

type item struct {
	value     interface{}
	createdAt int64
}

// MemoryCache - Структура для кэша
type MemoryCache struct {
	cache map[interface{}]*item
	sync.RWMutex
}

// NewMemoryCache - Создание нового кэша
func NewMemoryCache(ttl int64) *MemoryCache {
	c := &MemoryCache{
		cache: make(map[interface{}]*item),
	}

	go c.setTtlTimer(ttl)

	return c
}

func (c *MemoryCache) setTtlTimer(ttl int64) {
	for now := range time.Tick(time.Second) {
		c.Lock()
		for k, v := range c.cache {
			if now.Unix()-v.createdAt > ttl {
				delete(c.cache, k)
			}
		}

		c.Unlock()
	}
}

// Set - Добавление новой записи в кэш
func (c *MemoryCache) Set(key, value interface{}) error {
	c.Lock()
	c.cache[key] = &item{
		value:     value,
		createdAt: time.Now().Unix(),
	}
	c.Unlock()

	return nil
}

// Get - Получение записи из кэша по ключу
func (c *MemoryCache) Get(key interface{}) (interface{}, error) {
	c.RLock()
	item, ex := c.cache[key]
	c.RUnlock()

	if !ex {
		return nil, errors.New("not found")
	}

	return item.value, nil
}
