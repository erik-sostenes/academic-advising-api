package model

import "time"

type (
	AcademicAdvisoryIds struct {
		SubjectId             uint32 `json:"subject_id"`
		StudentTuition        uint32 `json:"student_tuition"`
		TeacherTuition        uint32 `json:"teacher_tuition"`
		UniversityCourseId    uint32 `json:"university_course_id"`
		SubCoordinatorTuition uint32 `json:"sub_coordinators_tuition"`
		CoordinatorTuition    uint32 `json:"coordinators_tuition"`
	}

	AcademicAdvisory struct {
		AdvisoryId          string    `json:"advice_id"`
		Description         string    `json:"description"`
		Reports             []byte    `json:"reports"`
		FromDate            time.Time `json:"from_date"`
		ToDate              time.Time `json:"to_date"`
		IsActive            bool      `json:"is_active"`
		AcademicAdvisoryIds AcademicAdvisoryIds
	}

	AcademicAdvisories []AcademicAdvisory
)

func NewAcademicAdvisoryIds(subjectId, studentTuition, teacherTuition, universityCourseId,
	subCoordinatorsTuition, coordinatorsTuition uint32) *AcademicAdvisoryIds {
	return &AcademicAdvisoryIds{
		SubjectId:             subjectId,
		StudentTuition:        studentTuition,
		TeacherTuition:        teacherTuition,
		UniversityCourseId:    universityCourseId,
		SubCoordinatorTuition: subCoordinatorsTuition,
		CoordinatorTuition:    coordinatorsTuition,
	}
}
func NewAcademicAdvisory(advisoryId, description string, reports []byte, fromDate,
	toDate time.Time, isActive bool, academicAdvisoryIds AcademicAdvisoryIds) *AcademicAdvisory {

	return &AcademicAdvisory{
		AdvisoryId:          advisoryId,
		Description:         description,
		Reports:             reports,
		FromDate:            fromDate,
		ToDate:              toDate,
		IsActive:            isActive,
		AcademicAdvisoryIds: academicAdvisoryIds,
	}
}
