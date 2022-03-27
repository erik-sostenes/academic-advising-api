package handlers

import (
	"net/http"
	"github.com/itsoeh/academic-advising-api/internal/model"
	"github.com/itsoeh/academic-advising-api/internal/services"
)

// AcademicAdvisory contains the methods to follow the notification flow when adding an advisory
type AcademicAdvisory interface {
	// AddAcademicAdvisory handler that receives the new advisory and will take care of adding
	// Note: only if the teacher agrees
	AddAcademicAdvisory(http.ResponseWriter, *http.Request)

	UpdateAcademicAdvisory(http.ResponseWriter, *http.Request)
}

type academicAdvisory struct {
	services services.AcademicAdvisoryAdministrator
	channels *model.Channels
}

// NewAcademicAdvisory implements the AcademicAdvisory interface
func NewAcademicAdvisory() AcademicAdvisory {
	return &academicAdvisory{
		services: services.NewAcademicAdvisingAdministrator(),
		channels: &model.Channels{
			ResponseTeacherStream: make(model.ResponseTeacherStream),
			NotifyTeacherStream:   make(model.NotifyTeacherStream),
		},
	}
}

func (a *academicAdvisory) AddAcademicAdvisory(w http.ResponseWriter, r *http.Request) {

}

func (a *academicAdvisory) UpdateAcademicAdvisory(w http.ResponseWriter, r *http.Request) {

}
