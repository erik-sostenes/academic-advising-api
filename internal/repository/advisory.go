package repository

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/itsoeh/academy-advising-api/internal/model"
)

// AdvisoryStorage contains all the methods to interact with the database
// and be able to interact with academic advising
type AdvisoryStorage interface {
	// InsertAdvisory method that has the task of adding a new academic advisory through a sql query
	InsertAdvisory(advisory *model.AcademicAdvisory) error
	// UpdateAdvisory method that has the task of updating a academic advisory
	UpdateAdvisory(isAccepted bool, advisoryId, teacherScheduleId string) error
	// DeleteAdvisory method that has the task of deleting a academic advisory
	DeleteAdvisory(advisoryId, teacherScheduleId string) error
}

// advisoryStorage implements AdvisoryStorage interface
type advisoryStorage struct {
	DB *sql.DB
}

// NewAdvisoryStorage implements the AdvisoryStorage interface
func NewAdvisoryStorage(DB *sql.DB) AdvisoryStorage {
	return &advisoryStorage{
		DB: DB,
	}
}

func (a *advisoryStorage) InsertAdvisory(advisory *model.AcademicAdvisory) (err error) {
	tx, err := a.DB.Begin()
	if err != nil {
		err = model.InternalServerError(err.Error())
		return
	}

	defer tx.Rollback()

	advisoryId := uuid.New().String()

	_, err = tx.Exec(sqlQueryInsertAdvisory,
		&advisoryId,
		&advisory.Description,
		&advisory.FromDate,
		&advisory.ToDate,
		&advisory.IsActive,
		&advisory.IsAccepted,
		&advisory.AcademicAdvisoryIds.SubjectId,
		&advisory.AcademicAdvisoryIds.StudentTuition,
		&advisory.AcademicAdvisoryIds.TeacherTuition,
		&advisory.AcademicAdvisoryIds.UniversityCourseId,
		&advisory.AcademicAdvisoryIds.SubCoordinatorTuition,
		&advisory.AcademicAdvisoryIds.CoordinatorTuition,
	)

	if code, ok := err.(*mysql.MySQLError); ok {
		//NOTE: Error Code: 1062. Duplicate entry "advisory_id" for key
		//NOTE: Error Code: 1452. Cannot add or update a child row: a foreign key constraint fails

		if code.Number == 1062 {
			err = model.StatusBadRequest(fmt.Sprintf("The advisory whith id %v already exist.", advisory.AdvisoryId))
			return
		}

		if code.Number == 1452 {
			err = model.StatusBadRequest("Check that all information fields of the advisory are correct.")
			return
		}

		err = model.InternalServerError("An error has occurred when adding a new advisory.")
		return
	}

	academicAdvisoryScheduleRecordId := uuid.New().String()

	_, err = tx.Exec(sqlQueryInserAcademicAdvisoryScheduleRecord,
		&academicAdvisoryScheduleRecordId,
		&advisory.TeacherScheduleId,
		&advisoryId,
	)

	if code, ok := err.(*mysql.MySQLError); ok {
		//NOTE: Error Code: 1062. Duplicate entry "academic_advisory_schedule_record_id" for key

		if code.Number == 1062 {
			err = model.StatusBadRequest(
				fmt.Sprintf("The academic advisory schedule record whith id %v already exist.", academicAdvisoryScheduleRecordId),
		)
			return
		}

		err = model.InternalServerError("An error has occurred when adding a new academic advisory schedule record.")
		return
	}

	if err = tx.Commit(); err != nil {
		err = model.InternalServerError(err.Error())
		return
	}

	return
}

func (a *advisoryStorage) UpdateAdvisory(isAccepted bool, advisoryId, teacherScheduleId string) (err error) {
	tx, err := a.DB.Begin()
	if err != nil {
		err = model.InternalServerError(err.Error())
		return
	}

	defer tx.Rollback()

	rows, err := tx.Exec(sqlQueryUpdateAdvisory,
		&isAccepted,
		&advisoryId,
	)

	if err != nil {
		err = model.InternalServerError("An error has ocurred when updating an advisory.")
		return
	}

	if rowAffect, _ := rows.RowsAffected(); rowAffect != 1 {
		err = model.NotFound(fmt.Sprintf("An advisory with id %v was not found", advisoryId))
		return
	}

	row, err := tx.Exec(slqQueryUpdateTeachersSchedules,
		&teacherScheduleId,
		&teacherScheduleId,
	)

	if err != nil {
		err = model.InternalServerError("An error has ocurred when updating an teacher schedules.")
		return
	}

	if rowAffect, _ := row.RowsAffected(); rowAffect != 1 {
		err = model.NotFound(fmt.Sprintf("An teacher schedules with id %v was not found", &teacherScheduleId,))
		return
	}
	
	if err = tx.Commit(); err != nil {
		err = model.InternalServerError(err.Error())
		return
	}
	return
}

func (a *advisoryStorage) DeleteAdvisory(advisoryId, teacherScheduleId string) (err error) {
	tx, err := a.DB.Begin()
	if err != nil {
		err = model.InternalServerError(err.Error())
		return
	}

	defer tx.Rollback()


	row, err := tx.Exec(sqlQueryDeleteAdvisory, &advisoryId)

	if err != nil {
		err = model.InternalServerError("An error has ocurred when deleting an advisory.")
		return
	}

	if rowAffect, _ := row.RowsAffected(); rowAffect != 1 {
		err = model.NotFound(fmt.Sprintf("An advisory with id %v was not found", advisoryId))
		return
	}

	row, err = tx.Exec(sqlQueryDeleteAcademicAdvisoryAcheduleRecord, &teacherScheduleId, &advisoryId)

	if err != nil {
		err = model.InternalServerError("An error has ocurred when deleting an academic Advisory Achedule Record.")
		return
	}

	if rowAffect, _ := row.RowsAffected(); rowAffect != 1 {
		err = model.NotFound(fmt.Sprintf("An Advisory Achedule Record with id %v was not found", advisoryId))
		return
	}

	
	if err = tx.Commit(); err != nil {
		err = model.InternalServerError(err.Error())
		return
	}
	return
}
