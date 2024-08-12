package master_data_repoimpl

import (
	"sync"
	"time"

	"github.com/averak/hbaas/app/domain/model"
)

type cacheEntry struct {
	revision  int
	data      model.MasterData
	lock      *sync.RWMutex
	createdAt time.Time
}

type cacheHolder struct {
	max   int
	cache map[int]*cacheEntry
	mu    sync.RWMutex
}

func newCacheHolder(max int) *cacheHolder {
	return &cacheHolder{
		max:   max,
		cache: map[int]*cacheEntry{},
	}
}

func (c *cacheHolder) get(revision int) (model.MasterData, bool) {
	c.mu.RLock()
	entry, ok := c.cache[revision]
	c.mu.RUnlock()

	if !ok {
		c.mu.Lock()
		if entry, ok = c.cache[revision]; !ok {
			entry = &cacheEntry{
				revision:  revision,
				lock:      &sync.RWMutex{},
				createdAt: time.Now(),
			}
			c.cache[revision] = entry
		}
		c.mu.Unlock()
	}

	entry.lock.RLock()
	defer entry.lock.RUnlock()

	return entry.data, ok
}

func (c *cacheHolder) set(revision int, data model.MasterData) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.cache[revision]
	if !ok {
		entry = &cacheEntry{
			revision:  revision,
			lock:      &sync.RWMutex{},
			createdAt: time.Now(),
		}
		c.cache[revision] = entry
	} else {
		entry.createdAt = time.Now() // 更新時にcreatedAtをリフレッシュ
	}

	// キャッシュが最大サイズに達している場合、最も古いエントリを削除
	if len(c.cache) > c.max {
		oldestRevision := c.findOldestRevision()
		delete(c.cache, oldestRevision)
	}

	entry.data = data
}

func (c *cacheHolder) findOldestRevision() int {
	var oldestRevision int
	var oldestTime time.Time

	for revision, entry := range c.cache {
		if oldestTime.IsZero() || entry.createdAt.Before(oldestTime) {
			oldestTime = entry.createdAt
			oldestRevision = revision
		}
	}
	return oldestRevision
}
