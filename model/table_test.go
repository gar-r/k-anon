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
				NewColumn("Identifier", &generalization.Suppressor{}),
			},
		})
		table.AddRow("test string")
		testutil.AssertEquals(1, len(table.rows), t)
	})

	t.Run("add row with non-identifier", func(t *testing.T) {
		table := NewTable(&Schema{
			Columns: []*Column{{name: "Non-Identifier"}},
		})
		table.AddRow("test string")
		testutil.AssertEquals(1, len(table.rows), t)
	})

	t.Run("add invalid sized row", func(t *testing.T) {
		table := NewTable(&Schema{
			Columns: []*Column{{name: "Col1"}},
		})
		table.AddRow("test string", "extra column")
		testutil.AssertEquals(1, len(table.rows), t)

	})

}

func TestTable_GetRows(t *testing.T) {
	table := NewTable(&Schema{
		Columns: []*Column{{name: "Col1"}},
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
		col := &Column{name: "identifier", g: &generalization.HierarchyGeneralizer{}}
		if !col.IsIdentifier() {
			t.Errorf("expectd identifier column")
		}
	})

	t.Run("non-identifier column", func(t *testing.T) {
		col := &Column{name: "non-identifier"}
		if col.IsIdentifier() {
			t.Errorf("expectd non-identifier column")
		}
	})

}

func TestTable_String(t *testing.T) {
	table := NewTable(&Schema{
		Columns: []*Column{
			{name: "Col1"},
			{name: "Col2"},
		},
	})
	table.AddRow("d1", "d2")
	actual := table.String()
	expected := "\nCol1\tCol2\t\nd1\td2\t\n"
	if expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestColumn_GetName(t *testing.T) {
	c := NewColumn("test", nil)
	if c.GetName() != "test" {
		t.Errorf("expected %v, got %v", "test", c.GetName())
	}
}

func TestColumn_GetGeneralizer(t *testing.T) {
	g := &generalization.Suppressor{}
	c := NewColumn("test", g)
	if c.GetGeneralizer() != g {
		t.Errorf("expected %v, got %v", g, c.GetGeneralizer())
	}
}

func TestColumn_GetWeight(t *testing.T) {
	c := NewWeightedColumn("test", nil, 3.4)
	if c.GetWeight() != 3.4 {
		t.Errorf("expected %v, got %v", 3.4, c.GetWeight())
	}
}

func TestNewColumn(t *testing.T) {

	t.Run("default weight", func(t *testing.T) {
		c := NewColumn("test", nil)
		actual := c.GetWeight()
		expected := 1.0
		if expected != actual {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

}

func TestNewWeightedColumn(t *testing.T) {

	t.Run("negative weight", func(t *testing.T) {
		c := NewWeightedColumn("test", nil, -1.5)
		actual := c.GetWeight()
		expected := 1.0
		if expected != actual {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

	t.Run("zero weight", func(t *testing.T) {
		c := NewWeightedColumn("test", nil, 0)
		actual := c.GetWeight()
		expected := 1.0
		if expected != actual {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

	t.Run("valid weight", func(t *testing.T) {
		c := NewWeightedColumn("test", nil, 0.4)
		actual := c.GetWeight()
		expected := 0.4
		if expected != actual {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

}
