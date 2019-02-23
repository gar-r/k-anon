package algorithm

import (
	"bitbucket.org/dargzero/k-anon/testutil"
	"fmt"
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
			testutil.AssertEquals(actual, test.t, t)
		})
	}
}

// 0 -- 1
// 2 -- 3 -- 4
// 5 -- 6
func TestPickComponent_NothingToPick(t *testing.T) {
	g := CreateNodesUndirected(7)
	AddEdge(g, 0, 1)
	AddEdge(g, 2, 3)
	AddEdge(g, 3, 4)
	AddEdge(g, 5, 6)
	d := NewDecomposer(g, 2)
	c := d.pickComponent(d.getThreshold())
	testutil.AssertNil(c, t)
}

// 0 -- 1
// 2 -- 3 -- 4 -- 5
// 6 -- 7
func TestPickComponent_AboveThreshold(t *testing.T) {
	g := CreateNodesUndirected(8)
	AddEdge(g, 0, 1)
	AddEdge(g, 2, 3)
	AddEdge(g, 3, 4)
	AddEdge(g, 4, 5)
	AddEdge(g, 6, 7)
	d := NewDecomposer(g, 2)
	c := d.pickComponent(d.getThreshold())
	testutil.AssertContains(t, c, 2, 3, 4, 5)
}

// 0 -- 1 -- s -- 2
func TestPickComponent_SteinersVertexSkipped(t *testing.T) {
	g := CreateNodesUndirected(3)
	AddEdge(g, 0, 1)
	d := NewDecomposer(g, 2)
	g.AddNode(g.NewNode()) // this will be a Steiner's vertex
	AddEdge(g, 3, 1)
	AddEdge(g, 3, 2)
	c := d.pickComponent(d.getThreshold())
	testutil.AssertNil(c, t)
}

func TestDecomposer_Decompose_TerminatesWhenFinished(t *testing.T) {
	g := CreateNodesUndirected(2)
	d := NewDecomposer(g, 2)
	d.Decompose()
}

func TestDecomposer_Decompose_ComponentSizes(t *testing.T) {
	g := GetUndirectedTestGraph1()
	k := 3
	l := g.Nodes().Len()
	d := NewDecomposer(g, k)
	d.Decompose()
	components := topo.ConnectedComponents(g)
	for _, c := range components {
		compLen := 0
		for _, n := range c {
			if n.ID() < int64(l) { // original node
				compLen++
			}
		}
		if compLen < k {
			t.Errorf("component size < k")
		}
	}
}

//   -- 0 --
//  |       |
//  1       3
//  |
//  2
//
//  u := 0, v := 1, k := 2, s > 2k-1
//  t >= k && s-t >= k
func TestPartition_TypeACut(t *testing.T) {
	g := CreateNodesUndirected(4)
	AddEdge(g, 0, 1)
	AddEdge(g, 0, 3)
	AddEdge(g, 1, 2)
	d := NewDecomposer(g, 2)
	u := g.Node(0)
	v := g.Node(1)
	d.performCutTypeA(u, v)
	if g.HasEdgeBetween(0, 1) {
		t.Errorf("no edge should be present between nodes 0 and 1")
	}
}

//            ------- 0 ------
//           |                |
//      ---- 1 ----           6
//     |     |     |          |
//     2     3     4          7
//           |
//           5
//
//  u := 0, v := 1, k := 4, s > 2k-1
//  s-t == k-1
func TestPartition_TypeBCut(t *testing.T) {
	g := GetUndirectedTestGraph2()
	d := NewDecomposer(g, 4)
	u := g.Node(0)
	v := g.Node(1)
	d.performCutTypeB(u, v)
	testutil.AssertVertexReplaced(t, g, 1, 8, 2, 3, 4)
}

//      ---- 1 ----------
//     |     |     |     |
//     2     3     4     0
//           |           |
//           5           6
//                       |
//						 7
//
// Same as graph2, just drawn with a different root
//
//  u := 1, v := 0, k := 4, s > 2k-1
//  t == k-1
func TestPartition_TypeCCut(t *testing.T) {
	g := GetUndirectedTestGraph2()
	d := NewDecomposer(g, 4)
	u := g.Node(1)
	v := g.Node(0)
	d.performCutTypeC(u, v)
	testutil.AssertVertexReplaced(t, g, 1, 8, 2, 3, 4)
}

//         ------ 0 -------
//        |    |      |    |
//        1    3      5    7
//        |    |      |    |
//        2    4      6    8
//
//  u := 1, v := 0, k := 4, s > 2k-1
//  (t < k || s-t < k), s-t != k-1, t != k-1
func TestPartition_TypeDCut(t *testing.T) {
	g := GetUndirectedTestGraph3()
	d := NewDecomposer(g, 4)
	u := g.Node(0)
	v := g.Node(1)
	component := topo.ConnectedComponents(g)[0]
	d.performCutTypeD(u, v, component)
	newComps := topo.ConnectedComponents(g)
	if len(newComps[0]) != 5 || len(newComps[1]) != 5 {
		t.Errorf("invalid component length after split")
	}
	if !(containsNode(newComps[0], g.Node(9)) ||
		containsNode(newComps[1], g.Node(9))) {
		t.Errorf("missing Steiner's vertex")
	}
	if 2 != g.From(9).Len() {
		t.Errorf("invalid connection count for new vertex")
	}
}
