package repository

import "github.com/itsoeh/academic-advising-api/internal/model"

const (
	sqlQueryAddAdvisory = `INSERT INTO advisories(
	advisory_id,
	description,
	reports,
	from_date,
	to_date,
	record_time,
	is_active,
	is_acepted,
	subject_id,
	student_tuition,
	teachers_tuition,
	university_course_id,
	subcoordinator_tuition,
	coordinator_tuition,
	) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

	sqlQueryUpdateAdisory = `
		UPDATE 
			advisories a
		SET a.is_acepted = ?
		WHERE a.advisory_id = ?	`

	sqlQueryDeleteAdvisory = `
		DELETE FROM 
			advisories a
		WHERE a.advisory_id = ? AND a.is_acepted = ?`
)

const (
	InvalidFieldsError = model.StatusBadRequest("Check that all information fields of the advisory are correct.")
)
