package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/qrave1/BlazingMoneyControlBot/internal/domain"
	"log/slog"
)

type WalletPostgresRepository struct {
	db  *sql.DB
	log *slog.Logger
}

func NewWalletPostgresRepository(db *sql.DB, log *slog.Logger) *WalletPostgresRepository {
	return &WalletPostgresRepository{db: db, log: log}
}

func (w *WalletPostgresRepository) Create(ctx context.Context, id int, userName string, balance int) error {
	fail := func(err error) error {
		return fmt.Errorf("failed to create wallet: %w", err)
	}

	tx, err := w.db.BeginTx(ctx, nil)
	if err != nil {
		return fail(err)
	}

	defer func(tx *sql.Tx) {
		_ = tx.Rollback()
	}(tx)

	result, err := tx.ExecContext(ctx, "INSERT INTO wallets (id, user_name, balance) VALUES ($1, $2, $3)", id, userName, balance)
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

func (w *WalletPostgresRepository) Read(ctx context.Context, id int) (domain.Wallet, error) {
	fail := func(err error) (domain.Wallet, error) {
		return domain.Wallet{}, fmt.Errorf("failed to read wallet: %w", err)
	}

	tx, err := w.db.BeginTx(ctx, nil)
	if err != nil {
		return fail(err)
	}

	defer func(tx *sql.Tx) {
		_ = tx.Rollback()
	}(tx)

	var wallet domain.Wallet
	if err = tx.QueryRowContext(ctx, "SELECT * FROM wallets WHERE id = $1", id).Scan(&wallet.Id, &wallet.UserName, &wallet.Balance); err != nil {
		return fail(err)
	}

	if err = tx.Commit(); err != nil {
		return fail(err)
	}

	return wallet, nil
}

func (w *WalletPostgresRepository) UpdateBalance(ctx context.Context, id, value int) error {
	fail := func(err error) error {
		return fmt.Errorf("failed to update wallet: %w", err)
	}

	tx, err := w.db.BeginTx(ctx, nil)
	if err != nil {
		return fail(err)
	}

	defer func(tx *sql.Tx) {
		_ = tx.Rollback()
	}(tx)

	result, err := tx.ExecContext(ctx, "UPDATE wallets SET balance = balance + $1 WHERE id = $2", value, id)
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

func (w *WalletPostgresRepository) Delete(ctx context.Context, id int) error {
	fail := func(err error) error {
		return fmt.Errorf("failed to delete wallet: %w", err)
	}

	tx, err := w.db.BeginTx(ctx, nil)
	if err != nil {
		return fail(err)
	}

	defer func(tx *sql.Tx) {
		_ = tx.Rollback()
	}(tx)

	result, err := tx.ExecContext(ctx, "DELETE FROM wallets WHERE id = $1", id)
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
