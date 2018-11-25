package algorithm

import (
	"gonum.org/v1/gonum/graph/topo"
	"k-anon/generalization"
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

func TestGeneralizeIdentifier(t *testing.T) {
	gen := generalization.GetIntGeneralizer1()
	data := []*model.Data{
		model.NewIdentifier(1, gen),
		model.NewIdentifier(2, gen),
		model.NewIdentifier(7, gen),
	}
	partitions := generalizeIdentifier(data)
	expected := generalization.NewPartition(1, 2, 3, 4, 5, 6, 7, 8, 9)
	for _, p := range partitions {
		if !expected.Equals(p) {
			t.Errorf("incorrect partition: %v", p)
		}
	}
}

func TestGeneralizeNonIdentifier(t *testing.T) {
	data := []*model.Data{
		model.NewNonIdentifier("test1"),
		model.NewNonIdentifier("test2"),
		model.NewNonIdentifier("test3"),
	}
	partitions := generalizeNonIdentifier(data)
	p1 := generalization.NewPartition("test1")
	p2 := generalization.NewPartition("test2")
	p3 := generalization.NewPartition("test3")
	for _, p := range partitions {
		if !(p.Equals(p1) || p.Equals(p2) || p.Equals(p3)) {
			t.Errorf("incorrect partition: %v", p)
		}
	}
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
