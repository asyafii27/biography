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

type TechnicalExperienceController struct {
	service services.TechnicalExperienceService
}

func NewTechnicalExperienceController(service services.TechnicalExperienceService) *TechnicalExperienceController {
	return &TechnicalExperienceController{service}
}

func (c *TechnicalExperienceController) GetAll(ctx *gin.Context) {
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

	items, totalItems, err := c.service.GetAll(page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Error(http.StatusInternalServerError, "Failed to retrieve TechnicalExperiences", err.Error()))
		return
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))

	meta := helpers.PaginationMeta{
		CurrentPage: page,
		PageSize:    limit,
		TotalPages:  totalPages,
		TotalItems:  totalItems,
	}

	ctx.JSON(http.StatusOK, helpers.Pagination(http.StatusOK, "TechnicalExperiences retrieved successfully", items, meta))
}

func (c *TechnicalExperienceController) GetByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.Error(http.StatusBadRequest, "Invalid ID format", nil))
		return
	}

	item, err := c.service.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, helpers.Error(http.StatusNotFound, "TechnicalExperience not found", nil))
		return
	}

	ctx.JSON(http.StatusOK, helpers.Success(http.StatusOK, "TechnicalExperience retrieved successfully", item))
}

func (c *TechnicalExperienceController) Create(ctx *gin.Context) {
	var input models.TechnicalExperience
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.Error(http.StatusBadRequest, "Invalid input data", err.Error()))
		return
	}

	created, err := c.service.Create(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Error(http.StatusInternalServerError, "Failed to create TechnicalExperience", err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, helpers.Success(http.StatusCreated, "TechnicalExperience created successfully", created))
}

func (c *TechnicalExperienceController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.Error(http.StatusBadRequest, "Invalid ID format", nil))
		return
	}

	var input models.TechnicalExperience
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.Error(http.StatusBadRequest, "Invalid input data", err.Error()))
		return
	}

	updated, err := c.service.Update(uint(id), input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Error(http.StatusInternalServerError, "Failed to update TechnicalExperience", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, helpers.Success(http.StatusOK, "TechnicalExperience updated successfully", updated))
}

func (c *TechnicalExperienceController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.Error(http.StatusBadRequest, "Invalid ID format", nil))
		return
	}

	err = c.service.Delete(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Error(http.StatusInternalServerError, "Failed to delete TechnicalExperience", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, helpers.Success(http.StatusOK, "TechnicalExperience deleted successfully", nil))
}
