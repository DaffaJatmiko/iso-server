package main

import (
	"log"
	"os"
	"time"

	"github.com/DaffaJatmiko/project-iso/config"
	"github.com/DaffaJatmiko/project-iso/internal/controller"
	"github.com/DaffaJatmiko/project-iso/internal/db"
	"github.com/DaffaJatmiko/project-iso/internal/model"
	"github.com/DaffaJatmiko/project-iso/internal/repository"
	"github.com/DaffaJatmiko/project-iso/internal/router"
	"github.com/DaffaJatmiko/project-iso/internal/service"
	"github.com/DaffaJatmiko/project-iso/pkg/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load Configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	// Log the configuration
	log.Printf("Loaded configuration: %+v", cfg)

	// Initialize Database
	database, err := db.InitDB(cfg.Database.DSN)
	if err != nil {
		panic(err)
	}

	// Initialize Redis
	//redisClient := db.InitRedis(cfg.Redis.Addr, cfg.Redis.Password, cfg.Redis.DB)

	// Auto Migrate the database
	if err := database.AutoMigrate(&model.Document{}, &model.Admin{}); err != nil {
		log.Fatalf("failed to auto migrate database: %v", err)
	}

		// Seed the database
		db.Seed(database)

	// Initialize repositories services, and controllers
	documentRepository := repository.NewDocumentRepository(database)
	documentService := service.NewDocumentService(documentRepository)
	documentController := controller.NewDocumentController(documentService)

	adminRepository := repository.NewUserRepository(database)
	smtpConfig := util.SMTPConfig{
		Host:     cfg.SMTP.Host,
		Port:     cfg.SMTP.Port,
		Username: cfg.SMTP.Username,
		Password: cfg.SMTP.Password,
		From:     cfg.SMTP.From,
	}
	adminService := service.NewUserService(adminRepository, cfg.JWT.SecretKey, smtpConfig)
	adminController := controller.NewUserController(adminService)

	// Ensure the upload directory exists
	if _, err := os.Stat("uploads"); os.IsNotExist(err) {
		if err := os.Mkdir("uploads", 0777); err != nil {
			log.Fatalf("failed to create upload directory: %v", err)
		}
	}

	// Initialize Gin Router
	r := gin.Default()

	// Set up CORS middleware with specific configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Static("/uploads", "./uploads")
	r.Static("/static", "./static")
	r.GET("/upload", func(c *gin.Context) {
		c.File("./static/upload.html")
	})
	router.SetupRoutes(r, documentController, adminController, cfg.JWT.SecretKey)

	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
