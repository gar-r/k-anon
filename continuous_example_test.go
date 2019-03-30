package k_anon

import (
	"bitbucket.org/dargzero/k-anon/generalization"
	"bitbucket.org/dargzero/k-anon/model"
	"fmt"
)

func ExampleAnonymizer_Continuous() {

	// define the schema & table
	table := model.NewTable(&model.Schema{
		Columns: []*model.Column{
			{"Name", &generalization.Suppressor{}},
			{"Status", nil},
			{"Gender", &generalization.Suppressor{}},
			{"Age", generalization.NewIntRangeGeneralizer(0, 150)},
			{"Kids", generalization.NewIntRangeGeneralizer(0, 2)},
			{"Income", generalization.NewIntRangeGeneralizer(10000, 50000)},
			{"Grade", generalization.ExampleGradeGeneralizer()},
			{"Motto", &generalization.PrefixGeneralizer{MaxWords: 100}},
		},
	})

	anon := &Anonymizer{
		table: table,
		k:     2,
	}

	// anonymize first chunk of data
	table.AddRow("Joe", "employee", "male", 25, 0, 16700, "A", "cats are wonderful little beings")
	table.AddRow("Jane", "client", "female", 25, 0, 15250, "A+", "cats are my favorite kind of animals ")
	table.AddRow("Jack", "employee", "male", 30, 2, 31400, "B+", "cats are very unique")
	table.AddRow("Janet", "employee", "female", 30, 1, 38900, "A+", "cats are interesting")
	table.AddRow("Steve", "client", "male", 28, 2, 44350, "B", "cats are my only pets")
	table.AddRow("Sarah", "client", "female", 28, 1, 15580, "A+", "cats are my favorite!")
	table.AddRow("Ben", "employee", "male", 25, 2, 40250, "B-", "cats are interesting, but sometimes also egoistic")
	table.AddRow("Anne", "client", "female", 30, 2, 35700, "A", "cats are my favorite kind of animals")

	anon.Anonymize()

	// anonymize second chunk of data
	table.AddRow("Michelle", "employee", "female", 27, 1, 22400, "B", "cats are secretly extraterrestrials")
	table.AddRow("Perseus", "client", "male", 31, 2, 38600, "A+", "dogs are mischievous")
	table.AddRow("Odysseus", "client", "male", 39, 0, 39250, "A", "dogs are war bringers")
	table.AddRow("Helene", "employee", "female", 29, 2, 21900, "B+", "dogs are silky and furry")

	anon.Anonymize()

	// anonymize third chunk of data
	table.AddRow("Donald", "client", "male", 26, 2, 24550, "A", "cats are secretly extraterrestrials")
	table.AddRow("Hillary", "client", "female", 70, 2, 38700, "B-", "dogs are loyal")
	table.AddRow("George", "client", "male", 65, 1, 39990, "B", "dogs are war bringers")
	table.AddRow("Victor", "employee", "male", 45, 0, 21000, "A-", "cats are silky and furry")

	anon.Anonymize()

	// process & print results
	fmt.Printf("%v", anon.table)

	// the resulting table will be similar to the below:

	// Name	Status		Gender	Age		Kids	Income			Grade		Motto
	// *	employee	*		[25]	[0]		[15000..17499]	[A, A+, A-]	cats are
	// *	client		*		[25]	[0]		[15000..17499]	[A, A+, A-]	cats are
	// *	employee	male	[0..74]	[0..2]	[30000..50000]	[A, A+, A-, B, B+, B-, C, C+, C-]	*
	// *	employee	female	[27..31][1..2]	[10000..50000]	[A, A+, A-]	cats are
	// *	client		male	[0..74]	[1..2]	[30000..50000]	[B, B+, B-]	*
	// *	client		female	[27..31][1..2]	[10000..50000]	[A, A+, A-]	cats are
	// *	employee	male	[0..74]	[1..2]	[30000..50000]	[B, B+, B-]	*
	// *	client		female	[27..31][1..2]	[10000..50000]	[A, A+, A-]	cats are
	// *	employee	female	[0..74]	[1..2]	[10000..50000]	[B, B+, B-]	*
	// *	client		male	[0..74]	[0..2]	[30000..50000]	[A, A+, A-, B, B+, B-, C, C+, C-]	*
	// *	client		male	[0..74]	[0..2]	[30000..50000]	[A, A+, A-, B, B+, B-, C, C+, C-]	*
	// *	employee	female	[0..74]	[1..2]	[10000..50000]	[B, B+, B-]	*
	// *	client		male	[0..74]	[0..2]	[20000..24999]	[A, A+, A-]	cats are
	// *	client		female	[0..74]	[1..2]	[10000..50000]	[B, B+, B-]	*
	// *	client		male	[0..74]	[1..2]	[30000..50000]	[B, B+, B-]	*
	// *	employee	male	[0..74]	[0..2]	[20000..24999]	[A, A+, A-]	cats are
}
