package web

import "github.com/gin-gonic/gin"

func InitWeb() *gin.Engine {
	server := gin.Default()
	usersHandler := &UsersHandler{}

	usersHandler.RegisterRouters(server.Group("/v0.1/api/users"))

	return server
}
