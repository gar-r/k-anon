package partition

import "fmt"

// IntRange represents an interval of integers with bounds.
type IntRange struct {
	min, max int
}

func NewIntRange(min, max int) *IntRange {
	if min > max {
		return &IntRange{min, min}
	}
	return &IntRange{min, max}
}

func (r *IntRange) Contains(item interface{}) bool {
	i, success := item.(int)
	if !success {
		return false
	}
	return r.min <= i && i <= r.max
}

func (r *IntRange) ContainsPartition(other Partition) bool {
	r2, success := other.(*IntRange)
	if success {
		return r.containsIntRange(r2)
	}
	itemSet, success := other.(*Set)
	if success {
		return r.containsItemSet(itemSet)
	}
	return false
}

func (r *IntRange) Equals(other Partition) bool {
	r2, success := other.(*IntRange)
	if !success {
		return false
	}
	return r.min == r2.min && r.max == r2.max
}

func (r *IntRange) String() string {
	if r.min == r.max {
		return fmt.Sprintf("[%d]", r.min)
	}
	return fmt.Sprintf("[%d..%d]", r.min, r.max)
}

// CanSplit returns true if the range can be split further
func (r *IntRange) CanSplit() bool {
	return r.max > r.min
}

// Split creates two new IntRanges from the original one by splitting it at the median
func (r *IntRange) Split() (r1, r2 *IntRange) {
	l := r.max - r.min
	if l < 2 {
		r1 = NewIntRange(r.min, r.min)
		r2 = NewIntRange(r.max, r.max)
	} else {
		cut := (l + 1) / 2
		r1 = NewIntRange(r.min, r.min+cut-1)
		r2 = NewIntRange(r.min+cut, r.max)
	}
	return
}

func (r *IntRange) containsIntRange(other *IntRange) bool {
	return r.min <= other.min && other.max <= r.max
}

func (r *IntRange) containsItemSet(other *Set) bool {
	for item := range other.Items {
		if !r.Contains(item) {
			return false
		}
	}
	return true
}
