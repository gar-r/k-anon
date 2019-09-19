# About this project

This is a Data Anonymization library written in Go. Anonymization is a process by which personal data is irreversibly altered in such a way that a data subject can no longer be identified directly or indirectly, either by the data controller alone or in collaboration with any other party.

# Why should I care about anonymizing data?

In our modern world data is a commodity. Big software companies have recognized this, and offer convenient, easy to use software and services “for free.” In reality however, nothing is free. Users of these services are paying with their personal data. As a result, we are now subject to a greater level of surveillance than in any point in history, and we hand over most of our data willingly. 

Modern societies, like the European Union now have stricter regulations on the handling and disclosure of personal data. One example is the “General Data Protection Regulation” (GDPR), which demands, that the disclosed data undergo an anonymization or pseudonymization process.  

# How to use the library?

## Import the library

The library supports __go modules__. In order to use it, add the following dependency in your `go.mod` file:

```
require (
	bitbucket.org/dargzero/k-anon v1.2.6
    // ...
)
```

The above line will import version `1.2.6`.

## Fetch the source

To obtain the source you can clone this repository, or run `go get bitbucket.org/dargzero/k-anon`.

## Usage:

In order to start anonymizing data, you will need to add the following steps to your code:

  1. Define a schema (friendly name and generalizer for each column):  
     * use a __nil__ generalizer for non-identifier (skipped) columns
     * see the Generalizer interface to implement a custom generalizer

  2. Supply the rows conforming to the above schema:  
     * number of columns should match
     * data type should be compatible with the assigned generalizer for the column

  3. Create an `Anonymizer` instance, and supply the `Table` and `K` parameters

  4. Call the `Anonymize()` function on the `Anonymizer` instance

The supplied Table will be anonymized in-place. Note, that items in partitions in the resulting Table are not ordered (always treat them as sets).

## Basic usage

The below example demonstrates the basic usage:

```go
func ExampleAnonymization() {

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
```

## Continuous mode

In continuous mode you can keep adding data to the table and call the anonymizer multiple times. It will reuse existing partitions generated in the previous iterations.

Example:

```go

func ExampleContinuousAnonymization() {

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

	err := anon.Anonymize()
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	// anonymize second chunk of data
	table.AddRow("Michelle", "employee", "female", 27, 1, 22400, 0.5, -0.2, "B", "cats are secretly extraterrestrials")
	table.AddRow("Perseus", "client", "male", 31, 2, 38600, 0.6, -0.15, "A+", "dogs are mischievous")
	table.AddRow("Odysseus", "client", "male", 39, 0, 39250, 0.5, -0.15, "A", "dogs are war bringers")
	table.AddRow("Helene", "employee", "female", 29, 2, 21900, 0.6, -0.2, "B+", "dogs are silky and furry")

	err = anon.Anonymize()
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	// anonymize third chunk of data
	table.AddRow("Donald", "client", "male", 26, 2, 24550, 0.9, 0.15, "A", "cats are secretly extraterrestrials")
	table.AddRow("Hillary", "client", "female", 70, 2, 38700, 0.5, -0.05, "B-", "dogs are loyal")
	table.AddRow("George", "client", "male", 65, 1, 39990, 0.5, 0.15, "B", "dogs are war bringers")
	table.AddRow("Victor", "employee", "male", 45, 0, 21000, 0.85, -0.3, "A-", "cats are silky and furry")

	err = anon.Anonymize()
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

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
```