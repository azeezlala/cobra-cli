package customer

import (
	"context"

	customerModel "github.com/azeezlala/cobra-cli/model/customer"
)

type CustomerInterface interface {
	CreateCustomer(ctx context.Context, customer customerModel.CustomerDetail) error
}

type Customer struct {
}

func (c Customer) CreateCustomer(ctx context.Context, customer customerModel.CustomerDetail) error {
	return nil
}
