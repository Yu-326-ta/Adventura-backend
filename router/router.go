package router

import (
	"echo-todo-api/controller"

	"github.com/labstack/echo/v4"
)

func NerRouter(uc controller.IUserController) *echo.Echo {
	// echoのインスタンスを作成
	e := echo.New()
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	return e
}
