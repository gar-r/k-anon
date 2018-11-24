package algorithm

import (
	"k-anon/model"
	"k-anon/testutil"
	"testing"
)

// 0 -- 0.500 -→ 1
// 0 -- 1.750 -→ 2
// 0 -- 2.750 -→ 3
// 1 -- 2.750 -→ 3
// 1 -- 2.250 -→ 2
// 2 -- 3.750 -→ 3
func TestBuildCostGraph1(t *testing.T) {
	table := model.GetIntTable1()
	g := BuildCostGraph(table)
	testutil.AssertEdgeCost(t, g, 0, 1, 0.5)
	testutil.AssertEdgeCost(t, g, 0, 2, 1.75)
	testutil.AssertEdgeCost(t, g, 0, 3, 2.75)
	testutil.AssertEdgeCost(t, g, 1, 3, 2.75)
	testutil.AssertEdgeCost(t, g, 1, 2, 2.25)
	testutil.AssertEdgeCost(t, g, 2, 3, 3.75)
}

// 0 -- 0.750 -→ 1
// 0 -- 1.750 -→ 2
func TestBuildCostGraph2(t *testing.T) {
	table := model.GetMixedTable1()
	g := BuildCostGraph(table)
	testutil.AssertEdgeCost(t, g, 0, 1, 0.75)
	testutil.AssertEdgeCost(t, g, 0, 2, 1.75)
}

func TestBuildEmptyCostGraph_Count(t *testing.T) {
	table := model.GetEmptyTable()
	g := buildEmptyCostGraph(table)
	nodeCount := g.Nodes().Len()
	edgeCount := g.Edges().Len()
	if nodeCount != len(table.Rows) {
		t.Errorf("incorrect node count")
	}
	testutil.AssertEquals(0, edgeCount, t)
}

func TestBuildEmptyCostGraph_NodeNames(t *testing.T) {
	table := model.GetEmptyTable()
	g := buildEmptyCostGraph(table)
	for i := range table.Rows {
		node := g.Node(int64(i))
		testutil.AssertNotNil(node, t)
	}
}
