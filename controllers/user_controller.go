package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"tugasKedua/model"

	"github.com/gorilla/mux"
)

func InsertUser(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	err := r.ParseForm()

	if err != nil {
		return
	}

	name := r.Form.Get("name")
	age, _ := strconv.Atoi(r.Form.Get("age"))
	address := r.Form.Get("address")

	_, errQuery := db.Exec("Insert into table_user(name,age,address) values(?,?,?)",
		name, age, address)

	var response model.UserResponse
	if errQuery != nil {
		response.Status = 400
		response.Message = "Couldn't insert the data"
	} else {
		response.Status = 200
		response.Message = "Success"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	err := r.ParseForm()

	if err != nil {
		return
	}

	vars := mux.Vars(r)
	userId := vars["user_id"]

	_, errQuery := db.Exec("Delete from table_user where id=?",
		userId)

	var response model.UserResponse
	if errQuery != nil {
		response.Status = 400
		response.Message = "Couldn't delete the data"
	} else {
		response.Status = 200
		response.Message = "Success"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	err := r.ParseForm()

	if err != nil {
		return
	}
	vars := mux.Vars(r)
	userId := vars["user_id"]

	name := r.Form.Get("name")

	_, errQuery := db.Exec("Update table_user set name=? where id=?",
		name, userId)
	var response model.UserResponse
	if errQuery != nil {
		response.Status = 400
		response.Message = "Couldn't update the data"
	} else {
		response.Status = 200
		response.Message = "Success"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	query := "SELECT * FROM table_user"

	rows, err := db.Query(query)
	if err != nil {
		var response model.ErrorResponse
		response.Status = 500
		response.Message = "Internal server error"
		log.Println(err)
		return
	}

	var user model.User
	var users []model.User
	for rows.Next() {
		if err != rows.Scan(&user.Id, &user.Name, &user.Age, &user.Address) {
			var response model.ErrorResponse
			response.Status = 500
			response.Message = "Internal server error"
			log.Fatal(err.Error())
		} else {
			users = append(users, user)
		}
	}

	var response model.UsersResponse
	response.Status = 200
	response.Message = "Request success"
	response.Data = users

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func InsertTransaction(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	err := r.ParseForm()

	if err != nil {
		return
	}

	user_id, _ := strconv.Atoi(r.Form.Get("user_id"))
	product_id, _ := strconv.Atoi(r.Form.Get("product_id"))
	quantity, _ := strconv.Atoi(r.Form.Get("quantity"))

	_, errQuery := db.Exec("Insert into table_transaction(user_id, product_id,quantity) values(?,?,?)",
		user_id, product_id, quantity)

	fmt.Println(errQuery)

	var response model.UserResponse
	if errQuery != nil {
		response.Status = 400
		response.Message = "Couldn't insert the data"
	} else {
		response.Status = 200
		response.Message = "Success"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	err := r.ParseForm()

	if err != nil {
		return
	}

	vars := mux.Vars(r)

	transaction_id := vars["transaction_id"]

	quantity, _ := strconv.Atoi(r.Form.Get("quantity"))

	_, errQuery := db.Exec("Update table_transaction set quantity=? where id=?",
		quantity, transaction_id)

	var response model.TransactionsResponse
	if errQuery != nil {
		response.Status = 400
		response.Message = "Couldn't update the data"
	} else {
		response.Status = 200
		response.Message = "Success"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	err := r.ParseForm()

	if err != nil {
		return
	}

	vars := mux.Vars(r)
	transaction_id := vars["transaction_id"]

	_, errQuery := db.Exec("Delete from table_transaction where id=?", transaction_id)

	var response model.TransactionsResponse
	if errQuery != nil {
		response.Status = 400
		response.Message = "Couldn't delete the data"
	} else {
		response.Status = 200
		response.Message = "Success"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	query := "Select * from table_transaction"

	rows, err := db.Query(query)

	if err != nil {
		var response model.ErrorResponse
		response.Status = 500
		response.Message = "Internal Server Error"
		log.Println(err)
		return
	}

	var transaction model.Transaction
	var transactions []model.Transaction
	for rows.Next() {
		if err := rows.Scan(&transaction.Id, &transaction.User_id, &transaction.Product_id, &transaction.Quantity); err != nil {
			log.Fatal(err.Error())
		} else {
			transactions = append(transactions, transaction)
		}
	}

	var response model.TransactionsResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = transactions

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	err := r.ParseForm()

	if err != nil {
		return
	}

	name_products := r.Form.Get("name")
	price, _ := strconv.Atoi(r.Form.Get("price"))

	_, errQuery := db.Exec("Insert into table_product(name, price) values(?,?)",
		name_products, price)

	var response model.ProductsResponse

	if errQuery != nil {
		response.Status = 400
		response.Message = "Couldn't insert the data"
	} else {
		response.Status = 200
		response.Message = "Success"
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

	_, errQuery := db.Exec("Delete from table_product where id=?",
		product_id)
	var response model.ProductsResponse
	if errQuery != nil {
		response.Status = 400
		response.Message = "Couldn't delete the data"
	} else {
		response.Status = 200
		response.Message = "Success"
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

	_, errQuery := db.Exec("Update table_product set price=? where id=?",
		price, product_id)
	fmt.Println(errQuery)
	var response model.ProductsResponse

	if errQuery != nil {
		response.Status = 400
		response.Message = "Couldn't update the data"
	} else {
		response.Status = 200
		response.Message = "Success"
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
		response.Message = "Internal Server Error"
		log.Println(err)
		return
	}

	var product model.Product
	var products []model.Product

	for rows.Next() {
		if err := rows.Scan(&product.Id, &product.Name, &product.Price); err != nil {
			log.Fatal(err.Error())
		} else {
			products = append(products, product)
		}
	}

	var response model.ProductsResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = products

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
