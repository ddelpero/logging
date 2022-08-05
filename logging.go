package logging

import (
	"os"
	"time"

	"github.com/nuveo/log"
)

var LogFileName string
var logDate time.Time

//SetLogFileName set name of file that will need to rotate
func SetLogFileName(fileName string) {
	LogFileName = fileName
	logDate = time.Now()
}

//rotate call this before logging to check if file needs to be rotated
//simple rotation based on day changing
func rotate() {
	if time.Now().Day() != logDate.Day() {
		newLogFileName := LogFileName + "." + time.Now().Format("2006-02-01")
		logDate = time.Now()
		os.Rename(LogFileName, newLogFileName)
	}
}

func Debugln(msg ...interface{}) {
	rotate()
	log.Debugln(msg...)
}

func Errorln(msg ...interface{}) {
	rotate()
	log.Errorln(msg...)
}

func Errorf(msg ...interface{}) {
	rotate()
	log.Errorf(msg...)
}

func Fatal(msg ...interface{}) {
	rotate()
	log.Fatal(msg...)
}

func Println(msg ...interface{}) {
	rotate()
	log.Println(msg...)
}






