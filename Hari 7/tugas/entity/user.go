package entity

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"nama"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type UserRequest struct {
	Name  string `json:"nama"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"nama"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type SuccessResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
}
