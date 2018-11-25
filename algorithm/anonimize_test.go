package algorithm

import (
	"gonum.org/v1/gonum/graph/topo"
	"k-anon/model"
	"k-anon/testutil"
	"testing"
)

func TestGetGroups(t *testing.T) {
	v0 := model.CreateVector([]int{}, nil)
	v1 := model.CreateVector([]int{}, nil)
	v2 := model.CreateVector([]int{}, nil)
	v3 := model.CreateVector([]int{}, nil)
	v4 := model.CreateVector([]int{}, nil)
	table := &model.Table{Rows: []*model.Vector{v0, v1, v2, v3, v4}}
	a := &Anonimizer{
		table: table,
		k:     2,
	}
	g := CreateNodesUndirected(5)
	AddEdge(g, 0, 3)
	AddEdge(g, 1, 2)
	groups := a.getGroups(topo.ConnectedComponents(g))
	testutil.AssertEquals(3, len(groups), t)
	assertGroup(groups, v0, v3)
	assertGroup(groups, v1, v2)
	assertGroup(groups, v4)
}

func assertGroup(groups [][]*model.Vector, items ...*model.Vector) bool {
	for _, group := range groups {
		if composedOf(group, items...) {
			return true
		}
	}
	return false
}

func composedOf(group []*model.Vector, items ...*model.Vector) bool {
	if len(group) != len(items) {
		return false
	}
	for _, item := range items {
		if !contains(group, item) {
			return false
		}
	}
	return true
}

func contains(group []*model.Vector, item *model.Vector) bool {
	for _, i := range group {
		if i == item {
			return true
		}
	}
	return false
}
