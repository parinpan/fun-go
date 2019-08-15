
# fun-go by Fachrin Aulia

Fun-go is the answer when you are sick of writing over and over again a looping block that seems giving much redundancy in your project. It's inspired by Javascript programming language to approach a better functional programming paradigm using Golang that at least currently doesn't support iteration methods.

```go
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
```

Benchmark Result tested on **MacBook Pro (13-inch, 2016, Two Thunderbolt 3 ports) 2 GHz Intel Core i5 8 GB 1867 MHz LPDDR3** with 1.000.000 slice elements of struct can be found in the test file.
```
goos: darwin
goarch: amd64
pkg: github.com/parinpan/fun-go
BenchmarkFunctionalIteratorMap/functionalIterator-4             2000000000               0.29 ns/op
BenchmarkFunctionalIteratorMap/normalIterator-4                 2000000000               0.13 ns/op
BenchmarkFunctionalIteratorFilter/functionalIterator-4          2000000000               0.14 ns/op
BenchmarkFunctionalIteratorFilter/normalIterator-4              2000000000               0.05 ns/op
BenchmarkFunctionalIteratorForEach/functionalIterator-4         2000000000               0.03 ns/op
BenchmarkFunctionalIteratorForEach/normalIterator-4             2000000000               0.00 ns/op
BenchmarkFunctionalIteratorReduce/functionalIterator-4          2000000000               0.03 ns/op
BenchmarkFunctionalIteratorReduce/normalIterator-4              2000000000               0.00 ns/op
```
