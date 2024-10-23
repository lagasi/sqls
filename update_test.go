package sqls

import (
	"reflect"
	"testing"
	"time"
)

func TestUpdate(t *testing.T) {
	SetDialect(PostgreSQL)

	t.Run("Update simple", func(t *testing.T) {
		sql, args := Update("users").
			Set("active", true).
			Where("id", "123").
			ToSql()

		if sql != "UPDATE users SET active=$1 WHERE id=$2" {
			t.Errorf("invalid sql: '%s'", sql)
		}
		if !reflect.DeepEqual(args, []any{true, "123"}) {
			t.Errorf("invalid args: '%v'", args)
		}
	})

	t.Run("Update simple 2", func(t *testing.T) {
		sql, args := Update("users").
			SetValues([]KeyVal{
				{"active", true},
				{"email", "fake@email.com"},
			}).
			Where("id", "123").
			ToSql()

		if sql != "UPDATE users SET active=$1,email=$2 WHERE id=$3" {
			t.Errorf("invalid sql: '%s'", sql)
		}
		if !reflect.DeepEqual(args, []any{true, "fake@email.com", "123"}) {
			t.Errorf("invalid args: '%v'", args)
		}
	})

	t.Run("Update complex", func(t *testing.T) {
		sql, args := Update("users").
			Set("active", false).
			WhereNotNull("password").
			WhereNull("email").
			WhereExp("created", "<", time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)).
			WhereIn("state", []any{"Washington", "Oregon"}).
			WhereRaw("first=last").
			ToSql()

		if sql != "UPDATE users SET active=$1 WHERE password IS NOT NULL AND email IS NULL AND created<$2 AND state IN ($3,$4) AND first=last" {
			t.Errorf("invalid sql: '%s'", sql)
		}
		if !reflect.DeepEqual(args, []any{false, time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC), "Washington", "Oregon"}) {
			t.Errorf("invalid args: '%v'", args)
		}
	})

}
