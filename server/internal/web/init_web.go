package web

import "github.com/gin-gonic/gin"

func InitWeb() *gin.Engine {
	server := gin.Default()
	// user
	InitUserHandler().RegisterRouters(server)

	return server
}
