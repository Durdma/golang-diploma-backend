package cache

// Cache - Интерфейс для работы с кэшем
type Cache interface {
	Set(key, value interface{}) error
	Get(key interface{}) (interface{}, error)
}
