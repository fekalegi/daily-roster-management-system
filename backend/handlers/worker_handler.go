package handlers

import (
	"daily-worker-roster-management-system/handlers/response"
	"daily-worker-roster-management-system/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WorkerHandler struct {
	Usecase usecase.WorkerUsecase
}

func NewWorkerHandler(u usecase.WorkerUsecase) *WorkerHandler {
	return &WorkerHandler{Usecase: u}
}

func (h *WorkerHandler) GetAll(c *gin.Context) {
	requests, err := h.Usecase.GetAll()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(c, http.StatusOK, requests)
}
