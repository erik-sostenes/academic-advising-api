package services

import (
	"github.com/itsoeh/academic-advising-api/internal/model"
	"github.com/itsoeh/academic-advising-api/internal/repository"
)

// AcademicAdvisoryAdministrator contains the method for administer the academic advisory 
type AcademicAdvisoryAdministrator interface {
	// AdministerAcademicAdvisory method to validate the input filds and submit data a the repository.
	// Note: Only the advise is saved, if the teacher accepts the advise
	AdministerAcademicAdvisory(*model.AcademicAdvisory) error
}

type academicAdvisoryAdministrator struct {
	repository.AdvisoryAggregator	
}

func NewAcademicAdvisingAdministrator() AcademicAdvisoryAdministrator {
	return &academicAdvisoryAdministrator{
		repository.NewAdvisoryAggregator(),
	}
}

func (a *academicAdvisoryAdministrator) AdministerAcademicAdvisory(advisory *model.AcademicAdvisory) error {
	return a.AdvisoryAggregator.AddAdvisory(advisory)
}
