package handlers

// Hadlers structure that manages the handlers
type Hadlers struct {
	AcademicAdvisory
}

// NewHandlers returns a handlers struct that contains all the handlers from AcademicAdvisory
func NewHandlers() Hadlers{
	return Hadlers{
		NewAcademicAdvisory(),
	}
}  
