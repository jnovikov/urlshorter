package keystorage

type CacheStorage struct {
	dataStorage KeyStorage
	cache       KeyStorage
}

func (c *CacheStorage) Get(key string) interface{} {
	if cacheValue := c.cache.Get(key); cacheValue != nil {
		return cacheValue
	}
	result := c.dataStorage.Get(key)
	if result != nil {
		c.cache.Set(key, result)
	}
	return result
}

func (c *CacheStorage) Set(key string, value interface{}) {
	c.dataStorage.Set(key, value)
}

func NewCacheStorage(data, cache KeyStorage) *CacheStorage {
	return &CacheStorage{data, cache}
}
