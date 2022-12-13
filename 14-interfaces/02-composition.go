package main

import "fmt"

type Product struct {
	Id   int
	Name string
	Cost float32
}

/*
func (product Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %.2f", product.Id, product.Name, product.Cost)
}
*/

/* fmt.Stringer interface implementation */
func (product Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %.2f", product.Id, product.Name, product.Cost)
}

func (product *Product) ApplyDiscount(discount float32) {
	product.Cost = product.Cost * ((100 - discount) / 100)
}

type PerishableProduct struct {
	Product
	Expiry string
}

/*
func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%s, Expiry = %q", pp.Product.Format(), pp.Expiry)
}
*/
/* fmt.Stringer interface implementation */
func (pp PerishableProduct) String() string {
	return fmt.Sprintf("%s, Expiry = %q", pp.Product, pp.Expiry)
}

func NewPerishableProduct(id int, name string, cost float32, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			Id:   id,
			Name: name,
			Cost: cost,
		},
		Expiry: expiry,
	}
}

func main() {

	grapes := NewPerishableProduct(100, "Grapes", 50, "2 Days")
	fmt.Println("Before applying discount....")
	/*
		fmt.Println(Format(grapes.Product))
		ApplyDiscount(&grapes.Product, 10)
	*/

	// fmt.Println(grapes.Product.Format())
	// fmt.Println(grapes.Format())
	fmt.Println(grapes)

	fmt.Println("After applying discount....")
	grapes.ApplyDiscount(10)
	// fmt.Println(grapes.Format())
	fmt.Println(grapes)

}
