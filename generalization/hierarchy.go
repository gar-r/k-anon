package generalization

import (
	"strings"
)

// Hierarchy represents a generalization hierarchy for a set of elements. Elements within the Hierarchy are from a given domain.
// GetLevel 0 of the hierarchy represents the elements separated into standalone partitions.
// Each subsequent level partitions the original elements into more generalized sets of elements.
// On the highest level of the Hierarchy all elements are grouped into a single partition.
// A given element can only appear once per generalization level across all partitions.
type Hierarchy struct {
	Partitions [][]*Partition
}

// Find locates the single Partition on the given level, that 'item' belongs to.
// If 'item' is not contained within any of the partitions, it returns nil.
func (h *Hierarchy) Find(item interface{}, level int) *Partition {
	l := h.GetLevel(level)
	if l != nil {
		for _, partition := range l {
			if partition.Contains(item) {
				return partition
			}
		}
	}
	return nil
}

// GetLevelCount returns the number of levels in the Hierarchy
func (h *Hierarchy) GetLevelCount() int {
	return len(h.Partitions)
}

func (h *Hierarchy) GetLevel(level int) []*Partition {
	if 0 <= level && level < h.GetLevelCount() {
		return h.Partitions[level]
	}
	return nil
}

// IsValid checks if all levels of the hierarchy are valid.
// In practice this means, that each element should only appear once per hierarchy level.
// In addition, combining all the partitions in each level should yield the same set as the highest generalization level.
func (h *Hierarchy) IsValid() bool {
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

// String returns a string representation of the hierarchy
func (h *Hierarchy) String() string {
	b := strings.Builder{}
	for _, level := range h.Partitions {
		b.WriteString("[")
		for _, partition := range level {
			b.WriteString(partition.String())
		}
		b.WriteString("]\n")
	}
	return strings.TrimSpace(b.String())
}

// Equals returns true, if the given two hierarchies are equal.
func Equals(h1, h2 *Hierarchy) bool {
	if h1.GetLevelCount() != h2.GetLevelCount() {
		return false
	}
	for i := 0; i < h1.GetLevelCount(); i++ {
		l1 := h1.GetLevel(i)
		l2 := h2.GetLevel(i)
	partition:
		for _, p1 := range l1 {
			for _, p2 := range l2 {
				if p1.Equals(p2) {
					continue partition
				}
			}
			return false
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
