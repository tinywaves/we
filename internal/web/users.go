package web

import (
	"errors"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"net/http"
	"wehub/internal/domain"
	"wehub/internal/service"
)

type UsersHandler struct {
	svc                        *service.UsersService
	emailCompiledExpression    *regexp.Regexp
	passwordCompiledExpression *regexp.Regexp
}

func InitUsersHandler(svc *service.UsersService) *UsersHandler {
	const (
		emailRegexpPattern    = `^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`
		passwordRegexpPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
	)
	return &UsersHandler{
		svc:                        svc,
		emailCompiledExpression:    regexp.MustCompile(emailRegexpPattern, regexp.None),
		passwordCompiledExpression: regexp.MustCompile(passwordRegexpPattern, regexp.None),
	}
}

func (u *UsersHandler) RegisterRouters(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/sign-up", u.SignUp)
	routerGroup.POST("/sign-in", u.SignIn)
	routerGroup.POST("/edit", u.Edit)
	routerGroup.GET("/profile", u.Profile)
}

func (u *UsersHandler) SignUp(ctx *gin.Context) {
	type Request struct {
		Email             string `json:"email"`
		Password          string `json:"password"`
		ConfirmedPassword string `json:"confirmedPassword"`
	}
	var request Request

	// Bind will parse your json by Request, if failed, return a 4xx error automatically; If succeeded, get data from request.
	if err := ctx.Bind(&request); err != nil {
		return
	}

	// check.
	if request.Password != request.ConfirmedPassword {
		ctx.String(http.StatusOK, "Password and confirmedPassword is different")
		return
	}
	matchedEmail, errEmail := u.emailCompiledExpression.MatchString(request.Email)
	matchedPassword, errPassword := u.passwordCompiledExpression.MatchString(request.Password)
	if errEmail != nil || errPassword != nil {
		ctx.String(http.StatusOK, "System error")
		return
	} else if !matchedEmail {
		ctx.String(http.StatusOK, "Your email format is wrong")
		return
	} else if !matchedPassword {
		ctx.String(http.StatusOK, "Your password format is wrong")
		return
	}

	// call svc to use service to process the request.
	err := u.svc.SignUp(
		ctx,
		domain.User{
			Email:    request.Email,
			Password: request.Password,
		},
	)
	if errors.Is(err, service.ErrorUserDuplicateEmail) {
		ctx.String(http.StatusOK, "Duplicate email address")
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "System error")
		return
	}

	ctx.String(http.StatusOK, "Sign up successful")
}

func (u *UsersHandler) SignIn(context *gin.Context) {

}

func (u *UsersHandler) Edit(context *gin.Context) {

}

func (u *UsersHandler) Profile(context *gin.Context) {

}
