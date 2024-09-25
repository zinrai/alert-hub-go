package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zinrai/alert-hub-go/internal/domain"
	"github.com/zinrai/alert-hub-go/internal/usecase"
)

type AlertHandler struct {
	AlertUsecase usecase.AlertUsecase
}

func NewAlertHandler(u usecase.AlertUsecase) *AlertHandler {
	return &AlertHandler{AlertUsecase: u}
}

func (h *AlertHandler) GetAlerts(c *gin.Context) {
	alerts, err := h.AlertUsecase.GetAlerts(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, alerts)
}

func (h *AlertHandler) CreateAlert(c *gin.Context) {
	var alert domain.Alert
	if err := c.ShouldBindJSON(&alert); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.AlertUsecase.CreateAlert(c, alert); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, alert)
}

func (h *AlertHandler) GetAlert(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid alert ID"})
		return
	}

	alert, err := h.AlertUsecase.GetAlert(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, alert)
}

func (h *AlertHandler) UpdateAlert(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid alert ID"})
		return
	}

	existingAlert, err := h.AlertUsecase.GetAlert(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve existing alert"})
		return
	}

	var updateData struct {
		Subject    *string `json:"subject"`
		Body       *string `json:"body"`
		Identifier *string `json:"identifier"`
		Urgency    *string `json:"urgency"`
		Status     *string `json:"status"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if updateData.Subject != nil {
		existingAlert.Subject = *updateData.Subject
	}
	if updateData.Body != nil {
		existingAlert.Body = *updateData.Body
	}
	if updateData.Identifier != nil {
		existingAlert.Identifier = *updateData.Identifier
	}
	if updateData.Urgency != nil {
		existingAlert.Urgency = *updateData.Urgency
	}
	if updateData.Status != nil {
		existingAlert.Status = *updateData.Status
	}

	if err := h.AlertUsecase.UpdateAlert(c, existingAlert); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, existingAlert)
}
