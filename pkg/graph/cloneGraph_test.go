package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCloneGraph(t *testing.T) {
	graph1 := &Node{
		Val: 1,
	}
	graph2 := &Node{
		Val: 2,
	}
	graph3 := &Node{
		Val: 3,
	}
	graph4 := &Node{
		Val: 4,
	}
	graph1.Neighbors = []*Node{
		graph2,
		graph4,
	}
	graph2.Neighbors = []*Node{
		graph1,
		graph3,
	}
	graph3.Neighbors = []*Node{
		graph2,
		graph4,
	}
	graph4.Neighbors = []*Node{
		graph1,
		graph3,
	}
	graphs := []*Node{
		graph1,
		nil,
	}
	for _, r := range graphs {
		rr := CloneGraph(r)
		assert.Equal(t, rr, r)
	}
}
