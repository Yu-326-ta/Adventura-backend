package controller

import (
	"echo-todo-api/model"
	"echo-todo-api/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ITaskController interface {
	GetAllTasks(c echo.Context) error
	GetTaskById(c echo.Context) error
	CreateTask(c echo.Context) error
	UpdateTask(c echo.Context) error
	DeleteTask(c echo.Context) error
}

type taskController struct {
	tu usecase.ITaskusecase
}

func NewTaskController(tu usecase.ITaskusecase) ITaskController {
	return &taskController{tu}
}

func (tc *taskController) GetAllTasks(c echo.Context) error {
	// userから送られてきたjwttokenをrouterでデコードし、echoのcontext内にuserというフィールドで格納されているので、そこからuserIdを取得
	user := c.Get("user").(jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	// GetAlltasksの引数はuint型なので（取得したUserIdはany型）float64に変換し、unitにアサーション
	taskRes, err := tc.tu.GetAllTasks(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)
}

func (tc *taskController) GetTaskById(c echo.Context) error {
	user := c.Get("user").(jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	// リクエストパラメータからtaskIdを取得(取得した値はstring型なのでAtoiでintに変換)
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)
	taskRes, err := tc.tu.GetTaskById(uint(userId.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)
}

func (tc *taskController) CreateTask(c echo.Context) error {
	user := c.Get("user").(jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	task := model.Task{}
	// c.Bindでリクエストボディの内容をtask構造体に変換し、アドレス先に格納
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	task.UserId = uint(userId.(float64))
	taskRes, err := tc.tu.CreateTask(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, taskRes)
}

func (tc *taskController) UpdateTask(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)

	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	taskRes, err := tc.tu.UpdateTask(task, uint(userId.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)
}

func (tc *taskController) DeleteTask(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)

	err := tc.tu.DeleteTask(uint(userId.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
