package util

import "log"

func CheckError(err error) bool {
	if err != nil {
		log.Fatal(err);
		panic(err);
	}
	return err != nil;
}
