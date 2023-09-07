package router

import (
	"echo-todo-api/controller"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func NerRouter(uc controller.IUserController, tc controller.ITaskController) *echo.Echo {
	// echoのインスタンスを作成
	e := echo.New()
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	t := e.Group("/tasks")
	t.Use(echojwt.WithConfig(echojwt.Config{
		// jwtを生成した時と同じシークレットキーをSigningKeyに指定
		SigningKey: []byte(os.Getenv("SECRET")),
		// クライアントから送られてくるjwtトークンの格納場所をTokenLookupに指定
		TokenLookup: "cookie:token",
	}))
	t.GET("", tc.GetAllTasks)
	t.GET("/:taskId", tc.GetTaskById)
	t.POST("", tc.CreateTask)
	t.PUT("/:taskId", tc.UpdateTask)
	t.DELETE("/:taskId", tc.DeleteTask)
	return e
}