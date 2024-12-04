package sqls

import (
	"reflect"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	SetDialect(PostgreSQL)

	t.Run("select all", func(t *testing.T) {
		sql, args := From("users").ToSql()

		if sql != `SELECT * FROM users` {
			t.Error("Invalid sql: " + sql)
		}
		if len(args) != 0 {
			t.Errorf("Invalid args: %v", args)
		}
	})

	t.Run("select simple", func(t *testing.T) {
		s := From("users")
		sql, args := s.Where("email", "test@email.com").Where("active", true).Limit(1).ToSql()

		if sql != `SELECT * FROM users WHERE email=$1 AND active=$2 LIMIT 1` {
			t.Error("Invalid sql: " + sql)
		}
		if !reflect.DeepEqual(args, []any{"test@email.com", true}) {
			t.Errorf("Invalid args: %v", args)
		}
	})

	t.Run("select complex", func(t *testing.T) {
		sql, args := From("users").
			Select("id", "name", "email").
			Join("roles", "users.id", "roles.user_id").
			Where("active", true).
			WhereIn("state", []any{"Washington", "Oregon"}).
			WhereExp("dob", ">", time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)).
			WhereNull("token").
			WhereNotNull("sso").
			WhereRaw("first=last").
			OrderBy("id DESC").
			Limit(20).Offset(100).
			ToSql()

		if sql != `SELECT id,name,email FROM users JOIN roles ON users.id=roles.user_id WHERE active=$1 AND state IN ($2,$3) AND dob>$4 AND token IS NULL AND sso IS NOT NULL AND first=last ORDER BY id DESC LIMIT 20 OFFSET 100` {
			t.Error("Invalid sql: " + sql)
		}
		if !reflect.DeepEqual(args, []any{true, "Washington", "Oregon", time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)}) {
			t.Errorf("Invalid args: %v", args)
		}
	})

	t.Run("select group by having", func(t *testing.T) {
		sql, args := From("users").
			Select("state", "COUNT(state) as count").
			GroupBy("state").
			Having("COUNT(state) > 5").
			ToSql()

		if sql != `SELECT state,COUNT(state) as count FROM users GROUP BY state HAVING COUNT(state) > 5` {
			t.Error("Invalid sql: " + sql)
		}
		if len(args) != 0 {
			t.Errorf("Invalid args: %v", args)
		}
	})

	t.Run("clear", func(t *testing.T) {
		query := From("users").
			Select("id", "name", "email").
			Join("roles", "users.id", "roles.user_id").
			Where("active", true).
			WhereIn("state", []any{"Washington", "Oregon"}).
			WhereExp("dob", ">", time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)).
			WhereNull("token").
			WhereNotNull("sso").
			WhereRaw("first=last").
			OrderBy("id DESC").
			Limit(20).Offset(100)

		query.ClearOrderBy()
		sql, args := query.ToSql()

		if sql != `SELECT id,name,email FROM users JOIN roles ON users.id=roles.user_id WHERE active=$1 AND state IN ($2,$3) AND dob>$4 AND token IS NULL AND sso IS NOT NULL AND first=last LIMIT 20 OFFSET 100` {
			t.Error("Invalid sql: " + sql)
		}
		if !reflect.DeepEqual(args, []any{true, "Washington", "Oregon", time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)}) {
			t.Errorf("Invalid args: %v", args)
		}

		query.ClearWhere()
		sql, args = query.ToSql()

		if sql != `SELECT id,name,email FROM users JOIN roles ON users.id=roles.user_id LIMIT 20 OFFSET 100` {
			t.Error("Invalid sql: " + sql)
		}
		if len(args) != 0 {
			t.Errorf("Invalid args: %v", args)
		}

		query.ClearJoin()
		sql, args = query.ToSql()
		if sql != `SELECT id,name,email FROM users LIMIT 20 OFFSET 100` {
			t.Error("Invalid sql: " + sql)
		}
		if len(args) != 0 {
			t.Errorf("Invalid args: %v", args)
		}

		query.Limit(0)
		query.Offset(0)
		sql, args = query.ToSql()
		if sql != `SELECT id,name,email FROM users` {
			t.Error("Invalid sql: " + sql)
		}
		if len(args) != 0 {
			t.Errorf("Invalid args: %v", args)
		}

		query.ClearSelect()
		sql, args = query.ToSql()
		if sql != `SELECT * FROM users` {
			t.Error("Invalid sql: " + sql)
		}
		if len(args) != 0 {
			t.Errorf("Invalid args: %v", args)
		}

		query.GroupBy("state").Having("COUNT(state) > 5").ClearGroupBy().ClearHaving()
		sql, args = query.ToSql()
		if sql != `SELECT * FROM users` {
			t.Error("Invalid sql: " + sql)
		}
		if len(args) != 0 {
			t.Errorf("Invalid args: %v", args)
		}
	})
}
