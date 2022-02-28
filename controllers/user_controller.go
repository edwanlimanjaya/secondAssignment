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
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	result, errQuery := db.Exec("Insert into table_user(name,age,address) values(?,?,?,?,?)",
		name, age, address, email, password)
	var user model.User
	temp, _ := result.LastInsertId()
	user.IdUser = int(temp)
	user.Name_user = name
	user.Age = age
	user.Address = address
	user.Email = email
	user.Password = password

	var response model.UserResponse
	if errQuery != nil {
		response.Status = 500
		response.Message = "An internal error occurred in the server"
		log.Fatal(errQuery.Error())
	} else {
		response.Status = 201
		response.Message = "New resource was created successfully"
		response.Data = user
		fmt.Println(user.IdUser)
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

	var response model.UserResponse

	result, errQuery := db.Exec("Delete from table_user where id=?",
		userId)

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
	var response model.UserResponse

	result, errQuery := db.Exec("Update table_user set name=? where id=?",
		name, userId)

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
		if err != rows.Scan(&user.IdUser, &user.Name_user, &user.Age, &user.Address, &user.Email, &user.Password) {
			var response model.ErrorResponse
			response.Status = 500
			response.Message = "Internal error occurred in the server"
			log.Fatal(err.Error())
		} else {
			users = append(users, user)
		}
	}

	var response model.UsersResponse
	response.Status = 200
	response.Message = "Request was successful"
	response.Data = users

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	err := r.ParseForm()

	if err != nil {
		var response model.ErrorResponse
		response.Status = 500
		response.Message = "Internal server error"
		log.Println(err)
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	_, errQuery := db.Exec(`SELECT * FROM table_user 
	WHERE email = ? 
	AND password = ?`, email, password)

	var response model.LoginResponse
	if errQuery != nil {
		response.Status = 500
		response.Message = "Internal server error"
		log.Println(err)
	} else {
		response.Status = 200
		response.Message = "The request was successful"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
