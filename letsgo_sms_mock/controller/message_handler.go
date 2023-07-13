package controller

import (
	"net/http"
	

	"github.com/cyneptic/letsgo_smspanel_mockapi/internal/core/entities"
	"github.com/cyneptic/letsgo_smspanel_mockapi/internal/core/ports"
	"github.com/cyneptic/letsgo_smspanel_mockapi/internal/core/services"
	"github.com/cyneptic/letsgo_smspanel_mockapi/controller/validators"
	"github.com/labstack/echo/v4"
)
type MessageHandler struct {
	svc ports.MessageServiceContracts
}

func NewMessageHnadler() *MessageHandler {
	svc := service.NewMessageService()
	return &MessageHandler{
		svc : svc,
	}
}

func AddMessageRoutes(e *echo.Echo) {
	messageHandler := NewMessageHnadler()
	e.POST("/create-message" , messageHandler.SaveMessage)
	e.GET("/give-messages" , messageHandler.GiveAllMessages)
}
func (h *MessageHandler) SaveMessage(c echo.Context) error {
	newMessage := new(entities.Message)
	

	err := c.Bind(&newMessage); if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}

	err = validators.SaveMessageValidator(*newMessage)

	if err != nil { 
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.svc.SaveMessageDB(*newMessage)
	if err != nil { 
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK , "message saved successfully")
}
func (h *MessageHandler) GiveAllMessages(c echo.Context) error {
	messages , err := h.svc.GiveAllMessages()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK , messages)
}