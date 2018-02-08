package input

import (
	"testing"
	"time"
)

const (
	MESSAGE = "Message Example"
	START_DAY = "05/03/17"
	START_TIME = "11:30"
	END_DAY = "25/12/18"
	END_TIME = "16:59"
)

var stubArgs = []string{MESSAGE, START_DAY, START_TIME, END_DAY, END_TIME}

func TestShouldThrownExceptionOnIncorrectArgsLength(t *testing.T) {

	args := []string{"arg1", "arg2"};
	_, err := ParseArgs(args)
	if err == nil {
		t.Errorf("expected error, but didn't receive it")
	}
}

func TestShouldParseMessageFromFirstArg(t *testing.T) {

	task, err := ParseArgs(stubArgs);

	if err != nil {
		t.Errorf("Error while parsing")
	}

	if task.Message() != MESSAGE {
		t.Errorf("Wrong Message argument, expect: %s but got :%s", MESSAGE, task.Message());
	}
}

func TestShouldParseStartDayFromLiteralValue(t *testing.T) {

	task, err := ParseArgs(stubArgs);

	if err != nil {
		t.Errorf("Error while parsing")
	}

	expectedDay := 5;
	if task.Start().Day() != expectedDay {
		t.Errorf("Wrong day argument, expect: %d but got :%d", expectedDay, task.Start().Day());
	}
}

func TestShouldParseStartTimeFromLiteralValue(t *testing.T) {

	task, err := ParseArgs(stubArgs);

	if err != nil {
		t.Errorf("Error while parsing")
	}

	expectedHour := 11;
	expectedMinute := 30;

	if task.Start().Hour() != expectedHour {
		t.Errorf("Wrong hour argument, expect: %d but got :%d", expectedHour, task.Start().Hour());
	}
	if task.Start().Minute() != expectedMinute {
		t.Errorf("Wrong minute argument, expect: %d but got :%d", expectedMinute, task.Start().Minute());
	}
}

func TestShouldParseStartDateTimeFromNowStringValue(t *testing.T) {

	args := []string{MESSAGE, "n", "n", END_DAY, END_TIME}

	task, err := ParseArgs(args);

	if err != nil {
		t.Errorf("Error while parsing")
	}

	now := time.Now();
	expectedDay := now.Day();
	expectedHour := now.Hour()
	expectedMinute := now.Minute()

	if task.Start().Day() != expectedDay {
		t.Errorf("Wrong day argument, expect: %d but got :%d", expectedDay, task.Start().Day());
	}
	if task.Start().Hour() != expectedHour {
		t.Errorf("Wrong hour argument, expect: %d but got :%d", expectedHour, task.Start().Hour());
	}
	if task.Start().Minute() != expectedMinute {
		t.Errorf("Wrong minute argument, expect: %d but got :%d", expectedMinute, task.Start().Minute());
	}

}


