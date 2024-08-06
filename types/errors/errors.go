package errors

import (
	"fmt"
)

const (
	NotFoundUser = iota
	DBError
)

var errMessage = map[int64]string{
	NotFoundUser: "Not Found User",
	DBError:      "db err : ",
}

// 서버에서 요청했을 때, 에러를 핸들링하는 곳
func Errorf(code int64, args ...interface{}) error {
	if message, ok := errMessage[code]; ok {
		return fmt.Errorf("%s : %v", message, args)
	} else {
		return fmt.Errorf("not found err code")
	}
}
