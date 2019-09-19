package kanon

import (
	"bitbucket.org/dargzero/k-anon/generalization"
	"bitbucket.org/dargzero/k-anon/model"
	"fmt"
)

// This is a testable example to demonstrate the usage of the anonymizer.
// The outline of the steps you need to take is the following:
// 1) Define a Table
//    a) define the schema (friendly name and generalizer for each column)
//          * use a 'nil' generalizer for non-identifier (skipped) columns
//          * see the Generalizer interface on how to implement a custom generalizer
//    b) define the rows, conforming to the above schema:
//          * number of columns should match
//          * data type should be compatible with the assigned generalizer for the column
// 2) create an Anonymizer instance, and supply the Table and K parameters
// 3) call the Anonymize() function
// The supplied Table will be anonymized in-place. Note, that items in partitions
// in the resulting Table are not ordered (always treat them as sets).
func ExampleAnonymizer_AnonymizeData() {

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
	table.AddRow("Joe", "employee", "male", 25, 0, 16700, 0.2, -0.35, "A", "cats are wonderful little beings")
	table.AddRow("Jane", "client", "female", 25, 0, 15250, 0.25, -0.3, "A-", "cats are my favorite kind of animals ")
	table.AddRow("Jack", "employee", "male", 30, 2, 31400, 0.6, -0.3, "A-", "cats are very unique")
	table.AddRow("Janet", "employee", "female", 30, 1, 38900, 0.9, -0.3, "A", "cats are interesting")
	table.AddRow("Steve", "client", "male", 28, 2, 44350, 0.9, -0.35, "A", "cats are my only pets")
	table.AddRow("Sarah", "client", "female", 28, 1, 15580, 0.56, -0.25, "A-", "cats are my favorite!")
	table.AddRow("Ben", "employee", "male", 25, 2, 40250, 0.9, -0.25, "A+", "cats are interesting, but sometimes also egoistic")
	table.AddRow("Anne", "client", "female", 30, 2, 35700, 0.6, -0.20, "A+", "cats are my favorite kind of animals")

	// create an anonymizer and run the anonymization
	anon := &Anonymizer{
		Table: table,
		K:     2,
	}
	err := anon.Anonymize()
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	// access & print the data
	fmt.Printf("%v", anon.Table)

	// the above will produce a similar Table:

	//	Name	Status		Gender	Age			Kids	Income			A-Index					Z-Index					Grade		Motto
	//	*		employee	*		[18..36]	[0..2]	[10000..50000]	(0.000000..1.000000)	(-0.375000..-0.250000)	[A]			cats are
	//	*		client		female	[18..36]	[0..2]	[15000..15624]	(0.000000..1.000000)	(-0.312500..-0.250000)	[A-]		cats are my
	//	*		employee	*		[30]		[2]		[30000..39999]	(0.600000)				(-0.500000..0.000000)	[A, A+, A-]	cats are
	//	*		employee	*		[18..36]	[0..2]	[10000..50000]	(0.000000..1.000000)	(-0.375000..-0.250000)	[A]			cats are
	//	*		client		male	[18..36]	[2]		[40000..44999]	(0.900000)				(-0.375000..-0.250000)	[A, A+, A-]	cats are
	//	*		client		female	[18..36]	[0..2]	[15000..15624]	(0.000000..1.000000)	(-0.312500..-0.250000)	[A-]		cats are my
	//	*		employee	male	[18..36]	[2]		[40000..44999]	(0.900000)				(-0.375000..-0.250000)	[A, A+, A-]	cats are
	//	*		client		*		[30]		[2]		[30000..39999]	(0.600000)				(-0.500000..0.000000)	[A, A+, A-]	cats are

}
