package main

import (
	"creatingwebservices/database"
	"creatingwebservices/product"
	"creatingwebservices/receipt"
	"net/http"

	// imported for side effects
	_ "github.com/go-sql-driver/mysql"
)

const apiBasePath = "/api"

func main() {
	database.SetupDatabase()
	receipt.SetupRoutes(apiBasePath)
	product.SetupRoutes(apiBasePath)
	http.ListenAndServe(":5000", nil)
}
