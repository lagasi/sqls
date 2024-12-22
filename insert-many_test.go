package sqls

import (
	"reflect"
	"testing"
)

func TestInsertMany(t *testing.T) {
	SetDialect(DefaultDialect)

	t.Run("Insert many simple", func(t *testing.T) {
		sql, args := InsertMany("users").Columns("name", "age").
			Values("John", 30).
			Values("Jane", 25).
			Values("Mary", 35).
			ToSql()

		if sql != "INSERT INTO users(name,age) VALUES (@1,@2),(@3,@4),(@5,@6)" {
			t.Errorf("invalid sql: '%s'", sql)
		}
		if !reflect.DeepEqual(args, []any{"John", 30, "Jane", 25, "Mary", 35}) {
			t.Errorf("invalid args: '%v'", args)
		}
	})

	t.Run("Insert many clear", func(t *testing.T) {
		sql, args := InsertMany("users").Columns("name", "age").
			Values("John", 30).
			Clear().
			Values("Jane", 25).
			Values("Mary", 35).
			ToSql()

		if sql != "INSERT INTO users(name,age) VALUES (@1,@2),(@3,@4)" {
			t.Errorf("invalid sql: '%s'", sql)
		}
		if !reflect.DeepEqual(args, []any{"Jane", 25, "Mary", 35}) {
			t.Errorf("invalid args: '%v'", args)
		}
	})

	t.Run("Insert many complex", func(t *testing.T) {
		sql, args := InsertMany("users").Columns("name", "age").
			Values("John", 30).
			Values("Jane", 25).
			Values("Mary", 35).
			OnConflict("DO NOTHING").
			Returning("id", "name").
			ToSql()

		if sql != "INSERT INTO users(name,age) VALUES (@1,@2),(@3,@4),(@5,@6) ON CONFLICT DO NOTHING RETURNING id,name" {
			t.Errorf("invalid sql: '%s'", sql)
		}
		if !reflect.DeepEqual(args, []any{"John", 30, "Jane", 25, "Mary", 35}) {
			t.Errorf("invalid args: '%v'", args)
		}
	})
}
