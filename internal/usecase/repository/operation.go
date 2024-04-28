package repository

import (
	"context"
	"github.com/qrave1/BlazingMoneyControlBot/internal/domain"
)

type Operation interface {
	Create(ctx context.Context, operation domain.Operation) error
	ReadByBatch(ctx context.Context, walletId, page int) ([]domain.Operation, error)
}
