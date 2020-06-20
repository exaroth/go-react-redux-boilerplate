package logger

import (
	"os"

	"github.com/exaroth/go-react-redux-boilerplate/pkg/config"
	"github.com/getsentry/sentry-go"

	logrus_stack "github.com/Gurpartap/logrus-stack"
	log "github.com/sirupsen/logrus"
)

var Logger *log.Logger

func init() {
	Logger = NewLogger()
}

func getLoggerHooks() log.LevelHooks {
	cfg := config.Config
	hooks := log.LevelHooks{}

	// dont log enything for test env.
	if cfg.ServiceEnv == config.ServiceEnvTest {
		return hooks
	}

	hooks.Add(logrus_stack.StandardHook())

	if cfg.SentryDSN != "" {
		sentryClient, err := sentry.NewClient(sentry.ClientOptions{
			Dsn:              cfg.SentryDSN,
			Environment:      string(cfg.ServiceEnv),
			AttachStacktrace: true,
		})
		if err != nil {
			panic(err)
		}
		hooks.Add(NewSentryHook(sentryClient, cfg, nil))
	}
	return hooks
}

func NewLogger() *log.Logger {
	cfg := config.Config
	return &log.Logger{
		Out:       os.Stdout,
		Formatter: new(log.JSONFormatter),
		Hooks:     getLoggerHooks(),
		Level:     cfg.LogLevel,
	}
}
