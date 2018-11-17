package algorithm

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
	"math"
	"testing"
)

func TestBuildAnonGraph(t *testing.T) {
	table := getTestTable(getExampleGeneralizer())
	g := BuildAnonGraph(table, 2)
	components := topo.ConnectedComponents(graph.Undirect{g})
	for _, c := range components {
		if len(c) < 2 {
			t.Errorf("component size smaller than 2")
		}
		for _, n := range c {
			outgoing := g.From(n.ID())
			if outgoing.Len() > 1 {
				t.Errorf("outgoing edge count greater than 1")
			}
		}
	}
}

// Component 1: 0 --> 1 <-- 2
// Component 2: 3 --> 4
// Weights: 1 -[4]-> 3; 1 -[2]->4; all others = 1
func TestPickTargetVertex(t *testing.T) {
	costGraph := simple.NewWeightedUndirectedGraph(0, math.MaxFloat64)
	costGraph.AddNode(costGraph.NewNode())
	costGraph.AddNode(costGraph.NewNode())
	costGraph.AddNode(costGraph.NewNode())
	costGraph.AddNode(costGraph.NewNode())
	costGraph.AddNode(costGraph.NewNode())
	costGraph.SetWeightedEdge(costGraph.NewWeightedEdge(costGraph.Node(0), costGraph.Node(1), 1))
	costGraph.SetWeightedEdge(costGraph.NewWeightedEdge(costGraph.Node(0), costGraph.Node(2), 1))
	costGraph.SetWeightedEdge(costGraph.NewWeightedEdge(costGraph.Node(0), costGraph.Node(3), 1))
	costGraph.SetWeightedEdge(costGraph.NewWeightedEdge(costGraph.Node(0), costGraph.Node(4), 1))
	costGraph.SetWeightedEdge(costGraph.NewWeightedEdge(costGraph.Node(1), costGraph.Node(2), 1))
	costGraph.SetWeightedEdge(costGraph.NewWeightedEdge(costGraph.Node(1), costGraph.Node(3), 4))
	costGraph.SetWeightedEdge(costGraph.NewWeightedEdge(costGraph.Node(1), costGraph.Node(4), 2))
	costGraph.SetWeightedEdge(costGraph.NewWeightedEdge(costGraph.Node(2), costGraph.Node(3), 1))
	costGraph.SetWeightedEdge(costGraph.NewWeightedEdge(costGraph.Node(2), costGraph.Node(4), 1))
	costGraph.SetWeightedEdge(costGraph.NewWeightedEdge(costGraph.Node(3), costGraph.Node(4), 1))
	g := simple.NewDirectedGraph()
	g.AddNode(g.NewNode())
	g.AddNode(g.NewNode())
	g.AddNode(g.NewNode())
	g.AddNode(g.NewNode())
	g.AddNode(g.NewNode())
	g.SetEdge(g.NewEdge(g.Node(0), g.Node(1)))
	g.SetEdge(g.NewEdge(g.Node(2), g.Node(1)))
	g.SetEdge(g.NewEdge(g.Node(3), g.Node(4)))
	component := []graph.Node{g.Node(0), g.Node(1), g.Node(2)}
	v := pickTargetVertex(g, component, g.Node(1), costGraph)
	if v.ID() != 4 {
		t.Errorf("expected target vertex with ID=4")
	}
}

// 0 --> 1 <-- 2
func TestPickSourceVertex(t *testing.T) {
	g := simple.NewDirectedGraph()
	g.AddNode(g.NewNode())
	g.AddNode(g.NewNode())
	g.AddNode(g.NewNode())
	g.SetEdge(g.NewEdge(g.Node(0), g.Node(1)))
	g.SetEdge(g.NewEdge(g.Node(2), g.Node(1)))
	u := pickSourceVertex(g, getComponents(g)[0])
	if u.ID() != 1 {
		t.Errorf("expected source vertex with ID=1")
	}
}

func TestPickComponent_1(t *testing.T) {
	components := [][]graph.Node{
		{simple.Node(0)},
		{simple.Node(1)},
	}
	if pickComponent(components, 1) != nil {
		t.Error()
	}
}

func TestPickComponent_2(t *testing.T) {
	components := [][]graph.Node{
		{simple.Node(0)},
		{simple.Node(1)},
	}
	if pickComponent(components, 2) == nil {
		t.Error()
	}
}

func TestPickComponent_3(t *testing.T) {
	components := [][]graph.Node{
		{simple.Node(0), simple.Node(1)},
	}
	if pickComponent(components, 2) != nil {
		t.Error()
	}
}

func TestPickComponent_4(t *testing.T) {
	components := [][]graph.Node{
		{simple.Node(0), simple.Node(1), simple.Node(2)},
		{simple.Node(3), simple.Node(4)},
	}
	if pickComponent(components, 3) == nil {
		t.Error()
	}
}

func TestPickComponent_5(t *testing.T) {
	components := [][]graph.Node{
		{simple.Node(0), simple.Node(1), simple.Node(2)},
		{simple.Node(3), simple.Node(4), simple.Node(5)},
	}
	if pickComponent(components, 3) != nil {
		t.Error()
	}
}
