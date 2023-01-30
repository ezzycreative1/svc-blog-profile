package dtos

type RegisterRequestBody struct {
	Firstname       string `json:"first_name"`
	Lastname        string `json:"last_name"`
	Email           string `json:"email"`
	PhoneNumber     string `json:"phone_number"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequestBody struct {
	Firstname   string `json:"first_name"`
	Lastname    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type ForgotPasswordRequestBody struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type LoginResponseBody struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserResponseBody struct {
	Firstname   string `json:"first_name"`
	Lastname    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	IsActive    int    `json:"is_active"`
}

type RegisterResponseBody struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}
