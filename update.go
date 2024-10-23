package sqls

import (
	"strings"
)

// UpdateStmt represents an SQL UPDATE statement.
type UpdateStmt struct {
	table   string
	columns []string
	whereClause
}

// Update creates a new UPDATE statement.
func Update(table string) *UpdateStmt {
	return &UpdateStmt{
		table: table,
	}
}

// Set adds a column and its corresponding value to the UPDATE statement.
func (s *UpdateStmt) Set(column string, value any) *UpdateStmt {
	s.args = append(s.args, value)
	p := params(len(s.args), 1)
	s.columns = append(s.columns, column+"="+p)
	return s
}

// SetValues adds multiple columns and their corresponding values to the UPDATE statement.
func (s *UpdateStmt) SetValues(values []KeyVal) *UpdateStmt {
	for _, kv := range values {
		s.args = append(s.args, kv.val)
		p := params(len(s.args), 1)
		s.columns = append(s.columns, kv.key+"="+p)
	}
	return s
}

// Where adds a WHERE clause to the UPDATE statement.
func (s *UpdateStmt) Where(column string, value any) *UpdateStmt {
	s.whereEquals(column, value)
	return s
}

// WhereNull adds a WHERE clause checking for NULL to the UPDATE statement.
func (s *UpdateStmt) WhereNull(column string) *UpdateStmt {
	s.whereNull(column)
	return s
}

// WhereNotNull adds a WHERE clause checking for NOT NULL to the UPDATE statement.
func (s *UpdateStmt) WhereNotNull(column string) *UpdateStmt {
	s.whereNotNull(column)
	return s
}

// WhereExp adds a WHERE clause with a custom expression to the UPDATE statement.
func (s *UpdateStmt) WhereExp(column string, eq string, value any) *UpdateStmt {
	s.whereExp(column, eq, value)
	return s
}

// WhereIn adds a WHERE IN clause to the UPDATE statement.
func (s *UpdateStmt) WhereIn(column string, in []any) *UpdateStmt {
	s.whereIn(column, in)
	return s
}

// WhereRaw adds a raw WHERE clause to the UPDATE statement.
func (s *UpdateStmt) WhereRaw(raw string) *UpdateStmt {
	s.whereRaw(raw)
	return s
}

// ToSql generates the SQL query string and the corresponding arguments for the UPDATE statement.
func (s *UpdateStmt) ToSql() (string, []any) {
	query := "UPDATE " + s.table + " SET " + strings.Join(s.columns, ",")

	if s.where != nil {
		query += " WHERE " + strings.Join(s.where, " AND ")
	}

	return query, s.args
}
