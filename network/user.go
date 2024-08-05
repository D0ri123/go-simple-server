package network

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-server/service"
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
	userService *service.User
}

func newUserRouter(router *Network, userService *service.User) *userRouter {
	userRouterInit.Do(func() {
		userRouterInstance = &userRouter{
			router:      router,
			userService: userService,
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
	u.userService.Create(nil)

	u.router.okResponse(c, types.NewApiResponse("성공입니다.", 1))
}

func (u *userRouter) get(c *gin.Context) {
	fmt.Println("get 입니다.")
	//u.router.okResponse(c, &types.ApiResponse{
	//	Result:      200,
	//	Description: "성공입니다.",
	//})

	u.router.okResponse(c, &types.GetUserResponse{
		ApiResponse: types.NewApiResponse("성공입니다.", 1),
		Users:       u.userService.Get(),
	})
}

func (u *userRouter) update(c *gin.Context) {
	fmt.Println("update 입니다.")

	u.userService.Update(nil, nil)

	u.router.okResponse(c, &types.UpdateUserResponse{
		ApiResponse: types.NewApiResponse("성공입니다.", 1),
		User: &types.User{
			Name: "변경된 이름",
			Age:  24,
		},
	})
}

func (u *userRouter) delete(c *gin.Context) {
	fmt.Println("delete 입니다.")

	u.userService.Delete(nil)

	u.router.okResponse(c, &types.DeleteUserResponse{
		ApiResponse: types.NewApiResponse("성공적으로 삭제했습니다.", 1),
	})
}
