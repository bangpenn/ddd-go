package services

import (
	"errors"
	"github.com/bangpenn/ddd-go/aggregate"
	"github.com/bangpenn/ddd-go/domain/customer"
	"github.com/bangpenn/ddd-go/domain/customer/memory"
	"github.com/bangpenn/ddd-go/domain/product"
	prodmem "github.com/bangpenn/ddd-go/domain/product/memory"
	"github.com/google/uuid"
	"log"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository
	products  product.ProductRepository
}

// NewOrderService(WithCustomerRepository, WithMemoryProductRepository)

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}
	// loop through all the cfgs and apply them
	for _, cfg := range cfgs {
		err := cfg(os)

		if err != nil {
			return nil, err
		}

	}
	return os, nil
}

// NewOrderService(
// 	WithMongoCustomerRepository()
// 	WithLogging("debug"),
// )

// WithCustomer Repository applieas a customer repoository to the OrderService
func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	// Return a function that matches the orderconfiguration alias
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := prodmem.New()

		for _, p := range products{
			if err := pr.Add(p); err i= nil {
				return err
			}
		}

		os.products = pr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)

}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productsIDs []uuid.UUID) error {
	// Fetch the Customer
	c, err := o.customers.Get(customerID)
	if err != nil {
		return err
	}

	// Get each Product, Ouchie no productrepository
	// log.Println(c)
	var products []aggregate.Product
	var total float64

	for _, id := range productsIDs{
		p, err := o.products.GetByID(id)

		if err != nil {
			return err
		}

		products = append(products, p)
		total += p.GetPrice()
	}
	log.Printf("Customer: %s has ordered %d products", c.GetID(), len(products))
	return nil
}
