package repository

const (
	sqlQueryAddAdvisory = `INSERT INTO advisories(
																						advisory_id,
																						description,
																						reports,
																						from_date,
																						to_date,
																						is_active,
																						subject_id,
																						student_tuition,
																						teachers_tuition,
																						university_course_id,
																						subcoordinator_tuition,
																						coordinator_tuition
																						) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`
)
