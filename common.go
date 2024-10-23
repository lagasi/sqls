package sqls

import (
	"strconv"
	"strings"
)

type whereClause struct {
	where []string
	args  []any
}

// KeyVal is a key-value pair
type KeyVal struct {
	key string
	val any
}

// Dialect is a SQL dialect
type Dialect struct {
	placeholder string
	paramCache  string
}

var (
	// PostgreSQL dialect
	PostgreSQL = Dialect{
		placeholder: "$",
	}
	// Dialect that uses @ placeholder
	DefaultDialect = Dialect{
		placeholder: "@",
	}
)

var curDialect Dialect

const MAX_PARAM_COUNT = 841 // 4096 characters

func init() {
	SetDialect(DefaultDialect)
}

// SetDialect sets the SQL dialect to use for queries.
func SetDialect(dialect Dialect) {
	curDialect = dialect
	if curDialect.paramCache == "" {
		curDialect.paramCache = generateParams(1, MAX_PARAM_COUNT)
	}
}

func generateParams(start int, count int) string {
	s := make([]string, 0, count)
	last := start + count

	for i := start; i < last; i++ {
		s = append(s, curDialect.placeholder+strconv.Itoa(i))
	}
	return strings.Join(s, ",")
}

func getIndex(n int) int {
	// faster manual calculation, supports up to 999
	if n <= 9 {
		return 3 * (n - 1)
	}
	if n <= 99 {
		return 3*9 + 4*(n-10)
	}
	return 3*9 + 4*90 + 5*(n-100)

	// slower dynamic calculation, supports any size
	// chars := 3 // @#, @##,  @###,
	// set := 9   // 1-9 10-99 100-999
	// max := 10
	// index := 0

	// for ; n > max; chars++ {
	// 	index += chars * set
	// 	set *= 10
	// 	max *= 10
	// }
	// index += chars * (n - max/10)
	// return index
}

func params(start int, count int) string {
	end := start + count

	if end >= MAX_PARAM_COUNT {
		return generateParams(start, count)
	}
	x := getIndex(start)
	y := getIndex(end) - 1

	return curDialect.paramCache[x:y]
}

func (s *whereClause) whereEquals(column string, value any) {
	s.args = append(s.args, value)
	p := params(len(s.args), 1)
	s.where = append(s.where, column+`=`+p)
}

func (s *whereClause) whereNull(column string) {
	s.where = append(s.where, column+` IS NULL`)
}

func (s *whereClause) whereNotNull(column string) {
	s.where = append(s.where, column+` IS NOT NULL`)
}

func (s *whereClause) whereExp(column string, ex string, value any) {
	s.args = append(s.args, value)
	p := params(len(s.args), 1)
	s.where = append(s.where, column+ex+p)
}

func (s *whereClause) whereIn(column string, values []any) {
	params := params(len(s.args)+1, len(values))

	s.where = append(s.where, column+` IN (`+params+`)`)
	s.args = append(s.args, values...)
}

func (s *whereClause) whereRaw(raw string) {
	s.where = append(s.where, raw)
}
