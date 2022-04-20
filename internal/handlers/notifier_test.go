package handlers

import (
	"github.com/itsoeh/academy-advising-api/internal/model"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

var responseTeacherStream = map[string]struct {
	channel    model.ChannelIsAccepted
	notifier   Notifier
	path       string
	statusCode int
	httpMethod string
}{
	"": {
		channel: model.ChannelIsAccepted{
			StudentId:  "19018125",
			IsAccepted: true,
			Message:    "His academic advising was accepted",
		},
		notifier:   NewNotifier(),
		path:       "/v1/itsoeh/academy-advising-api/sse",
		statusCode: http.StatusOK,
		httpMethod: http.MethodGet,
	},
	" ": {
		channel: model.ChannelIsAccepted{
			StudentId:  "19018126",
			IsAccepted: false,
			Message:    "your academic advisory was not accepted",
		},
		notifier:   NewNotifier(),
		path:       "/v1/itsoeh/academy-advising-api/sse",
		statusCode: http.StatusOK,
		httpMethod: http.MethodGet,
	},
}

func TestNotifier_Notify(t *testing.T) {
	for name, tt := range responseTeacherStream {
		tt := tt
		t.Run(name, func(t *testing.T) {
			e := echo.New()

			req := NewRequest(t, tt.httpMethod, tt.path, "")
			defer req.Body.Close()
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			n := NewNotifier()

			n.Notify(c)
			n.response.ResponseTeacherStream <- &tt.channel

			log.Println(rec.Code)
			log.Println(rec.Result().Body)
		})
	}
}

var academyAdvisoryQueryParams = map[string]struct {
	path       string
	advisoryId string
	isAccepted string
	notifier   Notifier
	statusCode int
	httpMethod string
}{
	"Academy Advising not found, StatusCode: 404": {
		path:       "/v1/itsoeh/academy-advising-api/update/:advisory_id/:is_accepted",
		advisoryId: "3476347",
		isAccepted: "false",
		statusCode: http.StatusNotFound,
		notifier:   NewNotifier(),
		httpMethod: http.MethodPut,
	},
	"isAcceptedQueryParam empty, StatusCode: 400": {
		path:       "/v1/itsoeh/academy-advising-api/update/:advisory_id/:is_accepted",
		advisoryId: "3476347",
		isAccepted: "  ",
		statusCode: http.StatusBadRequest,
		notifier:   NewNotifier(),
		httpMethod: http.MethodPut,
	},
	"advisoryIdQueryParam empty, StatusCode: 400": {
		path:       "/v1/itsoeh/academy-advising-api/update/:advisory_id/:is_accepted",
		advisoryId: "    ",
		isAccepted: "true",
		statusCode: http.StatusBadRequest,
		notifier:   NewNotifier(),
		httpMethod: http.MethodPut,
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
