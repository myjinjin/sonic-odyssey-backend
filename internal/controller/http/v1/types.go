package v1

type ErrorResponse struct {
	Error string `json:"error,omitempty"`
}

func NewErrorResponse(err error) ErrorResponse {
	return ErrorResponse{
		Error: err.Error(),
	}
}

type SignUpRequest struct {
	Email    string `json:"email" binding:"required,email" example:"user@example.com"`
	Password string `json:"password" binding:"required,min=8" example:"password123"`
	Name     string `json:"name" binding:"required" example:"John Doe"`
	Nickname string `json:"nickname" binding:"required" example:"johndoe"`
}

type SignUpResponse struct {
	UserID uint `json:"user_id" example:"1"`
}
