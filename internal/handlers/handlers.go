package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)
// Handlers structure that manages the handlers
type handlers struct {
	AdvisoryHandler
	Notifier
}

// NewHandlers returns a handler struct that contains all the handlers from AcademicAdvisory
func NewHandlers() *handlers {
	return  &handlers{
		NewAdvisoryHandler(),
		NewNotifier(),
	}
}

func NewRequest(t testing.TB, method, path string, dataJSON string) *http.Request {
	t.Helper()
	return httptest.NewRequest(method, path, strings.NewReader(dataJSON))
}
