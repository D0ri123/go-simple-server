package types

type ApiResponse struct {
	Result      int64       `json:"result"`
	Description string      `json:"description"`
	ErrCode     interface{} `json:"errCode"`
}
