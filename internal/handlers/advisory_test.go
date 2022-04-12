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
	advisoryJSON    string
	handlerAdvisory HandlerAdvisory
	path            string
	statusCode      int
	httpMethod      string
}{
	"Format is incorrect, StatusCode: 400": {
		advisoryJSON:    academicAdvisoryOneJSON,
		handlerAdvisory: NewHandlerAdvisory(),
		path:            "/v1/itsoeh/academy-advising-api/create",
		statusCode:      400,
		httpMethod:      http.MethodPost,
	},
	"Incorrect data, StatusCode: 400": {
		advisoryJSON:    academicAdvisoryTwoJSON,
		handlerAdvisory: NewHandlerAdvisory(),
		path:            "/v1/itsoeh/academy-advising-api/create",
		statusCode:      400,
		httpMethod:      http.MethodPost,
	},
}

func TestHandlersAdvisory_HandlerCreateAdvisory(t *testing.T) {
	for name, tt := range academyAdvisory {
		tt := tt
		t.Run(name, func(t *testing.T) {
			e := echo.New()

			req := NewRequest(t, tt.httpMethod, tt.path, tt.advisoryJSON)
			defer req.Body.Close()
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			tt.handlerAdvisory.HandlerCreateAdvisory(c)

			if rec.Code != tt.statusCode {
				t.Errorf("expected error code %v, got error code %v", tt.statusCode, rec.Code)
			}
		})
	}
}

var academyAdvisoryQueryParams = map[string]struct {
	path            string
	advisoryId      string
	isAccepted      string
	handlerAdvisory HandlerAdvisory
	statusCode      int
	httpMethod      string
}{
	"Academy Advising not found, StatusCode: 404": {
		path:            "/v1/itsoeh/academy-advising-api/update/:advisory_id/:is_accepted",
		advisoryId:      "3476347",
		isAccepted:      "false",
		statusCode:      404,
		handlerAdvisory: NewHandlerAdvisory(),
		httpMethod:      http.MethodPut,
	},
	"isAceptedQueryParam empty, StatusCode: 400": {
		path:            "/v1/itsoeh/academy-advising-api/update/:advisory_id/:is_accepted",
		advisoryId:      "3476347",
		isAccepted:      "  ",
		statusCode:      400,
		handlerAdvisory: NewHandlerAdvisory(),
		httpMethod:      http.MethodPut,
	},
	"advisoryIdQueryParam empty, StatusCode: 400": {
		path:            "/v1/itsoeh/academy-advising-api/update/:advisory_id/:is_accepted",
		advisoryId:      "    ",
		isAccepted:      "true",
		statusCode:      400,
		handlerAdvisory: NewHandlerAdvisory(),
		httpMethod:      http.MethodPut,
	},
}

func TestHandlersAdvisory_HandlerUpdateAdvisory(t *testing.T) {
	for name, tt := range academyAdvisoryQueryParams {
		tt := tt
		t.Run(name, func(t *testing.T) {
			e := echo.New()

			req := NewRequest(t, tt.httpMethod, tt.path, "")
			defer req.Body.Close()
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			c.SetParamNames("advisory_id", "is_accepted")
			c.SetParamValues(tt.advisoryId, tt.isAccepted)

			tt.handlerAdvisory.HandlerUpdateAdvisory(c)

			if rec.Code != tt.statusCode {
				t.Errorf("expected error code %v, got error code %v", tt.statusCode, rec.Code)
			}
		})
	}
}

func NewRequest(t testing.TB, method, path string, dataJSON string) *http.Request {
	t.Helper()
	return httptest.NewRequest(method, path, strings.NewReader(dataJSON))
}
