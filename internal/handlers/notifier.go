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
	Notify(chan *model.ChannelIsAccepted) echo.HandlerFunc
	// UpdateAdvisory http controller that will receive as a request
	// to create update the status of the academic advising
	UpdateAdvisory(services.AdvisoryManager, chan <- *model.ChannelIsAccepted) echo.HandlerFunc
}

type notifier struct {}

// NewNotifier returns a notifier structure that implements the Notifier interface
func NewNotifier() Notifier {
	return &notifier{}
}

func (n *notifier) Notify(response chan *model.ChannelIsAccepted) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Response().Header().Set("Content-Type", "text/event-stream")
		c.Response().Header().Set("Cache-Control", "no-cache")
		c.Response().Header().Set("Connection", "keep-alive")
		c.Response().WriteHeader(http.StatusOK)

		defer func() {
			close(response)
			response = nil

			log.Println("Client close connection")
		}()

		flusher, _ := c.Response().Writer.(http.Flusher)

		timeout := time.After(1 * time.Second)

		for {
			select {
			case response := <-response:

				v, errJ := json.Marshal(response)
				if errJ != nil {
					c.Response().Writer.WriteHeader(http.StatusInternalServerError)
					io.WriteString(c.Response().Writer, errJ.Error())
					return 
				}

				id := rand.NewSource(time.Now().Unix()).Int63()

				io.WriteString(c.Response().Writer, fmt.Sprintf("id: %v\nevent: handshake\ndata: %v", id, string(v)))
				io.WriteString(c.Response().Writer, "\n\n")

				flusher.Flush()
			case <-c.Request().Context().Done():
				return
			case <-timeout:
			}
		}
	}
}

func (*notifier) UpdateAdvisory(services services.AdvisoryManager, response chan <- *model.ChannelIsAccepted) echo.HandlerFunc {
	return func(c echo.Context) error {
		isAccepted, err := strconv.ParseBool(c.Param("is_accepted"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, model.Map{"error: ": "please verify that the value of the ´is_accepte' field is correct."})
		}

		advisoryId := c.Param("advisory_id")
		if strings.TrimSpace(advisoryId) == "" {
			return c.JSON(http.StatusBadRequest, model.Map{"error: ": "please verify that the value of the ´advisory_id' field is correct."})
		}

		teacherScheduleId := c.Param("teacher_schedule_id")
		if strings.TrimSpace(teacherScheduleId) == "" {
			return c.JSON(http.StatusBadRequest, model.Map{"error: ": "please verify that the value of the ´teacher_schedule_id' field is correct."})
		}

		err = services.UpdateAdvisoryStatus(isAccepted, advisoryId, teacherScheduleId)

		if _, ok := err.(model.NotFound); ok {
			return c.JSON(http.StatusNotFound, model.Map{"error: ": err.Error()})
		}

		if _, ok := err.(model.InternalServerError); ok {
			return c.JSON(http.StatusInternalServerError, model.Map{"error: ": err.Error()})
		}

		stream := &model.ChannelIsAccepted{}
		if err := c.Bind(stream); err != nil {
			return c.JSON(http.StatusBadRequest, model.Map{"error: ": err})
		}

		stream.IsAccepted = isAccepted

		response <- stream

		return c.JSON(http.StatusOK, model.Map{"message: ": "The process has been completed successfully."})
	}
}
