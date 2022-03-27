package repository

const (
	sqlQueryInsertAdvisory = `INSERT INTO advisories(
	advisory_id,
	description,
	from_date,
	to_date,
	is_active,
	is_acepted,
	subject_id,
	student_tuition,
	teachers_tuition,
	university_course_id,
	subcoordinator_tuition,
	coordinator_tuition
	) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

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
