package web

import "github.com/gin-gonic/gin"

// UserHandler Routers about user.
type UserHandler struct{}

func (u *UserHandler) RegisterRouters(server *gin.Engine) {
	routerGroup := server.Group("/user")
	routerGroup.POST("/signup", u.Signup)
	routerGroup.POST("/login", u.Login)
	routerGroup.POST("/edit", u.Edit)
	routerGroup.GET("/profile", u.Profile)
}

func (u *UserHandler) Signup(ctx *gin.Context) {}

func (u *UserHandler) Login(ctx *gin.Context) {}

func (u *UserHandler) Edit(ctx *gin.Context) {}

func (u *UserHandler) Profile(ctx *gin.Context) {}
