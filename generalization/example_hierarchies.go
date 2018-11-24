package generalization

func GetGradeHierarchy1() *Hierarchy {
	return &Hierarchy{Partitions: [][]*Partition{
		{
			NewPartition("A+"),
			NewPartition("A"),
			NewPartition("A-"),
			NewPartition("B+"),
			NewPartition("B"),
			NewPartition("B-"),
			NewPartition("C+"),
			NewPartition("C"),
			NewPartition("C-"),
		},
		{
			NewPartition("A+", "A", "A-"),
			NewPartition("B+", "B", "B-"),
			NewPartition("C+", "C", "C-"),
		},
		{
			NewPartition("A+", "A", "A-", "B+", "B", "B-", "C+", "C", "C-"),
		},
	}}
}
