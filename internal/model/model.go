package model

import "time"

type(
	academicAdvisoryIds struct {
		SubjectId uint32 `json:"subject_id"`
		StudentTuition uint32 `json:"student_tuition"`
		TeacherTuition uint32 `json:"teacher_tuition"`
		UniversityCourseId uint32 `json:"university_course_id"`
		SubCoordinatorsTuition uint32 `json:"sub_coordinators_tuition"`
		CoordinatorsTuition uint32 `json:"coordinators_tuition"`
	}

	academicAdvisory struct {
		AdviceId string `json:"advice_id"`
		Description string `json:"description"`
		Reports []byte `json:"reports"`
		FromDate time.Time `json:"from_date"`
		ToDate time.Time `json:"to_date"`
		IsActive bool `json:"is_active"`
		AcademicAdvisoryIds academicAdvisoryIds
	}
	
	AcademicAdvisory []AcademicAdvisory
)

func NewAcademicAdvisoryIds(subjectId, studentTuition, teacherTuition, universityCourseId,
														subCoordinatorsTuition, coordinatorsTuition uint32) *academicAdvisoryIds {
	return &academicAdvisoryIds{
		SubjectId: subjectId,
		StudentTuition: studentTuition,
		TeacherTuition: teacherTuition,
		UniversityCourseId: universityCourseId,
		SubCoordinatorsTuition: subCoordinatorsTuition,
		CoordinatorsTuition: coordinatorsTuition,
	}
}
func NewAcademicAdvisory(adviceId, description string,reports []byte, fromDate, 
												toDate time.Time, isActive bool, academicAdvisoryIds academicAdvisoryIds) *academicAdvisory {

	return &academicAdvisory{
		AdviceId: adviceId,
		Description: description,
		Reports: reports,
		FromDate: fromDate,
		ToDate: toDate,
		IsActive: isActive,
		AcademicAdvisoryIds: academicAdvisoryIds,
	}
}
