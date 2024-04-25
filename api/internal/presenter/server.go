package presenter

import (
	"context"
	"go-rest/api/internal/controller/system"
	"go-rest/api/internal/controller/user"

	"github.com/gin-gonic/gin"
)

const latest = "/v1"

type Server struct{}

func (s *Server) Run(ctx context.Context) error {
	r := gin.Default()
	v1 := r.Group(latest)

	// 死活監視用
	{
		systemHandler := system.NewSystemHandler()
		v1.GET("/health", systemHandler.Health)
	}
	// ユーザー管理機能
	{
		userHandler := user.NewUserHandler()
		v1.GET("", userHandler.GetUsers)
		v1.GET("/:id", userHandler.GetUserByID)
		v1.POST("", userHandler.EditUser)
		v1.DELETE("/:id", userHandler.DeleteUser)
	}
	err := r.Run()
	if err != nil {
		return err
	}
	return nil
}

func NewServer() *Server {
	return &Server{}
}
