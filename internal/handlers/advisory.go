package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/itsoeh/academic-advising-api/internal/model"
	"github.com/itsoeh/academic-advising-api/internal/services"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// AcademicAdvisory contains the methods to follow the notification flow when adding an advisory
type AcademicAdvisory interface {
	// AddAcademicAdvisory handler that receives the new advisory and will take care of adding
	// Note: only if the teacher agrees
	AddAcademicAdvisory(http.ResponseWriter, *http.Request)
	// NotifyAcademicAdvisory handler who is in charge of notifying the teacher
	// that a student wants to reserve an advisory
	NotifyAcademicAdvisory(http.ResponseWriter, *http.Request)
}

type academicAdvisory struct {
	services services.AcademicAdvisoryAdministrator
	channels *model.Channels
}

// NewAcademicAdvisory implements the AcademicAdvisory interface
func NewAcademicAdvisory() AcademicAdvisory {
	return &academicAdvisory{
		services: services.NewAcademicAdvisingAdministrator(),
		channels: &model.Channels{
			ResponseTeacherStream: make(model.ResponseTeacherStream),
			NotifyTeacherStream:   make(model.NotifyTeacherStream),
		},
	}
}

func (a *academicAdvisory) AddAcademicAdvisory(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()

	go func(ws *websocket.Conn) {
		defer close(a.channels.ResponseTeacherStream)

		for {
			select {
			case msg := <-a.channels.ResponseTeacherStream:
				log.Println("Enviando respuesta al alumno")
				_ = ws.WriteJSON(msg)
				log.Println("Respuesta enviada")
				log.Printf("El profesor ha aceptado ¡Se ha guardado tu asesoría!")
			}
		}
	}(ws)

	for {
		data := model.ChannelIsAccepted{}

		err := ws.ReadJSON(&data)
		if err != nil {
			log.Println(err)
			break
		}

		log.Printf("Message received: %v", data)

	}
}

func (a *academicAdvisory) NotifyAcademicAdvisory(w http.ResponseWriter, r *http.Request) {

}
