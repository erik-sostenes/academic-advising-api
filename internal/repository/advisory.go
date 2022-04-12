package repository

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/itsoeh/academy-advising-api/internal/model"
)

// AdvisoryStorage contains all the methods to interact with the database
// and be able to interact with academic advising
type AdvisoryStorage interface {
	// InsertAdvisory method that has the task of adding a new academic advisory through a sql query
	InsertAdvisory(advisory *model.AcademicAdvisory) error
	// UpdateAdvisory method that has the task of updating a academic advisory
	UpdateAdvisory(isAccepted bool, advisoryId string) error
	// DeleteAdvisory method that has the task of deleting a academic advisory
	DeleteAdvisory(advisoryId string) error
}

// advisoryStorage implements AdvisoryStorage interface
type advisoryStorage struct {
	DB *sql.DB
}

// NewAdvisoryStorage implements the AdvisoryStorage interface
func NewAdvisoryStorage() AdvisoryStorage {
	return &advisoryStorage{
		DB: NewDB(),
	}
}

func (a *advisoryStorage) InsertAdvisory(advisory *model.AcademicAdvisory) (err error) {
	_, err = a.DB.Exec(sqlQueryInsertAdvisory,
		&advisory.AdvisoryId,
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
	return
}

func (a *advisoryStorage) UpdateAdvisory(isAccepted bool, advisoryId string) (err error) {
	rows, err := a.DB.Exec(sqlQueryUpdateAdvisory,
		&isAccepted,
		&advisoryId,
	)

	if err != nil {
		err = model.InternalServerError("An error has ocurred when deleting an advisory.")
		return
	}

	if rowAffect, _ := rows.RowsAffected(); rowAffect != 1 {
		err = model.NotFound(fmt.Sprintf("An advisory with id %v was not found", advisoryId))
		return
	}

	return
}

func (a *advisoryStorage) DeleteAdvisory(advisoryId string) (err error) {
	row, err := a.DB.Exec(sqlQueryDeleteAdvisory, &advisoryId)

	if err != nil {
		err = model.InternalServerError("An error has ocurred when deleting an advisory.")
		return
	}

	if rowAffect, _ := row.RowsAffected(); rowAffect != 1 {
		err = model.NotFound(fmt.Sprintf("An advisory with id %v was not found", advisoryId))
		return
	}
	return
}
