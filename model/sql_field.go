package model

//field
type Field struct {
	Name string
	Type string
}

// mutiple filed
type Fields []Field

//table
type Table struct {
	Name string
	Fields
}

// mutiple tables
type Tables []Table
