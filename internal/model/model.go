package model

// Models
type (
	// AcademicAdvisoryIds conatins all the ids that make up the academic advisory
	AcademicAdvisoryIds struct {
		SubjectId             string `json:"subject_id"`
		StudentTuition        string `json:"student_tuition"`
		TeacherTuition        string `json:"teacher_tuition"`
		UniversityCourseId    string `json:"university_course_id"`
		SubCoordinatorTuition string `json:"sub_coordinator_tuition"`
		CoordinatorTuition    string `json:"coordinator_tuition"`
	}
	// AcademicAdvisory contains the fields of the academic advisory
	AcademicAdvisory struct {
		AdvisoryId          string              `json:"advisory_id"`
		Description         string              `json:"description"`
		Reports             string              `json:"reports,omitempty"`
		FromDate            string              `json:"from_date"`
		ToDate              string              `json:"to_date,omitempty"`
		RecordTime          uint32              `json:"record_time,omitempty"`
		IsActive            bool                `json:"is_active"`
		IsAccepted          bool                `json:"is_accepted"`
		TeacherScheduleId   string              `json:"teacher_schedule_id"`
		AcademicAdvisoryIds AcademicAdvisoryIds `json:"academic_advisory_ids"`
	}
	// AcademicAdvisories slice of the AcademicAdvisory structure
	AcademicAdvisories []AcademicAdvisory

	// Response
	Map map[string]interface{}
)

// Channels
type (
	// ChannelIsAccepted teacher response structure
	ChannelIsAccepted struct {
		StudentId  string `json:"student_id"`
		IsAccepted bool   `json:"is_accepted"`
		Message    string `json:"message"`
	}
)
