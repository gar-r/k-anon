package partition

import "fmt"

// IntRange represents an interval of integers with bounds.
// For the interval the lower bound is always inclusive, while teh upper bound is exclusive.
type IntRange struct {
	lower, upper int
}

func NewIntRange(lower, upper int) *IntRange {
	return &IntRange{lower, upper}
}

func (r *IntRange) Contains(item interface{}) bool {
	i, success := item.(int)
	if !success {
		return false
	}
	return r.lower <= i && i < r.upper
}

func (r *IntRange) ContainsPartition(other Partition) bool {
	r2, success := other.(*IntRange)
	if success {
		return r.containsIntRange(r2)
	}
	itemSet, success := other.(*ItemSet)
	if success {
		return r.containsItemSet(itemSet)
	}
	return false
}

func (r *IntRange) containsIntRange(other *IntRange) bool {
	return r.lower <= other.lower && other.upper < r.upper
}

func (r *IntRange) containsItemSet(other *ItemSet) bool {
	for item := range other.Items {
		if !r.Contains(item) {
			return false
		}
	}
	return true
}

func (r *IntRange) Equals(other Partition) bool {
	r2, success := other.(*IntRange)
	if !success {
		return false
	}
	return r.lower == r2.lower && r.upper == r2.upper
}

func (r *IntRange) String() string {
	return fmt.Sprintf("[%d..%d)", r.lower, r.upper)
}
