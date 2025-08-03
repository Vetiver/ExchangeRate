package handlers

import (
	"exchangeRate/config"
	"exchangeRate/constants"
	"exchangeRate/internal/domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DomainService interface {
	ExchangeRate(date string, val string, config *config.Config) (domain.Record, error)
}

type ExchangeRateResp struct {
	Date      string `json:"date"`
	Nominal   int    `json:"dominal"`
	Value     string `json:"dalue"`
	VunitRate string `json:"vunitRate"`
}

type BaseHandler struct {
	config              *config.Config
	logger              *zap.Logger
	exchangeRateService DomainService
}

func NewBaseHandler(logger *zap.Logger, ExchangeRate DomainService, config *config.Config) *BaseHandler {
	return &BaseHandler{
		config:              config,
		logger:              logger,
		exchangeRateService: ExchangeRate,
	}
}

func (h *BaseHandler) Ping(c *gin.Context) {
	h.logger.Info("Ping request received")
	c.String(http.StatusOK, constants.ServiceOK)
}

func (h *BaseHandler) ExchangeRate(c *gin.Context) {
	h.logger.Info("Get ExchangeRate")
	date := c.Query("date")
	if date == "" {
		date = time.Now().Format("02/01/2006")
	}
	val := c.Query("val")
	if val == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.WrongParamFormat})
		return
	}
	record, err := h.exchangeRateService.ExchangeRate(date, val, h.config)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp := ExchangeRateResp{Date: record.Date, Nominal: record.Nominal, Value: record.Value, VunitRate: record.VunitRate}

	c.JSON(http.StatusOK, resp)
}
