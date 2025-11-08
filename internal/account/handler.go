package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateAccount(c *gin.Context) {
	var json struct {
		OwnerID string `json:"owner_id"`
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_account, err := h.service.OpenAccount(json.OwnerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, _account)
}

func (h *Handler) GetAccount(c *gin.Context) {
	_id := c.Param("id")

	_account, err := h.service.GetAccount(_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, _account)
}

func (h *Handler) DepositAmount(c *gin.Context) {
	var json struct {
		AccountID string  `json:"account_id"`
		Amount    float32 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedBalance, err := h.service.DepositAmount(json.AccountID, json.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"updated_balance": updatedBalance})
}

func (h *Handler) WithdrawAmount(c *gin.Context) {
	var json struct {
		AccountID string  `json:"account_id"`
		Amount    float32 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedBalance, err := h.service.WithdrawAmount(json.AccountID, json.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"updated_balance": updatedBalance})
}
