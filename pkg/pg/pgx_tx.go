package pg

import (
	"context"
	"errors"
	"example-api/pkg/logger"
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type pgxPool struct {
	pool *pgxpool.Pool
}

func NewWithPGXPool(pool *pgxpool.Pool) PG {
	return &pgxPool{
		pool: pool,
	}
}

// Exec исполняет query.
func (p *pgxPool) Exec(ctx context.Context, sqlizer Sqlizer) (pgconn.CommandTag, error) {
	return execFn(ctx, p.pool, sqlizer)
}

// Select может сканировать сразу несколько рядов в slice.
// Если рядов нет, возвращает nil.
func (p *pgxPool) Select(ctx context.Context, dst interface{}, sqlizer Sqlizer) error {
	return selectFn(ctx, p.pool, dst, sqlizer)
}

// Get сканирует один ряд.
// Если рядов нет, возвращает ошибку pgx.ErrNoRows.
func (p *pgxPool) Get(ctx context.Context, dst interface{}, sqlizer Sqlizer) error {
	return getFn(ctx, p.pool, dst, sqlizer)
}

func (p *pgxPool) QueryBatch(ctx context.Context, dst interface{}, sqlizers []Sqlizer) error {
	return queryBatchFn(ctx, p.pool, dst, sqlizers)
}

func (p *pgxPool) ExecBatch(ctx context.Context, sqlizers []Sqlizer) error {
	return execBatchFn(ctx, p.pool, sqlizers)
}

// pgxTx обертка над транзакцией.
type pgxTx struct {
	pgxTx pgx.Tx
}

// Exec исполняет query.
func (t *pgxTx) Exec(ctx context.Context, sqlizer Sqlizer) (pgconn.CommandTag, error) {
	return execFn(ctx, t.pgxTx, sqlizer)
}

// Select может сканировать сразу несколько рядов в slice.
// Если рядов нет, возвращает nil.
func (t *pgxTx) Select(ctx context.Context, dst interface{}, sqlizer Sqlizer) error {
	return selectFn(ctx, t.pgxTx, dst, sqlizer)
}

// Get сканирует один ряд.
// Если рядов нет, возвращает ошибку pgx.ErrNoRows.
func (t *pgxTx) Get(ctx context.Context, dst interface{}, sqlizer Sqlizer) error {
	return getFn(ctx, t.pgxTx, dst, sqlizer)
}

// Commit завершает транзакцию.
func (t *pgxTx) Commit(ctx context.Context) error {
	return t.pgxTx.Commit(ctx)
}

// Rollback откатывает транзакцию.
func (t *pgxTx) Rollback(ctx context.Context) {
	err := t.pgxTx.Rollback(ctx)
	if errors.Is(err, pgx.ErrTxClosed) {
		return
	}

	if err != nil {
		logger.Errorf(ctx, "rollback transaction: %s", err.Error())
	}
}

func (t *pgxTx) QueryBatch(ctx context.Context, dst interface{}, sqlizers []Sqlizer) error {
	return queryBatchFn(ctx, t.pgxTx, dst, sqlizers)
}

func (t *pgxTx) ExecBatch(ctx context.Context, sqlizers []Sqlizer) error {
	return execBatchFn(ctx, t.pgxTx, sqlizers)
}

func execFn(ctx context.Context, e execer, sqlizer Sqlizer) (pgconn.CommandTag, error) {
	query, args, err := sqlizer.ToSql()
	if err != nil {
		return pgconn.CommandTag{}, fmt.Errorf("toSql: %w", err)
	}

	return e.Exec(ctx, query, args...)
}

// BeginTx транзакцию.
func (p *pgxPool) BeginTx(ctx context.Context, txOptions *pgx.TxOptions) (Tx, error) {
	var txOpts pgx.TxOptions
	if txOptions != nil {
		txOpts = *txOptions
	}

	tx, err := p.pool.BeginTx(ctx, txOpts)
	if err != nil {
		return nil, fmt.Errorf("can't start transaction: %w", err)
	}

	return &pgxTx{pgxTx: tx}, nil
}

func selectFn(ctx context.Context, q pgxscan.Querier, dst interface{}, sqlizer Sqlizer) error {
	query, args, err := sqlizer.ToSql()
	if err != nil {
		return fmt.Errorf("toSql: %w", err)
	}

	return pgxscan.Select(ctx, q, dst, query, args...)
}

func getFn(ctx context.Context, q pgxscan.Querier, dst interface{}, sqlizer Sqlizer) error {
	query, args, err := sqlizer.ToSql()
	if err != nil {
		return fmt.Errorf("toSql: %w", err)
	}

	return pgxscan.Get(ctx, q, dst, query, args...)
}

func sendBatchFn(ctx context.Context, sb sendBatcher, sqlizers []Sqlizer) (pgx.BatchResults, error) {
	batch := &pgx.Batch{}
	for _, s := range sqlizers {
		query, args, err := s.ToSql()
		if err != nil {
			return nil, fmt.Errorf("toSql: %w", err)
		}

		batch.Queue(query, args...)
	}

	return sb.SendBatch(ctx, batch), nil
}

func queryBatchFn(ctx context.Context, sb sendBatcher, dst interface{}, sqlizers []Sqlizer) error {
	batchResults, err := sendBatchFn(ctx, sb, sqlizers)
	if err != nil {
		return fmt.Errorf("send batch: %w", err)
	}
	defer func() {
		if err := batchResults.Close(); err != nil {
			logger.Errorf(ctx, "close batch results: %s", err.Error())
		}
	}()

	rows, err := batchResults.Query()
	if err != nil {
		return fmt.Errorf("batch query: %w", err)
	}

	if err := pgxscan.ScanAll(dst, rows); err != nil {
		return fmt.Errorf("scan rows: %w", err)
	}

	return nil
}

func execBatchFn(ctx context.Context, sb sendBatcher, sqlizers []Sqlizer) error {
	batchResults, err := sendBatchFn(ctx, sb, sqlizers)
	if err != nil {
		return fmt.Errorf("send batch: %w", err)
	}
	defer func() {
		if err := batchResults.Close(); err != nil {
			logger.Errorf(ctx, "close batch results: %s", err.Error())
		}
	}()

	_, err = batchResults.Exec()
	if err != nil {
		return fmt.Errorf("batch exec: %w", err)
	}

	return nil
}
