package filesystem

import (
	"github.com/zamariola/time-tracker-golang/input"
	"fmt"
	log "github.com/sirupsen/logrus"
)


type FileSystemHandler struct {
	trackingPath string;
}

func (fsh *FileSystemHandler) TrackingPath() string {
	return fsh.trackingPath;
}

func NewFileSystemHandler(path string) *FileSystemHandler {
	return &FileSystemHandler{path}
}

func (fsh FileSystemHandler) Write(task *input.Task) {

	task.Message()

}

func (fsh FileSystemHandler) ReadLast() *input.Task {

	return nil;
}

func (fsh FileSystemHandler) Format(task *input.Task) string {

	log.Debugf("Formatting %s %s %s", task.Message(), task.Start(), task.End())
	return fmt.Sprint(task.Message(), ",",
		task.Start().Format("2006-01-02 15:04"), ",",
		task.End().Format("2006-01-02 15:04"))
}


