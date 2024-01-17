package pg

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// PG содержит основные операции для работы с базой данных.
type PG interface {
	Queryable
	BeginTx(ctx context.Context, txOptions *pgx.TxOptions) (Tx, error)
}

// Tx - транзакция
type Tx interface {
	Queryable
	Commit(ctx context.Context) error
	Rollback(ctx context.Context)
}

// Queryable содержит основные операции для query-инга db.
type Queryable interface {
	Exec(ctx context.Context, sqlizer Sqlizer) (pgconn.CommandTag, error)
	Get(ctx context.Context, dst interface{}, sqlizer Sqlizer) error
	Select(ctx context.Context, dst interface{}, sqlizer Sqlizer) error
	QueryBatch(ctx context.Context, dst interface{}, sqlizers []Sqlizer) error
	ExecBatch(ctx context.Context, sqlizers []Sqlizer) error
}

// Sqlizer ...
type Sqlizer interface {
	ToSql() (sql string, args []interface{}, err error)
}

type execer interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}

type sendBatcher interface {
	SendBatch(context.Context, *pgx.Batch) pgx.BatchResults
}
