package algorithm

import (
	"fmt"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
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
	g := createTestGraph(7)
	addEdge(g, 0, 1)
	addEdge(g, 2, 3)
	addEdge(g, 3, 4)
	addEdge(g, 5, 6)
	d := NewDecomposer(g, 2)
	c := d.pickComponent(d.getThreshold())
	if c != nil {
		t.Errorf("expected nil, got %v", c)
	}
}

// 0 -- 1
// 2 -- 3 -- 4 -- 5
// 6 -- 7
func TestPickComponent_AboveThreshold(t *testing.T) {
	g := createTestGraph(8)
	addEdge(g, 0, 1)
	addEdge(g, 2, 3)
	addEdge(g, 3, 4)
	addEdge(g, 4, 5)
	addEdge(g, 6, 7)
	d := NewDecomposer(g, 2)
	c := d.pickComponent(d.getThreshold())
	assertContains(t, c, 2, 3, 4, 5)
}

// 0 -- 1 -- s -- 2
func TestPickComponent_SteinersVertexSkipped(t *testing.T) {
	g := createTestGraph(3)
	addEdge(g, 0, 1)
	d := NewDecomposer(g, 2)
	g.AddNode(g.NewNode()) // this will be a Steiner's vertex
	addEdge(g, 3, 1)
	addEdge(g, 3, 2)
	c := d.pickComponent(d.getThreshold())
	if c != nil {
		t.Errorf("expected nil, got %v", c)
	}
}

func TestDecomposer_Decompose_TerminatesWhenFinished(t *testing.T) {
	g := createTestGraph(2)
	d := NewDecomposer(g, 2)
	d.Decompose()
}

//                   ---------- 0 ---------------------------------
//                   |                      |          |           |
//         --------- 1 ---------          - 5 -      - 6 -       - 7 -
//         |         |          |        |     |     |     |    |     |
//       - 2 -     - 3 -      - 4 -     14    15    16     17  18     19
//      |     |   |     |    |     |
//      8     9  10     11  12     13
func TestDecomposer_Decompose(t *testing.T) {
	g := createTestGraph(20)
	addEdge(g, 0, 1)
	addEdge(g, 1, 2)
	addEdge(g, 1, 3)
	addEdge(g, 1, 4)
	addEdge(g, 2, 8)
	addEdge(g, 2, 9)
	addEdge(g, 3, 10)
	addEdge(g, 3, 11)
	addEdge(g, 4, 12)
	addEdge(g, 4, 13)
	addEdge(g, 0, 5)
	addEdge(g, 0, 6)
	addEdge(g, 0, 7)
	addEdge(g, 5, 14)
	addEdge(g, 5, 15)
	addEdge(g, 6, 16)
	addEdge(g, 6, 17)
	addEdge(g, 7, 18)
	addEdge(g, 7, 19)
	d := NewDecomposer(g, 3)
	d.Decompose()
	components := topo.ConnectedComponents(g)
	t.Errorf("%v", components)
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

func createTestGraph(nodeCount int) *simple.UndirectedGraph {
	g := simple.NewUndirectedGraph()
	for i := 0; i < nodeCount; i++ {
		g.AddNode(g.NewNode())
	}
	return g
}

func addEdge(g *simple.UndirectedGraph, u, v int64) {
	g.SetEdge(g.NewEdge(g.Node(u), g.Node(v)))
}
