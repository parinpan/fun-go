package main

import (
	"fmt"

	fungo "github.com/parinpan/fun-go"
)

/*
	Software Engineer: Fachrin Aulia Nasution
	An always in beta software engineer
*/

func main() {
	// ProductList -> []ProductList{...}
	productSlice := fungo.NewSlice(ProductList)

	// get all product which has official store shop
	productsFromOS := productSlice.Filter(func(value interface{}, index int) bool {
		product := value.(Product)
		return product.Shop.IsOfficialStore
	}).([]Product)

	// get slice of product names string
	productNames := productSlice.Map([]string{}, func(value interface{}, index int) interface{} {
		return value.(Product).Name
	}).([]string)

	// get total of product price in the list
	productsPriceTotal := productSlice.Reduce(func(accumulator int64, value interface{}) int64 {
		return accumulator + int64(value.(Product).Price)
	})

	// print all product names
	productNameSlice := fungo.NewSlice(productNames)
	/*
		Output:
		  Kemeja
		  Vas
		  Dompet
		  Laptop
		  Dispenser
	*/
	productNameSlice.ForEach(func(value interface{}, index int) {
		fmt.Println(value.(string))
	})

	// output: Product Price Total: 15000
	fmt.Println(fmt.Sprint("Product Price Total: ", productsPriceTotal))

	// output: []main.Product{main.Product{ID:1, Name:"Kemeja", Price:1000, Shop:main.Shop{Name:"shop #1", IsOfficialStore:true}}, main.Product{ID:3, Name:"Dompet", Price:3000, Shop:main.Shop{Name:"shop #3", IsOfficialStore:true}}, main.Product{ID:5, Name:"Dispenser", Price:5000, Shop:main.Shop{Name:"shop #5", IsOfficialStore:true}}}
	fmt.Printf("%#v\n", productsFromOS)
}
