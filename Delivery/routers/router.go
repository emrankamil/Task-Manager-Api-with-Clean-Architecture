package routers

import (
	"task-manger_with_clean_architecture/Delivery/controllers"
	domain "task-manger_with_clean_architecture/Domain"
	infrastructure "task-manger_with_clean_architecture/Infrastructure"
	repositories "task-manger_with_clean_architecture/Repositories"
	usecases "task-manger_with_clean_architecture/Usecases"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(timeout time.Duration, db *mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	PublicUserRouter(timeout, db, publicRouter)
	PublicTaskRouter(timeout, db, publicRouter)

	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(infrastructure.AuthMiddleware())
	// All Private APIs
	PrivateTaskRouter(timeout, db, protectedRouter)
	PromoteRouter(timeout, db, protectedRouter)
}


func PublicTaskRouter(timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	taskRepo := repositories.NewTaskRepository(db, domain.CollectionTask)
	taskUsecase := usecases.NewTaskUsecase(taskRepo, timeout)
	taskController := &controllers.TaskController{
		TaskUsecase : taskUsecase,
	}

	group.GET("/tasks", taskController.FetchAll)
	group.GET("/tasks/:id", taskController.FetchByTaskID)
}

func PrivateTaskRouter(timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	taskRepo := repositories.NewTaskRepository(db, domain.CollectionTask)
	taskUsecase := usecases.NewTaskUsecase(taskRepo, timeout)
	taskController := &controllers.TaskController{
		TaskUsecase : taskUsecase,
	}

	group.POST("/tasks", taskController.Create)
	group.PUT("/tasks/:id", taskController.Update)
	group.DELETE("/tasks/:id", taskController.Delete)
}

func PublicUserRouter(timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	userRepo := repositories.NewUserRepository(db, domain.CollectionUser)
	userUsecase := usecases.NewUserUsecase(userRepo, timeout)
	userController := &controllers.UserController{
		UserUsecase: userUsecase,
	}

	group.POST("/register", userController.Signup)
	group.POST("/login", userController.Login)
}

func PromoteRouter(timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	userRepo := repositories.NewUserRepository(db, domain.CollectionUser)
	userUsecase := usecases.NewUserUsecase(userRepo, timeout)
	userController := &controllers.UserController{
		UserUsecase: userUsecase,
	}

	group.PUT("/promote/:id", userController.PromoteUser)
}