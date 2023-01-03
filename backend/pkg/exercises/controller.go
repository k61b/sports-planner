package exercises

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kayraberktuncer/sports-planner/pkg/common/models"
)

type ExerciseRequestBody struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (h Handlers) GetExercises(ctx *gin.Context) {
	var exercises []models.Exercise

	if result := h.DB.Find(&exercises); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &exercises)
}

func (h Handlers) GetExercise(ctx *gin.Context) {
	id := ctx.Param("id")
	var exercise models.Exercise

	if result := h.DB.First(&exercise, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &exercise)
}

func (h Handlers) AddExercise(ctx *gin.Context) {
	body := ExerciseRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var exercise models.Exercise

	exercise.Title = body.Title
	exercise.Content = body.Content

	if result := h.DB.Create(&exercise); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &exercise)
}

func (h Handlers) UpdateExercise(ctx *gin.Context) {
	id := ctx.Param("id")
	body := ExerciseRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var exercise models.Exercise

	if result := h.DB.First(&exercise, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	exercise.Title = body.Title
	exercise.Content = body.Content

	h.DB.Save(&exercise)

	ctx.JSON(http.StatusOK, &exercise)
}

func (h Handlers) DeleteExercise(ctx *gin.Context) {
	id := ctx.Param("id")

	var exercise models.Exercise

	if result := h.DB.First(&exercise, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&exercise)

	ctx.Status(http.StatusOK)
}
