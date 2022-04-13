package notifier

import "log"

// Subscribers contains all the methods to keep track of notify to
// any object that implements the Notifier interface
type Subscribers interface {
	// AddNotifier method that is responsible for adding to a map all notifier
	AddNotifier(message string, n Notifier)
	// NotifyNotifier method that is responsible for notifying all the notifiers that are on the map
	NotifyNotifier()
}

type subscribers struct {
	notifiers map[string]Notifier
	Message   string
}

// NewSubscribers returns a subscribers structure that implements the Subscribers interface
func NewSubscribers() Subscribers {
	return &subscribers{}
}

func (s *subscribers) AddNotifier(message string, n Notifier) {
	if s.notifiers == nil {
		s.notifiers = make(map[string]Notifier)
	}

	s.notifiers[message] = n
}

func (s *subscribers) NotifyNotifier() {
	for k, v := range s.notifiers {
		log.Println(k)
		v.Notify()
	}
}
