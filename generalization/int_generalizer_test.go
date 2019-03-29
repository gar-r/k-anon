package generalization

import (
	"bitbucket.org/dargzero/k-anon/partition"
	"testing"
)

func TestNewIntGeneralizer(t *testing.T) {

	t.Run("empty range", func(t *testing.T) {
		g := NewIntGeneralizer(0, 0, 1)
		actual := g.hierarchy
		expected := &Hierarchy{}
		assertHierarchyEquals(expected, actual, t)
	})

	t.Run("single item range", func(t *testing.T) {
		g := NewIntGeneralizer(0, 1, 1)
		actual := g.hierarchy
		expected := &Hierarchy{
			Partitions: [][]*partition.ItemSet{
				{partition.NewItemSet(0)},
			},
		}
		assertHierarchyEquals(expected, actual, t)
	})

	t.Run("normal range", func(t *testing.T) {
		g := NewIntGeneralizer(-10, 10, 5)
		actual := g.hierarchy
		expected := &Hierarchy{
			Partitions: [][]*partition.ItemSet{
				{
					partition.NewItemSet(-10),
					partition.NewItemSet(-5),
					partition.NewItemSet(0),
					partition.NewItemSet(5),
				},
				{
					partition.NewItemSet(-10, -5),
					partition.NewItemSet(0, 5),
				},
				{
					partition.NewItemSet(-10, -5, 0, 5),
				},
			},
		}
		assertHierarchyEquals(expected, actual, t)
	})

	t.Run("test hierarchy is valid", func(t *testing.T) {
		g := NewIntGeneralizer(0, 10000, 1)
		if !g.hierarchy.IsValid() {
			t.Errorf("hierarchy is invalid: %v", g.hierarchy)
		}
	})

}

func TestNewIntGeneralizerFromItems(t *testing.T) {

	t.Run("empty set", func(t *testing.T) {
		g := NewIntGeneralizerFromItems()
		actual := g.hierarchy
		expected := &Hierarchy{}
		assertHierarchyEquals(expected, actual, t)
	})

	t.Run("single item", func(t *testing.T) {
		g := NewIntGeneralizerFromItems(1)
		actual := g.hierarchy
		expected := &Hierarchy{
			Partitions: [][]*partition.ItemSet{
				{partition.NewItemSet(1)},
			},
		}
		assertHierarchyEquals(expected, actual, t)
	})

	t.Run("items deduplicated", func(t *testing.T) {
		g := NewIntGeneralizerFromItems(1, 1, 1)
		actual := g.hierarchy
		expected := &Hierarchy{
			Partitions: [][]*partition.ItemSet{
				{partition.NewItemSet(1)},
			},
		}
		assertHierarchyEquals(expected, actual, t)
	})

	t.Run("multiple items", func(t *testing.T) {
		g := NewIntGeneralizerFromItems(1, 2, 3, 4)
		actual := g.hierarchy
		expected := &Hierarchy{
			Partitions: [][]*partition.ItemSet{
				{
					partition.NewItemSet(1),
					partition.NewItemSet(2),
					partition.NewItemSet(3),
					partition.NewItemSet(4),
				},
				{
					partition.NewItemSet(1, 2),
					partition.NewItemSet(3, 4),
				},
				{
					partition.NewItemSet(1, 2, 3, 4),
				},
			},
		}
		assertHierarchyEquals(expected, actual, t)
	})

	t.Run("median cut", func(t *testing.T) {
		g := NewIntGeneralizerFromItems(1, 2, 3)
		actual := g.hierarchy
		expected := &Hierarchy{
			Partitions: [][]*partition.ItemSet{
				{
					partition.NewItemSet(1),
					partition.NewItemSet(1),
					partition.NewItemSet(1),
				},
				{
					partition.NewItemSet(1),
					partition.NewItemSet(2, 3),
				},
				{
					partition.NewItemSet(1, 2, 3),
				},
			},
		}
		assertHierarchyEquals(expected, actual, t)
	})

}

//func Test_IntegerHierarchyBuilderPartitionCutAtMedian(t *testing.T) {
//	actual := buildIntHierarchy(1, 2, 3)
//	expected := &Hierarchy{
//		Partitions: [][]*ItemSet{
//			{NewItemSet(1), NewItemSet(2), NewItemSet(3)},
//			{NewItemSet(1), NewItemSet(2, 3)},
//			{NewItemSet(1, 2, 3)},
//		},
//	}
//	assertHierarchyEquals(expected, actual, t)
//}