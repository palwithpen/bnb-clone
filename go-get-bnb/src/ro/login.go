package ro

type Login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

type Reset struct {
	Email string `json:"email" validate:"required,email"`
}
