package model

import (
	"bitbucket.org/dargzero/k-anon/generalization"
	"bitbucket.org/dargzero/k-anon/partition"
	"fmt"
	"strings"
)

// Table contains data organized in a rectangular shape. It has a fixed
// amount of columns, which is constrained by a schema.
// Each row must in the table must conform to the table schema.
type Table struct {
	schema *Schema
	rows   []*Row
}

func NewTable(schema *Schema) *Table {
	return &Table{schema: schema}
}

func (t *Table) AddRow(items ...interface{}) {
	var data []partition.Partition
	for i, item := range items {
		if i >= len(t.schema.Columns) {
			break
		}
		col := t.schema.Columns[i]
		var p partition.Partition
		if col.g != nil {
			p = t.schema.Columns[i].g.InitItem(item)
		} else {
			p = partition.NewItem(item)
		}
		data = append(data, p)
	}
	t.rows = append(t.rows, &Row{Data: data})
}

func (t *Table) GetSchema() *Schema {
	return t.schema
}

func (t *Table) GetRows() []*Row {
	return t.rows
}

func (t *Table) String() string {
	sb := &strings.Builder{}
	t.appendHeader(sb)
	for _, row := range t.rows {
		for colIdx := range t.schema.Columns {
			p := row.Data[colIdx]
			sb.WriteString(p.String())
			sb.WriteString("\t")
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (t *Table) appendHeader(sb *strings.Builder) {
	sb.WriteString("\n")
	for _, col := range t.schema.Columns {
		sb.WriteString(fmt.Sprintf("%s\t", col.name))
	}
	sb.WriteString("\n")
}

// Schema defines the number of columns, and their attributes in a table.
type Schema struct {
	Columns []*Column
}

// Column represents a column definition in a table schema.
// If the generalizer is set to nil, the column will be treated as non-identifier.
// Weight is a positive floating point number, which adjusts the cost of a column
// when picked for generalization (default is 1.0).
type Column struct {
	name   string
	g      generalization.Generalizer
	weight float64
}

func NewColumn(name string, g generalization.Generalizer) *Column {
	return NewWeightedColumn(name, g, 1.0)
}

func NewWeightedColumn(name string, g generalization.Generalizer, w float64) *Column {
	var adjustedWeight float64
	if w < 0 || w == 0 {
		adjustedWeight = 1.0
	} else {
		adjustedWeight = w
	}
	return &Column{name, g, adjustedWeight}
}

func (c *Column) GetName() string {
	return c.name
}

func (c *Column) GetGeneralizer() generalization.Generalizer {
	return c.g
}

func (c *Column) GetWeight() float64 {
	return c.weight
}

func (c *Column) IsIdentifier() bool {
	return c.g != nil
}

// Row represents a row of data in a table.
type Row struct {
	Data []partition.Partition
}
