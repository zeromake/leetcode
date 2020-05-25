package links

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLRU(t *testing.T) {
	cache := LRUConstructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	assert.Equal(t, cache.Get(1), 1)  // 返回  1
	cache.Put(3, 3)                   // 该操作会使得密钥 2 作废
	assert.Equal(t, cache.Get(2), -1) // 返回 -1 (未找到)
	cache.Put(4, 4)                   // 该操作会使得密钥 1 作废
	assert.Equal(t, cache.Get(1), -1) // 返回 -1 (未找到)
	assert.Equal(t, cache.Get(3), 3)  // 返回  3
	assert.Equal(t, cache.Get(4), 4)  // 返回  4
	cache.Put(4, 8)
	assert.Equal(t, cache.Get(4), 8)
}
