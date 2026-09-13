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

type ExperienceController struct {
	service services.ExperienceService
}

func NewExperienceController(service services.ExperienceService) *ExperienceController {
	return &ExperienceController{service}
}

func (c *ExperienceController) GetAll(ctx *gin.Context) {
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

	experiences, totalItems, err := c.service.GetAll(page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Error("Failed to retrieve experiences", err.Error()))
		return
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))

	meta := helpers.PaginationMeta{
		CurrentPage: page,
		PageSize:    limit,
		TotalPages:  totalPages,
		TotalItems:  totalItems,
	}

	ctx.JSON(http.StatusOK, helpers.Pagination("Experiences retrieved successfully", experiences, meta))
}

func (c *ExperienceController) GetByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.Error("Invalid ID format", nil))
		return
	}

	experience, err := c.service.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, helpers.Error("Experience not found", nil))
		return
	}

	ctx.JSON(http.StatusOK, helpers.Success("Experience retrieved successfully", experience))
}

func (c *ExperienceController) Create(ctx *gin.Context) {
	var input models.Experience
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.Error("Invalid input data", err.Error()))
		return
	}

	if input.StartDate != nil && input.EndDate != nil {
		if input.StartDate.After(*input.EndDate) {
			ctx.JSON(http.StatusBadRequest, helpers.Error("start_date cannot be greater than end_date", nil))
			return
		}
	}

	experience, err := c.service.Create(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Error("Failed to create experience", err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, helpers.Success("Experience created successfully", experience))
}

func (c *ExperienceController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.Error("Invalid ID format", nil))
		return
	}

	var input models.Experience
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.Error("Invalid input data", err.Error()))
		return
	}

	if input.StartDate != nil && input.EndDate != nil {
		if input.StartDate.After(*input.EndDate) {
			ctx.JSON(http.StatusBadRequest, helpers.Error("start_date cannot be greater than end_date", nil))
			return
		}
	}

	experience, err := c.service.Update(uint(id), input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Error("Failed to update experience", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, helpers.Success("Experience updated successfully", experience))
}

func (c *ExperienceController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.Error("Invalid ID format", nil))
		return
	}

	err = c.service.Delete(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Error("Failed to delete experience", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, helpers.Success("Experience deleted successfully", nil))
}
