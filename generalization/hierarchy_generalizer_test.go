package generalization

import (
	"bitbucket.org/dargzero/k-anon/hierarchy"
	"bitbucket.org/dargzero/k-anon/partition"
	"bitbucket.org/dargzero/k-anon/testutil"
	"testing"
)

func TestHierarchyGeneralizer_Levels(t *testing.T) {
	h, _ := hierarchy.Build(partition.NewSet(),
		hierarchy.N(partition.NewSet(),
			hierarchy.N(partition.NewSet())))

	generalizer := &HierarchyGeneralizer{Hierarchy: h}
	testutil.AssertEquals(3, generalizer.Levels(), t)
}

func TestHierarchyGeneralizer_Generalize(t *testing.T) {

	hierarchy := hierarchy.GetGradeHierarchy()
	generalizer := &HierarchyGeneralizer{Hierarchy: hierarchy}

	t.Run("partition not in Hierarchy", func(t *testing.T) {
		p := partition.NewSet("missing")
		actual := generalizer.Generalize(p, 0)
		testutil.AssertNil(actual, t)
	})

	t.Run("generalize to level 0", func(t *testing.T) {
		p := partition.NewSet("B")
		actual := generalizer.Generalize(p, 0)
		expected := partition.NewSet("B")
		if !expected.Equals(actual) {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

	t.Run("generalize to level 1", func(t *testing.T) {
		p := partition.NewSet("B")
		actual := generalizer.Generalize(p, 1)
		expected := partition.NewSet("B+", "B", "B-")
		if !expected.Equals(actual) {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

	t.Run("generalize to level 2", func(t *testing.T) {
		p := partition.NewSet("B")
		actual := generalizer.Generalize(p, 2)
		expected := partition.NewSet("A+", "A", "A-", "B+", "B", "B-", "C+", "C", "C-")
		if !expected.Equals(actual) {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

	t.Run("partition cannot be generalized further", func(t *testing.T) {
		p := partition.NewSet("B")
		actual := generalizer.Generalize(p, 3)
		testutil.AssertNil(actual, t)
	})

	t.Run("continue generalization", func(t *testing.T) {
		p := partition.NewSet("B+", "B", "B-")
		actual := generalizer.Generalize(p, 2)
		expected := partition.NewSet("A+", "A", "A-", "B+", "B", "B-", "C+", "C", "C-")
		if !expected.Equals(actual) {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})
}
