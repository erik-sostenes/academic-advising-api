package handlers

// Handlers structure that manages the handlers
type handlers struct {
	HandlerAdvisory
}

// NewHandlers returns a handler struct that contains all the handlers from AcademicAdvisory
func NewHandlers() *handlers {
	return &handlers{
		NewHandlerAdvisory(),
	}
}
