package pokecache

import (
	"time"
	"sync"
)


type Cache struct {
	count map[string]CacheEntry
	mu *sync.Mutex
	interval time.Duration
} 

type CacheEntry struct {
	createdAt time.Time
	val []byte
}

func (cache *Cache) Add(key string, value []byte){
	cache.mu.Lock()
	defer cache.mu.Unlock()
	var cacheEn CacheEntry
	cacheEn.createdAt = time.Now()
	cacheEn.val = value
	cache.count[key] = cacheEn 
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	value, exist := cache.count[key]
	if !exist {
		return nil, exist
	}
	return value.val, exist
}

func (cache *Cache) reapLoop(){
	c := time.Tick(cache.interval)
	for range c{
		cache.mu.Lock()
		
		for key, value := range cache.count{ 
			elapsed := time.Since(value.createdAt)
			if elapsed >= cache.interval{
				delete(cache.count, key)
			}
		}
		cache.mu.Unlock()
		

	}
	
}



func NewCache(interval time.Duration) *Cache{
	var cache Cache
	m := make(map[string]CacheEntry)
	cache.mu = &sync.Mutex{}
	cache.count = m
	cache.interval = interval
	go cache.reapLoop()
	return &cache
	
}

