package network

import "github.com/gin-gonic/gin"

// API에 대해 라우터를 설정하는 부분이 될 것이다.
type Network struct {
	engine *gin.Engine
}

func NewNetwork() *Network {
	// mux,echo,gin.. 프레임워크가 있는데, gin을 많이 주로 사용한다.
	r := &Network{
		engine: gin.New(), //New는 실제 프로덕트 환경에서, Default는 개발 환경에서
	}

	newUserRouter(r)
	return r
}

func (n *Network) ServerStart(port string) error {
	return n.engine.Run(port)
}

// register 유틸 함수들
func (n *Network) registerGET(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.GET(path, handler...)
}

func (n *Network) registerPOST(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.POST(path, handler...)
}

func (n *Network) registerUPDATE(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.PUT(path, handler...)
}

func (n *Network) registerDELETE(path string, handler ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.DELETE(path, handler...)
}
