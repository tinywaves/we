package web

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func InitWeb() *gin.Engine {
	server := gin.Default()

	// CORS control.
	server.Use(cors.New(cors.Config{
		AllowHeaders: []string{"Content-Type", "Authorization"},
		// AllowOrigins:     []string{"http://localhost:3000"},
		// AllowMethods:     []string{"PUT", "PATCH"}, // Default value is simple methods (GET, POST, PUT, PATCH, DELETE, HEAD, and OPTIONS).
		AllowCredentials: true, // Allow link cookie and so on.
		AllowOriginFunc: func(origin string) bool {
			// Development.
			if strings.HasPrefix(origin, "http://localhost") {
				return true
			}
			return strings.Contains(origin, "personal website url")
		},
		MaxAge: 12 * time.Hour,
	}))

	// user
	InitUserHandler().RegisterRouters(server)

	return server
}
