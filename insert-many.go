package sqls

import "strings"

type InsertManyStmt struct {
	table     string
	columns   []string
	args      []any
	count     int
	returning string
	conflict  string
}

// InsertMany creates a new INSERT statement for multiple rows.
func InsertMany(table string) *InsertManyStmt {
	return &InsertManyStmt{
		table: table,
	}
}

// Columns specifies the columns to be inserted in the INSERT statement.
func (s *InsertManyStmt) Columns(columns ...string) *InsertManyStmt {
	s.columns = columns
	return s
}

// Values specifies the values to be inserted in the INSERT statement.
func (s *InsertManyStmt) Values(values ...any) *InsertManyStmt {
	s.args = append(s.args, values...)
	s.count++
	return s
}

// Clear resets the values to be inserted in the INSERT statement.
func (s *InsertManyStmt) Clear() *InsertManyStmt {
	s.args = []any{}
	s.count = 0
	return s
}

// OnConflict specifies the conflict resolution strategy in the INSERT statement.
func (s *InsertManyStmt) OnConflict(expression string) *InsertManyStmt {
	s.conflict = " ON CONFLICT " + expression
	return s
}

// Returning specifies the columns to be returned in the INSERT statement.
func (s *InsertManyStmt) Returning(columns ...string) *InsertManyStmt {
	s.returning = " RETURNING " + strings.Join(columns, ",")
	return s
}

// ToSql generates the SQL and returns the parameters.
func (s *InsertManyStmt) ToSql() (string, []any) {
	var values []string
	length := len(s.columns)
	end := s.count * length

	for i := 1; i <= end; i += length {
		values = append(values, params(i, length))
	}

	query := "INSERT INTO " + s.table + "(" + strings.Join(s.columns, ",") + ") VALUES (" + strings.Join(values, "),(") + ")" + s.conflict + s.returning
	return query, s.args
}
