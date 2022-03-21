package handlers

type Hadlers struct {
	AcademicAdvisory
}

func NewHandlers() Hadlers{
	return Hadlers{
		NewAcademicAdvisory(),
	}
}  
