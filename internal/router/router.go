package router

import (
	"github.com/DaffaJatmiko/project-iso/internal/controller"
	"github.com/DaffaJatmiko/project-iso/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, documentController controller.DocumentController, userController controller.UserController, jwtSecret string) {
	// Public Routes
	r.POST("/admin/register", userController.Register)
	r.POST("/admin/login", userController.Login)
	r.POST("/admin/request-password-reset", userController.RequestPasswordReset)
	r.POST("/admin/reset-password", userController.ResetPassword)

	r.GET("/api/documents", documentController.GetDocuments)

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware(jwtSecret))
	api.POST("/document", documentController.CreateDocument)
	// api.GET("/documents", documentController.GetDocuments)
	api.GET("/document/:id", documentController.GetDocumentByID)
	api.PUT("/document", documentController.UpdateDocument)
	api.DELETE("/document/:id", documentController.DeleteDocument)
	api.POST("/logout", userController.Logout)

}
