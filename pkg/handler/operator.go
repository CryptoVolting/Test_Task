package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testProject"
)

func (h *Handler) createOperator(c *gin.Context) {
	var input testProject.Operator
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := newPassword(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.services.Oper.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, idResponse{
		Id: id,
	})
}

func (h *Handler) getAllOperators(c *gin.Context) {
	list, err := h.services.Oper.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: list,
	})
}

func (h *Handler) getOperator(c *gin.Context) {
	id, err := h.services.Oper.GetById(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, id)
}

func (h *Handler) updateOperator(c *gin.Context) {
	var input testProject.UpdateOperatorInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Oper.UpdateById(c.Param("id"), input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok"})
}

func (h *Handler) deleteOperator(c *gin.Context) {
	err := h.services.Oper.DeleteById(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
