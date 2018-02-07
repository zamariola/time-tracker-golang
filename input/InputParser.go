package input

import (
	"time"
	"errors"
	"fmt"
	"regexp"
	"github.com/zamariola/time-tracker-golang/util"
)

const (
	ARGS_INDEX_MESSAGE = 0;
	ARGS_INDEX_START_TIME = 2;
	ARGS_INDEX_END_TIME = 4;

	SHORTCUT_LETTER_NOW = "n";
	SHORTCUT_LETTER_LAST = "l";
	SHORTCUT_LETTER_YESTERDAY = "y";
	SHORTCUT_LETTER_TODAY = "t";
)

type Task struct {
	message string
	start   time.Time
	end     time.Time
}

func ParseArgs(args []string) (*Task, error) {

	if len(args) < 5 {
		return &Task{}, errors.New(fmt.Sprintf("Invalid parameters length, expected 5 received %d", len(args)))
	}

	return parseArgsContent(args), nil;
}

func parseArgsContent(args []string) *Task {

	task := Task{}
	var err error = nil

	for i := range args {

		switch i {
		case ARGS_INDEX_MESSAGE:
			task.message = args[i];
		case ARGS_INDEX_START_TIME:
			task.start, err = parseDateTime(args[i - 1], args[i]);
			util.CheckError(err);
		case ARGS_INDEX_END_TIME:
			task.end, err = parseDateTime(args[i - 1], args[i]);
			util.CheckError(err)
		}
	}

	return &task;

}
func parseDateTime(dateString string, timeString string) (time.Time, error) {

	alphaNumPattern := regexp.MustCompile(`^[A-Za-z]+$`)
	var day, hour time.Time

	//TODO: Parse start and end datetimes using letters to short (today, now, yesterday and etc)
	if alphaNumPattern.MatchString(dateString) {
		day = convertShortcutToDateTime(dateString);
	} else {
		//TODO: Insert logic here
		day, _ = time.Parse("Mon Jan 2 15:04:05 -0700 MST 2006", dateString)
	}
	if alphaNumPattern.MatchString(dateString) {
		hour = convertShortcutToDateTime(timeString);
	} else {
		//TODO: Insert logic here
	}

	//TODO: Combine this two infos
	fmt.Print(day);
	fmt.Print(hour);

	return *new(time.Time), nil;
}

func convertShortcutToDateTime(text string) time.Time {

	switch text {

	case SHORTCUT_LETTER_NOW:
		return time.Now();
	case SHORTCUT_LETTER_LAST:
		return *new(time.Time);
	case SHORTCUT_LETTER_YESTERDAY:
		dur, err := time.ParseDuration("-24h");
		util.CheckError(err);
		return time.Now().Add(dur);
	case SHORTCUT_LETTER_TODAY:
		return time.Now();
	default:
		return time.Now();
	}
}