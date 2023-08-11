package main

import (
	"echo-todo-api/controller"
	"echo-todo-api/db"
	"echo-todo-api/repository"
	"echo-todo-api/router"
	"echo-todo-api/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRipository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NerRouter(userController)
	// e.Start()でサーバースタートしてe.Loggerでエラー時にログ情報を出力してプログラム強制終了
	e.Logger.Fatal(e.Start(":8080"))
}
