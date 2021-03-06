// package product

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"strconv"
// 	"strings"

// 	"creatingwebservices/cors"
// )

// const productsBasePath = "products"

// func SetupRoutes(apiBasePath string) {
// 	handleProducts := http.HandlerFunc(productsHandler)
// 	handleProduct := http.HandlerFunc(productHandler)
// 	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, productsBasePath), cors.Middleware(handleProducts))
// 	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, productsBasePath), cors.Middleware(handleProduct))
// }

// func productsHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodGet:
// 		productList := getProductList()
// 		productsJSON, err := json.Marshal(productList)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write(productsJSON)
// 	case http.MethodPost:
// 		var newProduct Product
// 		bodyBytes, err := ioutil.ReadAll(r.Body)
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}
// 		err = json.Unmarshal(bodyBytes, &newProduct)
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}
// 		if newProduct.ProductID != 0 {
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}

// 		_, err = addOrUpdateProduct(newProduct)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
// 		w.WriteHeader(http.StatusCreated)
// 		return
// 	case http.MethodOptions:
// 		return
// 	}
// }

// func productHandler(w http.ResponseWriter, r *http.Request) {

// 	urlPathSegments := strings.Split(r.URL.Path, "products/")
// 	productID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
// 	if err != nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		return
// 	}
// 	product := getProduct(productID)
// 	if product == nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		return
// 	}

// 	switch r.Method {
// 	case http.MethodGet:
// 		prodJSON, err := json.Marshal(product)
// 		if err != nil {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write(prodJSON)
// 	case http.MethodPut:
// 		var updatedProduct Product
// 		bodyBytes, err := ioutil.ReadAll(r.Body)
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}
// 		err = json.Unmarshal(bodyBytes, &updatedProduct)
// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}
// 		if updatedProduct.ProductID != productID {
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}

// 		addOrUpdateProduct(updatedProduct)

// 		w.WriteHeader(http.StatusCreated)
// 		return
// 	case http.MethodDelete:
// 		removeProduct(productID)
// 	case http.MethodOptions:
// 		return
// 	default:
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 	}
// }
package product
