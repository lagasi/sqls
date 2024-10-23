package sqls

import "strings"

// DeleteStmt represents a SQL DELETE statement.
type DeleteStmt struct {
	table string
	whereClause
}

// Delete creates a new DELETE statement.
func Delete(table string) *DeleteStmt {
	return &DeleteStmt{
		table: table,
	}
}

// Where adds a WHERE column=value clause to the DELETE statement.
func (s *DeleteStmt) Where(column string, value any) *DeleteStmt {
	s.whereEquals(column, value)
	return s
}

// WhereNull adds a WHERE column IS NULL clause to the DELETE statement.
func (s *DeleteStmt) WhereNull(column string) *DeleteStmt {
	s.whereNull(column)
	return s
}

// WhereNotNull adds a WHERE column IS NOT NULL clause to the DELETE statement.
func (s *DeleteStmt) WhereNotNull(column string) *DeleteStmt {
	s.whereNotNull(column)
	return s
}

// WhereExp adds a WHERE column expression value clause to the DELETE statement.
func (s *DeleteStmt) WhereExp(column string, ex string, value any) *DeleteStmt {
	s.whereExp(column, ex, value)
	return s
}

// WhereIn adds a WHERE column IN (values) clause to the DELETE statement.
func (s *DeleteStmt) WhereIn(column string, in []any) *DeleteStmt {
	s.whereIn(column, in)
	return s
}

// WhereRaw adds a raw WHERE clause to the DELETE statement.
func (s *DeleteStmt) WhereRaw(raw string) *DeleteStmt {
	s.whereRaw(raw)
	return s
}

// ToSql generates the SQL DELETE statement and returns it along with any arguments.
func (s *DeleteStmt) ToSql() (string, []any) {
	query := "DELETE FROM " + s.table

	if s.where != nil {
		query += " WHERE " + strings.Join(s.where, " AND ")
	}

	return query, s.args
}
