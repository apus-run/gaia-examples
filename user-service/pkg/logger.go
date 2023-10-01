package bootstrap

import (
	"os"

	"github.com/apus-run/sea-kit/log"
)

func NewLoggerProvider() log.Logger {
	l := log.NewStdLogger(os.Stdout)
	return log.With(
		l,
		"service.id", "",
		"service.name", "",
		"service.version", "",
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
	)
}
