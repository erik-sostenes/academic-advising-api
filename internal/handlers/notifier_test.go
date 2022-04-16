package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)


var academyAdvisoryQueryParams = map[string]struct {
	path            string
	advisoryId      string
	isAccepted      string
	notifier 				Notifier
	statusCode      int
	httpMethod      string
}{
	"Academy Advising not found, StatusCode: 404": {
		path:            "/v1/itsoeh/academy-advising-api/update/:advisory_id/:is_accepted",
		advisoryId:      "3476347",
		isAccepted:      "false",
		statusCode:      404,
		notifier: 			 NewNotifier(),
		httpMethod:      http.MethodPut,
	},
	"isAcceptedQueryParam empty, StatusCode: 400": {
		path:            "/v1/itsoeh/academy-advising-api/update/:advisory_id/:is_accepted",
		advisoryId:      "3476347",
		isAccepted:      "  ",
		statusCode:      400,
		notifier: 			 NewNotifier(),
		httpMethod:      http.MethodPut,
	},
	"advisoryIdQueryParam empty, StatusCode: 400": {
		path:            "/v1/itsoeh/academy-advising-api/update/:advisory_id/:is_accepted",
		advisoryId:      "    ",
		isAccepted:      "true",
		statusCode:      400,
		notifier: 			 NewNotifier(),
		httpMethod:      http.MethodPut,
	},
}

func TestNotifier_UpdateAdvisory(t *testing.T) {
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

			tt.notifier.UpdateAdvisory(c)

			if rec.Code != tt.statusCode {
				t.Errorf("expected error code %v, got error code %v", tt.statusCode, rec.Code)
			}
		})
	}
}
