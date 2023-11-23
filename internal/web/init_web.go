package web

import "github.com/gin-gonic/gin"

func InitWeb() *gin.Engine {
	server := gin.Default()
	usersHandler := InitUserHandler()

	usersHandler.RegisterRouters(server.Group("/v0.1/api/users"))

	return server
}
