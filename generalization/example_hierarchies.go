package generalization

import "bitbucket.org/dargzero/k-anon/partition"

func GetGradeHierarchy() *Hierarchy {
	return &Hierarchy{Partitions: [][]*partition.ItemSet{
		{
			partition.NewItemSet("A+"),
			partition.NewItemSet("A"),
			partition.NewItemSet("A-"),
			partition.NewItemSet("B+"),
			partition.NewItemSet("B"),
			partition.NewItemSet("B-"),
			partition.NewItemSet("C+"),
			partition.NewItemSet("C"),
			partition.NewItemSet("C-"),
		},
		{
			partition.NewItemSet("A+", "A", "A-"),
			partition.NewItemSet("B+", "B", "B-"),
			partition.NewItemSet("C+", "C", "C-"),
		},
		{
			partition.NewItemSet("A+", "A", "A-", "B+", "B", "B-", "C+", "C", "C-"),
		},
	}}
}
