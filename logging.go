package logging

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
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
func rotate() string {
	newLogFileName := ""
	if time.Now().Day() != logDate.Day() {
		newLogFileName = fmt.Sprintf("%s.%s.log", strings.Replace(LogFileName, ".log", "", 1), time.Now().Format("2006-01-02"))
		logDate = time.Now()
		os.Rename(LogFileName, newLogFileName)
	}
	return newLogFileName;
}

func ZipLog(newLogFileName string) {
	//archive log file
	archiveLogFileName := LogFileName + "." + time.Now().Format("2006-02-01")
	archive, err := os.Create(archiveLogFileName + ".zip")
	if err != nil {
		log.Errorln("Error creating archive log file: ", err)
	}
	defer archive.Close()

	zipWriter := zip.NewWriter(archive)
	defer zipWriter.Close()
	f, err := os.Open(newLogFileName)
	if err != nil {
		log.Errorln("Error opening log file: ", err)
	}

	_, filename := filepath.Split(newLogFileName)
	zipW, err := zipWriter.Create(filename)
	if err != nil {
		log.Errorln("Error adding zip file: ", err)
	}
	if _, err = io.Copy(zipW, f); err != nil {
		log.Errorln("Error copying file to zip: ", err)
	}
	f.Close()
	zipWriter.Close()
	os.Remove(newLogFileName)
	// don't log the error; otherwise, we'll get an infinite loop
	// if err != nil {
	// 	log.Errorln("Error removing log file: ", err)
	// }
}

func Debugln(msg ...interface{}) {
	newLogFileName := rotate()
	log.Debugln(msg...)
	if newLogFileName != "" {
		ZipLog(newLogFileName)
	}
}

func Errorln(msg ...interface{}) {
	newLogFileName := rotate()
	log.Errorln(msg...)
	if newLogFileName != "" {
		ZipLog(newLogFileName)
	}
}

func Errorf(msg ...interface{}) {
	newLogFileName := rotate()
	log.Errorf(msg...)
	if newLogFileName != "" {
		ZipLog(newLogFileName)
	}
}

func Fatal(msg ...interface{}) {
	newLogFileName := rotate()
	log.Fatal(msg...)
	if newLogFileName != "" {
		ZipLog(newLogFileName)
	}
}

func Println(msg ...interface{}) {
	newLogFileName := rotate()
	log.Println(msg...)
	if newLogFileName != "" {
		ZipLog(newLogFileName)
	}
}

