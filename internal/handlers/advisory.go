package handlers

import (
	"net/http"
	"github.com/itsoeh/academic-advising-api/internal/services"
)

type AcademicAdvisory interface {
	AddAcademicAdvisory(http.ResponseWriter, *http.Request)
	NotifyTeacher(w http.ResponseWriter, r *http.Request)
}

type academicAdvisory struct {
	services.AcademicAdvisoryAdministrator
}

func NewAcademicAdvisory() AcademicAdvisory {
	return &academicAdvisory{
		services.NewAcademicAdvisingAdministrator(),
	}
}

func (a *academicAdvisory) AddAcademicAdvisory(w http.ResponseWriter, r *http.Request) {

}

func (a *academicAdvisory) NotifyTeacher(w http.ResponseWriter, r *http.Request){

}


