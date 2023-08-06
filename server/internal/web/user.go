package web

import "github.com/gin-gonic/gin"

// UserHandler Routers about user.
type UserHandler struct{}

func (u *UserHandler) RegisterRouters(server *gin.Engine) {
	server.POST("/user/signup", u.Signup)
	server.POST("/user/login", u.Login)
	server.POST("/user/edit", u.Edit)
	server.GET("/user/profile", u.Profile)
}

func (u *UserHandler) Signup(ctx *gin.Context) {}

func (u *UserHandler) Login(ctx *gin.Context) {}

func (u *UserHandler) Edit(ctx *gin.Context) {}

func (u *UserHandler) Profile(ctx *gin.Context) {}
