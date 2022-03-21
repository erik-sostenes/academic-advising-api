package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/itsoeh/academic-advising-api/internal/model"
)

// AdvisoryAggregator contains the method to add a new academic advisory
type AdvisoryAggregator interface {
	// AddAdvisory method that has the task of adding a new academic advisory through a sql query
	AddAdvisory(*model.AcademicAdvisory) error
}

// advisoryAggregator implements AdvisoryAggregator interface
type advisoryAggregator struct {
	*sql.DB
}

func NewAdvisoryAggregator() AdvisoryAggregator {
	return &advisoryAggregator{
		DB: NewDB(),
	}
}

func (a *advisoryAggregator) AddAdvisory(advisory *model.AcademicAdvisory) (err error) {
	_, err = a.Exec(sqlQueryAddAdvisory,
		&advisory.AdvisoryId,
		&advisory.Description,
		&advisory.Reports,
		&advisory.FromDate,
		&advisory.ToDate,
		&advisory.IsActive,
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
			log.Println(err)
			return
		}

		if code.Number == 1452 {
			err = InvalidFieldsError
			log.Println(err)
			return
		}

		err = model.InternalServerError("An error has occurred when adding a new advisory.")
		log.Println(err)
		return
	}
	return
}
