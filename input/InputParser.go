package input

import (
	"time"
	"errors"
	"fmt"
)

const (
	ARGS_INDEX_MESSAGE = 0;
	ARGS_INDEX_START_DATE = 1;
	ARGS_INDEX_START_TIME = 2;
	ARGS_INDEX_END_DATE = 3;
	ARGS_INDEX_END_TIME = 4;
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

	return parseArgsContent(args);
}

func parseArgsContent(args []string) (*Task, error) {

	task := Task{}
	task.message = args[ARGS_INDEX_MESSAGE]

	//TODO: Parse start and end datetimes using letters to short (today, now, yesterday and etc)
	task.start,_ = time.Parse("Mon Jan 2 15:04:05 -0700 MST 2006",args[ARGS_INDEX_START_DATE])

	return &task,nil

}