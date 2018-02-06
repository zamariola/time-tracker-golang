package main

import (
	"github.com/zamariola/time-tracker-golang/filesystem"
	"fmt"
	"os"
	"github.com/zamariola/time-tracker-golang/input"
)

func main() {


	//TODO: Please myself, use test instead of main
	argsWithoutProg := os.Args[1:]
	fmt.Print(argsWithoutProg);

	task, erro := input.ParseArgs(argsWithoutProg);
	fmt.Print(task);
	fmt.Print("\n");
	fmt.Print(erro);

	config, _ := filesystem.LoadConfig("")
	path := config[filesystem.KEY_CONFIG_TRACKING_PATH];
	fmt.Print(path);
}