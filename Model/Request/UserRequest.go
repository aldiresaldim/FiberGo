package Request

type UserCreateRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserUpdateRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type UserEmailRequest struct {
	Email string `json:"email" validate:"required"`
}
