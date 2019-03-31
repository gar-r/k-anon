package partition

// Partition is a single partition or node in a generalization hierarchy
type Partition interface {

	// Contains returns true if the partition contains the given item
	Contains(item interface{}) bool

	// ContainsPartition returns true if this partition contains the given partition
	ContainsPartition(other Partition) bool

	// Equals returns true of this partition is equal to the other partition
	Equals(other Partition) bool

	// String returns a string representation of the partition
	String() string
}

type Range interface {
	Partition

	InitItem(item interface{}) Range
	Min() float64
	Max() float64
	CanSplit() bool
	Split() (Range, Range)
	MaxSplit() int
}

func countSplit(r Range) int {
	if !r.CanSplit() {
		return 0
	}
	_, r2 := r.Split()
	return countSplit(r2) + 1
}
