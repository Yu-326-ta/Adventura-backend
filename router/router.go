package router

import (
	"echo-todo-api/controller"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NerRouter(uc controller.IUserController, tc controller.ITaskController) *echo.Echo {
	// echoのインスタンスを作成
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// アクセスを許可するフロントのドメインを設定
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		// リクエストヘッダーに含めるパラメータを設定、echo.HeaderXCSRFTokenでcsrfトークンを送信可能に
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		// クッキーの送受信を可能にするための設定
		AllowCredentials: true,
	}))
	// csrfトークンを格納するクーキーの設定
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		// CookieSameSite: http.SameSiteNoneMode,
		// ポストマンで動作確認するためのsamesite
		CookieSameSite: http.SameSiteDefaultMode,
		// csrfトークンの有効期限の設定
		//CookieMaxAge:   60,
	}))
	healthResponse := map[string]string{
		"status": "ok",
		"version": "4",
	}
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, healthResponse)
	})
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)
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