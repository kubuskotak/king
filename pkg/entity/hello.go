package entity

type Hello struct {
	Message string `json:"message" validate:"required"`
}
