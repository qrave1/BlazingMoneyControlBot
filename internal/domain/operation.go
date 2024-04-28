package domain

type Operation struct {
	WalletId int
	Type     string
	Amount   int
	Reason   string
}

const (
	TypeDeposit = "deposit"
	TypeDebit   = "debit"
)
