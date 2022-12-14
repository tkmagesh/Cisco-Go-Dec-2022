package main

import (
	"errors"
	"fmt"
	"sort"
)

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (p Product) String() string {
	return fmt.Sprintf("Id=%d, Name=%s, Cost=%v, Units=%d, Category=%s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

/*
Write the apis for the following
*/
type Products []Product

func (products Products) String() string {
	result := ""
	for _, p := range products {
		result += fmt.Sprintf("%s\n", p)
	}
	return result
}

/*
IndexOf => return the index of the given product

	ex:  returns the index of the given product
*/

func (products Products) IndexOf(product Product) (int, error) {
	for i, val := range products {
		if val == product {
			return i, nil
		}
	}
	err := errors.New("Product not found!")
	return 0, err
}

/*
Includes => return true if the given product is present in the collection else return false
	ex:  returns true if the given product is present in the collection
*/

func (products Products) Includes(product Product) bool {
	for _, val := range products {
		if product == val {
			return true
		}
	}
	return false
}

/*
Filter => return a new collection of products that satisfy the given condition
	use cases:
		1. filter all costly products (cost > 1000)
			OR
		2. filter all stationary products (category = "Stationary")
			OR
		3. filter all costly stationary products
		etc
*/

type Specification interface {
	IsSatisfied(product *Product) bool
}

type CostSpecification struct {
	Cost float32
}

func (c CostSpecification) IsSatisfied(product *Product) bool {
	/* if product.Cost > c.Cost {
		return true
	} else {
		return false
	} */
	return product.Cost > c.Cost
}

type CategorySpecification struct {
	Category string
}

func (c CategorySpecification) IsSatisfied(product *Product) bool {
	if product.Category == c.Category {
		return true
	} else {
		return false
	}
}

type CostAndCategorySpecification struct {
	Cost     float32
	Category string
}

func (cc CostAndCategorySpecification) IsSatisfied(product *Product) bool {
	if (product.Cost > cc.Cost) && (product.Category == cc.Category) {
		return true
	} else {
		return false
	}
}

type ProductPredicate func(product Product) bool

func (products Products) Filter(predicate ProductPredicate) []Product {
	result := make([]Product, 0)
	for _, value := range products {
		if predicate(value) {
			result = append(result, value)
		}
	}
	return result
}

/*
	Any => return true if any of the product in the collections satifies the given criteria
		use cases:
			1. are there any costly product (cost > 1000)?
			2. are there any stationary product (category = "Stationary")?
			3. are there any costly stationary product?
			etc
*/

func (products Products) Any(spec Specification) bool {
	for _, v := range products {
		if spec.IsSatisfied(&v) {
			return true
		}
	}
	return false
}

/*
	All => return true if all the products in the collections satifies the given criteria
		use cases:
			1. are all the products costly products (cost > 1000)?
			2. are all the products stationary products (category = "Stationary")?
			3. are all the products costly stationary products?
			etc
*/

func (products Products) All(spec Specification) bool {
	for _, v := range products {
		if !spec.IsSatisfied(&v) {
			return false
		}
	}
	return true
}

/*

	Write the apis for the following
		Sort => Sort the products collection by any attribute
			IMPORTANT : MUST Use sort.Sort() function
            use cases:
                1. sort the products collection by cost
                2. sort the products collection by name
                3. sort the products collection by units
				etc

*/

/* sort.Interface implementation for Products (default sorting) */
func (a Products) Len() int           { return len(a) }
func (a Products) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Products) Less(i, j int) bool { return a[i].Id < a[j].Id }

/* sort by name */
type ByName struct {
	Products
}

func (a ByName) Less(i, j int) bool { return a.Products[i].Name < a.Products[j].Name }

/* sort by cost */
type ByCost struct {
	Products
}

func (a ByCost) Less(i, j int) bool { return a.Products[i].Cost < a.Products[j].Cost }

/* sort by units */
type ByUnits struct {
	Products
}

func (a ByUnits) Less(i, j int) bool { return a.Products[i].Units < a.Products[j].Units }

/*
type ByCost []Product

func (a ByCost) Len() int           { return len(a) }
func (a ByCost) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCost) Less(i, j int) bool { return a[i].Cost < a[j].Cost }

type ByName []Product

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

type ByUnits []Product

func (a ByUnits) Len() int           { return len(a) }
func (a ByUnits) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByUnits) Less(i, j int) bool { return a[i].Units < a[j].Units }
*/

func main() {
	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
		Product{109, "Golden Pen", 2000, 20, "Stationary"},
	}

	//IndexOf
	PencilProduct := Product{107, "Pencil", 2, 100, "Stationary"}
	// index, err := IndexOf(products, PencilProduct)
	index, err := products.IndexOf(PencilProduct)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Index of Pencil Product=", index)
	}

	//Includes
	StoveProduct := Product{102, "Stove", 5000, 5, "Utencil"}
	includes := products.Includes(StoveProduct)
	if includes {
		fmt.Println("Stove Product is included")
	} else {
		fmt.Println("Stove Product is NOT included")
	}

	//Filter (using ISpecification)
	/*

		costSpec := CostSpecification{1000}
		costlyProducts := products.Filter(costSpec)

		fmt.Println("\nProducts with Cost > 1000:")
		fmt.Println(costlyProducts)

		categorySpec := CategorySpecification{"Stationary"}
		StationaryProducts := products.Filter(categorySpec)

		fmt.Println("\nProducts with Stationary Category:")
		fmt.Println(StationaryProducts)

		costAndCategorySpec := CostAndCategorySpecification{1000, "Stationary"}
		CostlyStationaryProducts := products.Filter(costAndCategorySpec)

		fmt.Println("\nProducts with Cost > 1000 and having Stationary Category:")
		fmt.Println(CostlyStationaryProducts)
	*/
	//Filter (using ProductPredicate)
	costlyProductPredicate := func(product Product) bool {
		return product.Cost > 1000
	}
	costlyProducts := products.Filter(costlyProductPredicate)

	fmt.Println("\nProducts with Cost > 1000:")
	fmt.Println(costlyProducts)

	stationaryProductPredicate := func(product Product) bool {
		return product.Category == "Stationary"
	}
	StationaryProducts := products.Filter(stationaryProductPredicate)

	fmt.Println("\nProducts with Stationary Category:")
	fmt.Println(StationaryProducts)

	costlyStationaryProductPredicate := func(product Product) bool {
		return costlyProductPredicate(product) && stationaryProductPredicate(product)
	}
	CostlyStationaryProducts := products.Filter(costlyStationaryProductPredicate)
	fmt.Println("\nProducts with Cost > 1000 and having Stationary Category:")
	fmt.Println(CostlyStationaryProducts)

	//Any (Can be modified to use the ProductPredicate)
	costSpec := CostSpecification{1000}
	anyCostlyProducts := products.Any(costSpec)
	fmt.Print("\nAre there any costly(>1000) Products:")
	fmt.Println(anyCostlyProducts)

	categorySpec := CategorySpecification{"Stationary"}
	anyStationaryProducts := products.Any(categorySpec)
	fmt.Print("\nAre there any Stationary Products:")
	fmt.Println(anyStationaryProducts)

	costAndCategorySpec := CostAndCategorySpecification{1000, "Stationary"}
	anyCostlyStationaryProducts := products.Any(costAndCategorySpec)
	fmt.Print("\nAre there any Costly Stationary Products:")
	fmt.Println(anyCostlyStationaryProducts)

	// All
	allCostlyProducts := products.All(costSpec)
	fmt.Print("\nAre there all costly(>1000) Products:")
	fmt.Println(allCostlyProducts)

	allStationaryProducts := products.All(categorySpec)
	fmt.Print("\nAre there all Stationary Products:")
	fmt.Println(allStationaryProducts)

	allCostlyStationaryProducts := products.All(costAndCategorySpec)
	fmt.Print("\nAre there all Costly Stationary Products:")
	fmt.Println(allCostlyStationaryProducts)

	//Sort
	sort.Sort(ByCost{products})
	fmt.Println("After sorting by Cost")

	fmt.Println(products)

	sort.Sort(ByName{products})
	fmt.Println("After sorting by Name")

	fmt.Println(products)

	sort.Sort(ByUnits{products})
	fmt.Println("After sorting by Units")

	fmt.Println(products)

}
