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

type AwardeeController struct {
	service services.AwardeeService
}

func NewAwardeeController(service services.AwardeeService) *AwardeeController {
	return &AwardeeController{service}
}

func (c *AwardeeController) GetAll(ctx *gin.Context) {
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
		ctx.JSON(http.StatusInternalServerError, helpers.Error(http.StatusInternalServerError, "Failed to retrieve Awardees", err.Error()))
		return
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(limit)))

	meta := helpers.PaginationMeta{
		CurrentPage: page,
		PageSize:    limit,
		TotalPages:  totalPages,
		TotalItems:  totalItems,
	}

	ctx.JSON(http.StatusOK, helpers.Pagination(http.StatusOK, "Awardees retrieved successfully", items, meta))
}

func (c *AwardeeController) GetByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.Error(http.StatusBadRequest, "Invalid ID format", nil))
		return
	}

	item, err := c.service.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, helpers.Error(http.StatusNotFound, "Awardee not found", nil))
		return
	}

	ctx.JSON(http.StatusOK, helpers.Success(http.StatusOK, "Awardee retrieved successfully", item))
}

func (c *AwardeeController) Create(ctx *gin.Context) {
	var input models.Awardee
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.Error(http.StatusBadRequest, "Invalid input data", err.Error()))
		return
	}

	created, err := c.service.Create(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Error(http.StatusInternalServerError, "Failed to create Awardee", err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, helpers.Success(http.StatusCreated, "Awardee created successfully", created))
}

func (c *AwardeeController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.Error(http.StatusBadRequest, "Invalid ID format", nil))
		return
	}

	var input models.Awardee
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.Error(http.StatusBadRequest, "Invalid input data", err.Error()))
		return
	}

	updated, err := c.service.Update(uint(id), input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Error(http.StatusInternalServerError, "Failed to update Awardee", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, helpers.Success(http.StatusOK, "Awardee updated successfully", updated))
}

func (c *AwardeeController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.Error(http.StatusBadRequest, "Invalid ID format", nil))
		return
	}

	err = c.service.Delete(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.Error(http.StatusInternalServerError, "Failed to delete Awardee", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, helpers.Success(http.StatusOK, "Awardee deleted successfully", nil))
}
