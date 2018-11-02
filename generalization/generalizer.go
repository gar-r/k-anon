package generalization

// Generalizer represents a value generalization procedure.
// Generalization means, that a value from a given domain is replaced with a less specific,
// but semantically consistent value from the same domain.
type Generalizer interface {
	// Generalize takes a value and generalizes it 'n' levels further.
	// It returns a generalized Partition containing the item, and other items generalized into the same partition.
	// This method will return nil if the value cannot be generalized to the given level.
	Generalize(item interface{}, n int) *Partition

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

func (g *HierarchyGeneralizer) Generalize(item interface{}, n int) *Partition {
	return g.hierarchy.Find(item, n)
}

func (g *HierarchyGeneralizer) Levels() int {
	return g.hierarchy.GetLevelCount()
}

// Suppressor is a special kind of Generalizer, which only has a single generalization level, suppress.
// Suppressing a value will simply replace it with the '*' token.
type Suppressor struct {
}

// Generalize returns a Partition containing either the value itself (n=0), or the '*' token
// representing the suppressed value (n=1).
// In all other cases it returns nil.
func (s *Suppressor) Generalize(item interface{}, n int) *Partition {
	if n == 0 {
		return NewPartition(item)
	}
	if n == 1 {
		return NewPartition("*")
	}
	return nil
}

func (s *Suppressor) Levels() int {
	return 1
}
