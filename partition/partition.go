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

	Min() float64
	Max() float64
	CanSplit() bool
	Split() (Range, Range)
	MaxSplit() int
}

func countSplit(r Range, n int) int {
	if !r.CanSplit() {
		return n
	}
	r1, r2 := r.Split()
	n1 := countSplit(r1, n+1)
	n2 := countSplit(r2, n+1)
	if n1 > n2 {
		return n1
	}
	return n2
}
