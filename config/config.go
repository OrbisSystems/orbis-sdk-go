package config

type Level uint32

const (
	PanicLogLevel Level = iota
	FatalLogLevel
	ErrorLogLevel
	WarnLogLevel
	InfoLogLevel
	DebugLogLevel
	TraceLogLevel
)

type Config struct {
	LogLevel Level
	Host     string // JUST hostname, without setting schema. Example: localhost, NOT https://localhost
}
