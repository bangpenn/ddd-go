package memory

import (
	"errors"
	"github.com/bangpenn/ddd-go/aggregate"
	"github.com/bangpenn/ddd-go/domain/customer"
	"github.com/google/uuid"
	"testing"
)

func TestMemory_GetCustom(t *testing.T) {
	type testCase struct {
		name         string
		id           uuid.UUID
		exepectedErr error
	}

	cust, err := aggregate.NewCustomer("rizal")
	if err != nil {
		t.Fatal(err)
	}

	id := cust.GetID()

	repo := MemoryRepository{
		customers: map[uuid.UUID]aggregate.Customer{
			id: cust,
		},
	}

	testCases := []testCase{
		{
			name:         "no customer by id",
			id:           uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			exepectedErr: customer.ErrCustomerNotFound,
		}, {
			name:         "customer by id",
			id:           id,
			exepectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)

			if !errors.Is(err, tc.exepectedErr) {
				t.Errorf("expected error %v, got %v", tc.exepectedErr, err)
			}
		})
	}

}
