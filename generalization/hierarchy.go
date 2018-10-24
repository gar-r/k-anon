package generalization

// Hierarchy represents a generalization hierarchy for a set of elements. Elements within the Hierarchy are from a given domain.
// Level 0 of the hierarchy represents the elements separated into standalone partitions.
// Each subsequent level partitions the original elements into more generalized sets of elements.
// On the highest level of the Hierarchy all elements are grouped into a single partition.
// A given element can only appear once per generalization level across all partitions.
type Hierarchy struct {
	Partitions [][]*Partition
}

// Levels returns the number of levels in the Hierarchy
func (h *Hierarchy) Levels() int {
	return len(h.Partitions)
}

// GetLevel returns the partitions on a given level
func (h *Hierarchy) GetLevel(level int) []*Partition {
	return h.Partitions[level]
}
