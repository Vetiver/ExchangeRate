package cbr

import "sync"

type Cache struct {
	mu      sync.RWMutex
	records map[string]map[string][]Record
}

func NewCache() *Cache {
	return &Cache{
		records: make(map[string]map[string][]Record),
	}
}

func (c *Cache) Get(date string, val string) ([]Record, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if records, ok := c.records[date]; ok {
		if recordsForVal, exists := records[val]; exists {
			return recordsForVal, true
		}
	}
	return nil, false
}

func (c *Cache) Set(date string, val string, records []Record) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.records[date]; !ok {
		c.records[date] = make(map[string][]Record)
	}
	c.records[date][val] = records
}