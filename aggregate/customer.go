package aggregate

import "github.com/bangpenn/ddd-go/entity"

type Customer struct {
	person   *entity.Person
	products []*entity.Item
}
