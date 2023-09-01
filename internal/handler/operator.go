package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testProject/pkg"
)

func (h *Handler) createOperator(c *gin.Context) {
	var input pkg.Operator
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := newPassword(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.usecases.Oper.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, idResponse{
		Id: id,
	})
}

func (h *Handler) getAllOperators(c *gin.Context) {
	list, err := h.usecases.Oper.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: list,
	})
}

func (h *Handler) getOperator(c *gin.Context) {
	id, err := h.usecases.Oper.GetById(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, id)
}

func (h *Handler) updateOperator(c *gin.Context) {
	var input pkg.UpdateOperatorInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.usecases.Oper.UpdateById(c.Param("id"), input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok"})
}

func (h *Handler) deleteOperator(c *gin.Context) {
	err := h.usecases.Oper.DeleteById(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
