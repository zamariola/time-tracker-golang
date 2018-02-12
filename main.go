package main

import (
	"github.com/zamariola/time-tracker-golang/filesystem"
	"os"
	"github.com/zamariola/time-tracker-golang/input"
	"github.com/zamariola/time-tracker-golang/util"
	log "github.com/sirupsen/logrus"
)

func main() {

	initLog()

	argsWithoutProg := os.Args[1:]
	log.Debug(argsWithoutProg);

	taskPtr, err := input.ParseArgs(argsWithoutProg);
	util.CheckError(err);


	config, _ := filesystem.LoadConfig("")
	path := config[filesystem.KEY_CONFIG_TRACKING_PATH];
	log.Debug(path)

	fileSystemHandlerPtr := filesystem.NewFileSystemHandler(path);
	err = fileSystemHandlerPtr.Write(taskPtr)
	util.CheckError(err)

	fileSystemHandlerPtr.ReadLast();

}

func initLog() {

	env := os.Getenv("APP_ENV")
	if env == "production" {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.DebugLevel)
	}
}