package sqls

import (
	"strconv"
	"strings"
)

// SelectStmt represents an SQL SELECT statement.
type SelectStmt struct {
	limit   int
	offset  int
	table   string
	columns []string
	whereClause
	joins   []string
	groupBy string
	having  string
	orderBy string
}

// From creates a new SELECT statement.
func From(table string) *SelectStmt {
	return &SelectStmt{
		table: table,
	}
}

// Select specifies the columns to be selected in the SELECT statement.
func (s *SelectStmt) Select(columns ...string) *SelectStmt {
	s.columns = columns
	return s
}

// Where adds a WHERE column=value clause to the SELECT statement.
func (s *SelectStmt) Where(column string, value any) *SelectStmt {
	s.whereEquals(column, value)
	return s
}

// WhereNull adds a WHERE column IS NULL clause to the SELECT statement.
func (s *SelectStmt) WhereNull(column string) *SelectStmt {
	s.whereNull(column)
	return s
}

// WhereNotNull adds a WHERE column IS NOT NULL clause to the SELECT statement.
func (s *SelectStmt) WhereNotNull(column string) *SelectStmt {
	s.whereNotNull(column)
	return s
}

// WhereExp adds a WHERE column expression value clause to the SELECT statement.
func (s *SelectStmt) WhereExp(column string, ex string, value any) *SelectStmt {
	s.whereExp(column, ex, value)
	return s
}

// WhereIn adds a WHERE column IN (values) clause to the SELECT statement.
func (s *SelectStmt) WhereIn(column string, values []any) *SelectStmt {
	s.whereIn(column, values)
	return s
}

// WhereRaw adds a raw WHERE clause to the SELECT statement.
func (s *SelectStmt) WhereRaw(raw string) *SelectStmt {
	s.whereRaw(raw)
	return s
}

// OrderBy adds an ORDER BY clause to the SELECT statement.
func (s *SelectStmt) OrderBy(columns ...string) *SelectStmt {
	s.orderBy = " ORDER BY " + strings.Join(columns, ",")
	return s
}

// Limit adds a LIMIT clause to the SELECT statement.
func (s *SelectStmt) Limit(limit int) *SelectStmt {
	s.limit = limit
	return s
}

// Offset adds an OFFSET clause to the SELECT statement.
func (s *SelectStmt) Offset(offset int) *SelectStmt {
	s.offset = offset
	return s
}

// Join adds a JOIN clause to the SELECT statement.
func (s *SelectStmt) Join(table string, on1 string, on2 string) *SelectStmt {
	s.joins = append(s.joins, "JOIN "+table+" ON "+on1+"="+on2)
	return s
}

// GroupBy adds a GROUP BY clause to the SELECT statement.
func (s *SelectStmt) GroupBy(columns ...string) *SelectStmt {
	s.groupBy = " GROUP BY " + strings.Join(columns, ",")
	return s
}

// Having adds a HAVING clause to the SELECT statement.
func (s *SelectStmt) Having(conditions ...string) *SelectStmt {
	s.having = " HAVING " + strings.Join(conditions, ",")
	return s
}

// ClearSelect clears the selected columns in the SELECT statement.
func (s *SelectStmt) ClearSelect() *SelectStmt {
	s.columns = nil
	return s
}

// ClearWhere clears the WHERE clause and its arguments in the SELECT statement.
func (s *SelectStmt) ClearWhere() *SelectStmt {
	s.where = nil
	s.args = nil
	return s
}

// ClearJoin clears the JOIN clauses in the SELECT statement.
func (s *SelectStmt) ClearJoin() *SelectStmt {
	s.joins = nil
	return s
}

// ClearGroupBy clears the GROUP BY clause in the SELECT statement.
func (s *SelectStmt) ClearGroupBy() *SelectStmt {
	s.groupBy = ""
	return s
}

// ClearHaving clears the HAVING clause in the SELECT statement.
func (s *SelectStmt) ClearHaving() *SelectStmt {
	s.having = ""
	return s
}

// ClearOrderBy clears the ORDER BY clause in the SELECT statement.
func (s *SelectStmt) ClearOrderBy() *SelectStmt {
	s.orderBy = ""
	return s
}

// ToSql generates the SQL query string and the corresponding arguments for the SELECT statement.
func (s *SelectStmt) ToSql() (string, []any) {
	var query string

	if s.columns == nil {
		query = "SELECT * FROM " + s.table
	} else {
		query = "SELECT " + strings.Join(s.columns, `,`) + " FROM " + s.table
	}

	if s.joins != nil {
		query += " " + strings.Join(s.joins, " ")
	}

	if s.where != nil {
		query += " WHERE " + strings.Join(s.where, " AND ")
	}

	query += s.orderBy + s.groupBy + s.having

	if s.limit > 0 {
		query += " LIMIT " + strconv.Itoa(s.limit)
	}

	if s.offset > 0 {
		query += " OFFSET " + strconv.Itoa(s.offset)
	}

	return query, s.args
}
