package notifier

type Subscribers interface {
	AddNotifier(isAccepted bool, n *Notifier)
	NotifyNotifier()
}

type subscribers struct {
	notifiers map[string]Notifier
	Message   string
}

func (s *subscribers) AddNotifier(message string, n *Notifier) {
	if s.notifiers == nil {
		s.notifiers = make(map[string]Notifier)
	}

	s.notifiers[message] = *n
}

func (s *subscribers) NotifyNotifier() {
	for _, v := range s.notifiers {
		v.Notify(s.Message)
	}
}
