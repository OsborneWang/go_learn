package routes

import (
	"github.com/gin-gonic/gin"

	"goServer/internal/handlers"
)

// Register 将所有 HTTP 路由注册到 gin Engine。
func Register(r *gin.Engine, authHandler *handlers.AuthHandler, mailHandler *handlers.MailHandler) {
	api := r.Group("/api/v1")
	{
		api.POST("/register", authHandler.Register)
		api.POST("/login", authHandler.Login)

		api.POST("/mail/test", mailHandler.SendTest)
	}
}
