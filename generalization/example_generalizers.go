package generalization

// Level 4: (1, 2, 3, 4, 5, 6, 7, 8, 9)
// Level 3: (1, 2, 3, 4) (5, 6, 7, 8, 9)
// Level 2: (1, 2) (3, 4) (5, 6) (7, 8, 9)
// Level 1: (1) (2) (3) (4)	(5) (6) (7) (8, 9)
// Level 0: (1) (2) (3) (4)	(5) (6) (7) (8) (9)
func GetIntGeneralizer() Generalizer {
	return NewIntGeneralizer(1, 10, 1)
}

// Level 2: (A+ A A- B+ B B-)
// Level 1: (A+ A A-) (B+ B B-)
// Level 0: (A+) (A) (A-) (B+) (B) (B-)
func GetGradeGeneralizer() *HierarchyGeneralizer {
	generalizer := NewHierarchyGeneralizer(&Hierarchy{
		Partitions: [][]*Partition{
			{
				NewPartition("A+"),
				NewPartition("A"),
				NewPartition("A-"),
				NewPartition("B+"),
				NewPartition("B"),
				NewPartition("B-"),
			},
			{
				NewPartition("A+", "A", "A-"),
				NewPartition("B+", "B", "B-"),
			},
			{
				NewPartition("A+", "A", "A-", "B+", "B", "B-"),
			},
		},
	})
	return generalizer
}
