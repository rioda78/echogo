package Dto

type RequestBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`

	DisplayName string `json:"display_name" validate:"required"`
}
