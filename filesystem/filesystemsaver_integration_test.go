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
	START_TIME, _ = time.Parse(time.RFC3339, "2018-03-01T10:22:00Z")
	END_TIME, _ = time.Parse(time.RFC3339, "2018-03-02T11:33:00Z")
)

var fsh = NewFileSystemHandler("/home/CIT/leonardoz/.time-tracker/time-tracker_TEST.log")

func TestShouldCorrectlyFormatTask(t *testing.T) {

	expectedMessage := MESSAGE + "," + "2018-03-01 10:22" + "," + "2018-03-02 11:33"
	task := entity.NewTask(MESSAGE, START_TIME, END_TIME)
	realMessage := fsh.Format(task)

	if (expectedMessage != realMessage) {
		t.Errorf("Wrong Message Format, expected: %s but got: %s", expectedMessage, realMessage)
	}

}

func TestShouldUnmarshallFormattedTask(t *testing.T) {

	textToUnmarshall := MESSAGE + "," + "2018-03-01 10:22" + "," + "2018-05-02 11:33";
	taskPtr := Unmarshall(textToUnmarshall);

	if (taskPtr.Message() != MESSAGE) {
		t.Errorf("Wrong Unmarshal, expected: %s but got: %s", MESSAGE, taskPtr.Message())
	}

	if (!compareMicroDates(taskPtr.Start(), 2018, 3, 1, 10, 22)) {
		t.Errorf("Wrong Unmarshal, expected: %s but got: %s", "2018-03-01 10:22", taskPtr.Start())
	}

	if (!compareMicroDates(taskPtr.End(), 2018, 5, 2, 11, 33)) {
		t.Errorf("Wrong Unmarshal, expected: %s but got: %s", "2018-05-02 11:33", taskPtr.End())
	}

}

func TestShouldAppendTaskOnEndOfFile(t *testing.T){

	task := entity.NewTask(MESSAGE, START_TIME, END_TIME);
	fsh.Write(task)
}

func compareMicroDates(timeToCompare time.Time, year, month, day, hour, minute int) bool {
	return timeToCompare.Year() == year &&
		int(timeToCompare.Month()) == month &&
		timeToCompare.Day() == day &&
		timeToCompare.Hour() == hour &&
		timeToCompare.Minute() == minute
}





