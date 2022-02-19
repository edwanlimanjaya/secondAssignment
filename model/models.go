package model

type User struct {
	Id      int    `json : "id"`
	Name    string `json : "name"`
	Age     int    `json : "age"`
	Address string `json : "address"`
}

type Transaction struct {
	Id         int `json : "id"`
	User_id    int `json : "user_id"`
	Product_id int `json : "product_id"`
	Quantity   int `json : "quanity"`
}

type Product struct {
	Id    int    `json : "id"`
	Name  string `json : "name"`
	Price int    `json : "price"`
}

type ErrorResponse struct {
	Status  int    `json : "status"`
	Message string `json : "message"`
}

type UserResponse struct {
	Status  int    `json : "status"`
	Message string `json : "message"`
	Data    User   `json : "data"`
}

type UsersResponse struct {
	Status  int    `json : "Status"`
	Message string `json : "Message"'`
	Data    []User `json : "Data"`
}

type TransactionResponse struct {
	Status  int         `json : "Status"`
	Message string      `json : "Message"'`
	Data    Transaction `json : "Data"`
}

type TransactionsResponse struct {
	Status  int           `json : "Status"`
	Message string        `json : "Message"'`
	Data    []Transaction `json : "Data"`
}

type ProductResponse struct {
	Status  int     `json : "Status"`
	Message string  `json : "Message"'`
	Data    Product `json : "Data"`
}

type ProductsResponse struct {
	Status  int       `json : "Status"`
	Message string    `json : "Message"'`
	Data    []Product `json : "Data"`
}
