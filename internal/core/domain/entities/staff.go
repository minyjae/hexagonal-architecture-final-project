package entities

type Staff struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

type StaffLoginRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type StaffLoginResponse struct {
	Token string `json:"token"`
	Staff Staff  `json:"staff"`
}
