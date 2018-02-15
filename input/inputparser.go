package input

import (
	"time"
	"fmt"
	"regexp"
	"github.com/zamariola/time-tracker-golang/util"
	log "github.com/sirupsen/logrus"
	"github.com/zamariola/time-tracker-golang/entity"
	"github.com/zamariola/time-tracker-golang/filesystem"
)

const (
	ARGS_INDEX_MESSAGE = 0;
	ARGS_INDEX_START_TIME = 2;
	ARGS_INDEX_END_TIME = 4;

	SHORTCUT_LETTER_NOW = "n";
	SHORTCUT_LETTER_LAST = "l";
	SHORTCUT_LETTER_YESTERDAY = "y";
	SHORTCUT_LETTER_TODAY = "t";

	//todo: get it from config
	TIME_INPUT_PATTERN = "15:04";
	DATE_INPUT_PATTERN = "02/01/06";
)



func ParseArgs(args []string) (*entity.Task, error) {

	if len(args) < 5 {
		return &entity.Task{}, fmt.Errorf("Invalid parameters length, expected 5 received %d", len(args))
	}

	return parseArgsContent(args), nil;
}

func parseArgsContent(args []string) *entity.Task {

	var err error;
	var message string;
	var start, end time.Time;

	for i := range args {

		switch i {
		case ARGS_INDEX_MESSAGE:
			message = args[i];
		case ARGS_INDEX_START_TIME:
			start, err = parseDateTime(args[i - 1], args[i]);
			util.CheckError(err);
		case ARGS_INDEX_END_TIME:
			end, err = parseDateTime(args[i - 1], args[i]);
			util.CheckError(err)
		}
	}

	return entity.NewTask(message, start, end);

}
func parseDateTime(dateString string, timeString string) (time.Time, error) {

	alphaNumPattern := regexp.MustCompile(`^[A-Za-z_]+$`)
	var day, hour time.Time

	if alphaNumPattern.MatchString(dateString) {
		day = convertShortcutToDateTime(dateString);
	} else {
		day, _ = time.Parse(DATE_INPUT_PATTERN, dateString)
	}
	if alphaNumPattern.MatchString(timeString) {
		hour = convertShortcutToDateTime(timeString);
	} else {
		hour, _ = time.Parse(TIME_INPUT_PATTERN, timeString)
	}

	return time.Date(day.Year(), day.Month(), day.Day(),
		hour.Hour(), hour.Minute(), 0, 0, day.Location()), nil;
}

func convertShortcutToDateTime(text string) time.Time {

	switch text {

	case SHORTCUT_LETTER_NOW:
		return time.Now();
	case SHORTCUT_LETTER_LAST:
		//todo: make flexible destination
		fsh := filesystem.NewFileSystemHandlerFromDefaultConfig();

		return fsh.ReadLast().End();
	case SHORTCUT_LETTER_YESTERDAY:
		dur, err := time.ParseDuration("-24h");
		util.CheckError(err);
		return time.Now().Add(dur);
	case SHORTCUT_LETTER_TODAY:
		return time.Now();
	default:
		log.Warn("Unknown shortcut %s, using now()", text)
		return time.Now();
	}
}