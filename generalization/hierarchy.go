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

// Valid checks if all levels of the hierarchy are valid.
// In practice this means, that each element should only appear once per hierarchy level.
// In addition, combining all the partitions in each level should yield the same set as the highest generalization level.
func (h *Hierarchy) Valid() bool {
	last := h.Partitions[len(h.Partitions)-1][0]
	for _, level := range h.Partitions {
		occurrences := makeOccurrenceMap(level)
		if len(occurrences) != len(last.items) {
			return false
		}
		for item, value := range occurrences {
			if value > 1 || !last.Contains(item) {
				return false
			}
		}
	}
	return true
}

func makeOccurrenceMap(level []*Partition) map[interface{}]int {
	occurrences := make(map[interface{}]int)
	for _, partition := range level {
		for item := range partition.items {
			occurrences[item] += 1
		}
	}
	return occurrences
}
