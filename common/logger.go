package common

import (
	"log"
	"os"
)

var (
	Debug *log.Logger
)

func InitLogger(debugMode bool) {
	if debugMode {
		Debug = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		Debug = log.New(nil, "", 0)
	}
}
