package algorithm

import (
	"fmt"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
	"k-anon/testutil"
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
	g := testutil.CreateNodes(7)
	testutil.AddEdge(g, 0, 1)
	testutil.AddEdge(g, 2, 3)
	testutil.AddEdge(g, 3, 4)
	testutil.AddEdge(g, 5, 6)
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
	g := testutil.CreateNodes(8)
	testutil.AddEdge(g, 0, 1)
	testutil.AddEdge(g, 2, 3)
	testutil.AddEdge(g, 3, 4)
	testutil.AddEdge(g, 4, 5)
	testutil.AddEdge(g, 6, 7)
	d := NewDecomposer(g, 2)
	c := d.pickComponent(d.getThreshold())
	testutil.AssertContains(t, c, 2, 3, 4, 5)
}

// 0 -- 1 -- s -- 2
func TestPickComponent_SteinersVertexSkipped(t *testing.T) {
	g := testutil.CreateNodes(3)
	testutil.AddEdge(g, 0, 1)
	d := NewDecomposer(g, 2)
	g.AddNode(g.NewNode()) // this will be a Steiner's vertex
	testutil.AddEdge(g, 3, 1)
	testutil.AddEdge(g, 3, 2)
	c := d.pickComponent(d.getThreshold())
	if c != nil {
		t.Errorf("expected nil, got %v", c)
	}
}

func TestDecomposer_Decompose_TerminatesWhenFinished(t *testing.T) {
	g := testutil.CreateNodes(2)
	d := NewDecomposer(g, 2)
	d.Decompose()
}

func TestDecomposer_Decompose_ComponentSizes(t *testing.T) {
	g := testutil.GetUndirectedTestGraph1()
	d := NewDecomposer(g, 3)
	d.Decompose()
	components := topo.ConnectedComponents(g)
	for _, c := range components {
		if len(c) > (2*3 - 1) {
			t.Errorf("invalid component size: %v", c)
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
	g := testutil.CreateNodes(4)
	testutil.AddEdge(g, 0, 1)
	testutil.AddEdge(g, 0, 3)
	testutil.AddEdge(g, 1, 2)
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
	g := testutil.GetUndirectedTestGraph2()
	d := NewDecomposer(g, 4)
	u := g.Node(0)
	v := g.Node(1)
	d.performCutTypeB(u, v)
	assertVertexReplaced(t, g, 1, 8, 2, 3, 4)
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
	g := testutil.GetUndirectedTestGraph2()
	d := NewDecomposer(g, 4)
	u := g.Node(1)
	v := g.Node(0)
	d.performCutTypeC(u, v)
	assertVertexReplaced(t, g, 1, 8, 2, 3, 4)
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
	g := testutil.GetUndirectedTestGraph3()
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

func assertVertexReplaced(t *testing.T, g graph.Undirected, original, new int64, connections ...int64) {
	for _, conn := range connections {
		if g.HasEdgeBetween(original, conn) {
			t.Errorf("unexpected edge between %v and %v", original, conn)
		}
		if !g.HasEdgeBetween(new, conn) {
			t.Errorf("expected edge between %v and %v", new, conn)
		}
	}
}
