package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	DatabaseName = "taskmanager"
	CollectionTask = "tasks"
	CollectionUser = "users"
)

type Task struct {
 ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
 Title       string    `json:"title"`
 Description string    `json:"description"`
 DueDate     time.Time `json:"due_date"`
 Status      string    `json:"status"`
}

type User struct{
	ID				primitive.ObjectID		`bson:"_id"`
	Name		    *string			`json:"name" validate:"required,min=2,max=100"`
	Username		*string			`json:"username" validate:"required,min=2,max=100"`
	Password		*string			`json:"Password" validate:"required,min=6"`
	Email			*string			`json:"email" validate:"email,required"`
	User_type		string			`json:"user_type"`
	Created_at		time.Time		`json:"created_at"`
	Updated_at		time.Time		`json:"updated_at"`
	User_id			string			`json:"user_id"`
}

type TaskRepository interface {
	Create(c context.Context, task *Task) error
	FetchAll(c context.Context) ([]Task, error)
	FetchByTaskID(c context.Context, taskID string) (Task, error)
	Update(c context.Context, taskID string, updatedTask Task) error
	Delete(c context.Context, taskID string) error
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	FindByUsername(c context.Context, usrname string) (User, error)
	Update(c context.Context, userID string) error
}

type TaskUsecase interface {
	Create(c context.Context, task *Task) error
	FetchAll(c context.Context) ([]Task, error)
	FetchByTaskID(c context.Context, taskID string) (Task, error)
	Update(c context.Context, taskID string, updatedTask Task) error
	Delete(c context.Context, taskID string) error
}

type UserUsecase interface {
	Create(c context.Context, user *User) error
	HandleLogin(c context.Context, username *User) (string, string, error)
	Update(c context.Context, userID string) error
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}
