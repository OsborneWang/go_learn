package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"goServer/internal/config"
	"goServer/internal/database"
	"goServer/internal/handlers"
	"goServer/internal/models"
	"goServer/internal/routes"
	"goServer/internal/services"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("未加载到 .env 文件，采用系统环境变量: %v", err)
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	db, err := database.New(cfg)
	if err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}

	// 自动迁移模型，确保初次运行即可生成数据表
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}

	mailService := services.NewMailService(cfg)
	authService := services.NewAuthService(db, mailService)
	authHandler := handlers.NewAuthHandler(authService)
	mailHandler := handlers.NewMailHandler(mailService)

	router := gin.Default()
	routes.Register(router, authHandler, mailHandler)

	log.Printf("Server listening on :%s", cfg.ServerPort)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
