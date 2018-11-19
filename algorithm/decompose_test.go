package algorithm

import (
	"fmt"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"testing"
)

func TestGetThreshold(t *testing.T) {
	tests := []struct {
		k int
		t int
	}{
		{1, 1},
		{2, 3},
		{3, 5},
		{4, 7},
		{5, 10},
		{6, 13},
		{7, 16},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%d -> %d", test.k, test.t), func(t *testing.T) {
			actual := getThreshold(test.k)
			if actual != test.t {
				t.Errorf("Expected %v, got %v", test.t, actual)
			}
		})
	}
}

// 0 -- 1
// 2 -- 3 -- 4 -- 5
// 6
func TestPickComponentToSplit(t *testing.T) {
	g := simple.NewUndirectedGraph()
	g.AddNode(g.NewNode())
	g.AddNode(g.NewNode())
	g.AddNode(g.NewNode())
	g.AddNode(g.NewNode())
	g.AddNode(g.NewNode())
	g.AddNode(g.NewNode())
	g.AddNode(g.NewNode())
	g.SetEdge(g.NewEdge(g.Node(0), g.Node(1)))
	g.SetEdge(g.NewEdge(g.Node(2), g.Node(3)))
	g.SetEdge(g.NewEdge(g.Node(3), g.Node(4)))
	g.SetEdge(g.NewEdge(g.Node(4), g.Node(5)))
	c := pickComponentToSplit(g, 2)
	if len(c) != 4 {
		t.Errorf("component length mismatch")
	}
	assertComponentHasVertex(t, c, 2)
	assertComponentHasVertex(t, c, 3)
	assertComponentHasVertex(t, c, 4)
	assertComponentHasVertex(t, c, 5)
}

// 0 -- 1 -- 2
// 3 -- 4 -- 5
// 6 -- 7
func TestPickComponentToSplit_Finished(t *testing.T) {
	g := simple.NewUndirectedGraph()
	g.AddNode(g.NewNode())
	g.AddNode(g.NewNode())
	g.AddNode(g.NewNode())
	g.AddNode(g.NewNode())
	g.AddNode(g.NewNode())
	g.AddNode(g.NewNode())
	g.AddNode(g.NewNode())
	g.AddNode(g.NewNode())
	g.SetEdge(g.NewEdge(g.Node(0), g.Node(1)))
	g.SetEdge(g.NewEdge(g.Node(1), g.Node(2)))
	g.SetEdge(g.NewEdge(g.Node(3), g.Node(4)))
	g.SetEdge(g.NewEdge(g.Node(4), g.Node(5)))
	g.SetEdge(g.NewEdge(g.Node(6), g.Node(7)))
	c := pickComponentToSplit(g, 2)
	if c != nil {
		t.Errorf("expected nil, got %v", c)
	}
}

func assertComponentHasVertex(t *testing.T, component []graph.Node, vertex int64) {
	for _, n := range component {
		if n.ID() == vertex {
			return
		}
	}
	t.Errorf("component does not contain vertex(%v)", vertex)
}
