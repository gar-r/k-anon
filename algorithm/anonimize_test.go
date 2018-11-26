package algorithm

import (
	"gonum.org/v1/gonum/graph/topo"
	"k-anon/generalization"
	"k-anon/model"
	"k-anon/testutil"
	"testing"
)

func TestAnonimizer_AnonimizeData(t *testing.T) {
	gen1 := generalization.GetIntGeneralizer()
	gen2 := generalization.GetGradeGeneralizer()
	table := &model.Table{
		Rows: []*model.Vector{
			{
				Items: []*model.Data{
					model.NewIdentifier(9, gen1),
					model.NewIdentifier("A+", gen2),
				},
			},
			{
				Items: []*model.Data{
					model.NewIdentifier(2, gen1),
					model.NewIdentifier("B-", gen2),
				},
			},
			{
				Items: []*model.Data{
					model.NewIdentifier(6, gen1),
					model.NewIdentifier("A-", gen2),
				},
			},
			{
				Items: []*model.Data{
					model.NewIdentifier(4, gen1),
					model.NewIdentifier("B+", gen2),
				},
			},
		},
	}
	anon := &Anonimizer{
		table: table,
		k:     2,
	}
	result := anon.AnonimizeData()
	assertKAnonimity(table, result, 2, t)
}

func assertKAnonimity(table *model.Table, data [][]*generalization.Partition, k int, t *testing.T) {
	for i, r1 := range data {
		count := 0
		for _, r2 := range data {
			if inSamePartition(r1, r2, func(col int) bool {
				return table.Rows[0].Items[col].IsIdentifier()
			}) {
				count++
			}
		}
		if count < k {
			t.Errorf("k-anonimity violated in row %v", i)
		}
	}
}

func inSamePartition(r1, r2 []*generalization.Partition, isIdColumn func(int) bool) bool {
	for col := 0; col < len(r1); col++ {
		p1 := r1[col]
		p2 := r2[col]
		if isIdColumn(col) && !p1.Equals(p2) {
			return false
		}
	}
	return true
}

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
	gen := generalization.GetIntGeneralizer()
	data := []*model.Data{
		model.NewIdentifier(1, gen),
		model.NewIdentifier(2, gen),
		model.NewIdentifier(7, gen),
	}
	partitions := generalize(data)
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
	partitions := generalize(data)
	p1 := generalization.NewPartition("test1")
	p2 := generalization.NewPartition("test2")
	p3 := generalization.NewPartition("test3")
	for _, p := range partitions {
		if !(p.Equals(p1) || p.Equals(p2) || p.Equals(p3)) {
			t.Errorf("incorrect partition: %v", p)
		}
	}
}

func TestAnonimize(t *testing.T) {
	gen1 := generalization.GetIntGeneralizer()
	gen2 := generalization.GetGradeGeneralizer()
	groups := []*model.Vector{
		{
			Items: []*model.Data{
				model.NewIdentifier(9, gen1),
				model.NewIdentifier("A+", gen2),
				model.NewNonIdentifier("data1"),
			},
		},
		{
			Items: []*model.Data{
				model.NewIdentifier(8, gen1),
				model.NewIdentifier("A", gen2),
				model.NewNonIdentifier("data2"),
			},
		},
		{
			Items: []*model.Data{
				model.NewIdentifier(6, gen1),
				model.NewIdentifier("A-", gen2),
				model.NewNonIdentifier("data3"),
			},
		},
	}
	partitions := anonimize(groups)
	testutil.AssertEquals(3, len(partitions), t)
	assertSamePartition([]*generalization.Partition{
		partitions[0][0],
		partitions[1][0],
		partitions[2][0]}, t)
	assertSamePartition([]*generalization.Partition{
		partitions[0][1],
		partitions[1][1],
		partitions[2][1]}, t)
	partitions[0][2].Equals(generalization.NewPartition("data1"))
	partitions[1][2].Equals(generalization.NewPartition("data2"))
	partitions[2][2].Equals(generalization.NewPartition("data3"))
}

func assertSamePartition(p []*generalization.Partition, t *testing.T) {
	first := p[0]
	for i := 1; i < len(p); i++ {
		if !first.Equals(p[i]) {
			t.Errorf("partitions are not equal: %v, %v", first, p[i])
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
