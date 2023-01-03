package exercises

import (
	"github.com/gin-gonic/gin"
	"github.com/kayraberktuncer/sports-planner/pkg/common/models"
)

type Handlers struct {
	DB models.Store
}

func RegisterRoutes(router *gin.Engine, db models.Store) {
	h := &Handlers{
		DB: db,
	}

	routes := router.Group("/exercises")
	routes.POST("/", h.AddExercise)
	routes.GET("/", h.GetExercises)
	routes.GET("/:id", h.GetExercise)
	routes.PUT("/:id", h.UpdateExercise)
	routes.DELETE("/:id", h.DeleteExercise)
}
