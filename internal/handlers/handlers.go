package handlers

// Hadlers structure that manages the handlers
type hadlers struct {
	HandlersAdvisory
}

// NewHandlers returns a handlers struct that contains all the handlers from AcademicAdvisory
func NewHandlers() *hadlers{
	return &hadlers{
		NewHandlersAdvisory(),
	}
}  
