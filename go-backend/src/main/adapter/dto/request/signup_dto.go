package request

type SignUpDTO struct {
	PhoneNumber     string `json:"phoneNumber"`
	Email           string `json:"email"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	Source          string `json:"source"`
}
