package services

import (
	"github.com/bangpenn/ddd-go/aggregate"
	"github.com/google/uuid"
	"testing"
)

func init_products(t *testing.T) []aggregate.Product {
	beer, err := aggregate.NewProduct("Beer", "Healthy Beverage", 2.99)
	if err != nil {
		t.Fatal(err)
	}

	peenuts, err := aggregate.NewProduct("Peanuts", "Snacks", 0.99)
	if err != nil {
		t.Fatal(err)
	}

	wine, err := aggregate.NewProduct("Wine", "Nasty Drink", 0.99)
	if err != nil {
		t.Fatal(err)
	}

	return []aggregate.Product{
		beer, peenuts, wine,
	}

}

func TestOrder_NewOrderSeriv(t *testing.T) {

	products := init_products(t)

	os, err := NewOrderService(
		WithCustomerRepository(),
		WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Fatal(err)
	}

	cust, err := aggregate.NewCustomer("Rizal")
	if err != nil {
		t.Error(err)
	}

	err = os.customers.Add(cust)
	if err != nil {
		t.Error(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = os.CreateOrder(cust.GetID(), order)

	if err != nil {
		t.Error(err)
	}

}
