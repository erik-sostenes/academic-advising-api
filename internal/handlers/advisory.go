package handlers

import (
	"net/http"

	"github.com/itsoeh/academy-advising-api/internal/model"
	"github.com/itsoeh/academy-advising-api/internal/services"
	"github.com/labstack/echo/v4"
)

// HandlerAdvisory contains all http handlers to receive requests and responses from academic advising
type Advisory interface {
	// CreateAdvisory http handler that you will
	// receive as a request to create a new academic advisory
	CreateAdvisory(echo.Context) error
}

type advisory struct {
	services services.AdvisoryManager
}

// NewAdvisory
func NewAdvisory() Advisory {
	return &advisory{
		services: services.NewAdvisoryManager(),
	}
}

func (h *advisory) CreateAdvisory(c echo.Context) error {
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
