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
		d := NewDecomposer(simple.NewUndirectedGraph(), test.k)
		t.Run(fmt.Sprintf("%d -> %d", test.k, test.t), func(t *testing.T) {
			actual := d.getThreshold()
			if actual != test.t {
				t.Errorf("Expected %v, got %v", test.t, actual)
			}
		})
	}
}

// 0 -- 1
// 2 -- 3 -- 4
// 5 -- 6
func TestPickComponent_NothingToPick(t *testing.T) {
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
	g.SetEdge(g.NewEdge(g.Node(5), g.Node(6)))
	d := NewDecomposer(g, 2)
	c := d.pickComponent()
	if c != nil {
		t.Errorf("expected nil, got %v", c)
	}
}

// 0 -- 1
// 2 -- 3 -- 4 -- 5
// 6 -- 7
func TestPickComponent_AboveThreshold(t *testing.T) {
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
	g.SetEdge(g.NewEdge(g.Node(2), g.Node(3)))
	g.SetEdge(g.NewEdge(g.Node(3), g.Node(4)))
	g.SetEdge(g.NewEdge(g.Node(4), g.Node(5)))
	g.SetEdge(g.NewEdge(g.Node(6), g.Node(7)))
	d := NewDecomposer(g, 2)
	c := d.pickComponent()
	assertContains(t, c, 2, 3, 4, 5)
}

func assertContains(t *testing.T, component []graph.Node, ids ...int64) {
	for _, id := range ids {
		for _, n := range component {
			if n.ID() == id {
				return
			}
		}
		t.Errorf("component %v does not contain node %v", component, id)
	}
}
