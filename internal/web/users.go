package web

import "github.com/gin-gonic/gin"

type UsersHandler struct{}

func (u *UsersHandler) RegisterRouters(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/sign-up", u.SignUp)
	routerGroup.POST("/sign-in", u.SignIn)
	routerGroup.POST("/edit", u.Edit)
	routerGroup.GET("/profile", u.Profile)
}

func (u *UsersHandler) SignUp(context *gin.Context) {

}

func (u *UsersHandler) SignIn(context *gin.Context) {

}

func (u *UsersHandler) Edit(context *gin.Context) {

}

func (u *UsersHandler) Profile(context *gin.Context) {

}
