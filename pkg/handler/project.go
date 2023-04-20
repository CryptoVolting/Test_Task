package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"testProject"
)

func (h *Handler) createProject(c *gin.Context) {
	var input testProject.Project
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Proj.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, idResponse{
		Id: id,
	})
}

func (h *Handler) getAllProjects(c *gin.Context) {
	list, err := h.services.Proj.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: list,
	})
}

func (h *Handler) getProject(c *gin.Context) {
	id, err := h.services.Proj.GetById(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, id)
}

func (h *Handler) updateOProject(c *gin.Context) {
	var input testProject.UpdateProjectInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Proj.UpdateById(c.Param("id"), input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok"})
}

func (h *Handler) deleteProject(c *gin.Context) {
	err := h.services.Proj.DeleteById(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) operatorToProject(c *gin.Context) {
	var input testProject.IdOperatorAndProject
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Proj.CreateAssign(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, idResponse{
		Id: id,
	})
}

func (h *Handler) delOperatorToProject(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := h.services.Proj.DeleteByIdAssign(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
