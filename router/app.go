package router

import (
	"gowebproject1/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	route := gin.Default()
	route.GET("/index", service.GetApi)
	UserRouter(route)
	return route

}

// UserRouter 注册路由
func UserRouter(engine *gin.Engine) {
	// 查找用户路由
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())
	// 查找用户列表
	engine.GET("/userList", service.GetUser)

	// 创建用户
	engine.POST("/user", service.CreateUser)

	// 删除用户
	engine.DELETE("/user/:id", service.DeleteUser)

	// 更新用户
	engine.PUT("/user/:id", service.UpdateUser)
}
