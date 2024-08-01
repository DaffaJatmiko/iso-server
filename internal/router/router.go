package router

import (
	"github.com/DaffaJatmiko/project-iso/internal/controller"
	"github.com/DaffaJatmiko/project-iso/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	r *gin.Engine,
	documentController controller.DocumentController,
	userController controller.UserController,
	galleryController controller.GalleryController,
	auditController controller.AuditController,
	jwtSecret string,
) {
	// Public Routes
	r.POST("/admin/register", userController.Register)
	r.POST("/admin/login", userController.Login)
	r.POST("/admin/request-password-reset", userController.RequestPasswordReset)
	r.POST("/admin/reset-password", userController.ResetPassword)

	r.GET("/api/documents", documentController.GetDocuments)
	r.GET("/api/galleries", galleryController.GetGalleries)

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware(jwtSecret))

	// Document Routes
	api.POST("/document", documentController.CreateDocument)
	// api.GET("/documents", documentController.GetDocuments)
	api.GET("/document/:id", documentController.GetDocumentByID)
	api.PUT("/document", documentController.UpdateDocument)
	api.DELETE("/document/:id", documentController.DeleteDocument)
	api.POST("/logout", userController.Logout)

	// Gallery Routes
	api.POST("/gallery", galleryController.CreateGallery)
	api.GET("/gallery/:id", galleryController.GetGalleryByID)
	api.PUT("/gallery", galleryController.UpdateGallery)
	api.DELETE("/gallery/:id", galleryController.DeleteGallery)

	// Audit Routes
	api.POST("/audit", auditController.CreateAudit)
	api.POST("/audit/kesesuaian", auditController.CreateKesesuaian)
	api.POST("/audit/audit-with-kesesuaian", auditController.CreateAuditWithKesesuaian)
	api.GET("/audit/:id", auditController.GetAuditByID)
	api.GET("/audits", auditController.GetAllAudits)
	api.PUT("/audit/:id", auditController.UpdateAudit)
	api.DELETE("/audit/:id", auditController.DeleteAudit)
	api.GET("/audit/persentase-dokumen/:auditID", auditController.CalculatePersentaseKesesuaianDokumen)
	api.GET("/audit/persentase-kategori/:kategori", auditController.CalculatePersentaseKesesuaianPerKategori)
	api.GET("/audit/persentase-poin-audit/:poin_audit", auditController.CalculatePersentaseKesesuaianPerPoinAudit)
	api.GET("/audit/persentase/poin/:poinAudit/kategori/:kategori", auditController.CalculatePersentaseKesesuaianPerPoinAuditPerKategori)

}
