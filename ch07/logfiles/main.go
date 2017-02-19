package main

import (
	"log"
	"os"
)

var (
	Notice  *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func main() {
	noticeFile, err := os.OpenFile("notice.log", os.O_RDWR|os.O_APPEND, 0660)
	defer noticeFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	warnFile, err := os.OpenFile("warnings.log", os.O_RDWR|os.O_APPEND, 0660)
	defer warnFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	errorFile, err := os.OpenFile("error.log", os.O_RDWR|os.O_APPEND, 0660)
	defer errorFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	Notice = log.New(noticeFile, "NOTICE: ", log.Ldate|log.Ltime)
	Notice.SetOutput(noticeFile)
	Notice.Println("This is basically F.Y.I.")

	Warning = log.New(warnFile, "WARNING: ", log.Ldate|log.Ltime)
	Warning.SetOutput(warnFile)
	Warning.Println("Perhaps this needs your attention?")

	Error = log.New(errorFile, "ERROR: ", log.Ldate|log.Ltime)
	Error.SetOutput(errorFile)
	Error.Println("You REALLY should fix this!")
}
