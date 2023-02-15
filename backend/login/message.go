package login

type LoginRequest struct {
	UserId   string `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	UserId   string `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
}
