package web

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
	"wehub/internal/repository"
	"wehub/internal/repository/dao"
	"wehub/internal/service"
)

func initDatabase() *gorm.DB {
	database, err := gorm.Open(mysql.Open("root:root@tcp(localhost:13316)/wehub"))
	if err != nil {
		panic(err)
	}
	err = dao.InitTables(database)
	if err != nil {
		panic(err)
	}
	return database
}

func initUser(database *gorm.DB, server *gin.Engine) {
	usersDao := dao.InitUsersDao(database)
	usersRepository := repository.InitUsersRepository(usersDao)
	usersService := service.InitUsersService(usersRepository)
	usersHandler := InitUsersHandler(usersService)
	usersHandler.RegisterRouters(server.Group("/v0.1/api/users"))
}

func InitWeb() *gin.Engine {
	database := initDatabase()

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

	// users.
	initUser(database, server)

	return server
}
