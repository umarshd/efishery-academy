package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"Name"`
}

var (
	products = map[int]*Product{}
	nomor    = 1
)

func CreateProduct(c echo.Context) error {
	p := &Product{
		ID: nomor,
	}

	if err := c.Bind(p); err != nil {
		return err
	}
	products[p.ID] = p
	nomor++
	return c.JSON(http.StatusCreated, products)

}

func GetProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func ShowProduct(c echo.Context) error {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	product := products[id]
	return c.JSON(http.StatusOK, product)
}

func updateProduct(c echo.Context) error {
	idParam := c.Param("id")
	idProduct, err := strconv.Atoi(idParam)

	if err != nil {
		errorMessage := "Invalid request parameter"
		return c.JSON(http.StatusBadRequest, errorMessage)
	}

	var result *Product

	for _, v := range products {
		if v.ID == idProduct {
			if err := c.Bind(v); err != nil {
				errorMessage := "Invalid request body"
				return c.JSON(http.StatusBadRequest, errorMessage)
			}
			result = v
		}
	}

	if result == nil {
		errorMessage := "Product not found"
		return c.JSON(http.StatusNotFound, errorMessage)
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteProduct(c echo.Context) error {
	idParam := c.Param("id")
	idProduct, err := strconv.Atoi(idParam)

	if err != nil {
		errorMessage := "Invalid request parameter"
		return c.JSON(http.StatusBadRequest, errorMessage)
	}

	delete(products, idProduct)
	successMessage := "Product with id " + idParam + " deleted successfully"

	return c.JSON(http.StatusOK, successMessage)
}

func main() {
	r := echo.New()

	r.POST("/", CreateProduct)
	r.GET("/products", GetProducts)
	r.GET("/products/:id", ShowProduct)
	r.PUT("/products/:id", updateProduct)
	r.DELETE("/products/:id", DeleteProduct)

	r.Logger.Fatal(r.Start(":1323"))
}
