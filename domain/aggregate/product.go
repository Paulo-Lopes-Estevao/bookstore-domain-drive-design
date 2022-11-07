package aggregate

import "fmt"

type Product struct {
	CategoryID string
	Name       string
	Price      float64
}

func NewProduct(categoryID, name string, price float64) (*Product, error) {
	product := &Product{
		CategoryID: categoryID,
		Name:       name,
		Price:      price,
	}
	if err := product.Validate(); err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) Validate() error {
	if p.CategoryID == "" || p.Name == "" {
		return fmt.Errorf("categoryID and name are required")
	}
	if p.Price == 0 {
		return fmt.Errorf("price is required")
	}
	return nil
}
