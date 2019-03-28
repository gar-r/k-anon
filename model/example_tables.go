package model

import "bitbucket.org/dargzero/k-anon/generalization"

func GetIntTable1() *Table {
	g := generalization.GetIntGeneralizer()
	return &Table{
		Schema: &Schema{
			Columns: []*Column{
				{"Col1", g},
				{"Col2", g},
				{"Col3", g},
				{"Col4", g},
			},
		},
		Rows: []*Row{
			NewRow(1, 1, 1, 1),
			NewRow(1, 1, 1, 2),
			NewRow(4, 5, 1, 1),
			NewRow(1, 3, 5, 7),
		},
	}
}

func GetStudentTable() *Table {
	dim2 := generalization.NewIntGeneralizerFromItems(25, 27, 28, 30)
	dim3 := generalization.NewIntGeneralizerFromItems(0, 1, 2)
	dim4 := generalization.NewIntGeneralizerFromItems(10, 15, 30, 35, 40, 45)
	dim5 := generalization.GetGradeGeneralizer()
	return &Table{
		Schema: &Schema{
			Columns: []*Column{
				{"Gender", &generalization.Suppressor{}},
				{"Col 2", dim2},
				{"Col 3", dim3},
				{"Col 4", dim4},
				{"Col 5", dim5},
			},
		},
		Rows: []*Row{
			NewRow("Male", 25, 0, 35, "A"),
			NewRow("Female", 25, 0, 45, "A+"),
			NewRow("Male", 30, 2, 30, "B"),
			NewRow("Female", 30, 1, 35, "B+"),
			NewRow("Male", 28, 1, 40, "A-"),
			NewRow("Female", 28, 1, 15, "B"),
			NewRow("Male", 27, 0, 15, "B-"),
			NewRow("Female", 27, 2, 30, "B"),
		},
	}
}

func GetMixedTable1() *Table {
	return &Table{
		Schema: &Schema{
			Columns: []*Column{
				{"Score", generalization.GetIntGeneralizer()},
				{"Grade", generalization.GetGradeGeneralizer()},
			},
		},
		Rows: []*Row{
			NewRow(9, "A+"),
			NewRow(8, "A"),
			NewRow(5, "B-"),
		},
	}
}

func GetMixedTable2() *Table {
	return &Table{
		Schema: &Schema{
			Columns: []*Column{
				{"Score", generalization.GetIntGeneralizer()},
				{"Grade", generalization.GetGradeGeneralizer()},
			},
		},
		Rows: []*Row{
			NewRow(9, "A+"),
			NewRow(2, "B-"),
			NewRow(6, "A-"),
			NewRow(4, "B+"),
		},
	}
}

func GetMixedTable3() *Table {
	return &Table{
		Schema: &Schema{
			Columns: []*Column{
				{"Score", generalization.GetIntGeneralizer()},
				{"Grade", generalization.GetGradeGeneralizer()},
				{"Motto", &generalization.PrefixGeneralizer{MaxWords: 5}},
				{"Remark", nil},
			},
		},
		Rows: []*Row{
			NewRow(9, "A+", "cats are wild", "data1"),
			NewRow(8, "A", "cats are evil", "data2"),
			NewRow(6, "A-", "cats are fluffy", "data3"),
		},
	}
}

// GetEmptyTable return an empty table with 3 rows
func GetEmptyTable() *Table {
	table := &Table{
		Rows: []*Row{
			{},
			{},
			{},
		},
	}
	return table
}
