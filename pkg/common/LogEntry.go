package common

import "time"

type LogEntry struct {
	Time    time.Time `json:"time"`
	Phase   string    `json:"phase"`
	Message string    `json:"message"`
	Line    int       `json:"line"`
}
