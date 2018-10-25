package generalization

// Generalizer represents a value generalization procedure.
// Generalization means, that a value from a given domain is replaced with a less specific,
// but semantically consistent value from the same domain.
type Generalizer interface {
	// Generalize takes a value and generalizes it 'n' levels further.
	// It returns a generalized Partition containing the item, and other items generalized into the same partition.
	Generalize(item interface{}, n int) *Partition
}

// HierarchyGeneralizer is an implementation of Generalizer which uses a generalization Hierarchy to calculate generalized values.
type HierarchyGeneralizer struct {
	hierarchy *Hierarchy
}

// Creates a new HierarchyGeneralizer from the given h Hierarchy.
// In case the hierarchy is not valid, it returns nil.
func NewHierarchyGeneralizer(h *Hierarchy) *HierarchyGeneralizer {
	if !h.Valid() {
		return nil
	}
	return &HierarchyGeneralizer{
		hierarchy: h,
	}
}

// Generalize takes a value from the domain, and generalizes it n levels using the Hierarchy.
// It returns a generalized Partition containing the item, and other items generalized into the same partition,
// or nil if the value cannot be generalized to the given level with the hierarchy.
func (g *HierarchyGeneralizer) Generalize(item interface{}, n int) *Partition {
	return g.hierarchy.Find(item, n)
}

// StringGeneralizer is an implementation of Generalizer which works on strings only.
// When generalizing n levels on a string, it will truncate n characters off the end of the string,
// until it is reduced to the fully suppressed value: '*'.
type StringGeneralizer struct {
}

// Generalize returns a partition containing a single string, which is the n-generalized
// version if the input parameter 'item'. The result will contain '*' if the value is fully suppressed.
// The function returns nil if the value cannot be generalized further, or if it is not a string.
func (g *StringGeneralizer) Generalize(item interface{}, n int) *Partition {
	s, ok := item.(string)
	if !ok || n > len(s) || n < 0 {
		return nil
	}
	if n == len(s) {
		return NewPartition("*")
	}
	return NewPartition(s[0 : len(s)-n])
}
