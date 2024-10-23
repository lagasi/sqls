package sqls

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	SetDialect(DefaultDialect)

	t.Run("Insert simple", func(t *testing.T) {
		sql, args := Insert("users").
			Set("email", "fake@email.com").
			Set("password", "l33tP@$$w0rd").
			ToSql()

		if sql != "INSERT INTO users (email,password) VALUES (@1,@2)" {
			t.Errorf("invalid sql: '%s'", sql)
		}
		if !reflect.DeepEqual(args, []any{"fake@email.com", "l33tP@$$w0rd"}) {
			t.Errorf("invalid args: '%v'", args)
		}
	})

	t.Run("Insert simple 2", func(t *testing.T) {
		sql, args := Insert("users").
			SetValues([]KeyVal{
				{"email", "fake@email.com"},
				{"password", "l33tP@$$w0rd"},
				{"active", true},
			}).
			ToSql()

		if sql != "INSERT INTO users (email,password,active) VALUES (@1,@2,@3)" {
			t.Errorf("invalid sql: '%s'", sql)
		}
		if !reflect.DeepEqual(args, []any{"fake@email.com", "l33tP@$$w0rd", true}) {
			t.Errorf("invalid args: '%v'", args)
		}
	})

	t.Run("Insert complex", func(t *testing.T) {
		sql, args := Insert("users").
			Set("email", "fake@email.com").
			Set("password", "l33tP@$$w0rd").
			Set("active", true).
			Returning("id").
			OnConflict("IGNORE").
			ToSql()

		if sql != "INSERT INTO users (email,password,active) VALUES (@1,@2,@3) ON CONFLICT IGNORE RETURNING id" {
			t.Errorf("invalid sql: '%s'", sql)
		}
		if !reflect.DeepEqual(args, []any{"fake@email.com", "l33tP@$$w0rd", true}) {
			t.Errorf("invalid args: '%v'", args)
		}
	})

}
