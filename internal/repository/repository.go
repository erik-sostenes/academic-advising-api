package repository

const (
	sqlQueryInsertAdvisory = `INSERT INTO advisories(
	advisory_id,
	description,
	from_date,
	to_date,
	is_active,
	is_accepted,
	subject_id,
	student_tuition,
	teachers_tuition,
	university_course_id,
	subcoordinator_tuition,
	coordinator_tuition
	) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	sqlQueryUpdateAdvisory = `
		UPDATE 
			advisories a
		SET a.is_accepted = ?
		WHERE a.advisory_id = ?	`

	sqlQueryDeleteAdvisory = `
		DELETE FROM 
			advisories a
		WHERE a.advisory_id = ?`
)
