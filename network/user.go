package network

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-server/types"
	"sync"
)

var (
	userRouterInit     sync.Once // null 포인터 체크,단 한번만 호출되어야 하는 API인 경우
	userRouterInstance *userRouter
)

type userRouter struct {
	router *Network
	// service
}

func newUserRouter(router *Network) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{
			router: router,
		}
	})

	router.registerGET("/", userRouterInstance.get)
	router.registerPOST("/", userRouterInstance.create)
	router.registerUPDATE("/", userRouterInstance.update)
	router.registerDELETE("/", userRouterInstance.delete)

	return userRouterInstance
}

// CRUD
func (u *userRouter) create(c *gin.Context) {
	fmt.Println("create 입니다.")
}

func (u *userRouter) get(c *gin.Context) {
	fmt.Println("get 입니다.")
	//u.router.okResponse(c, &types.ApiResponse{
	//	Result:      200,
	//	Description: "성공입니다.",
	//})

	u.router.okResponse(c, &types.UserResponse{
		ApiResponse: &types.ApiResponse{
			Result:      200,
			Description: "성공입니다.",
		},
		User: &types.User{
			Name: "이설희",
			Age:  24,
		},
	})
}

func (u *userRouter) update(c *gin.Context) {
	fmt.Println("update 입니다.")
}

func (u *userRouter) delete(c *gin.Context) {
	fmt.Println("delete 입니다.")
}
