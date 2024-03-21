package generalization

import (
	"github.com/gar-r/k-anon/hierarchy"
	"github.com/gar-r/k-anon/partition"
)

// HierarchyGeneralizer is an implementation of Generalizer which uses a
// generalization Hierarchy to calculate generalized values.
type HierarchyGeneralizer struct {
	Hierarchy hierarchy.Hierarchy
}

// Generalize generalizes the given partition to level (n-l) of the Hierarchy, where n is the
// total number of levels in the Hierarchy and l is the level of p in the Hierarchy.
// Consequently for leaf partitions this method will always generalize them to level n.
func (g *HierarchyGeneralizer) Generalize(p partition.Partition, n int) partition.Partition {
	if n >= g.Levels() {
		return nil
	}
	h := g.Hierarchy.Find(p)
	if h == nil {
		return nil
	}
	l := h.Levels() - 1
	for i := l; i < n; i++ {
		h = h.Parent()
	}
	return h.Partition()
}

// Levels returns the number of levels in the hierarchy.
func (g *HierarchyGeneralizer) Levels() int {
	return g.Hierarchy.Levels()
}

// InitItem wraps the item in a new set containing only the item itself.
func (g *HierarchyGeneralizer) InitItem(item interface{}) partition.Partition {
	return partition.NewSet(item)
}
