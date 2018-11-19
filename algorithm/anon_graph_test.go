package algorithm

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
	"k-anon/generalization"
	"k-anon/model"
	"math"
	"testing"
)

// Table data:
// 		(each column  -> int generalizer)
// 1 1 1 1
// 1 1 1 2
// 4 5 1 1
// 1 3 5 7
func TestBuildAnonGraph_1(t *testing.T) {
	table := getTestTable(getExampleGeneralizer())
	k := 2
	g := BuildAnonGraph(table, k)
	verifyForestProperties(g, t, k)
}

// Table data:
// 		(col1 -> suppress)
// 		(col2 -> int)
// 		(col3 -> int)
// 		(col4 -> int)
// 		(col5 -> grade)
// Male		25 	0 	35 	A
// Female	25	0	45	A+
// Male		30	2	30	B
// Female	30	1	35	B+
// Male		28	1	40	A-
// Female	28	1	15	B
// Male		27	0	15	B-
// Female	27	2	30	B
func TestBuildAnonGraph_2(t *testing.T) {
	dim1 := &generalization.Suppressor{}
	dim2 := generalization.NewHierarchyGeneralizer((&generalization.IntegerHierarchyBuilder{Items: []int{25, 27, 28, 30}}).NewIntegerHierarchy())
	dim3 := generalization.NewHierarchyGeneralizer((&generalization.IntegerHierarchyBuilder{Items: []int{0, 1, 2}}).NewIntegerHierarchy())
	dim4 := generalization.NewHierarchyGeneralizer((&generalization.IntegerHierarchyBuilder{Items: []int{10, 15, 30, 35, 40, 45}}).NewIntegerHierarchy())
	dim5 := getExampleGeneralizer2()
	table := &model.Table{
		Rows: []*model.Vector{
			{Items: []*model.Data{model.NewData("Male", dim1), model.NewData(25, dim2), model.NewData(0, dim3), model.NewData(35, dim4), model.NewData("A", dim5)}},
			{Items: []*model.Data{model.NewData("Female", dim1), model.NewData(25, dim2), model.NewData(0, dim3), model.NewData(45, dim4), model.NewData("A+", dim5)}},
			{Items: []*model.Data{model.NewData("Male", dim1), model.NewData(30, dim2), model.NewData(2, dim3), model.NewData(30, dim4), model.NewData("B", dim5)}},
			{Items: []*model.Data{model.NewData("Female", dim1), model.NewData(30, dim2), model.NewData(1, dim3), model.NewData(35, dim4), model.NewData("B+", dim5)}},
			{Items: []*model.Data{model.NewData("Male", dim1), model.NewData(28, dim2), model.NewData(1, dim3), model.NewData(40, dim4), model.NewData("A-", dim5)}},
			{Items: []*model.Data{model.NewData("Female", dim1), model.NewData(28, dim2), model.NewData(1, dim3), model.NewData(15, dim4), model.NewData("B", dim5)}},
			{Items: []*model.Data{model.NewData("Male", dim1), model.NewData(27, dim2), model.NewData(0, dim3), model.NewData(15, dim4), model.NewData("B-", dim5)}},
			{Items: []*model.Data{model.NewData("Female", dim1), model.NewData(27, dim2), model.NewData(2, dim3), model.NewData(30, dim4), model.NewData("B", dim5)}},
		},
	}
	k := 3
	g := BuildAnonGraph(table, k)
	verifyForestProperties(g, t, k)
}

func verifyForestProperties(g graph.Directed, t *testing.T, k int) {
	components := topo.ConnectedComponents(graph.Undirect{g})
	for _, c := range components {
		if len(c) < k {
			t.Errorf("component size smaller than %d", k)
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

func TestPickComponentToExtend_1(t *testing.T) {
	components := [][]graph.Node{
		{simple.Node(0)},
		{simple.Node(1)},
	}
	if pickComponentToExtend(components, 1) != nil {
		t.Error()
	}
}

func TestPickComponentToExtend_2(t *testing.T) {
	components := [][]graph.Node{
		{simple.Node(0)},
		{simple.Node(1)},
	}
	if pickComponentToExtend(components, 2) == nil {
		t.Error()
	}
}

func TestPickComponentToExtend_3(t *testing.T) {
	components := [][]graph.Node{
		{simple.Node(0), simple.Node(1)},
	}
	if pickComponentToExtend(components, 2) != nil {
		t.Error()
	}
}

func TestPickComponentToExtend_4(t *testing.T) {
	components := [][]graph.Node{
		{simple.Node(0), simple.Node(1), simple.Node(2)},
		{simple.Node(3), simple.Node(4)},
	}
	if pickComponentToExtend(components, 3) == nil {
		t.Error()
	}
}

func TestPickComponentToExtend_5(t *testing.T) {
	components := [][]graph.Node{
		{simple.Node(0), simple.Node(1), simple.Node(2)},
		{simple.Node(3), simple.Node(4), simple.Node(5)},
	}
	if pickComponentToExtend(components, 3) != nil {
		t.Error()
	}
}
