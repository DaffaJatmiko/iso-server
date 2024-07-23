package controller

import (
	"github.com/DaffaJatmiko/project-iso/internal/model"
	"github.com/DaffaJatmiko/project-iso/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type AuditController interface {
	CreateAudit(ctx *gin.Context)
	CreateKesesuaian(ctx *gin.Context)
	CreateAuditWithKesesuaian(ctx *gin.Context)
	GetAuditByID(ctx *gin.Context)
	GetAllAudits(ctx *gin.Context)
	DeleteAudit(ctx *gin.Context)
	CalculatePersentaseKesesuaianDokumen(ctx *gin.Context)
	CalculatePersentaseKesesuaianPerKategori(ctx *gin.Context)
	CalculatePersentaseKesesuaianPerPoinAudit(ctx *gin.Context)
}

type auditController struct {
	service service.AuditService
}

func NewAuditController(service service.AuditService) AuditController {
	return &auditController{service: service}
}

func (c *auditController) CreateAudit(ctx *gin.Context) {
	var audit model.Audit
	if err := ctx.ShouldBindJSON(&audit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateAudit(&audit); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, audit)
}

func (c *auditController) CreateKesesuaian(ctx *gin.Context) {
	var kesesuaian model.Kesesuaian
	if err := ctx.ShouldBindJSON(&kesesuaian); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fetch the associated Audit to ensure it's loaded
	audit, err := c.service.GetAuditByID(kesesuaian.AuditID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	kesesuaian.Audit = *audit

	if err := c.service.CreateKesesuaian(&kesesuaian); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, kesesuaian)
}

func (c *auditController) CreateAuditWithKesesuaian(ctx *gin.Context) {
	var audit model.Audit
	if err := ctx.ShouldBindJSON(&audit); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdAudit, err := c.service.CreateAuditWithKesesuaian(&audit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdAudit)
}

func (c *auditController) GetAuditByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid audit ID"})
		return
	}

	audit, err := c.service.GetAuditByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, audit)
}

func (c *auditController) GetAllAudits(ctx *gin.Context) {
	audits, err := c.service.GetAllAudits()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, audits)
}

func (c *auditController) DeleteAudit(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid audit ID"})
	}

	audit, err := c.service.GetAuditByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Audit not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = c.service.DeleteAudit(audit.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Delete audit successfully"})
}

func (c *auditController) CalculatePersentaseKesesuaianDokumen(ctx *gin.Context) {
	idStr := ctx.Param("auditID")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid audit ID"})
		return
	}

	persentase, err := c.service.CalculatePersentaseKesesuaianDokumen(uint(id))
	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Audit not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"persentase": persentase})
}

func (c *auditController) CalculatePersentaseKesesuaianPerKategori(ctx *gin.Context) {
	kategori := ctx.Param("kategori")
	persentase, err := c.service.CalculatePersentaseKesesuaianPerKategori(kategori)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, persentase)
}

func (c *auditController) CalculatePersentaseKesesuaianPerPoinAudit(ctx *gin.Context) {
	poinAudit := ctx.Param("poin_audit")
	persentase, err := c.service.CalculatePersentaseKesesuaianPerPoinAudit(poinAudit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, persentase)
}
