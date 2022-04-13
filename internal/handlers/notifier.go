package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/itsoeh/academy-advising-api/internal/model"
	"github.com/itsoeh/academy-advising-api/internal/services"
	"github.com/labstack/echo/v4"
)

// Notifier contains the method to notify subscribers
type Notifier interface {
	// Notify method that is responsible for notifying the message
	Notify(c echo.Context) error
	// HandlerUpdateAdvisory http controller that will receive as a request
	// to create update the status of the academic advising
	HandlerUpdateAdvisory(echo.Context) error

}

type notifier struct {
	response *model.Channels
	services services.AdvisoryManager
}

// NewNotifier returns a notifier structure that implements the Notifier interface
func NewNotifier() Notifier {
	return &notifier{
		&model.Channels{
			ResponseTeacherStream: make(model.ResponseTeacherStream),
		},
		services.NewAdvisoryManager(),
	}
}

func (n *notifier) Notify(c echo.Context) (err error) {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Response().Header().Set("Content-Type", "applicationn/stream+json")
	c.Response().Header().Set("Cache-Control", "no-cache")
	c.Response().Header().Set("Connection", "keep-alive")

	defer func() {
		close(n.response.ResponseTeacherStream)
		n.response.ResponseTeacherStream = nil

		log.Println("Client close connection")
	}()


	flusher, ok := c.Response().Writer.(http.Flusher);
	if !ok {
		log.Println("Could not init http.Flusher")
		return c.JSON(http.StatusInternalServerError, model.Map{"error": "An error ocurred on the server."})
	}

	for {
		select {
		case response := <- n.response.ResponseTeacherStream:
			if err := json.NewEncoder(c.Response().Writer).Encode(response); err != nil {
				return c.JSON(http.StatusBadRequest, model.Map{"error": err})
			}
			flusher.Flush()
		case <- c.Request().Context().Done():
			return
		}
	}
}

func (h *notifier) HandlerUpdateAdvisory(c echo.Context) error {
	isAccepted, err := strconv.ParseBool(c.Param("is_accepted"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Map{"error: ": "Check the path param 'is_accepted', it is empty."})
	}

	advisoryId := c.Param("advisory_id")
	if strings.TrimSpace(advisoryId) == "" {
		return c.JSON(http.StatusBadRequest, model.Map{"error: ": "Check the path param 'advisory_id', it is empty."})
	}

	err = h.services.UpdateAdvisoryStatus(isAccepted, advisoryId)

	if _, ok := err.(model.NotFound); ok {
		return c.JSON(http.StatusNotFound, model.Map{"error: ": err.Error()})
	}

	if _, ok := err.(model.InternalServerError); ok {
		return c.JSON(http.StatusInternalServerError, model.Map{"error: ": err.Error()})
	}
	
	h.response.ResponseTeacherStream <-&model.ChannelIsAccepted{
		IsAccepted: isAccepted,
		Message: "Tu asesoría academica ha sido aceptada",
	}

	return c.JSON(http.StatusOK, model.Map{"message: ": "The process has been completed successfully."})
}
