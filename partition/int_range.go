package partition

import (
	"fmt"
	"math"
)

// IntRange represents an interval of integers with bounds min and max.
type IntRange struct {
	min, max int
}

// NewIntRange creates a new instance of IntRange with min and max bounds.
func NewIntRange(min, max int) *IntRange {
	if min > max {
		return &IntRange{min: min, max: min}
	}
	return &IntRange{min: min, max: max}
}

// Min returns the min bound of the range.
func (r *IntRange) Min() float64 {
	return float64(r.min)
}

// Max returns the max bound of the range.
func (r *IntRange) Max() float64 {
	return float64(r.max)
}

// Contains returns true when the range contains the item.
// Note, that the item must be of type int, otherwise the result is always false.
func (r *IntRange) Contains(item interface{}) bool {
	i, success := item.(int)
	if !success {
		return false
	}
	return r.min <= i && i <= r.max
}

// ContainsPartition returns true, when this partition contains the other partition.
// Note, that the other partition must be of type IntRange, otherwise the result is always false.
func (r *IntRange) ContainsPartition(other Partition) bool {
	r2, success := other.(*IntRange)
	if success {
		return r.containsIntRange(r2)
	}
	set, success := other.(*Set)
	if success {
		return r.containsSet(set)
	}
	return false
}

// Equals returns true when this partition equals the other partition.
// Note, that the other partition must be of type IntRange, otherwise the result is always false.
func (r *IntRange) Equals(other Partition) bool {
	r2, success := other.(*IntRange)
	if !success {
		return false
	}
	return r.min == r2.min && r.max == r2.max
}

// String returns the string representation of the partition.
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
func (r *IntRange) Split() (r1, r2 Range) {
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

// MaxSplit returns the number of times this partition can be split.
func (r *IntRange) MaxSplit() int {
	return int(math.Ceil(math.Log2(float64(r.max - r.min + 1))))
}

// InitItem returns a newly initalized Range from the given item.
func (r *IntRange) InitItem(item interface{}) Range {
	intVal, ok := item.(int)
	if ok {
		return NewIntRange(intVal, intVal)
	}
	uintVal, ok := item.(uint)
	if ok {
		return NewIntRange(int(uintVal), int(uintVal))
	}
	floatVal, _ := item.(float64)
	return NewIntRange(int(floatVal), int(floatVal))
}

func (r *IntRange) containsIntRange(other *IntRange) bool {
	return r.min <= other.min && other.max <= r.max
}

func (r *IntRange) containsSet(other *Set) bool {
	for item := range other.Items {
		if !r.Contains(item) {
			return false
		}
	}
	return true
}
