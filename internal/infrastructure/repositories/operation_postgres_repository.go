package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/qrave1/BlazingMoneyControlBot/internal/domain"
)

const operationLimit = 5

type OperationPostgresRepository struct {
	db *sql.DB
}

func NewOperationPostgresRepository(db *sql.DB) *OperationPostgresRepository {
	return &OperationPostgresRepository{db: db}
}

func (o *OperationPostgresRepository) Create(ctx context.Context, operation domain.Operation) error {
	fail := func(err error) error {
		return fmt.Errorf("failed to create operation: %w", err)
	}

	tx, err := o.db.BeginTx(ctx, nil)
	if err != nil {
		return fail(err)
	}

	defer func(tx *sql.Tx) {
		_ = tx.Rollback()
	}(tx)

	result, err := tx.ExecContext(
		ctx,
		"INSERT INTO operations (wallet_id, type, amount, reason) VALUES ($1, $2, $3, $4)",
		operation.WalletId,
		operation.Type,
		operation.Amount,
		operation.Reason,
	)
	if err != nil {
		return fail(err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return fail(err)
	}

	if affected == 0 {
		return fail(err)
	}

	if err = tx.Commit(); err != nil {
		return fail(err)
	}

	return nil
}

func (o *OperationPostgresRepository) ReadByBatch(ctx context.Context, walletId, page int) ([]domain.Operation, error) {
	fail := func(err error) ([]domain.Operation, error) {
		return nil, fmt.Errorf("failed to read operation: %w", err)
	}

	tx, err := o.db.BeginTx(ctx, nil)
	if err != nil {
		return fail(err)
	}

	defer func(tx *sql.Tx) {
		_ = tx.Rollback()
	}(tx)

	offset := operationLimit * (page - 1)

	rows, err := tx.QueryContext(
		ctx,
		"SELECT wallet_id, type, amount, reason FROM operations WHERE wallet_id = $1 ORDER BY created_at LIMIT $2 OFFSET $3",
		walletId,
		operationLimit,
		offset,
	)
	if err != nil {
		return fail(err)
	}

	var operations []domain.Operation

	for rows.Next() {
		op := domain.Operation{}
		err = rows.Scan(&op.WalletId, &op.Type, &op.Amount, &op.Reason)
		if err != nil {
			return nil, err
		}
		operations = append(operations, op)
	}

	if err = tx.Commit(); err != nil {
		return fail(err)
	}

	return operations, nil
}
