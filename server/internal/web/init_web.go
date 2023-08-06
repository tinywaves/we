package web

import "github.com/gin-gonic/gin"

func InitWeb() *gin.Engine {
	server := gin.Default()
	// user
	userHandler := &UserHandler{}
	userHandler.RegisterRouters(server)

	return server
}
