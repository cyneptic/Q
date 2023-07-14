package controller

import (
	"net/http"

	"github.com/cyneptic/letsgo_smspanel_mockapi/controller/validators"
	"github.com/cyneptic/letsgo_smspanel_mockapi/internal/core/entities"
	"github.com/cyneptic/letsgo_smspanel_mockapi/internal/core/ports"
	service "github.com/cyneptic/letsgo_smspanel_mockapi/internal/core/services"
	"github.com/labstack/echo/v4"
)

type MessageHandler struct {
	svc ports.MessageServiceContracts
}

func NewMessageHnadler() *MessageHandler {
	svc := service.NewMessageService()
	return &MessageHandler{
		svc: svc,
	}
}

func AddMessageRoutes(e *echo.Echo) {
	messageHandler := NewMessageHnadler()
	e.POST("/create-message", messageHandler.SaveMessage)
	e.GET("/give-messages", messageHandler.GiveAllMessages)
}

func (h *MessageHandler) SaveMessage(c echo.Context) error {
	n := new(entities.DTOMessage)

	err := c.Bind(&n)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	err = validators.SaveMessageValidator(*n)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	newMessage := &entities.Message{
		Content:  n.Content,
		Receiver: n.Receiver,
		Sender:   n.Sender,
	}

	err = h.svc.SaveMessageDB(*newMessage)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "message saved successfully")
}
func (h *MessageHandler) GiveAllMessages(c echo.Context) error {
	messages, err := h.svc.GiveAllMessages()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, messages)
}
