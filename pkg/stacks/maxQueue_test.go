package stacks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxQueue(t *testing.T) {
	queue := Constructor()
	assert.Equal(t, queue.Max_value(), -1)
	assert.Equal(t, queue.Pop_front(), -1)
	assert.Equal(t, queue.Max_value(), -1)
	queue.Push_back(46)
	assert.Equal(t, queue.Max_value(), 46)
	assert.Equal(t, queue.Pop_front(), 46)
	assert.Equal(t, queue.Max_value(), -1)
	assert.Equal(t, queue.Pop_front(), -1)
	queue.Push_back(868)
	assert.Equal(t, queue.Pop_front(), 868)
	assert.Equal(t, queue.Pop_front(), -1)
	assert.Equal(t, queue.Pop_front(), -1)
	queue.Push_back(525)
	assert.Equal(t, queue.Pop_front(), 525)
	assert.Equal(t, queue.Max_value(), -1)
	queue.Push_back(666)
	queue.Push_back(868)
	assert.Equal(t, queue.Max_value(), 868)
	assert.Equal(t, queue.Pop_front(), 666)
	assert.Equal(t, queue.Max_value(), 868)
	assert.Equal(t, queue.Pop_front(), 868)
	assert.Equal(t, queue.Max_value(), -1)

	queue.Push_back(868)
	queue.Push_back(666)
	queue.Push_back(866)
	assert.Equal(t, queue.Max_value(), 868)
	assert.Equal(t, queue.Pop_front(), 868)
	assert.Equal(t, queue.Max_value(), 866)
	assert.Equal(t, queue.Pop_front(), 666)
	assert.Equal(t, queue.Max_value(), 866)
	assert.Equal(t, queue.Pop_front(), 866)
	assert.Equal(t, queue.Max_value(), -1)
}
