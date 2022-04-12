package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/itsoeh/academic-advising-api/internal/model"
	"github.com/itsoeh/academic-advising-api/internal/services"
	"github.com/labstack/echo/v4"
)

// HandlersAdvisory contains all http handlers to receive requests and responses from academic advising
type HandlersAdvisory interface {
	// HandlerCreateAdvisory http handler that you will
	// receive as a request to create a new academic advisory
	HandlerCreateAdvisory(echo.Context) error
	// HandlerUpdateAdvisory http controller that will receive as a request
	// to create update the status of the academic advising
	HandlerUpdateAdvisory(echo.Context) error
}

type handlerAdvisory struct {
	services services.AdvisoryManager
}

// NewHandlersAdvisory
func NewHandlersAdvisory() HandlersAdvisory {
	return &handlerAdvisory{
		services: services.NewAdvisoryManager(),
	}
}

func (h *handlerAdvisory) HandlerCreateAdvisory(c echo.Context) error {
	advisory := &model.AcademicAdvisory{}

	if err := c.Bind(advisory); err != nil {
		return c.JSON(http.StatusBadRequest, model.Map{"error: ": "The academic advising format is incorrect."})
	}

	err := h.services.CreateAdvisory(advisory)

	if _, ok := err.(model.StatusBadRequest); ok {
		return c.JSON(http.StatusBadRequest, model.Map{"error: ": err.Error()})
	}

	if _, ok := err.(model.InternalServerError); ok {
		return c.JSON(http.StatusInternalServerError, model.Map{"error: ": err.Error()})
	}

	return c.JSON(http.StatusCreated, model.Map{"message": "Wait for the notification of the professor."})
}

func (a *handlerAdvisory) HandlerUpdateAdvisory(c echo.Context) error {
	isAcepted, err := strconv.ParseBool(c.Param("is_acepted"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Map{"error: ": "Check the path param 'is_acepted', it is empty."})
	}

	advisoryId := c.Param("advisory_id")
	if strings.TrimSpace(advisoryId) == "" {
		return c.JSON(http.StatusBadRequest, model.Map{"error: ": "Check the path param 'advisory_id', it is empty."})
	}

	err = a.services.UpdateAdvisoryStatus(isAcepted, advisoryId)

	if _, ok := err.(model.NotFound); ok {
		return c.JSON(http.StatusNotFound, model.Map{"error: ": err.Error()})
	}

	if _, ok := err.(model.InternalServerError); ok {
		return c.JSON(http.StatusInternalServerError, model.Map{"error: ": err.Error()})
	}

	return c.JSON(http.StatusOK, model.Map{"message: ": "The process has been completed successfully."})
}
