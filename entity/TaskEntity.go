package entity

import "time"

type Task struct {
	message string
	start   time.Time
	end     time.Time
}

func NewTask(message string, start time.Time, end time.Time) *Task {
	return &Task{message, start, end}
}

func (t Task) Message() string {
	return t.message;
}

func (t Task) Start() time.Time {
	return t.start;
}

func (t Task) End() time.Time {
	return t.end;
}
