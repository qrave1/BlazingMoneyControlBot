package repository

import (
	"context"
	"github.com/qrave1/BlazingMoneyControlBot/internal/domain"
)

type Wallet interface {
	Create(ctx context.Context, id int, userName string, balance int) error
	Read(ctx context.Context, id int) (domain.Wallet, error)
	UpdateBalance(ctx context.Context, id, value int) error
	Delete(ctx context.Context, id int) error
}
