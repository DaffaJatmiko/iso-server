package controller

import (
	"github.com/DaffaJatmiko/project-iso/internal/model"
	"github.com/DaffaJatmiko/project-iso/internal/service"
	"github.com/DaffaJatmiko/project-iso/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DocumentController interface {
	CreateDocument(c *gin.Context)
	GetDocuments(c *gin.Context)
	GetDocumentByID(c *gin.Context)
	UpdateDocument(c *gin.Context)
	DeleteDocument(c *gin.Context)
}

type DocumentControllerImpl struct {
	service service.DocumentService
}

func NewDocumentController(service service.DocumentService) *DocumentControllerImpl {
	return &DocumentControllerImpl{service: service}
}
func (d *DocumentControllerImpl) CreateDocument(c *gin.Context) {
	// Bind form fields manually
	document := model.Document{
		FileName:    c.PostForm("file_name"),
		Link:        c.PostForm("link"),
		Description: c.PostForm("description"),
	}

	// Retrieve the file
	file, err := c.FormFile("image")
	if err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, "Failed to upload image")
		return
	}

	// Call the service layer to create the document
	err = d.service.CreateDocument(&document, file)
	if err != nil {
		util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Document uploaded successfully",
		"document": document,
	})
}

func (d *DocumentControllerImpl) GetDocuments(c *gin.Context) {
	documents, err := d.service.GetDocuments()
	if err != nil {
		util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	util.JSONResponse(c, http.StatusOK, documents)
}

func (d *DocumentControllerImpl) GetDocumentByID(c *gin.Context) {
	id := c.Param("id")
	documentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	document, err := d.service.GetDocumentByID(uint(documentID))
	if err != nil {
		util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"document": document})
}

func (d *DocumentControllerImpl) UpdateDocument(c *gin.Context) {
	var document model.Document

	// Bind form fields manually
	idStr := c.PostForm("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		util.ErrorResponse(c, http.StatusBadRequest, "Invalid document ID")
		return
	}
	document.ID = uint(id) // Set the ID correctly

	document.FileName = c.PostForm("file_name")
	document.Link = c.PostForm("link")
	document.Description = c.PostForm("description")

	// Retrieve the file
	file, err := c.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		util.ErrorResponse(c, http.StatusBadRequest, "Failed to upload image")
		return
	}

	// Call the service layer to update the document
	err = d.service.UpdateDocument(&document, file)
	if err != nil {
		util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, model.NewSuccessResp("Update Document Success"))
}

func (d *DocumentControllerImpl) DeleteDocument(c *gin.Context) {
	id := c.Param("id")

	documentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = d.service.DeleteDocument(uint(documentID))
	if err != nil {
		util.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, model.NewSuccessResp("Delete Document Success"))
}
