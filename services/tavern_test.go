package services

import (
	// "errors"
	"github.com/bangpenn/ddd-go/aggregate"
	"github.com/bangpenn/ddd-go/domain/customer/mongo"
	// "github.com/bangpenn/ddd-go/services"
	"github.com/google/uuid"
	// "go.mongodb.org/mongo-driver/bson"
	// "\go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	"context"
	"testing"
)

func Test_Tavern(t *testing.T) {
	products := init_products(t)

	os, err := NewOrderService(
		WithMongoCustomerRepository(context.Background(), "mongodb://localhost:8080"),
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

	if err = os.customers.Add(cust); err != nil {
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
