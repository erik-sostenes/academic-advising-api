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

	sqlQueryInserAcademicAdvisoryScheduleRecord = `
	INSERT INTO academic_advisory_schedule_record(
		academic_advisory_schedule_record_id,
		teacher_schedule_id,
		advisory_id
	) VALUES(?, ?, ?);`
	
	slqQueryUpdateTeachersSchedules  = `
	UPDATE teachers_schedules ts SET student_accountant =
  	(SELECT count(*) FROM academic_advisory_schedule_record ar WHERE ar.teacher_schedule_id = ?)
	WHERE ts.teacher_schedule_id = ?;`

	sqlQueryUpdateAdvisory = `
		UPDATE 
			advisories a
		SET a.is_accepted = ?
		WHERE a.advisory_id = ?	`

	sqlQueryDeleteAdvisory = `
		DELETE FROM 
			advisories a
		WHERE a.advisory_id = ?`

	sqlQueryDeleteAcademicAdvisoryAcheduleRecord  = `
		DELETE FROM
			academic_advisory_schedule_record a 
		WHERE a.teacher_schedule_id = ? AND advisory_id = ?;
	`
)
