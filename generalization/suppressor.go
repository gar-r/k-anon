package generalization

import "github.com/gar-r/k-anon/partition"

// Suppressor is a special kind of g, which only has a single generalization level, suppress.
// Suppressing a value will simply replace it with the '*' token.
type Suppressor struct {
}

// Generalize returns either the value itself (n=0), or the '*' token representing a suppressed value (n=1).
// In all other cases it returns nil.
func (s *Suppressor) Generalize(p partition.Partition, n int) partition.Partition {
	if n == 0 {
		return p
	}
	if n == 1 {
		return s.InitItem("*")
	}
	return nil
}

// Levels returns the number of levels of the generalizer.
func (s *Suppressor) Levels() int {
	return 2
}

// InitItem initializes the given item into a new partition.
func (s *Suppressor) InitItem(item interface{}) partition.Partition {
	return partition.NewItem(item)
}
