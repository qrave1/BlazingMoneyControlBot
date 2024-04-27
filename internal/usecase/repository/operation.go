package repository

import (
	"context"
	"github.com/qrave1/BlazingMoneyControlBot/internal/domain"
)

type Operation interface {
	Create(ctx context.Context, userId int, operation string, amount int, reason string) error
	Read(ctx context.Context, id int) (domain.Operation, error)
	ReadAll(ctx context.Context, userId int) ([]domain.Operation, error)
	Delete(ctx context.Context, id int) error
}
