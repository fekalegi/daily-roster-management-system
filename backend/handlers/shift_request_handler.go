package handlers

import (
	"daily-worker-roster-management-system/handlers/response"
	model "daily-worker-roster-management-system/models"
	"daily-worker-roster-management-system/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ShiftRequestHandler struct {
	Usecase usecase.ShiftRequestUsecase
}

func NewShiftRequestHandler(u usecase.ShiftRequestUsecase) *ShiftRequestHandler {
	return &ShiftRequestHandler{Usecase: u}
}

func (h *ShiftRequestHandler) Create(c *gin.Context) {
	var req model.ShiftRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	req.Status = "pending"
	if err := h.Usecase.CreateRequest(&req); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Created(c, req)
}

func (h *ShiftRequestHandler) ListPending(c *gin.Context) {
	requests, err := h.Usecase.GetPendingRequests()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(c, http.StatusOK, requests)
}

func (h *ShiftRequestHandler) Approve(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Usecase.Approve(id); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(c, http.StatusOK, nil)
}

func (h *ShiftRequestHandler) Reject(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Usecase.Reject(id); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(c, http.StatusOK, nil)
}

func (h *ShiftRequestHandler) List(c *gin.Context) {
	status := c.Query("status")

	var results []model.ShiftRequestDetail
	var err error

	if status != "" {
		results, err = h.Usecase.GetByStatus(status)
	} else {
		results, err = h.Usecase.GetAll()
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error_message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"status": "success",
		"data":   results,
	})
}
