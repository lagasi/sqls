package sqls

import (
	"reflect"
	"testing"
	"time"
)

func TestDelete(t *testing.T) {
	SetDialect(PostgreSQL)

	t.Run("Delete simple", func(t *testing.T) {
		sql, args := Delete("users").
			Where("active", false).
			ToSql()

		if sql != "DELETE FROM users WHERE active=$1" {
			t.Errorf("invalid sql: '%s'", sql)
		}
		if !reflect.DeepEqual(args, []any{false}) {
			t.Errorf("invalid args: '%v'", args)
		}
	})

	t.Run("Delete complex", func(t *testing.T) {
		sql, args := Delete("users").
			Where("active", false).
			WhereNotNull("password").
			WhereNull("email").
			WhereExp("created", "<", time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)).
			WhereIn("state", []any{"Washington", "Oregon"}).
			WhereRaw("first=last").
			ToSql()

		if sql != "DELETE FROM users WHERE active=$1 AND password IS NOT NULL AND email IS NULL AND created<$2 AND state IN ($3,$4) AND first=last" {
			t.Errorf("invalid sql: '%s'", sql)
		}
		if !reflect.DeepEqual(args, []any{false, time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC), "Washington", "Oregon"}) {
			t.Errorf("invalid args: '%v'", args)
		}
	})

}
