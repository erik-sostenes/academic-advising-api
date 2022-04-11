package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

// academicAdvisoryOneJSON Format is incorrect
const academicAdvisoryOneJSON = `{
		"description": "This is test.",
		"reports": "cddffhuyugt43664543g4hv43gruygdvyg43fv",
		"from_date":    "2022-04-01T08:41:50Z",
		"to_date":      "2022-04-01T08:41:50Z",
		"is_active":    false,
		"is_acepted": false,
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
		"record_time": "2022-04-01T08:41:50Z",
		"is_active":    false,
		"is_acepted": false,
		"academic_advisory_ids": {
			"subject_id" :             1,
			"student_tuition" :        1,
			"teacher_tuition":        	1,
			"university_course_id":    	1,
			"sub_coordinator_tuition": 	1,
			"coordinator_tuition":    	1
		}
}`

var academicAdvisory = map[string]struct {
	advisoryJSON string
	path         string
	statusCode   uint16
	httpMethod   string
}{
	"StatusBadRequest: Format is incorrect": {
		advisoryJSON: academicAdvisoryOneJSON,
		path:         "/v1/itsoeh/academy-advising-api/create",
		statusCode:   400,
	},
	"StatusBadRequest: Incorrect data": {
		advisoryJSON: academicAdvisoryTwoJSON,
		path:         "/v1/itsoeh/academy-advising-api/create",
		statusCode:   400,
	},
}

func TestHandlersAdvisory_HandlerCreateAdvisory(t *testing.T) {
	for name, tt := range academicAdvisory {
		tt := tt
		t.Run(name, func(t *testing.T) {
			e := echo.New()

			req := NewRequest(t, http.MethodPost, tt.path, tt.advisoryJSON)
			defer req.Body.Close()
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			NewHandlersAdvisory().HandlerCreateAdvisory(c)

			if rec.Code != int(tt.statusCode) {
				t.Errorf("expected error code %v, got error code %v", tt.statusCode, rec.Code)
			}
		})
	}
}

func NewRequest(t testing.TB, method, path string, dataJSON string) *http.Request {
	t.Helper()
	return httptest.NewRequest(method, path, strings.NewReader(dataJSON))
}
