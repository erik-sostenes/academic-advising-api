package notifier

import (
	"fmt"
)

// Notifier contains the method to notify subscribers
type Notifier interface {
	// Notify method that is responsible for notifying the message
	Notify()
}

type notifier struct {
	IsAccepted bool
}

// New returns a notifier structure that implements the Notifier interface
// receive the response
func New(IsAccepted bool) Notifier {
	return &notifier{IsAccepted: IsAccepted}
}

func (n *notifier) Notify() {
	fmt.Printf("Se le acaba de notificar al alumno %v", n.IsAccepted)
}
