package filesystem

import (
	"testing"
	"time"
	"github.com/zamariola/time-tracker-golang/entity"
)

const (
	MESSAGE = "Stub Message";
)

var (
	START_TIME, _ = time.Parse(time.RFC3339,   "2018-03-01T10:22:00Z")
	END_TIME, _ = time.Parse(time.RFC3339,   "2018-03-02T11:33:00Z")
)

var fsh = NewFileSystemHandler("/stub/path")

func TestShouldCorrectlyFormatTask(t *testing.T) {

	expectedMessage := MESSAGE + "," + "2018-03-01 10:22"+","+"2018-03-02 11:33"
	task := entity.NewTask(MESSAGE, START_TIME, END_TIME)
	realMessage := fsh.Format(task)

	if(expectedMessage != realMessage) {
		t.Errorf("Wrong Message Format, expected: %s but got :%s", expectedMessage, realMessage)
	}

}

