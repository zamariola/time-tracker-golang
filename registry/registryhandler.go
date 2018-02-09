package registry

import "github.com/zamariola/time-tracker-golang/input"

type Repository interface {
	Write(task *input.Task)
	ReadLast() *input.Task
}
