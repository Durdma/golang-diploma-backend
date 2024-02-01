package cache

import (
	"errors"
	"sync"
)

// MemoryCache - Структура для кэша
type MemoryCache struct {
	cache map[interface{}]interface{}
	sync.RWMutex
}

// NewMemoryCache - Создание нового кэша
func NewMemoryCache() *MemoryCache {
	return &MemoryCache{
		cache: make(map[interface{}]interface{}),
	}
}

// Set - Добавление новой записи в кэш
func (c *MemoryCache) Set(key, value interface{}) error {
	c.Lock()
	c.cache[key] = value
	c.Unlock()

	return nil
}

// Get - Получение записи из кэша по ключу
func (c *MemoryCache) Get(key interface{}) (interface{}, error) {
	c.RLock()
	value, ex := c.cache[key]
	c.RUnlock()

	if !ex {
		return nil, errors.New("not found")
	}

	return value, nil
}
