package types

type User struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

type GetUserResponse struct {
	*ApiResponse
	Users []*User `json:"result"`
}

type CreateRequest struct {
	Name string `json:"name" binding:"required"`
	Age  int64  `json:"age" binding:"required"`
}

func (c *CreateRequest) ToUser() *User {
	return &User{
		Name: c.Name,
		Age:  c.Age,
	}
}

type CreateUserResponse struct {
	*ApiResponse
}

// gin 프레임워크 깃허브에 아래옵션들을 많이 볼 수 잇다.
type UpdateUserRequest struct {
	Name string `json:"name" binding:"required"`
	Age  int64  `json:"updateAge" binding:"required"`
}

type UpdateUserResponse struct {
	*ApiResponse
	*User
}

type DeleteRequest struct {
	Name string `json:"name" binding:"required"`
}

type DeleteUserResponse struct {
	*ApiResponse
}

func NewApiResponse(message string, responseCode int, errCode interface{}) *ApiResponse {
	return &ApiResponse{
		Description: message,
		Result:      int64(responseCode),
		ErrCode:     errCode,
	}
}

func (c *DeleteRequest) ToUser() *User {
	return &User{
		Name: c.Name,
	}
}
