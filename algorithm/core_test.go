package algorithm

import (
	"k-anon/model"
	"testing"
)

func TestBuildCoreGraph_Count(t *testing.T) {
	table := &model.Table{
		Rows: []*model.Vector{
			{},
			{},
			{},
		},
	}
	g := BuildCoreGraph(table)
	nodeCount := g.Nodes().Len()
	edgeCount := g.Edges().Len()
	if nodeCount != len(table.Rows) {
		t.Errorf("Incorrect node count")
	}
	if edgeCount != 0 {
		t.Errorf("Core graph should not contain edges")
	}
}

func TestBuildCoreGraph_NodeNames(t *testing.T) {
	table := &model.Table{
		Rows: []*model.Vector{
			{},
			{},
			{},
		},
	}
	g := BuildCoreGraph(table)
	for i := range table.Rows {
		node := g.Node(int64(i))
		if node == nil {
			t.Errorf("Node index %d was not in the graph", i)
		}
	}
}
