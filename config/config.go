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
	LogLevel           Level
	Protocol           string // Optional protocol to use for requests. https is used by default. Example: http, https
	Host               string // JUST hostname, without setting schema. Example: localhost, NOT http://localhost
	ManualTokenRefresh bool
}
