package controllers

import (
	"net/http"
	domain "task-manger_with_clean_architecture/Domain"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskUsecase domain.TaskUsecase
}

type UserController struct {
	UserUsecase domain.UserUsecase
}

//user controllers
func (uc *UserController) Signup(c *gin.Context){
	var user domain.User

	if err := c.BindJSON(&user); err != nil{
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	result := uc.UserUsecase.Create(c, &user)
	if result != nil{
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: result.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "user registered successfully",
	})
}

func (uc *UserController) Login(c *gin.Context){
	var user domain.User
	if err := c.BindJSON(&user); err != nil{
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	token, refreshToken, err := uc.UserUsecase.HandleLogin(c, &user)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User loged in successfully!", "token":token, "refresh_token":refreshToken})
}

func (uc *UserController) PromoteUser(c *gin.Context){
	var userID = c.Param("id")
	err := uc.UserUsecase.Update(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "User promoted to ADMIN"})
}

// task controllers
func (tc *TaskController) Create(c *gin.Context){
	var task domain.Task

	err := c.ShouldBind(&task)
	if err != nil{
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = tc.TaskUsecase.Create(c, &task)
	if err != nil{

		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Task created successfully",
	})
}

func (u *TaskController) FetchAll(c *gin.Context) {
	tasks, err := u.TaskUsecase.FetchAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (u *TaskController) FetchByTaskID(c *gin.Context) {
	taskID := c.Param("id")

	tasks, err := u.TaskUsecase.FetchByTaskID(c, taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (u *TaskController) Update(c *gin.Context) {
	taskID := c.Param("id")
	var updatedTask domain.Task

	err := c.ShouldBind(&updatedTask)
	if err != nil{
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err = u.TaskUsecase.Update(c, taskID, updatedTask)

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Message Updated Succesfully."})
}

func (u *TaskController) Delete(c *gin.Context) {
	taskID := c.Param("id")

	err := u.TaskUsecase.Delete(c, taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Message Deleted."})
}