package web

import (
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserHandler Routers about user.
type UserHandler struct {
	emailRegex    *regexp.Regexp
	passwordRegex *regexp.Regexp
}

func InitUserHandler() *UserHandler {
	const (
		emailRegexPattern    = `^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`
		passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
	)
	emailRegex := regexp.MustCompile(emailRegexPattern, regexp.None)
	passwordRegex := regexp.MustCompile(passwordRegexPattern, regexp.None)
	return &UserHandler{
		emailRegex:    emailRegex,
		passwordRegex: passwordRegex,
	}
}

func (u *UserHandler) RegisterRouters(server *gin.Engine) {
	routerGroup := server.Group("/user")
	routerGroup.POST("/signup", u.Signup)
	routerGroup.POST("/login", u.Login)
	routerGroup.POST("/edit", u.Edit)
	routerGroup.GET("/profile", u.Profile)
}

func (u *UserHandler) Signup(ctx *gin.Context) {
	// Signup request internal struct.
	type SignupRequest struct {
		Email           string `json:"email"`
		ConfirmPassword string `json:"confirmPassword"`
		Password        string `json:"password"`
	}

	var request SignupRequest
	// Bind will parse data depends on Content-Type. If hit error, return a 4xx http error.
	if err := ctx.Bind(&request); err != nil {
		return
	}

	// Test.
	ok, err := u.emailRegex.MatchString(request.Email)
	if err != nil {
		// Regex pattern is wrong.
		ctx.String(http.StatusOK, "system error")
		return
	}
	if !ok {
		// Email format is wrong.
		ctx.String(http.StatusOK, "email format is wrong")
		return
	}
	if request.Password != request.ConfirmPassword {
		ctx.String(http.StatusOK, "password and confirmPassword is not same")
		return
	}
	ok, err = u.passwordRegex.MatchString(request.Password)
	if err != nil {
		ctx.String(http.StatusOK, "system error")
		return
	}
	if !ok {
		ctx.String(http.StatusOK, "password format is wrong")
		return
	}

	// database.

	ctx.String(http.StatusOK, "signup success")
}

func (u *UserHandler) Login(ctx *gin.Context) {}

func (u *UserHandler) Edit(ctx *gin.Context) {}

func (u *UserHandler) Profile(ctx *gin.Context) {}
