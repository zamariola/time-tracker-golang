package main

import (
	"github.com/zamariola/time-tracker-golang/filesystem"
	"fmt"
	"os"
	"github.com/zamariola/time-tracker-golang/input"
	"github.com/zamariola/time-tracker-golang/util"
)

func main() {


	//TODO: Please myself, use test instead of main
	argsWithoutProg := os.Args[1:]
	fmt.Println(argsWithoutProg);

	task, err := input.ParseArgs(argsWithoutProg);
	util.CheckError(err);
	fmt.Println(task);
	fmt.Println(task.End().Format("2006/01/02 15:04"));
	fmt.Println(task.Start().Format("2006/01/02 15:04"));



	config, _ := filesystem.LoadConfig("")
	path := config[filesystem.KEY_CONFIG_TRACKING_PATH];
	fmt.Print(path);
}