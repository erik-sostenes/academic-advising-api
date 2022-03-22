package model

import "time"

// Models
type (
	// AcademicAdvisoryIds conatins all the ids that make up the academic advisory
	AcademicAdvisoryIds struct {
		SubjectId             uint32 `json:"subject_id"`
		StudentTuition        uint32 `json:"student_tuition"`
		TeacherTuition        uint32 `json:"teacher_tuition"`
		UniversityCourseId    uint32 `json:"university_course_id"`
		SubCoordinatorTuition uint32 `json:"sub_coordinators_tuition"`
		CoordinatorTuition    uint32 `json:"coordinators_tuition"`
	}
	// AcademicAdvisory conatins the fields of the academic advisory
	AcademicAdvisory struct {
		AdvisoryId          string    `json:"advice_id"`
		Description         string    `json:"description"`
		Reports             []byte    `json:"reports"`
		FromDate            time.Time `json:"from_date"`
		ToDate              time.Time `json:"to_date"`
		RecordTime          time.Time `json:"record_time"`
		IsActive            bool      `json:"is_active"`
		AcademicAdvisoryIds AcademicAdvisoryIds
	}
	// AcademicAdvisories slice of the AcademicAdvisory structure
	AcademicAdvisories []AcademicAdvisory
)

// Channels
type (
	// ResponseTeacherStream ChannelIsAccepted channel
	ResponseTeacherStream chan *ChannelIsAccepted
	// NotifyTeacherStream ChannelAcademicAdvisory channel
	NotifyTeacherStream chan *ChannelAcademicAdvisory

	// ChannelAcademicAdvisory  structure of the message when notifying the teacher
	ChannelAcademicAdvisory struct {
		AcademicAdvisory AcademicAdvisory
		Message          string `json:"message"`
	}
	// ChannelIsAccepted teacher response structure
	ChannelIsAccepted struct {
		IsAccepted bool   `json:"is_accepted"`
		Message    string `json:"message"`
	}
	// Channels contains all channels
	Channels struct {
		responseTeacherStream ResponseTeacherStream
		notifyTeacherStream   NotifyTeacherStream
	}
)
