package main

import (
	"github.com/zamariola/time-tracker-golang/filesystem"
	"os"
	"github.com/zamariola/time-tracker-golang/input"
	"github.com/zamariola/time-tracker-golang/util"
	log "github.com/sirupsen/logrus"
	"fmt"
)

func main() {

	initLog()

	argsWithoutProg := os.Args[1:]
	log.Debugf("Parsing input args: %s", argsWithoutProg);

	taskPtr, err := input.ParseArgs(argsWithoutProg);
	util.CheckError(err);

	fileSystemHandlerPtr := filesystem.NewFileSystemHandlerFromDefaultConfig();

	log.Debugf("Trying to write task %s to filesystem", *fileSystemHandlerPtr);

	err = fileSystemHandlerPtr.Write(taskPtr)
	util.CheckError(err)

	fmt.Printf("Saved : %s - %s - %s \n", taskPtr.Message(), taskPtr.Start(), taskPtr.End())
}

func initLog() {

	env := os.Getenv("APP_ENV")
	if env == "production" {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.DebugLevel)
	}
}