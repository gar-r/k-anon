package generalization

import "bitbucket.org/dargzero/k-anon/partition"

// Generalizer represents a value generalization procedure.
// Generalization means, that a value from a given domain is replaced with a less specific,
// but semantically consistent value from the same domain.
type Generalizer interface {
	// Generalizes a partition by n levels and returns the generalized ItemSet.
	// The input partition can contain a single or multiple items.
	// The generalized partition contains all items from the input partition, and other items generalized into the same partition.
	// This function will return nil if the value cannot be generalized to the given level.
	Generalize(p partition.Partition, n int) partition.Partition

	// Levels returns the maximum level of generalization
	Levels() int
}

// HierarchyGeneralizer is an implementation of Generalizer which uses a generalization Hierarchy to calculate generalized values.
type HierarchyGeneralizer struct {
	hierarchy *Hierarchy
}

// Creates a new HierarchyGeneralizer from the given h Hierarchy.
// In case the hierarchy is not valid, it returns nil.
func NewHierarchyGeneralizer(h *Hierarchy) *HierarchyGeneralizer {
	if !h.IsValid() {
		return nil
	}
	return &HierarchyGeneralizer{
		hierarchy: h,
	}
}

func (g *HierarchyGeneralizer) Generalize(p partition.Partition, n int) partition.Partition {
	for l := n; l < g.Levels(); l++ { // continue searching in upper levels of the hierarchy
		level := g.hierarchy.GetLevel(l)
		for _, part := range level {
			if part.ContainsPartition(p) {
				return part
			}
		}
	}
	return nil
}

func (g *HierarchyGeneralizer) Levels() int {
	return g.hierarchy.GetLevelCount()
}

// Suppressor is a special kind of Generalizer, which only has a single generalization level, suppress.
// Suppressing a value will simply replace it with the '*' token.
type Suppressor struct {
}

// Generalize returns either the value itself (n=0), or the '*' token representing a suppressed value (n=1).
// In all other cases it returns nil.
func (s *Suppressor) Generalize(p partition.Partition, n int) partition.Partition {
	if n == 0 {
		return p
	}
	if n == 1 {
		return partition.NewItemSet("*")
	}
	return nil
}

func (s *Suppressor) Levels() int {
	return 2
}
