package domain

type Wallet struct {
	// Id is tg_user_id
	Id       int
	UserName string

	// Balance in RUB
	Balance int64

	//History []Operation
}
