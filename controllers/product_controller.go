package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"tugasKedua/model"

	"github.com/gorilla/mux"
)

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	err := r.ParseForm()

	if err != nil {
		return
	}

	name_product := r.Form.Get("name_product")
	price, _ := strconv.Atoi(r.Form.Get("price"))

	result, errQuery := db.Exec("Insert into table_product(name_product, price) values(?,?)",
		name_product, price)

	var response model.ProductResponse
	var product model.Product
	temp, _ := result.LastInsertId()

	product.IdProduct = int(temp)
	product.Name_product = name_product
	product.Price = price

	if errQuery != nil {
		response.Status = 500
		response.Message = "An internal error occurred in the server"
		response.Data = product
		log.Fatal(errQuery.Error())
	} else {
		response.Status = 201
		response.Message = "New resource was created successfully"
		response.Data = product
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	err := r.ParseForm()

	if err != nil {
		return
	}

	vars := mux.Vars(r)

	product_id := vars["product_id"]

	result, errQuery := db.Exec("Delete from table_product where idProduct=?",
		product_id)

	var response model.ProductResponse

	temp, _ := result.RowsAffected()

	if temp != 0 {
		if errQuery != nil {
			var response model.ErrorResponse
			response.Status = 500
			response.Message = "Internal error occurred in the server"
			log.Fatal(errQuery.Error())
		} else {
			response.Status = 200
			response.Message = "The request was successful"
		}
	} else {
		response.Status = 404
		response.Message = "Indicates that the targeted resource does not exist"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	err := r.ParseForm()

	if err != nil {
		return
	}

	vars := mux.Vars(r)

	product_id := vars["product_id"]

	price, _ := strconv.Atoi(r.Form.Get("price"))

	result, errQuery := db.Exec("Update table_product set price=? where id=?",
		price, product_id)

	var response model.ProductResponse

	temp, _ := result.RowsAffected()

	if temp != 0 {
		if errQuery != nil {
			response.Status = 500
			response.Message = "Internal error occurred in the server"
			log.Fatal(errQuery.Error())
		} else {
			response.Status = 200
			response.Message = "The request was successful"
		}
	} else {
		response.Status = 404
		response.Message = "Indicates that the targeted resource does not exist"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	query := "Select * from table_product"

	rows, err := db.Query(query)

	if err != nil {
		var response model.ErrorResponse
		response.Status = 500
		response.Message = "Internal error occurred in the server"
		log.Println(err)
		return
	}

	var product model.Product
	var products []model.Product

	for rows.Next() {
		if err := rows.Scan(&product.IdProduct, &product.Name_product, &product.Price); err != nil {
			var response model.ErrorResponse
			response.Status = 500
			response.Message = "Internal error occurred in the server"
			log.Fatal(err.Error())
		} else {
			products = append(products, product)
		}
	}

	var response model.ProductsResponse
	response.Status = 200
	response.Message = "request was successful"
	response.Data = products

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
