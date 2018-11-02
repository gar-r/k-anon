package algorithm

import (
	"github.com/gyuho/goraph"
	"k-anon/generalization"
	"k-anon/model"
	"strconv"
	"testing"
)

// 0 -- 0.500 -→ 1
// 0 -- 1.750 -→ 2
// 0 -- 2.750 -→ 3
// 1 -- 2.750 -→ 3
// 1 -- 2.250 -→ 2
// 2 -- 3.750 -→ 3
func TestBuildCostGraph(t *testing.T) {
	generalizer := getExampleGeneralizer()
	table := getTestTable(generalizer)
	graph := BuildCostGraph(table)
	if graph.GetNodeCount() != 4 {
		t.Errorf("Graph should contain 4 nodes")
	}
	assertEdgeCost(t, graph, 0, 1, 0.5)
	assertEdgeCost(t, graph, 0, 2, 1.75)
	assertEdgeCost(t, graph, 0, 3, 2.75)
	assertEdgeCost(t, graph, 1, 3, 2.75)
	assertEdgeCost(t, graph, 1, 2, 2.25)
	assertEdgeCost(t, graph, 2, 3, 3.75)
}

func assertEdgeCost(t *testing.T, graph goraph.Graph, node1, node2 int, expectedCost float64) {
	id1 := goraph.StringID(strconv.Itoa(node1))
	id2 := goraph.StringID(strconv.Itoa(node2))
	cost, err := graph.GetWeight(id1, id2)
	if err != nil {
		t.Errorf("graph weight error nodes: %d,%d graph: %v", node1, node2, graph)
	}
	if expectedCost != cost {
		t.Errorf("expected cost %v, got %v", expectedCost, cost)
	}
}

func getNode(t *testing.T, graph goraph.Graph, id int) goraph.Node {
	node, err := graph.GetNode(goraph.StringID(strconv.Itoa(id)))
	if err != nil {
		t.Errorf("Error getting node %d from graph %v", id, graph)
	}
	return node
}

func getTestTable(generalizer generalization.Generalizer) *model.Table {
	return &model.Table{
		Rows: []*model.Vector{
			createVector([]int{1, 1, 1, 1}, generalizer),
			createVector([]int{1, 1, 1, 2}, generalizer),
			createVector([]int{4, 5, 1, 1}, generalizer),
			createVector([]int{1, 3, 5, 7}, generalizer),
		},
	}
}
