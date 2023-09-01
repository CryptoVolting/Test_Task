package handler

import (
	"github.com/gin-gonic/gin"
	"testProject/internal/usecase"
)

type Handler struct {
	usecases *usecase.Usecase
}

func NewHandler(usecases *usecase.Usecase) *Handler {
	return &Handler{usecases: usecases}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	operator := router.Group("/operator")
	{
		operator.POST("/new", h.createOperator)
		operator.GET("/list", h.getAllOperators)
		operator.GET("/:id", h.getOperator)
		operator.PUT("/:id", h.updateOperator)
		operator.DELETE("/:id", h.deleteOperator)

	}

	project := router.Group("/project")
	{
		project.POST("/new", h.createProject)
		project.GET("/list", h.getAllProjects)
		project.GET("/:id", h.getProject)
		project.PUT("/:id", h.updateOProject)
		project.DELETE("/:id", h.deleteProject)
		project.POST("/assign", h.operatorToProject)
		project.DELETE("/remove/:id", h.delOperatorToProject)
	}
	return router
}
