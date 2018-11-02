package algorithm

import (
	"github.com/gyuho/goraph"
	"k-anon/model"
	"strconv"
	"testing"
)

func TestBuildCoreGraph_NodeCount(t *testing.T) {
	table := &model.Table{
		Rows: []*model.Vector{
			{},
			{},
			{},
		},
	}
	g := BuildCoreGraph(table)
	nodeCount := g.GetNodeCount()
	if nodeCount != len(table.Rows) {
		t.Errorf("Incorrect node count")
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
		_, err := g.GetNode(goraph.StringID(strconv.Itoa(i)))
		if err != nil {
			t.Errorf("Node index %d was not in the graph", i)
		}
	}
}
