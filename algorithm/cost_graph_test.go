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
func TestBuildCostGraph1(t *testing.T) {
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

// 0 -- 0.750 -→ 1
// 0 -- 1.750 -→ 2
func TestBuildCostGraph2(t *testing.T) {
	generalizer1 := getExampleGeneralizer()
	generalizer2 := getExampleGeneralizer2()
	table := &model.Table{
		Rows: []*model.Vector{
			{
				Items: []*model.Data{
					model.NewData(9, generalizer1),
					model.NewData("A+", generalizer2),
				},
			},
			{
				Items: []*model.Data{
					model.NewData(8, generalizer1),
					model.NewData("A", generalizer2),
				},
			},
			{
				Items: []*model.Data{
					model.NewData(5, generalizer1),
					model.NewData("B-", generalizer2),
				},
			},
		},
	}
	graph := BuildCostGraph(table)
	if graph.GetNodeCount() != 3 {
		t.Errorf("Graph should contain 3 nodes")
	}
	assertEdgeCost(t, graph, 0, 1, 0.75)
	assertEdgeCost(t, graph, 0, 2, 1.75)
}

func assertEdgeCost(t *testing.T, graph goraph.Graph, node1, node2 int, expectedCost float64) {
	id1 := goraph.StringID(strconv.Itoa(node1))
	id2 := goraph.StringID(strconv.Itoa(node2))
	cost := getCost(t, graph, id1, id2)
	cost2 := getCost(t, graph, id2, id1)
	if expectedCost != cost || expectedCost != cost2 {
		t.Errorf("expected cost %v, got %v", expectedCost, cost)
	}
}

func getCost(t *testing.T, graph goraph.Graph, id1 goraph.StringID, id2 goraph.StringID) float64 {
	cost, err := graph.GetWeight(id1, id2)
	if err != nil {
		t.Errorf("graph weight error nodes: %v,%v graph: %v", id1, id2, graph)
	}
	return cost
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

func getExampleGeneralizer2() *generalization.HierarchyGeneralizer {
	generalizer := generalization.NewHierarchyGeneralizer(&generalization.Hierarchy{
		Partitions: [][]*generalization.Partition{
			{
				generalization.NewPartition("A+"),
				generalization.NewPartition("A"),
				generalization.NewPartition("A-"),
				generalization.NewPartition("B+"),
				generalization.NewPartition("B"),
				generalization.NewPartition("B-"),
			},
			{
				generalization.NewPartition("A+", "A", "A-"),
				generalization.NewPartition("B+", "B", "B-"),
			},
			{
				generalization.NewPartition("A+", "A", "A-", "B+", "B", "B-"),
			},
		},
	})
	return generalizer
}
