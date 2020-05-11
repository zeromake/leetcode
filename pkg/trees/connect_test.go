package trees

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnect(t *testing.T) {
	trees := []*Node{
		{
			Val: 1,
			Left: &Node{
				Val: 2,
				Left: &Node{
					Val: 4,
				},
				Right: &Node{
					Val: 5,
				},
			},
			Right: &Node{
				Val: 3,
				Left: &Node{
					Val: 6,
				},
				Right: &Node{
					Val: 7,
				},
			},
		},
		{
			Val: 1,
			Left: &Node{
				Val: 2,
				Right: &Node{
					Val: 5,
				},
			},
			Right: &Node{
				Val: 3,
				Left: &Node{
					Val: 6,
				},
				Right: &Node{
					Val: 7,
				},
			},
		},
		nil,
	}
	result := []string{
		"[1,#,2,3,#,4,5,6,7,#]",
		"[1,#,2,3,#,5,6,7,#]",
		"[]",
	}
	for i, r := range trees {
		rr := Connect(r)
		assert.Equal(t, rr.String(), result[i])
	}
}
