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
		Partitions: [][]*ItemSet{
			{
				NewItemSet("A+"),
				NewItemSet("A"),
				NewItemSet("A-"),
				NewItemSet("B+"),
				NewItemSet("B"),
				NewItemSet("B-"),
			},
			{
				NewItemSet("A+", "A", "A-"),
				NewItemSet("B+", "B", "B-"),
			},
			{
				NewItemSet("A+", "A", "A-", "B+", "B", "B-"),
			},
		},
	})
	return generalizer
}
