package model

type Field struct {
	Name  string
	Type  string
	Value string
}

type Fields []Field

type Table struct {
	Name string
	Fields
}

type Tables []Table
