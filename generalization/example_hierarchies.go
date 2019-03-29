package generalization

func GetGradeHierarchy() *Hierarchy {
	return &Hierarchy{Partitions: [][]*ItemSet{
		{
			NewItemSet("A+"),
			NewItemSet("A"),
			NewItemSet("A-"),
			NewItemSet("B+"),
			NewItemSet("B"),
			NewItemSet("B-"),
			NewItemSet("C+"),
			NewItemSet("C"),
			NewItemSet("C-"),
		},
		{
			NewItemSet("A+", "A", "A-"),
			NewItemSet("B+", "B", "B-"),
			NewItemSet("C+", "C", "C-"),
		},
		{
			NewItemSet("A+", "A", "A-", "B+", "B", "B-", "C+", "C", "C-"),
		},
	}}
}
