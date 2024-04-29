package dbo

import (
	"github.com/qrave1/BlazingMoneyControlBot/internal/domain"
	"strconv"
)

type WalletRedisDBO struct {
	Id       string
	UserName string
	Balance  int64
}

func FromDomain(wallet domain.Wallet) WalletRedisDBO {
	return WalletRedisDBO{
		Id:       strconv.Itoa(wallet.Id),
		UserName: wallet.UserName,
		Balance:  wallet.Balance,
	}
}
