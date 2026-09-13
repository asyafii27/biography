package controllers

import (
	"biography-api/helpers"
	"biography-api/models"
	"biography-api/services"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BiographyController struct {
	service services.BiographyService
}

func NewBiographyController(service services.BiographyService) *BiographyController {
	return &BiographyController{service}
}

func (c *BiographyController) GetAll(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	limitStr := ctx.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	biographies, totalItems, err := c.service.GetAll(page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Error(http.StatusInternalServerError, "Failed to retrieve biographies", err.Error()))
		return
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))

	meta := helpers.PaginationMeta{
		CurrentPage: page,
		PageSize:    limit,
		TotalPages:  totalPages,
		TotalItems:  totalItems,
	}

	ctx.JSON(http.StatusOK, helpers.Pagination(http.StatusOK, "Biographies retrieved successfully", biographies, meta))
}

func (c *BiographyController) GetByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.Error(http.StatusBadRequest, "Invalid ID format", nil))
		return
	}

	biography, err := c.service.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, helpers.Error(http.StatusNotFound, "Biography not found", nil))
		return
	}

	ctx.JSON(http.StatusOK, helpers.Success(http.StatusOK, "Biography retrieved successfully", biography))
}

func (c *BiographyController) Create(ctx *gin.Context) {
	var input models.Biography
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.Error(http.StatusBadRequest, "Invalid input data", err.Error()))
		return
	}

	biography, err := c.service.Create(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Error(http.StatusInternalServerError, "Failed to create biography", err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, helpers.Success(http.StatusCreated, "Biography created successfully", biography))
}

func (c *BiographyController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.Error(http.StatusBadRequest, "Invalid ID format", nil))
		return
	}

	var input models.Biography
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.Error(http.StatusBadRequest, "Invalid input data", err.Error()))
		return
	}

	biography, err := c.service.Update(uint(id), input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Error(http.StatusInternalServerError, "Failed to update biography", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, helpers.Success(http.StatusOK, "Biography updated successfully", biography))
}

func (c *BiographyController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.Error(http.StatusBadRequest, "Invalid ID format", nil))
		return
	}

	err = c.service.Delete(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Error(http.StatusInternalServerError, "Failed to delete biography", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, helpers.Success(http.StatusOK, "Biography deleted successfully", nil))
}
