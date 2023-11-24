package web

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func InitWeb() *gin.Engine {
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "production.com")
		},
		MaxAge: 12 * time.Hour,
	}))

	usersHandler := InitUserHandler()
	usersHandler.RegisterRouters(server.Group("/v0.1/api/users"))

	return server
}
