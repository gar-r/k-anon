package partition

import (
	"fmt"

	"gonum.org/v1/gonum/floats/scalar"
)

const delta = 0.00001

// FloatRange encapsulates a range of float values between min and max.
type FloatRange struct {
	min, max float64
}

// NewFloatRange creates a new instance of FloatRange with given min and max values.
func NewFloatRange(min, max float64) *FloatRange {
	if min > max {
		return &FloatRange{min: min, max: min}
	}
	return &FloatRange{min: min, max: max}
}

// Min returns the min value of the range.
func (r *FloatRange) Min() float64 {
	return r.min
}

// Max returns the max value of the range.
func (r *FloatRange) Max() float64 {
	return r.max
}

// Contains returns true when the float range contains the given item.
// Note, that item must be a float value, otherwise the result is always false.
func (r *FloatRange) Contains(item interface{}) bool {
	f, success := item.(float64)
	if !success {
		return false
	}
	return r.min <= f && f <= r.max
}

// ContainsPartition returns true, when the float range contains the other partition.
// Note, that the other partition must be a float range as well.
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

// Equals returns true, when the float range equals the other partition.
// Note, that the other partition must be a float range as well, otherwise this method always returns false.
func (r *FloatRange) Equals(other Partition) bool {
	r2, success := other.(*FloatRange)
	if !success {
		return false
	}

	return scalar.EqualWithinAbs(r.min, r2.min, delta) &&
		scalar.EqualWithinAbs(r.max, r2.max, delta)
}

// String returns a string representation of the float range.
func (r *FloatRange) String() string {
	if scalar.EqualWithinAbs(r.min, r.max, delta) {
		return fmt.Sprintf("(%f)", r.min)
	}
	return fmt.Sprintf("(%f..%f)", r.min, r.max)
}

// CanSplit returns true when the float range can still be split into two float ranges.
func (r *FloatRange) CanSplit() bool {
	return !scalar.EqualWithinAbs(r.min, r.max, delta)
}

// Split creates two new IntRanges from the original one by splitting it at the median
func (r *FloatRange) Split() (r1, r2 Range) {
	l := r.max - r.min
	cut := l / 2
	r1 = NewFloatRange(r.min, r.min+cut)
	r2 = NewFloatRange(r.min+cut, r.max)
	return
}

// MaxSplit returns the number of times this range can be split.
func (r *FloatRange) MaxSplit() int {
	return countSplit(r)
}

// InitItem creates an initial range from the given item.
func (r *FloatRange) InitItem(item interface{}) Range {
	floatVal, ok := item.(float64)
	if ok {
		return NewFloatRange(floatVal, floatVal)
	}
	intVal, ok := item.(int)
	if ok {
		return NewFloatRange(float64(intVal), float64(intVal))
	}
	uintVal, _ := item.(uint)
	return NewFloatRange(float64(uintVal), float64(uintVal))
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
