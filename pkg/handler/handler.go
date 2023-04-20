package handler

import (
	"github.com/gin-gonic/gin"
	"testProject/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
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
