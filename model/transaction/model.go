package transaction

type SendMoney struct {
	Receiver string  `json:"receiver"`
	Sender   string  `json:"sender"`
	Amount   float64 `json:"amount"`
	Pin      int     `json:"pin"`
}

type TransactionHistory struct {
	UserID      int
	Transaction []SendMoney
}
