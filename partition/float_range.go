package partition

import (
	"fmt"
	"gonum.org/v1/gonum/floats"
)

const delta = 0.000000001

type FloatRange struct {
	min, max float64
}

func NewFloatRange(min, max float64) *FloatRange {
	if min > max {
		return &FloatRange{min: min, max: min}
	}
	return &FloatRange{min: min, max: max}
}

func (r *FloatRange) Min() float64 {
	return r.min
}

func (r *FloatRange) Max() float64 {
	return r.max
}

func (r *FloatRange) Contains(item interface{}) bool {
	f, success := item.(float64)
	if !success {
		return false
	}
	return r.min <= f && f <= r.max
}

func (r *FloatRange) ContainsPartition(other Partition) bool {
	r2, success := other.(*FloatRange)
	if success {
		return r.containsFloatRange(r2)
	}
	set, success := other.(*Set)
	if success {
		return r.containsSet(set)
	}
	return false
}

func (r *FloatRange) Equals(other Partition) bool {
	r2, success := other.(*FloatRange)
	if !success {
		return false
	}

	return floats.EqualWithinAbs(r.min, r2.min, delta) &&
		floats.EqualWithinAbs(r.max, r2.max, delta)
}

func (r *FloatRange) String() string {
	if floats.EqualWithinAbs(r.min, r.max, delta) {
		return fmt.Sprintf("(%f)", r.min)
	}
	return fmt.Sprintf("(%f..%f)", r.min, r.max)
}

func (r *FloatRange) CanSplit() bool {
	return !floats.EqualWithinAbs(r.min, r.max, delta)
}

// Split creates two new IntRanges from the original one by splitting it at the median
func (r *FloatRange) Split() (r1, r2 Range) {
	l := r.max - r.min
	cut := l / 2
	r1 = NewFloatRange(r.min, r.min+cut)
	r2 = NewFloatRange(r.min+cut, r.max)
	return
}

func (r *FloatRange) MaxSplit() int {
	return countSplit(r) + 1
}

func (r *FloatRange) InitItem(item interface{}) Range {
	val, _ := item.(float64)
	return NewFloatRange(val, val)
}

func (r *FloatRange) containsFloatRange(other *FloatRange) bool {
	return r.min <= other.min && other.max <= r.max
}

func (r *FloatRange) containsSet(other *Set) bool {
	for item := range other.Items {
		if !r.Contains(item) {
			return false
		}
	}
	return true
}
