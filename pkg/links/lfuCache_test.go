package links

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLFUCache(t *testing.T) {
	lfu := Constructor(2)
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	assert.Equal(t, lfu.Get(1), 1)
	lfu.Put(3, 3)
	assert.Equal(t, lfu.Get(2), -1)
	assert.Equal(t, lfu.Get(3), 3)
	lfu.Put(4, 4)
	assert.Equal(t, lfu.Get(1), -1)
	assert.Equal(t, lfu.Get(3), 3)
	assert.Equal(t, lfu.Get(4), 4)
	lfu.Put(4, 5)

	lfu = Constructor(0)
	lfu.Put(1, 1)
	assert.Equal(t, lfu.Get(1), -1)
}
