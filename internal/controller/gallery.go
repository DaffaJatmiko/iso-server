package controller

import (
	"github.com/DaffaJatmiko/project-iso/internal/service"
	"github.com/DaffaJatmiko/project-iso/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GalleryController interface {
	CreateGallery(c *gin.Context)
	GetGalleries(c *gin.Context)
	GetGalleryByID(c *gin.Context)
	UpdateGallery(c *gin.Context)
	DeleteGallery(c *gin.Context)
}

type GalleryControllerImpl struct {
	service service.GalleryService
}

func NewGalleryController(service service.GalleryService) *GalleryControllerImpl {
	return &GalleryControllerImpl{service: service}
}

func (g *GalleryControllerImpl) CreateGallery(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, "Failed to upload image")
		return
	}

	err = g.service.CreateGallery(file)
	if err != nil {
		util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Gallery uploaded successfully",
	})
}

func (g *GalleryControllerImpl) GetGalleries(c *gin.Context) {
	galleries, err := g.service.GetGalleries()
	if err != nil {
		util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	util.JSONResponse(c, http.StatusOK, galleries)
}

func (g *GalleryControllerImpl) GetGalleryByID(c *gin.Context) {
	id := c.Param("id")
	galleryID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	gallery, err := g.service.GetGalleryByID(uint(galleryID))
	if err != nil {
		util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"gallery": gallery})
}

func (g *GalleryControllerImpl) UpdateGallery(c *gin.Context) {
	idStr := c.PostForm("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, "Invalid gallery ID")
		return
	}

	file, err := c.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		util.ErrorResponse(c, http.StatusBadRequest, "Failed to upload image")
		return
	}

	err = g.service.UpdateGallery(uint(id), file)
	if err != nil {
		util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Gallery updated successfully"})
}

func (g *GalleryControllerImpl) DeleteGallery(c *gin.Context) {
	id := c.Param("id")

	galleryID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = g.service.DeleteGallery(uint(galleryID))
	if err != nil {
		util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Gallery deleted successfully"})
}
