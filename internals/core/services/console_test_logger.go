package services

import (
	"log"
	"time"
)

func PrintTimestamp(startTimestamp time.Time, method string, path string, status int) {
	log.Printf("Logging Middleware: Method: %s, Path: %s, Status: %d, Duration: %s",
		method, path, status, time.Since(startTimestamp))
}
