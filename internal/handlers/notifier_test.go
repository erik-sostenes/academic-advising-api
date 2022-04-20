package handlers

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/itsoeh/academy-advising-api/internal/repository"
	"github.com/itsoeh/academy-advising-api/internal/services"
	"github.com/labstack/echo/v4"
)

/*
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
			stream := make(chan *model.ChannelIsAccepted)

			go func() {
				defer close(stream)
				stream <- &tt.channel
			}()

			e.GET(tt.path, tt.notifier.Notify(stream))
			

			req := NewRequest(t, tt.httpMethod, tt.path, "")
			defer req.Body.Close()
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			e.ServeHTTP(rec, req)

			res := rec.Result()
			defer res.Body.Close()


			log.Println(rec.Body.String())
		})
	}
}
*/

var academyAdvisoryQueryParams = map[string]struct {
	notifier   Notifier
	path       string
	advisoryId string
	isAccepted string
	statusCode int
	httpMethod string
}{
	"Academy Advising not found, StatusCode: 404": {
		notifier:   NewNotifier(),
		path:       "/v1/itsoeh/academy-advising-api/update/%v/%v",
		advisoryId: "3476347",
		isAccepted: "false",
		statusCode: http.StatusNotFound,
		httpMethod: http.MethodPut,
	},
	"Bad query param 'is_accepted', StatusCode: 400": {
		notifier:   NewNotifier(),
		path:       "/v1/itsoeh/academy-advising-api/update/%v/%v",
		advisoryId: "3476347",
		isAccepted: "hk",
		statusCode: http.StatusBadRequest,
		httpMethod: http.MethodPut,
	},
	"Empty query param 'advisory_id', StatusCode: 404": {
		notifier:   NewNotifier(),
		path:       "/v1/itsoeh/academy-advising-api/update/%v/%v",
		advisoryId: "",
		isAccepted: "true",
		statusCode: http.StatusNotFound,
		httpMethod: http.MethodPut,
	},
}

func TestNotifier_UpdateAdvisory(t *testing.T) {
	repository := repository.NewAdvisoryStorage()
	services := services.NewAdvisoryManager(repository)

	for name, tt := range academyAdvisoryQueryParams {
		tt := tt
		t.Run(name, func(t *testing.T) {
			e := echo.New()
			
			e.PUT("/v1/itsoeh/academy-advising-api/update/:is_accepted/:advisory_id", tt.notifier.UpdateAdvisory(services, nil))		

			req := NewRequest(t, tt.httpMethod, fmt.Sprintf(tt.path, tt.isAccepted, tt.advisoryId), "")
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
