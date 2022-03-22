package handlers

import (
	"net/http"

	"github.com/itsoeh/academic-advising-api/internal/services"
)

// AcademicAdvisory contains the methods to follow the notification flow when adding an advisory
type AcademicAdvisory interface {
	// AddAcademicAdvisory handler that receives the new advisory and will take care of adding
	// Note: only if the teacher agrees
	AddAcademicAdvisory(http.ResponseWriter, *http.Request)
	// NotifyTeacher handler who is in charge of notifying the teacher
	// that a student wants to reserve an advisory
	NotifyTeacher(http.ResponseWriter, *http.Request)
}

type academicAdvisory struct {
	services.AcademicAdvisoryAdministrator
}

// NewAcademicAdvisory implements the AcademicAdvisory interface
func NewAcademicAdvisory() AcademicAdvisory {
	return &academicAdvisory{
		services.NewAcademicAdvisingAdministrator(),
	}
}

func (a *academicAdvisory) AddAcademicAdvisory(w http.ResponseWriter, r *http.Request) {

}

func (a *academicAdvisory) NotifyTeacher(w http.ResponseWriter, r *http.Request) {

}
