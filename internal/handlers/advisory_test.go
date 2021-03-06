package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/itsoeh/academy-advising-api/internal/repository"
	"github.com/itsoeh/academy-advising-api/internal/services"
	"github.com/labstack/echo/v4"
)

// academicAdvisoryOneJSON Format is incorrect
const academicAdvisoryOneJSON = `{
		"description": "This is test.",
		"reports": "cddffhuyugt43664543g4hv43gruygdvyg43fv",
		"from_date":    "2022-04-01T08:41:50Z",
		"to_date":      "2022-04-01T08:41:50Z",
		"is_active":    false,
		"is_accepted": false,
		"academic_advisory_ids": {
			"subject_id" :             1,
			"student_tuition" :        1,
			"teacher_tuition":         1,
			"university_course_id":    1,
			"sub_coordinator_tuition": 1,
		}
}`

// academicAdvisoryTwoJSON Incorrect data
const academicAdvisoryTwoJSON = `{
	"advisory_id":  "190HY5D",
		"description": "This is test.",
		"reports": "cddffhuyugt43664543g4hv43gruygdvyg43fv",
		"from_date":    "2022-04-01T08:41:50Z",
		"to_date":      "2022-04-01T08:41:50Z",
		"record_time": 	1,
		"is_active":    false,
		"is_accepted": false,
		"academic_advisory_ids": {
			"subject_id" :             1,
			"student_tuition" :        1,
			"teacher_tuition":        	1,
			"university_course_id":    	1,
			"sub_coordinator_tuition": 	1,
			"coordinator_tuition":    	1
		}
}`

var academyAdvisory = map[string]struct {
	advisoryJSON string
  handlers     AdvisoryHandler
	path         string
	statusCode   int
	httpMethod   string
}{
	"Format is incorrect, StatusCode: 400": {
		advisoryJSON: academicAdvisoryOneJSON,
		handlers:     NewAdvisoryHandler(),
		path:         "/v1/itsoeh/academy-advising-api/create",
		statusCode:   http.StatusBadRequest,
		httpMethod:   http.MethodPost,
	},
	"Incorrect data, StatusCode: 400": {
		advisoryJSON: academicAdvisoryTwoJSON,
		handlers: 		NewAdvisoryHandler(),
		path:         "/v1/itsoeh/academy-advising-api/create",
		statusCode:   http.StatusBadRequest,
		httpMethod:   http.MethodPost,
	},
}

func TestAdvisory_CreateAdvisory(t *testing.T) {
	DB := repository.NewDB()
	repository := repository.NewAdvisoryStorage(DB)
	services := services.NewAdvisoryManager(repository)

	for name, tt := range academyAdvisory {
		tt := tt
		t.Run(name, func(t *testing.T) {
			e := echo.New()

			e.POST(tt.path, tt.handlers.CreateAdvisory(services))
			
			req := NewRequest(t, tt.httpMethod, tt.path, tt.advisoryJSON)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			
			
			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)

			res := rec.Result()
			defer res.Body.Close()
				
			if res.StatusCode != tt.statusCode {
				t.Errorf("expected error code %v, got error code %v", tt.statusCode, rec.Code)
			}
		})
	}
}
