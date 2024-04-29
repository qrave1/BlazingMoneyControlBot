package usecase

import (
	"github.com/qrave1/BlazingMoneyControlBot/internal/usecase/repository"
	"log/slog"
)

type WalletUsecase struct {
	wr  repository.Wallet
	log *slog.Logger
}
