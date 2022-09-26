package main

import (
	"fmt"
	"tugas-week-1-1/products"

	"github.com/yudapc/go-rupiah"
)

func main() {

	dataProduct := products.ListProduct()
	saldo := 100000
	sisaSoldo := saldo

	// Poin A
	paidProducts := dataProduct.BeliProduct(&sisaSoldo)

	fmt.Println("Daftar belanjaan :")
	fmt.Println("----------------------------------------")

	for _, product := range paidProducts {
		fmt.Printf("%s  %s\n", product.Barang, rupiah.FormatRupiah(float64(product.Harga)))
	}
	fmt.Println("========================================")
	fmt.Println("Total Belanja :", rupiah.FormatRupiah(float64(saldo-sisaSoldo)))
	fmt.Println("Total Bayar :", rupiah.FormatRupiah(float64(saldo)))
	fmt.Println("Kembali:", rupiah.FormatRupiah(float64(sisaSoldo)))

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	// Poin B
	productTermurah := dataProduct.ProductTermurah()
	productTermahal := dataProduct.ProdcutTermahal()

	fmt.Println("Product Termurah dan Termahal")
	fmt.Printf("Produk Termurah : %s  %s\n", productTermurah.Barang, rupiah.FormatRupiah(float64(productTermurah.Harga)))
	fmt.Printf("Produk Termahal : %s  %s\n", productTermahal.Barang, rupiah.FormatRupiah(float64(productTermahal.Harga)))

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	harga := 10000

	// Poin C
	productInPrice := dataProduct.DataProductHarga(harga)

	fmt.Printf("Produk Dengan Harga %s :\n", rupiah.FormatRupiah(float64(harga)))
	fmt.Println("----------------------------------------")

	for _, product := range productInPrice {
		fmt.Printf("%s  %s\n", product.Barang, rupiah.FormatRupiah(float64(product.Harga)))
	}

}
