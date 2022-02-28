package model

type User struct {
	IdUser    int    `json : "idUser"`
	Name_user string `json : "name_user"`
	Age       int    `json : "age"`
	Address   string `json : "address"`
	Email     string `json :"email"`
	Password  string `json : "password"`
}

type Transaction struct {
	IdTransaction int `json : "idTranscation"`
	User_id       int `json : "user_id"`
	Product_id    int `json : "product_id"`
	Quantity      int `json : "quanity"`
}

type Product struct {
	IdProduct    int    `json : "idProduct"`
	Name_product string `json : "name_product"`
	Price        int    `json : "price"`
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

type AllResponse struct {
	Status             int         `json : "Status"`
	Message            string      `json : "Message"'`
	UserTransaction    User        `json : UserTransaction`
	ProductTransaction Product     `json : ProductTransaction`
	DataTransaction    Transaction `json : "DataTransaction"`
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
