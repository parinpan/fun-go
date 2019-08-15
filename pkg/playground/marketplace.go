package main

/*
	Software Engineer: Fachrin Aulia Nasution
	An always in beta software engineer
*/

// Shop struct
type Shop struct {
	Name            string
	IsOfficialStore bool
}

// Product struct
type Product struct {
	ID    int
	Name  string
	Price float64
	Shop  Shop
}

// ProductList data collection
var ProductList = []Product{
	Product{ID: 1, Name: "Kemeja", Price: 1000, Shop: Shop{Name: "shop #1", IsOfficialStore: true}},
	Product{ID: 2, Name: "Vas", Price: 2000, Shop: Shop{Name: "shop #2", IsOfficialStore: false}},
	Product{ID: 3, Name: "Dompet", Price: 3000, Shop: Shop{Name: "shop #3", IsOfficialStore: true}},
	Product{ID: 4, Name: "Laptop", Price: 4000, Shop: Shop{Name: "shop #4", IsOfficialStore: false}},
	Product{ID: 5, Name: "Dispenser", Price: 5000, Shop: Shop{Name: "shop #5", IsOfficialStore: true}},
}
