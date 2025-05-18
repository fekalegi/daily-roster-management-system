package handlers

import (
	"daily-worker-roster-management-system/handlers/response"
	model "daily-worker-roster-management-system/models"
	"daily-worker-roster-management-system/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ShiftHandler struct {
	Usecase usecase.ShiftUsecase
}

func NewShiftHandler(u usecase.ShiftUsecase) *ShiftHandler {
	return &ShiftHandler{Usecase: u}
}

func (h *ShiftHandler) GetShifts(c *gin.Context) {
	shifts, err := h.Usecase.GetShifts()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(c, http.StatusOK, shifts)
}

func (h *ShiftHandler) CreateShift(c *gin.Context) {
	var shift model.Shift
	if err := c.ShouldBindJSON(&shift); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.Usecase.CreateShift(&shift); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Created(c, shift)
}

func (h *ShiftHandler) GetUnassigned(c *gin.Context) {
	shifts, err := h.Usecase.GetUnassignedShift()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(c, http.StatusOK, shifts)
}
