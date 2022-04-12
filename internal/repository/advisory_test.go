package repository

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/itsoeh/academy-advising-api/internal/model"
)

var testAcademicAdvisory = &model.AcademicAdvisory{
	AdvisoryId:  "190HY5D",
	Description: "This is test.",
	FromDate:    time.Now(),
	ToDate:      time.Now().AddDate(0, 2, 7),
	IsActive:    false,
	AcademicAdvisoryIds: model.AcademicAdvisoryIds{
		SubjectId:             1,
		StudentTuition:        1,
		TeacherTuition:        1,
		UniversityCourseId:    1,
		SubCoordinatorTuition: 1,
		CoordinatorTuition:    1,
	},
}

var testInsertAdvisory = map[string]struct {
	advisoryStorage  AdvisoryStorage
	academicAdvisory model.AcademicAdvisory
	expectError      error
}{
	"Test 1. StatusBadRequest: Invalid Fields Error": {
		advisoryStorage:  NewAdvisoryStorage(),
		academicAdvisory: *testAcademicAdvisory,
		expectError:      model.StatusBadRequest("Check that all information fields of the advisory are correct."),
	},
	"Test 2. StatusBadRequest: Invalid Fields Error": {
		advisoryStorage:  NewAdvisoryStorage(),
		academicAdvisory: *testAcademicAdvisory,
		expectError:      model.StatusBadRequest("Check that all information fields of the advisory are correct."),
	},
}

func TestAdvisoryStorage_InsertAdvisory(t *testing.T) {
	for name, tt := range testInsertAdvisory {
		tt := tt
		t.Run(name, func(t *testing.T) {
			err := tt.advisoryStorage.InsertAdvisory(&tt.academicAdvisory)
			if !(errors.Is(err, tt.expectError)) {
				t.Errorf("\n expected error %v\n, got error %v\n", tt.expectError, err)
			}
		})
	}
}

var testParameters = map[string]struct {
	advisoryStorage AdvisoryStorage
	isAccepted      bool
	advisoryId      string
	expectError     error
}{
	"Test 1. StatusNotFound: Advisory not found": {
		advisoryStorage: NewAdvisoryStorage(),
		isAccepted:      true,
		advisoryId:      "2002ESSHTS",
		expectError:     model.NotFound(fmt.Sprintf("An advisory with id %v was not found", "2002ESSHTS")),
	},
	" Test 2. StarusNotFound: Advisory not fount found": {
		advisoryStorage: NewAdvisoryStorage(),
		isAccepted:      false,
		advisoryId:      "2001ESSHTE",
		expectError:     model.NotFound(fmt.Sprintf("An advisory with id %v was not found", "2001ESSHTE")),
	},
}

func TestAdvisoryStorage_UpdateAdvisory(t *testing.T) {
	for name, tt := range testParameters {
		tt := tt
		t.Run(name, func(t *testing.T) {
			err := tt.advisoryStorage.UpdateAdvisory(tt.isAccepted, tt.advisoryId)
			if !errors.Is(err, tt.expectError) {
				t.Fatalf("\n expected error %v\n, got error %v\n", tt.expectError, err)
			}
		})
	}
}

func TestAdvisoryStorage_DeleteAdvisory(t *testing.T) {
	for name, tt := range testParameters {
		tt := tt
		t.Run(name, func(t *testing.T) {
			err := tt.advisoryStorage.UpdateAdvisory(tt.isAccepted, tt.advisoryId)
			if !errors.Is(err, tt.expectError) {
				t.Fatalf("\n expected error %v\n, got error %v\n", tt.expectError, err)
			}
		})
	}
}
