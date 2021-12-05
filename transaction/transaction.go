package transaction

import (
	"context"

	transactionModel "github.com/azeezlala/cobra-cli/model/transaction"
)

type Transaction interface {
	SendMoney(ctx context.Context, transact transactionModel.SendMoney) error
	TransactionHistory(ctx context.Context, accountNumber string, pin int) ([]transactionModel.TransactionHistory, error)
}

type TransactionInit struct {
}

func (t TransactionInit) SendMoney(ctx context.Context, transact transactionModel.SendMoney) error {

	return nil
}

func (t TransactionInit) TransactionHistory(ctx context.Context, accountNumber string, pin int) ([]transactionModel.TransactionHistory, error) {
	var transactionHistory []transactionModel.TransactionHistory
	var err error

	return transactionHistory, err
}
