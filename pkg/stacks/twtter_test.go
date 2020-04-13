package stacks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwitter(t *testing.T) {
	twitter := TwitterConstructor()
	twitter.PostTweet(1, 5)
	assert.Equal(t, twitter.GetNewsFeed(1), []int{5})
	twitter.Follow(1, 1)
	twitter.Follow(1, 2)
	assert.Equal(t, twitter.GetNewsFeed(1), []int{5})
	twitter.PostTweet(2, 6)
	assert.Equal(t, twitter.GetNewsFeed(1), []int{6, 5})
	twitter.Unfollow(1, 2)
	assert.Equal(t, twitter.GetNewsFeed(1), []int{5})
}
