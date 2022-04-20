package handlers

import (
	"net/http"

	"github.com/itsoeh/academy-advising-api/internal/model"
	"github.com/itsoeh/academy-advising-api/internal/services"
	"github.com/labstack/echo/v4"
)

// AdvisoryHandler contains all http handlers to receive requests and responses from academic advising
type AdvisoryHandler interface {
	// CreateAdvisory http handler that you will
	// receive as a request to create a new academic advisory
	CreateAdvisory(services.AdvisoryManager) echo.HandlerFunc
}

type advisoryHandler struct {}

func NewAdvisoryHandler() AdvisoryHandler {
	return &advisoryHandler{}
}

func (*advisoryHandler) CreateAdvisory(services services.AdvisoryManager) echo.HandlerFunc {
	return func(c echo.Context) error {
		advisory := &model.AcademicAdvisory{}

		if err := c.Bind(advisory); err != nil {
			return c.JSON(http.StatusBadRequest, model.Map{"error: ": "The academic advising format is incorrect."})
		}

		err := services.CreateAdvisory(advisory)

		if _, ok := err.(model.StatusBadRequest); ok {
			return c.JSON(http.StatusBadRequest, model.Map{"error: ": err.Error()})
		}

		if _, ok := err.(model.InternalServerError); ok {
			return c.JSON(http.StatusInternalServerError, model.Map{"error: ": err.Error()})
		}

		return c.JSON(http.StatusCreated, model.Map{"message": "Wait for the notification of the professor."})
	}
}
