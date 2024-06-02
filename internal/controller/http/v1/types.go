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

type SendPasswordRecoveryEmailRequest struct {
	Email string `json:"email" binding:"required,email" example:"user@example.com"`
}

type SendPasswordRecoveryEmailResponse struct{}

type ResetPasswordRequest struct {
	Password string `json:"password" binding:"required,min=8" example:"password123"`
	FlowID   string `json:"flow_id" binding:"required" example:"cc833698-4519-4873-b9b4-67d6fef70dcb:1717170088"`
}

type ResetPasswordResponse struct{}

type GetMyUserInfoResponse struct {
	UserID          uint   `json:"user_id" example:"1"`
	Email           string `json:"email" example:"user@example.com"`
	Name            string `json:"name" example:"name"`
	Nickname        string `json:"nickname" example:"nickname"`
	ProfileImageURL string `json:"profile_image_url" example:"https://example.com/profile.png"`
	Bio             string `json:"bio" example:"bio..."`
	Website         string `json:"website" example:"https://example.com"`
}
