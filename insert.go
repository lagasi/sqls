package sqls

import (
	"strings"
)

// InsertStmt represents an SQL INSERT statement.
type InsertStmt struct {
	table     string
	columns   []string
	args      []any
	returning string
	conflict  string
}

// Insert creates a new INSERT statement.
func Insert(table string) *InsertStmt {
	return &InsertStmt{
		table: table,
	}
}

// Set adds a column and its corresponding value to the INSERT statement.
func (s *InsertStmt) Set(column string, value any) *InsertStmt {
	s.columns = append(s.columns, column)
	s.args = append(s.args, value)
	return s
}

// SetValues adds multiple columns and their corresponding values to the INSERT statement.
func (s *InsertStmt) SetValues(values []KeyVal) *InsertStmt {
	for _, kv := range values {
		s.columns = append(s.columns, kv.key)
		s.args = append(s.args, kv.val)
	}
	return s
}

// Returning specifies the columns to be returned after the INSERT statement is executed.
func (s *InsertStmt) Returning(columns ...string) *InsertStmt {
	s.returning = " RETURNING " + strings.Join(columns, ",")
	return s
}

// OnConflict specifies the conflict resolution strategy for the INSERT statement.
func (s *InsertStmt) OnConflict(expression string) *InsertStmt {
	s.conflict = " ON CONFLICT " + expression
	return s
}

// ToSql generates the SQL query string and the corresponding arguments for the INSERT statement.
func (s *InsertStmt) ToSql() (string, []any) {
	sql := "INSERT INTO " + s.table + " (" + strings.Join(s.columns, ",") + ") VALUES (" + params(1, len(s.columns)) + ")" + s.conflict + s.returning
	return sql, s.args
}
