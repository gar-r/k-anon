package k_anon

import (
	"bitbucket.org/dargzero/k-anon/generalization"
	"bitbucket.org/dargzero/k-anon/model"
	"fmt"
)

func ExampleAnonymizer_Continuous() {

	// define the schema & Table
	table := model.NewTable(&model.Schema{
		Columns: []*model.Column{
			model.NewColumn("Name", &generalization.Suppressor{}),
			model.NewColumn("Status", nil),
			model.NewColumn("Gender", &generalization.Suppressor{}),
			model.NewColumn("Age", generalization.NewIntRangeGeneralizer(0, 150)),
			model.NewColumn("Kids", generalization.NewIntRangeGeneralizer(0, 2)),
			model.NewColumn("Income", generalization.NewIntRangeGeneralizer(10000, 50000)),
			model.NewColumn("A-Index", generalization.NewFloatRangeGeneralizer(0.0, 1.0)),
			model.NewColumn("Z-Index", generalization.NewFloatRangeGeneralizer(-0.5, 0.5)),
			model.NewWeightedColumn("Grade", generalization.ExampleGradeGeneralizer(), 1.2),
			model.NewWeightedColumn("Motto", &generalization.PrefixGeneralizer{MaxWords: 100}, 0.1),
		},
	})

	anon := &Anonymizer{
		Table: table,
		K:     2,
	}

	// anonymize first chunk of data
	table.AddRow("Joe", "employee", "male", 25, 0, 16700, 0.2, -0.35, "A", "cats are wonderful little beings")
	table.AddRow("Jane", "client", "female", 25, 0, 15250, 0.25, -0.3, "A-", "cats are my favorite kind of animals ")
	table.AddRow("Jack", "employee", "male", 30, 2, 31400, 0.6, -0.3, "A-", "cats are very unique")
	table.AddRow("Janet", "employee", "female", 30, 1, 38900, 0.9, -0.3, "A", "cats are interesting")
	table.AddRow("Steve", "client", "male", 28, 2, 44350, 0.9, -0.35, "A", "cats are my only pets")
	table.AddRow("Sarah", "client", "female", 28, 1, 15580, 0.56, -0.25, "A-", "cats are my favorite!")
	table.AddRow("Ben", "employee", "male", 25, 2, 40250, 0.9, -0.25, "A+", "cats are interesting, but sometimes also egoistic")
	table.AddRow("Anne", "client", "female", 30, 2, 35700, 0.6, -0.20, "A+", "cats are my favorite kind of animals")

	anon.Anonymize()

	// anonymize second chunk of data
	table.AddRow("Michelle", "employee", "female", 27, 1, 22400, 0.5, -0.2, "B", "cats are secretly extraterrestrials")
	table.AddRow("Perseus", "client", "male", 31, 2, 38600, 0.6, -0.15, "A+", "dogs are mischievous")
	table.AddRow("Odysseus", "client", "male", 39, 0, 39250, 0.5, -0.15, "A", "dogs are war bringers")
	table.AddRow("Helene", "employee", "female", 29, 2, 21900, 0.6, -0.2, "B+", "dogs are silky and furry")

	anon.Anonymize()

	// anonymize third chunk of data
	table.AddRow("Donald", "client", "male", 26, 2, 24550, 0.9, 0.15, "A", "cats are secretly extraterrestrials")
	table.AddRow("Hillary", "client", "female", 70, 2, 38700, 0.5, -0.05, "B-", "dogs are loyal")
	table.AddRow("George", "client", "male", 65, 1, 39990, 0.5, 0.15, "B", "dogs are war bringers")
	table.AddRow("Victor", "employee", "male", 45, 0, 21000, 0.85, -0.3, "A-", "cats are silky and furry")

	anon.Anonymize()

	// process & print results
	fmt.Printf("%v", anon.Table)

	// the resulting Table will be similar to the below:

	//	Name	Status		Gender	Age			Kids	Income			A-Index					Z-Index					Grade		Motto
	//	*		employee	male	[18..36]	[0..2]	[10000..50000]	(0.000000..1.000000)	(-0.350000)				[A]			cats are
	//	*		client		female	[18..36]	[0..2]	[15000..15624]	(0.000000..1.000000)	(-0.312500..-0.250000)	[A-]		cats are my
	//	*		employee	*		[30]		[2]		[30000..39999]	(0.600000)				(-0.500000..0.000000)	[A, A+, A-]	cats are
	//	*		employee	*		[18..36]	[1..2]	[30000..50000]	(0.900000)				(-0.312500..-0.250000)	[A, A+, A-]	cats are
	//	*		client		male	[18..36]	[0..2]	[10000..50000]	(0.000000..1.000000)	(-0.350000)				[A]			cats are
	//	*		client		female	[18..36]	[0..2]	[15000..15624]	(0.000000..1.000000)	(-0.312500..-0.250000)	[A-]		cats are my
	//	*		employee	*		[18..36]	[1..2]	[30000..50000]	(0.900000)				(-0.312500..-0.250000)	[A, A+, A-]	cats are
	//	*		client		*		[30]		[2]		[30000..39999]	(0.600000)				(-0.500000..0.000000)	[A, A+, A-]	cats are
	//	*		employee	female	[27..31]	[1..2]	[21875..22499]	(0.000000..1.000000)	(-0.200000)				[B, B+, B-]	*
	//	*		client		male	[0..74]		[0..2]	[37500..39999]	(0.000000..1.000000)	(-0.150000)				[A, A+, A-]	dogs are
	//	*		client		male	[0..74]		[0..2]	[37500..39999]	(0.000000..1.000000)	(-0.150000)				[A, A+, A-]	dogs are
	//	*		employee	female	[27..31]	[1..2]	[21875..22499]	(0.000000..1.000000)	(-0.200000)				[B, B+, B-]	*
	//	*		client		male	[0..74]		[0..2]	[20000..24999]	(0.750000..1.000000)	(-0.500000..0.500000)	[A, A+, A-]	cats are
	//	*		client		*		[65..74]	[1..2]	[37500..39999]	(0.500000)				(-0.500000..0.500000)	[B, B+, B-]	dogs are
	//	*		client		*		[65..74]	[1..2]	[37500..39999]	(0.500000)				(-0.500000..0.500000)	[B, B+, B-]	dogs are
	//	*		employee	male	[0..74]		[0..2]	[20000..24999]	(0.750000..1.000000)	(-0.500000..0.500000)	[A, A+, A-]	cats are

}
