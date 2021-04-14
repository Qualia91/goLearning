package log_lib

const (
	DEBUG = iota
	WARNING
	FULL
	INFO
	ERROR
)

type LogData struct {
	Level         int
	Time          int
	CorrelationID int
	HostID        string
	Message       string
}
