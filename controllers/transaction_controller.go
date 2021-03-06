package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"strconv"
	"tugasKedua/model"

	"github.com/gorilla/mux"
)

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

	result, errQuery := db.Exec("Insert into table_transaction(user_id, product_id,quantity) values(?,?,?)",
		user_id, product_id, quantity)

	var response model.TransactionResponse
	var transaction model.Transaction
	temp, _ := result.LastInsertId()

	transaction.IdTransaction = int(temp)
	transaction.User_id = user_id
	transaction.Product_id = product_id
	transaction.Quantity = quantity

	if errQuery != nil {
		response.Status = 500
		response.Message = "An internal error occurred in the server"
		response.Data = transaction
		log.Fatal(errQuery.Error())
	} else {
		response.Status = 201
		response.Message = "New resource was created successfully"
		response.Data = transaction
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
// 	db := Connect()
// 	defer db.Close()

// 	err := r.ParseForm()

// 	if err != nil {
// 		return
// 	}

// 	vars := mux.Vars(r)
// 	transaction_id := vars["transaction_id"]

// 	quantity, _ := strconv.Atoi(r.Form.Get("quantity"))

// 	result, errQuery := db.Exec("Update table_transaction set quantity=? where id=?",
// 		quantity, transaction_id)

// 	var response model.TransactionResponse
// 	temp, _ := result.RowsAffected()

// 	if temp != 0 {
// 		if errQuery != nil {
// 			response.Status = 500
// 			response.Message = "Internal error occurred in the server"
// 			log.Fatal(errQuery.Error())
// 		} else {
// 			response.Status = 200
// 			response.Message = "The request was successful"
// 		}
// 	} else {
// 		response.Status = 404
// 		response.Message = "Indicates that the targeted resource does not exist"
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)

// }

// func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
// 	db := Connect()
// 	defer db.Close()

// 	err := r.ParseForm()

// 	if err != nil {
// 		return
// 	}

// 	vars := mux.Vars(r)
// 	transaction_id := vars["transaction_id"]

// 	result, errQuery := db.Exec("Delete from table_transaction where id=?", transaction_id)

// 	var response model.TransactionResponse
// 	temp, _ := result.RowsAffected()

// 	if temp != 0 {
// 		if errQuery != nil {
// 			response.Status = 500
// 			response.Message = "Internal error occurred in the server"
// 			log.Fatal(errQuery.Error())
// 		} else {
// 			response.Status = 200
// 			response.Message = "The request was successful"
// 		}
// 	} else {
// 		response.Status = 404
// 		response.Message = "Indicates that the targeted resource does not exist"
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }

// func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
// 	db := Connect()
// 	defer db.Close()

// 	query := "Select * from table_transaction"

// 	rows, err := db.Query(query)

// 	if err != nil {
// 		var response model.ErrorResponse
// 		response.Status = 500
// 		response.Message = "Internal error occurred in the server"
// 		log.Println(err)
// 		return
// 	}

// 	var transaction model.Transaction
// 	var transactions []model.Transaction
// 	for rows.Next() {
// 		if err := rows.Scan(&transaction.IdTransaction, &transaction.User_id, &transaction.Product_id, &transaction.Quantity); err != nil {
// 			var response model.ErrorResponse
// 			response.Status = 500
// 			response.Message = "Internal error occurred in the server"
// 			log.Fatal(err.Error())
// 		} else {
// 			transactions = append(transactions, transaction)
// 		}
// 	}

// 	var response model.TransactionsResponse
// 	response.Status = 200
// 	response.Message = "Request was successful"
// 	response.Data = transactions

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(response)
// }

func GetAllDetailTransactions(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	err := r.ParseForm()

	if err != nil {
		return
	}

	vars := mux.Vars(r)
	user_id := vars["user_id"]

	//user_id := r.Form.Get("user_id")

	var user model.User
	var transaction model.Transaction
	var product model.Product

	errQuery := db.QueryRow(`SELECT table_user.idUser, table_user.name_user, table_user.age, table_user.address, table_product.idProduct, table_product.name_product, table_product.price, table_transaction.quantity 
	FROM table_user 
	INNER JOIN table_transaction 
	ON table_user.idUser = table_transaction.user_id 
	INNER JOIN table_product 
	ON table_transaction.product_id = table_product.idProduct
	WHERE table_user.idUser =?`, user_id).Scan(&user.IdUser, &user.Name_user, &user.Age, &user.Address, &product.IdProduct, &product.Name_product, &product.Price, &transaction.Quantity)

	if errQuery != nil {
		var response model.ErrorResponse
		response.Status = 404
		response.Message = "Indicates that the targeted resource does not exist"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		var response model.AllResponse
		response.Status = 200
		response.Message = "Request was successful"
		response.UserTransaction = user
		response.ProductTransaction = product
		response.DataTransaction = transaction
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}

}
