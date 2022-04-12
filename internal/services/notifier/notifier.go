package notifier

import "log"

type Notifier interface {
	Notify(string)
}

type notifier struct{}

func (*notifier) Notify(message string) {
	log.Println(message)
}