package k_anon

import (
	"bitbucket.org/dargzero/k-anon/generalization"
	"bitbucket.org/dargzero/k-anon/model"
	"fmt"
)

// This is a testable example to demonstrate the usage of the anonymizer.
// The outline of the steps you need to take is the following:
// 1) Define a table
//    a) define the schema (friendly name and generalizer for each column)
//          * use a 'nil' generalizer for non-identifier (skipped) columns
//          * see the Generalizer interface to implement a custom generalizer
//    b) define the rows, conforming to the above schema:
//          * number of columns should match
//          * data type should be compatible with the assigned generalizer for the column
// 2) create an Anonymizer instance, and supply the table and k parameters
// 3) call the Anonymize() function
// The supplied table will be anonymized in-place. Note, that items in partitions
// in the resulting table are not ordered (always treat them as sets).
func ExampleAnonymizer_AnonymizeData() {

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
	table.AddRow("Joe", "employee", "male", 25, 0, 16700, "A", "cats are wonderful little beings")
	table.AddRow("Jane", "client", "female", 25, 0, 15250, "A+", "cats are my favorite kind of animals ")
	table.AddRow("Jack", "employee", "male", 30, 2, 31400, "B+", "cats are very unique")
	table.AddRow("Janet", "employee", "female", 30, 1, 38900, "A+", "cats are interesting")
	table.AddRow("Steve", "client", "male", 28, 2, 44350, "B", "cats are my only pets")
	table.AddRow("Sarah", "client", "female", 28, 1, 15580, "A+", "cats are my favorite!")
	table.AddRow("Ben", "employee", "male", 25, 2, 40250, "B-", "cats are interesting, but sometimes also egoistic")
	table.AddRow("Anne", "client", "female", 30, 2, 35700, "A", "cats are my favorite kind of animals")

	// create an anonymizer and run the anonymization
	anon := &Anonymizer{
		table: table,
		k:     2,
	}
	anon.Anonymize()

	// access & print the data
	fmt.Printf("%v", anon.table)

	// the above will produce a similar table:

	// Name	Status		Gender	Age			Kids	Income			Grade			Motto
	// *	employee	*		[18..36]	[0..2]	[10000..50000]	[A, A+, A-]		cats are
	// *	client		*		[18..36]	[0..2]	[10000..50000]	[A, A+, A-]		cats are
	// *	employee	male	[18..36]	[2]		[30000..50000]	[B, B+, B-]		cats are
	// *	employee	female	[27..31]	[1]		[10000..50000]	[A+]			cats are
	// *	client		male	[18..36]	[2]		[30000..50000]	[B, B+, B-]		cats are
	// *	client		female	[27..31]	[1]		[10000..50000]	[A+]			cats are
	// *	employee	male	[18..36]	[2]		[30000..50000]	[B, B+, B-]		cats are
	// * 	client		*		[18..36]	[0..2]	[10000..50000]	[A, A+, A-]		cats are
}
