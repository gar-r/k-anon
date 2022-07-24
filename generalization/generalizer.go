package generalization

import "git.okki.hu/garric/k-anon/partition"

// Generalizer encapsulates a value generalization procedure.
// Generalization means, that a value from a given domain is replaced with a less specific,
// but semantically consistent value from the same domain.
type Generalizer interface {
	// Generalizes a partition by n levels and returns the generalized partition.
	// The generalized partition must contain the input partition.
	// This function will return nil if the partition cannot be generalized to the given level.
	Generalize(p partition.Partition, n int) partition.Partition

	// InitItem is called on newly added data items to convert them into an initial partition.
	InitItem(item interface{}) partition.Partition

	// Levels returns the maximum level of generalization.
	Levels() int
}
