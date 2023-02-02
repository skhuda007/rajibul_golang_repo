package logger

import "log"

func LogInfo(msg string) {
	log.Print("INFO-", msg)
}

func LogWarn(msg string) {
	log.Print("WARN-", msg)
}

func LogError(msg string) {
	log.Printf("Error- %v", msg)
}
