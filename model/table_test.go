package model

import (
	"bitbucket.org/dargzero/k-anon/generalization"
	"bitbucket.org/dargzero/k-anon/partition"
	"bitbucket.org/dargzero/k-anon/testutil"
	"testing"
)

func TestNewTable(t *testing.T) {
	table := NewTable(&Schema{})
	testutil.AssertNotNil(table, t)
}

func TestTable_AddRow(t *testing.T) {

	t.Run("add row with identifier", func(t *testing.T) {
		table := NewTable(&Schema{
			Columns: []*Column{
				{"Identifier", &generalization.Suppressor{}},
			},
		})
		table.AddRow("test string")
		testutil.AssertEquals(1, len(table.rows), t)
	})

	t.Run("add row with non-identifier", func(t *testing.T) {
		table := NewTable(&Schema{
			Columns: []*Column{{Name: "Non-Identifier"}},
		})
		table.AddRow("test string")
		testutil.AssertEquals(1, len(table.rows), t)
	})

	t.Run("add invalid sized row", func(t *testing.T) {
		table := NewTable(&Schema{
			Columns: []*Column{{Name: "Col1"}},
		})
		table.AddRow("test string", "extra column")
		testutil.AssertEquals(1, len(table.rows), t)

	})

}

func TestTable_GetRows(t *testing.T) {
	table := NewTable(&Schema{
		Columns: []*Column{{Name: "Col1"}},
	})
	table.AddRow("test string")
	row := table.GetRows()[0]
	actual := row.Data[0]
	expected := partition.NewItem("test string")
	if !expected.Equals(actual) {
		t.Errorf("expected %v got %v", expected, actual)
	}
}

func TestTable_GetSchema(t *testing.T) {
	schema := &Schema{}
	table := NewTable(schema)
	testutil.AssertEquals(schema, table.GetSchema(), t)
}

func TestColumn_IsIdentifier(t *testing.T) {

	t.Run("identifier column", func(t *testing.T) {
		col := &Column{Name: "identifier", Generalizer: &generalization.HierarchyGeneralizer{}}
		if !col.IsIdentifier() {
			t.Errorf("expectd identifier column")
		}
	})

	t.Run("non-identifier column", func(t *testing.T) {
		col := &Column{Name: "non-identifier"}
		if col.IsIdentifier() {
			t.Errorf("expectd non-identifier column")
		}
	})

}

func TestTable_String(t *testing.T) {
	table := NewTable(&Schema{
		Columns: []*Column{
			{Name: "Col1"},
			{Name: "Col2"},
		},
	})
	table.AddRow("d1", "d2")
	actual := table.String()
	expected := "\nCol1\tCol2\t\nd1\td2\t\n"
	if expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}
