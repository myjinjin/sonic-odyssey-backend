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
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Name     string `json:"name" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
}

type SignUpResponse struct {
	UserID uint `json:"user_id"`
}
