package registry

import (
	"github.com/zamariola/time-tracker-golang/entity"
)

type Repository interface {
	Write(task *entity.Task)
	ReadLast() *entity.Task
	Format(task *entity.Task)
}
