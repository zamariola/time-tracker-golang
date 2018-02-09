package filesystem

import "github.com/zamariola/time-tracker-golang/input"

type FileSystemHandler struct {
	TrackingPath string
}

func NewFileSystemHandler(path string) * FileSystemHandler {
	return &FileSystemHandler{path}
}

func (fsh FileSystemHandler) Write(task *input.Task) {

}

func (fsh FileSystemHandler) ReadLast() *input.Task {

	return nil;
}


