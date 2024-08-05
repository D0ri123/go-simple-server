package types

type User struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

type GetUserResponse struct {
	*ApiResponse
	*User
}

type CreateUserResponse struct {
	*ApiResponse
}

type UpdateUserResponse struct {
	*ApiResponse
	*User
}

type DeleteUserResponse struct {
	*ApiResponse
}

func NewApiResponse(message string, responseCode int) *ApiResponse {
	return &ApiResponse{
		Description: message,
		Result:      int64(responseCode),
	}
}
