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
		SubCoordinatorTuition uint32 `json:"sub_coordinator_tuition"`
		CoordinatorTuition    uint32 `json:"coordinator_tuition"`
	}
	// AcademicAdvisory contains the fields of the academic advisory
	AcademicAdvisory struct {
		AdvisoryId          string              `json:"advisory_id"`
		Description         string              `json:"description"`
		Reports             string              `json:"reports,omitempty"`
		FromDate            time.Time           `json:"from_date"`
		ToDate              time.Time           `json:"to_date,omitempty"`
		RecordTime          uint32              `json:"record_time,omitempty"`
		IsActive            bool                `json:"is_active"`
		IsAccepted          bool                `json:"is_accepted"`
		AcademicAdvisoryIds AcademicAdvisoryIds `json:"academic_advisory_ids"`
	}
	// AcademicAdvisories slice of the AcademicAdvisory structure
	AcademicAdvisories []AcademicAdvisory

	// Response
	Map map[string]interface{}
)

// Channels
type (
	// ResponseTeacherStream ChannelIsAccepted channel
	ResponseTeacherStream chan *ChannelIsAccepted

	// ChannelIsAccepted teacher response structure
	ChannelIsAccepted struct {
		StudentId  string `json:"student_id"`
		IsAccepted bool   `json:"is_accepted"`
		Message    string `json:"message"`
	}

	// Channels contains all channels
	Channels struct {
		ResponseTeacherStream ResponseTeacherStream
	}
)
