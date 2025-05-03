package handler

import (
	"gpt-service-go/service"
	"net/http"
	"strings"
	"github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
)

type ChatHandler struct {
	openAIService *service.OpenAIService
}
var log *logrus.Logger

type ChatRequest struct {
	Message string `json:"message"`
}

type ChatResponse struct {
	Response string `json:"response"`
}

func NewChatHandler(openAIService *service.OpenAIService, logger *logrus.Logger) *ChatHandler {
	log = logger
	return &ChatHandler{openAIService: openAIService}	
}

func (h *ChatHandler) HandleChat(c echo.Context) error {
	req := new(ChatRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}
	
	if len(strings.TrimSpace(req.Message)) > 1000 {
		return echo.NewHTTPError(http.StatusBadRequest, "Message too long")
	}

	resp, err := h.openAIService.SendMessage(req.Message)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	log.Info("Request processed successfully")
	return c.JSON(http.StatusOK, ChatResponse{Response: resp})	
}
