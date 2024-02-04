package pkg

import (
    "fmt"
    "log"
    "os"
)

var (
    InfoLogger    *log.Logger
    WarningLogger *log.Logger
    ErrorLogger   *log.Logger
)

func init() {
    InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    WarningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
    ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(message string, args ...interface{}) {
    formattedMessage := fmt.Sprintf(message, args...)
    InfoLogger.Println(formattedMessage)
}

func Warning(message string, args ...interface{}) {
    formattedMessage := fmt.Sprintf(message, args...)
    WarningLogger.Println(formattedMessage)
}

func Error(message string, args ...interface{}) {
    formattedMessage := fmt.Sprintf(message, args...)
    ErrorLogger.Println(formattedMessage)
}
