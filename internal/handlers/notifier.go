package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

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
	HandlerUpdateAdvisory(ctx echo.Context) error
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

func (n *notifier) Notify(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Response().Header().Set("Content-Type", "text/event-stream")
	c.Response().Header().Set("Cache-Control", "no-cache")
	c.Response().Header().Set("Connection", "keep-alive")
	c.Response().WriteHeader(http.StatusOK)

	defer func() {
		close(n.response.ResponseTeacherStream)
		n.response.ResponseTeacherStream = nil

		log.Println("Client close connection")
	}()

	flusher, _ := c.Response().Writer.(http.Flusher)

	timeout := time.After(1 * time.Second)

	for {
		select {
		case response := <-n.response.ResponseTeacherStream:

			v, err := json.Marshal(response)
			if err != nil {
				c.Response().Writer.WriteHeader(http.StatusInternalServerError)
				io.WriteString(c.Response().Writer, err.Error())
				break
			}

			id := rand.NewSource(time.Now().Unix()).Int63()

			io.WriteString(c.Response().Writer, fmt.Sprintf("id: %v\nevent: eventSSE\ndata: %v", id, string(v)))
			io.WriteString(c.Response().Writer, "\n\n")

			flusher.Flush()
		case <-c.Request().Context().Done():
			break
		case <-timeout:
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

	message := make(map[string]string)
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, model.Map{"error: ": err})
	}

	h.response.ResponseTeacherStream <- &model.ChannelIsAccepted{
		IsAccepted: isAccepted,
		Message:    message["message"],
	}

	return c.JSON(http.StatusOK, model.Map{"message: ": "The process has been completed successfully."})
}
