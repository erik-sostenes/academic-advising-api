package services

import (
	"github.com/itsoeh/academy-advising-api/internal/model"
	"github.com/itsoeh/academy-advising-api/internal/repository"
)

// AdvisoryManager contains the methods to manage the creation of an advisory,
// and check if the teacher accepts
type AdvisoryManager interface {
	// CreateAdvisory create a new academic advisory
	CreateAdvisory(*model.AcademicAdvisory) error
	// UpdateAdvisoryStatus method that updates the status of academic advisory
	// NOTE: only if the teacher accepts the academic advisory
	UpdateAdvisoryStatus(isAccepted bool, advisoryId string) error
}

type advisoryManager struct {
	advisoryStorage repository.AdvisoryStorage
}

// NewAdvisoryManager returns the AdvisoryManager interface
func NewAdvisoryManager() AdvisoryManager {
	return &advisoryManager{
		advisoryStorage: repository.NewAdvisoryStorage(),
	}
}

func (a *advisoryManager) CreateAdvisory(advisory *model.AcademicAdvisory) error {
	return a.advisoryStorage.InsertAdvisory(advisory)
}

func (a *advisoryManager) UpdateAdvisoryStatus(isAccepted bool, advisoryId string) (err error) {
	// NOTE: if the status is false, academic advisory will be removed
	if !isAccepted {
		err = a.advisoryStorage.DeleteAdvisory(advisoryId)
		return
	}

	return a.advisoryStorage.UpdateAdvisory(isAccepted, advisoryId)
}
