package cflog

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

// Global variable
var (
	LogFilePath string
)

// Log...
type Log struct {
	Error *log.Logger
	Info  *log.Logger
	Trace *log.Logger
}

// SavePrintLog ...
// Print Log and Save Log
func SavePrintLog() *Log {

	var logFile *os.File
	logPath := "/logs/"

	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	Error := log.New(io.MultiWriter(logFile, os.Stderr),
		"\x1b[31;1m ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info := log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Trace := log.New(os.Stdout,
		"\x1b[34;1m TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	initCh := make(chan bool)
	go func() {
	LOOP:

		t := time.Now()
		sYear := strconv.Itoa(t.Year())
		sMonth := strconv.Itoa(int(t.Month()))
		day := t.Day()
		hour := strconv.Itoa(t.Hour())

		filename := fmt.Sprintf("%s.log", t.Format("2006010215"))
		st := sYear + sMonth

		LogFilePath = pwd + logPath
		filePath := pwd + logPath + st + "/" + strconv.Itoa(t.Day()) + "/" + hour + "/"

		err := os.MkdirAll(filePath, 0777)
		if err != nil {
			log.Println(err)
			return
		}

		ERfilePath := filePath + "ERROR/"
		os.Mkdir(ERfilePath, 0700)

		ERlogFile, err := os.OpenFile(ERfilePath+filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
		if err != nil {
			log.Println("ERlogFile : os.OpenFile : ", err)
		}
		Error.SetOutput(io.MultiWriter(ERlogFile, os.Stdout))

		INfilePath := filePath + "INFO/"
		os.Mkdir(INfilePath, 0700)

		INlogFile, err := os.OpenFile(INfilePath+filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
		if err != nil {
			log.Println("INlogFile : os.OpenFile : ", err)
		}

		Info.SetOutput(io.MultiWriter(INlogFile, os.Stdout))

		TRfilePath := filePath + "TRACE/"
		os.Mkdir(TRfilePath, 0700)

		TRlogFile, err := os.OpenFile(TRfilePath+filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
		if err != nil {
			log.Println("TRlogFile : os.OpenFile : ", err)
		}
		Trace.SetOutput(io.MultiWriter(TRlogFile, os.Stdout))

		initCh <- true
		for {
			if time.Now().Day() != day {
				logFile.Close()
				goto LOOP
			}
			time.Sleep(time.Millisecond * 1000)
			goto LOOP
		}
	}()

	<-initCh

	mylog := &Log{
		Error: Error,
		Info:  Info,
		Trace: Trace,
	}

	return mylog
}

// GetFilePath ...
// Get LogFile Path
func GetFilePath() string {

	return LogFilePath
}
