package hierarchy

import (
	"bitbucket.org/dargzero/k-anon/partition"
	"bitbucket.org/dargzero/k-anon/testutil"
	"testing"
)

func TestBuild(t *testing.T) {

	t.Run("single node", func(t *testing.T) {
		_, err := Build(partition.NewSet("test"))
		if err != nil {
			t.Errorf("%v", err)
		}
	})

	t.Run("non-full tree", func(t *testing.T) {
		_, err := Build(partition.NewSet(),
			N(partition.NewSet(),
				N(partition.NewSet()),
				N(partition.NewSet())),
			N(partition.NewSet()))
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("incorrect partitioning", func(t *testing.T) {
		_, err := Build(partition.NewSet(1, 2, 3, 4),
			N(partition.NewSet(1, 2),
				N(partition.NewSet(1)),
				N(partition.NewSet(2))),
			N(partition.NewSet(3, 4),
				N(partition.NewSet(3)),
				N(partition.NewSet("invalid"))))
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("build proper hierarchy", func(t *testing.T) {
		_, err := Build(partition.NewSet(1, 2, 3, 4),
			N(partition.NewSet(1, 2),
				N(partition.NewSet(1)),
				N(partition.NewSet(2))),
			N(partition.NewSet(3, 4),
				N(partition.NewSet(3)),
				N(partition.NewSet(4))))
		if err != nil {
			t.Errorf("%v", err)
		}
	})
}

func TestN(t *testing.T) {
	h := N(partition.NewSet(1, 2, 3))
	testutil.AssertNotNil(h, t)
}

func TestNode_Partition(t *testing.T) {
	expected := partition.NewSet(1, 2, 3)
	h := N(expected)
	if !expected.Equals(h.Partition()) {
		t.Errorf("expected %v, got %v", expected, h.Partition())
	}
}

func TestNode_Levels(t *testing.T) {
	h := GetGradeHierarchy()
	testutil.AssertEquals(3, h.Levels(), t)
}

func TestNode_Children(t *testing.T) {

	t.Run("children", func(t *testing.T) {
		h := GetGradeHierarchy()
		children := h.Children()
		expected1 := partition.NewSet("A+", "A", "A-")
		expected2 := partition.NewSet("B+", "B", "B-")
		actual1 := children[0].Partition()
		actual2 := children[1].Partition()
		if !expected1.Equals(actual1) {
			t.Errorf("expected %v, got %v", expected1, actual1)
		}
		if !expected2.Equals(actual2) {
			t.Errorf("expected %v, got %v", expected2, actual2)
		}
	})

	t.Run("leaf nodes", func(t *testing.T) {
		h := GetGradeHierarchy()
		n := h.Find(partition.NewSet("A"))
		testutil.AssertNil(n.Children(), t)
	})

}

func TestNode_Find(t *testing.T) {

	h := GetGradeHierarchy()

	t.Run("find non-existing partition", func(t *testing.T) {
		p := partition.NewSet("X")
		n := h.Find(p)
		testutil.AssertNil(n, t)
	})

	t.Run("find existing partition", func(t *testing.T) {
		p := partition.NewSet("A")
		n := h.Find(p)
		if !p.Equals(n.Partition()) {
			t.Errorf("expected %v, got %v", p, n.Partition())
		}
	})

}

func TestNode_Parent(t *testing.T) {

	h := GetGradeHierarchy()

	t.Run("parent of root", func(t *testing.T) {
		testutil.AssertNil(h.Parent(), t)
	})

	t.Run("parent of child node", func(t *testing.T) {
		part := partition.NewSet("A+", "A", "A-")
		parent := h.Find(part).Parent()
		testutil.AssertEquals(h.Partition(), parent.Partition(), t)
	})

	t.Run("parent of leaf node", func(t *testing.T) {
		part := partition.NewSet("A")
		parent := h.Find(part).Parent()
		expected := partition.NewSet("A+", "A", "A-")
		if !expected.Equals(parent.Partition()) {
			t.Errorf("expected %v, got %v", expected, parent.Partition())
		}
	})

}