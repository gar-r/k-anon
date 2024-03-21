package hierarchy

import (
	"testing"

	"github.com/gar-r/k-anon/partition"
)

func TestAutoBuild(t *testing.T) {

	t.Run("empty input items", func(t *testing.T) {
		_, err := AutoBuild(2)
		if err == nil {
			t.Error("expected error, got none")
		}
	})

	t.Run("too few items", func(t *testing.T) {
		_, err := AutoBuild(3, "a", "b")
		if err == nil {
			t.Error("expected error, got none")
		}
	})

	t.Run("auto build test 1", func(t *testing.T) {
		h, err := AutoBuild(2, "a", "b", "c", "d")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		expected, _ := Build(partition.NewSet("a", "b", "c", "d"),
			N(partition.NewSet("a", "b"),
				N(partition.NewSet("a")),
				N(partition.NewSet("b"))),
			N(partition.NewSet("c", "d"),
				N(partition.NewSet("c")),
				N(partition.NewSet("d"))))
		assertHierarchyEquals(expected, h, t)
	})

	t.Run("auto build test 2", func(t *testing.T) {
		h, err := AutoBuild(2, "a", "b", "c", "d", "e")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		expected, _ := Build(partition.NewSet("a", "b", "c", "d", "e"),
			N(partition.NewSet("a", "b", "c"),
				N(partition.NewSet("a", "b"),
					N(partition.NewSet("a")),
					N(partition.NewSet("b"))),
				N(partition.NewSet("c"),
					N(partition.NewSet("c")))),
			N(partition.NewSet("d", "e"),
				N(partition.NewSet("d"),
					N(partition.NewSet("d"))),
				N(partition.NewSet("e"),
					N(partition.NewSet("e")))))
		assertHierarchyEquals(expected, h, t)
	})
}

func assertHierarchyEquals(expected, actual Hierarchy, t *testing.T) {
	t.Helper()
	if !expected.Partition().Equals(actual.Partition()) {
		t.Errorf("expected: %v, got %v", expected.Partition(), actual.Partition())
	}
	ec := expected.Children()
	ac := actual.Children()
	if len(ec) != len(ac) {
		t.Error("child count error")
	}
	for i := 0; i < len(ec); i++ {
		assertHierarchyEquals(ec[i], ac[i], t)
	}
}
