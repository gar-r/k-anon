package algorithm

import (
	"bitbucket.org/dargzero/k-anon/model"
	"bitbucket.org/dargzero/k-anon/testutil"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
	"testing"
)

func TestBuildAnonGraph_1(t *testing.T) {
	table := model.GetIntTable1()
	k := 2
	g, _ := BuildAnonGraph(table, k)
	verifyForestProperties(g, t, k)
}

func TestBuildAnonGraph_2(t *testing.T) {
	table := model.GetStudentTable()
	k := 3
	g, _ := BuildAnonGraph(table, k)
	verifyForestProperties(g, t, k)
}

// Component 1: 0 --> 1 <-- 2
// Component 2: 3 --> 4
// Weights: 1 -[4]-> 3; 1 -[2]->4; all others = 1
func TestPickTargetVertex(t *testing.T) {
	costGraph := CreateNodesWeightedUndirected(5)
	AddWeightedEdge(costGraph, 0, 1, 1)
	AddWeightedEdge(costGraph, 0, 2, 1)
	AddWeightedEdge(costGraph, 0, 3, 1)
	AddWeightedEdge(costGraph, 0, 4, 1)
	AddWeightedEdge(costGraph, 1, 2, 1)
	AddWeightedEdge(costGraph, 1, 3, 4)
	AddWeightedEdge(costGraph, 1, 4, 2)
	AddWeightedEdge(costGraph, 2, 3, 1)
	AddWeightedEdge(costGraph, 2, 4, 1)
	AddWeightedEdge(costGraph, 3, 4, 1)
	g := CreateNodesDirected(5)
	AddEdge(g, 0, 1)
	AddEdge(g, 2, 1)
	AddEdge(g, 3, 4)
	component := []graph.Node{g.Node(0), g.Node(1), g.Node(2)}
	v := pickTargetVertex(g, component, g.Node(1), costGraph)
	testutil.AssertEquals(int64(4), v.ID(), t)
}

// 0 --> 1 <-- 2
func TestPickSourceVertex(t *testing.T) {
	g := CreateNodesDirected(3)
	AddEdge(g, 0, 1)
	AddEdge(g, 2, 1)
	u := pickSourceVertex(g, UndirectedConnectedComponents(g)[0])
	testutil.AssertEquals(int64(1), u.ID(), t)
}

func TestPickComponentToExtend_1(t *testing.T) {
	components := [][]graph.Node{
		{simple.Node(0)},
		{simple.Node(1)},
	}
	c := pickComponentToExtend(components, 1)
	testutil.AssertNil(c, t)
}

func TestPickComponentToExtend_2(t *testing.T) {
	components := [][]graph.Node{
		{simple.Node(0)},
		{simple.Node(1)},
	}
	c := pickComponentToExtend(components, 2)
	testutil.AssertNotNil(c, t)
}

func TestPickComponentToExtend_3(t *testing.T) {
	components := [][]graph.Node{
		{simple.Node(0), simple.Node(1)},
	}
	c := pickComponentToExtend(components, 2)
	testutil.AssertNil(c, t)
}

func TestPickComponentToExtend_4(t *testing.T) {
	components := [][]graph.Node{
		{simple.Node(0), simple.Node(1), simple.Node(2)},
		{simple.Node(3), simple.Node(4)},
	}
	c := pickComponentToExtend(components, 3)
	testutil.AssertNotNil(c, t)
}

func TestPickComponentToExtend_5(t *testing.T) {
	components := [][]graph.Node{
		{simple.Node(0), simple.Node(1), simple.Node(2)},
		{simple.Node(3), simple.Node(4), simple.Node(5)},
	}
	c := pickComponentToExtend(components, 3)
	testutil.AssertNil(c, t)
}

func verifyForestProperties(g graph.Directed, t *testing.T, k int) {
	components := topo.ConnectedComponents(graph.Undirect{g})
	for _, c := range components {
		if len(c) < k {
			t.Errorf("component size should not be smaller than %d", k)
		}
		for _, n := range c {
			outgoing := g.From(n.ID())
			if outgoing.Len() > 1 {
				t.Errorf("outgoing edge count should not be greater than 1")
			}
		}
	}
}
