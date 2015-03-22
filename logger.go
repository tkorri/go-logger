package logger

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

// Logger call depth
const callDepth = 2

var (
	VERBOSE = &LogLevel{tag: "VERBOSE", ordinal: 1}
	DEBUG   = &LogLevel{tag: "DEBUG  ", ordinal: 2}
	INFO    = &LogLevel{tag: "INFO   ", ordinal: 3}
	ERROR   = &LogLevel{tag: "ERROR  ", ordinal: 4}

	vLog = log.New(&proxy{level: VERBOSE}, "", 0)
	dLog = log.New(&proxy{level: DEBUG}, "", 0)
	iLog = log.New(&proxy{level: INFO}, "", 0)
	eLog = log.New(&proxy{level: ERROR}, "", log.Lshortfile)

	currentLevel = INFO
	timeFormat   = "2006-01-02 15:04:05"
	logFile      = os.Stderr
)

type LogLevel struct {
	tag     string
	ordinal int
}

type proxy struct {
	level *LogLevel
}

func (t *proxy) Write(p []byte) (n int, err error) {

	if currentLevel.ordinal <= t.level.ordinal {

		timeStamp := time.Now().Format(timeFormat)

		prefix := timeStamp + " " + t.level.tag + " "

		var prefixCount, logCount int

		prefixCount, err = logFile.Write([]byte(prefix))
		if err == nil {
			logCount, err = logFile.Write(p)
		}

		return (prefixCount + logCount), err
	} else {
		return 0, nil
	}
}

func getLevel(level string) *LogLevel {
	logLevel := strings.TrimSpace(strings.ToUpper(level))

	switch logLevel {
	case strings.TrimSpace(VERBOSE.tag):
		return VERBOSE
	case strings.TrimSpace(DEBUG.tag):
		return DEBUG
	case strings.TrimSpace(INFO.tag):
		return INFO
	case strings.TrimSpace(ERROR.tag):
		return ERROR
	}

	return INFO
}

func Init() {
	// Set the default log to use the same format as this library
	log.SetOutput(&proxy{level: INFO})
	log.SetFlags(0)
}

func SetLogLevel(newLevel string) {
	level := getLevel(newLevel)
	I("Setting log level to", level.tag)
	currentLevel = level
}

func SetTimeFormat(format string) {
	timeFormat = format
}

func SetLogFileLocation(fileLocation string) {

	// Try to create log directory
	logFilePath, _ := filepath.Abs(fileLocation)
	err := os.MkdirAll(path.Dir(logFilePath), 0755)

	if err != nil {
		E("Cannot create log directory:", err)
		os.Exit(-1)
	}

	f, err := os.OpenFile(fileLocation, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		E("Error opening log file:", err)
		os.Exit(-1)
	}

	log.Println("Redirecting logs to:", logFilePath)

	// Close previous log file
	if logFile != os.Stderr {
		logFile.Close()
	}

	logFile = f
}

func V(v ...interface{}) {
	if currentLevel.ordinal <= VERBOSE.ordinal {
		vLog.Output(callDepth, fmt.Sprintln(v...))
	}
}

func D(v ...interface{}) {
	if currentLevel.ordinal <= DEBUG.ordinal {
		dLog.Output(callDepth, fmt.Sprintln(v...))
	}
}

func I(v ...interface{}) {
	if currentLevel.ordinal <= INFO.ordinal {
		iLog.Output(callDepth, fmt.Sprintln(v...))
	}
}

func E(v ...interface{}) {
	if currentLevel.ordinal <= ERROR.ordinal {
		eLog.Output(callDepth, fmt.Sprintln(v...))
	}
}
