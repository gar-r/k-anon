package generalization

import (
	"bitbucket.org/dargzero/k-anon/hierarchy"
)

// Level 4: [1..9]
// Level 3: [1..4] [5..9]
// Level 2: [1..2] [3..4] [5..6][(7..9]
// Level 1: [1] [2] [3] [4] [5] [6] [7] [8..9]
// Level 0: [1] [2] [3] [4] [5] [6] [7] [8] [9]
func ExampleIntGeneralizer() Generalizer {
	return NewIntRangeGeneralizer(1, 9)
}

// Level 2: (A+ A A- B+ B B- C+ C C-)
// Level 1: (A+ A A-) (B+ B B-) (C+ C C-)
// Level 0: (A+) (A) (A-) (B+) (B) (B-) (C+) (C) (C-)
func ExampleGradeGeneralizer() Generalizer {
	return &HierarchyGeneralizer{
		Hierarchy: hierarchy.GetGradeHierarchy(),
	}
}
