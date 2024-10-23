# sqls

A simple SQL query builder for Go.

`sqls`

- Helps you build SQL queries at runtime.
- It keeps simplicity and performance in mind.

## Limitations

- It does not guarantee the generate SQL is correct as there are no checks for table names, column names, etc.
- It does not escape keywords for table and column names. You need to do that yourself.
- It does not support complex queries.
- It does not have a `WhereOr` function although you can use `WhereRaw` without parameters. It is recommend to use `UNION ALL` as there are usually performance issues with `OR` clauses.

## Usage

### Select Dialect

The default dialect works with most SQL such as MySQL, SQL Server, and SQLite. PostgreSQL is always supported. Currently, the only difference is placeholders.

```go
sqls.SetDialect(sqls.DefaultDialect)
sqls.SetDialect(sqls.PostgreSQL)

// custom dialect
sqls.SetDialect(sqls.Dialect{placeholder: "#"})
```

### SELECT

```go
sql, args := sqls.From("users as u").
  Select("u.id", "u.name", "u.email").
  Join("roles as r", "u.id", "r.user_id").
  Where("active", true).
  WhereIn("r.role", []any{"admin", "editor"}).
  WhereExp("dob", ">", time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)).
  OrderBy("u.id DESC").Limit(20).Offset(100).
  ToSql()

db, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
if err != nil {
  fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
  os.Exit(1)
}

rows, _ := db.Query(context.Background(), sql, args...)
users, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[User])
```

### INSERT

```go
sql, args := Insert("users").
  Set("email", "fake@email.com").
  Set("password", "l33tP@$$w0rd").
  ToSql()
// or
sql, args := Insert("users").
  SetValues([]sqls.KeyVal{
    {"email", "fake@email.com"},
    {"password", "l33tP@$$w0rd"},
    {"active", true},
  }).ToSql()
```

### UPDATE

```go
sql, args := Update("users").
  Set("active", true).
  Where("id", "123").
  ToSql()
```

### DELETE

```go
sql, args := Delete("users").
  Where("id", "123").
  ToSql()
```