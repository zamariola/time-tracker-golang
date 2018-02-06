package main

import (
	"github.com/zamariola/time-tracker-golang/filesystem"
	"fmt"
)

func main() {

	fmt.Print(filesystem.LoadConfig(""));
}