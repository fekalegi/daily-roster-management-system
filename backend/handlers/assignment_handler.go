package handlers

import (
	"daily-worker-roster-management-system/handlers/response"
	model "daily-worker-roster-management-system/models"
	"daily-worker-roster-management-system/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AssignmentHandler struct {
	Usecase usecase.AssignmentUsecase
}

func NewAssignmentHandler(u usecase.AssignmentUsecase) *AssignmentHandler {
	return &AssignmentHandler{Usecase: u}
}

func (h *AssignmentHandler) Create(c *gin.Context) {
	var a model.Assignment
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Usecase.Assign(&a); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Created(c, a)
}

func (h *AssignmentHandler) Get(c *gin.Context) {
	if workerID := c.Query("worker_id"); workerID != "" {
		id, _ := strconv.Atoi(workerID)
		data, err := h.Usecase.GetByWorker(id)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		response.JSON(c, http.StatusOK, data)
		return
	} else if date := c.Query("date"); date != "" {
		data, err := h.Usecase.GetByDate(date)
		if err != nil {
			response.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		response.JSON(c, http.StatusOK, data)
		return
	} else {
		data, err := h.Usecase.FindAll()
		if err != nil {
			response.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		response.JSON(c, http.StatusOK, data)
		return
	}
}

func (h *AssignmentHandler) FindAll(c *gin.Context) {
	data, err := h.Usecase.FindAll()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
	}

	response.JSON(c, http.StatusOK, data)
}
