package products

import "sort"

type Product struct {
	ID     int
	Barang string
	Harga  int
}

type Products []Product

var DataProduct = []Product{
	{ID: 1, Barang: "Benih Lele", Harga: 50000},
	{ID: 2, Barang: "Pakan Lele Cap Menara", Harga: 25000},
	{ID: 3, Barang: "Probiotik A", Harga: 75000},
	{ID: 4, Barang: "Probiotik Nila B", Harga: 10000},
	{ID: 5, Barang: "Pakan Nila", Harga: 20000},
	{ID: 6, Barang: "Benih Nila", Harga: 20000},
	{ID: 7, Barang: "Cupang", Harga: 5000},
	{ID: 8, Barang: "Benih Nila", Harga: 30000},
	{ID: 9, Barang: "Benih Cupang", Harga: 10000},
	{ID: 10, Barang: "Probiotik B", Harga: 10000},
}

func ListProduct() Products {
	return DataProduct
}

func (products Products) UrutkanProductHarga() Products {

	sort.Slice(products, func(x, y int) bool {
		return products[x].Harga < products[y].Harga
	})

	return products
}

func (products Products) BeliProduct(sisaSoldo *int) Products {
	var result Products

	products.UrutkanProductHarga()

	for _, product := range DataProduct {
		if *sisaSoldo >= product.Harga {
			*sisaSoldo -= product.Harga
			result = append(result, product)
		} else {
			break
		}
	}

	return result
}

func (products Products) ProdcutTermahal() Product {
	result := products[0]

	for i := 1; i < len(products); i++ {
		if products[i].Harga > result.Harga {
			result = products[i]
		}
	}

	return result
}

func (products Products) ProductTermurah() Product {
	result := products[0]

	for i := 1; i < len(products); i++ {
		if products[i].Harga < result.Harga {
			result = products[i]
		}
	}

	return result
}

func (products Products) DataProductHarga(harga int) Products {
	var result Products

	for _, product := range products {
		if product.Harga == harga {
			result = append(result, product)
		}
	}

	return result
}
