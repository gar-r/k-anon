package algorithm

import (
	"bitbucket.org/dargzero/k-anon/generalization"
	"bitbucket.org/dargzero/k-anon/model"
	"bitbucket.org/dargzero/k-anon/testutil"
	"testing"
)

func TestBuildCostGraph(t *testing.T) {

	// 0 -- 0.500 -→ 1
	// 0 -- 1.750 -→ 2
	// 0 -- 2.750 -→ 3
	// 1 -- 2.750 -→ 3
	// 1 -- 2.250 -→ 2
	// 2 -- 3.750 -→ 3
	t.Run("cost graph 1", func(t *testing.T) {
		table := model.GetIntTable1()
		g, _ := BuildCostGraph(table)
		testutil.AssertEdgeCost(t, g, 0, 1, 0.5)
		testutil.AssertEdgeCost(t, g, 0, 2, 1.75)
		testutil.AssertEdgeCost(t, g, 0, 3, 2.75)
		testutil.AssertEdgeCost(t, g, 1, 3, 2.75)
		testutil.AssertEdgeCost(t, g, 1, 2, 2.25)
		testutil.AssertEdgeCost(t, g, 2, 3, 3.75)
	})

	// 0 -- 0.750 -→ 1
	// 0 -- 1.750 -→ 2
	t.Run("cost graph 2", func(t *testing.T) {
		table := model.GetMixedTable1()
		g, _ := BuildCostGraph(table)
		testutil.AssertEdgeCost(t, g, 0, 1, 0.75)
		testutil.AssertEdgeCost(t, g, 0, 2, 1.75)
	})

	t.Run("empty cost graph count", func(t *testing.T) {
		table := model.GetEmptyTable()
		g := buildEmptyCostGraph(table)
		nodeCount := g.Nodes().Len()
		edgeCount := g.Edges().Len()
		if nodeCount != len(table.GetRows()) {
			t.Errorf("incorrect node count")
		}
		testutil.AssertEquals(0, edgeCount, t)
	})

	t.Run("empty cost graph node names", func(t *testing.T) {
		table := model.GetEmptyTable()
		g := buildEmptyCostGraph(table)
		for i := range table.GetRows() {
			node := g.Node(int64(i))
			testutil.AssertNotNil(node, t)
		}
	})

	t.Run("add costs error", func(t *testing.T) {
		table := model.NewTable(&model.Schema{
			Columns: []*model.Column{
				model.NewColumn("Score", generalization.ExampleIntGeneralizer()),
				model.NewColumn("Grade", generalization.ExampleGradeGeneralizer()),
			},
		})
		table.AddRow(999, "A+")
		table.AddRow(8, "A")
		table.AddRow(5, "B-")
		_, err := BuildCostGraph(table)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

}
