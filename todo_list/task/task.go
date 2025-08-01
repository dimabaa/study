package task

import (
	"time"
)

type TaskStruct struct {
	Subject    string
	Text       string
	TimeCreate time.Time
	Completed  bool
}

func Task(subject string, text string) TaskStruct {
	return TaskStruct{
		Subject:    subject,
		Text:       text,
		TimeCreate: time.Now(),
		Completed:  false,
	}

}
