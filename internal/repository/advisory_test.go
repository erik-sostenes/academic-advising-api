package repository

import (
	"errors"
	"testing"
	"time"

	"github.com/itsoeh/academic-advising-api/internal/model"
)

var testAcademicAdvisory = &model.AcademicAdvisory{
	AdvisoryId: "190HY5D",
	Description: "This is test.",
	Reports: []byte("This is test."),
	FromDate: time.Now(),
	ToDate: time.Now().AddDate(0, 2, 7),
	RecordTime: time.Now(),
	IsActive: true,
	AcademicAdvisoryIds: model.AcademicAdvisoryIds{
		SubjectId: 1,
		StudentTuition: 1,
		TeacherTuition: 1,
		UniversityCourseId: 1,
		SubCoordinatorTuition: 1,
		CoordinatorTuition: 1,
	},
}

var testAdvisoryAggregator = map[string]struct{
	advisoryAggregator AdvisoryAggregator
	academicAdvisory model.AcademicAdvisory
	expectError error
}{
	"Test 1. StatusBadRequest: Invalid Fields Error": {
		advisoryAggregator: NewAdvisoryAggregator(),
		academicAdvisory: *testAcademicAdvisory, 
		expectError: InvalidFieldsError,
	},
	"Test 2. StatusBadRequest: Invalid Fields Error": {
		advisoryAggregator: NewAdvisoryAggregator(),
		academicAdvisory: *testAcademicAdvisory, 
		expectError: InvalidFieldsError,
	},
}

func TestAdvisoryAggregator_AddAdvisory(t *testing.T) {
	for name, tt := range testAdvisoryAggregator {
		tt := tt
		t.Run(name, func(t *testing.T) {	
			err := tt.advisoryAggregator.AddAdvisory(&tt.academicAdvisory)
			if !(errors.Is(err, tt.expectError)) {
				t.Errorf("\n expect error %v\n, got error %v\n", tt.expectError, err)
			} 
		})
	}
}
