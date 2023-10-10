package services

import (
	"github.com/bangpenn/ddd-go/aggregate"
	"github.com/google/uuid"
	"testing"
)

func Test_Tavern(t *testing.T) {
	products := init_products(t)

	os, err = NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Fatal(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Fatal(err)
	}

	cust, err := aggregate.NewCustomer("Rizal")
	if err != nil {
		t.Fatal(err)
	}

	if err = os.customers.Add(); err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = tavern.Order(cust.GetID(), order)

	if err != nil {
		t.Fatal(err)
	}
}
