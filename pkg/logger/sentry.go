package logger

import (
	"net/http"
	"reflect"

	"github.com/exaroth/go-react-redux-boilerplate/pkg/config"
	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
)

var (
	levelMap = map[logrus.Level]sentry.Level{
		logrus.TraceLevel: sentry.LevelDebug,
		logrus.DebugLevel: sentry.LevelDebug,
		logrus.InfoLevel:  sentry.LevelInfo,
		logrus.WarnLevel:  sentry.LevelWarning,
		logrus.ErrorLevel: sentry.LevelError,
		logrus.FatalLevel: sentry.LevelFatal,
		logrus.PanicLevel: sentry.LevelFatal,
	}
)

// Converter is a a generator function which will prepare event instance to be
// sent to sentry.
type Converter func(entry *logrus.Entry, event *sentry.Event, eventHint *sentry.EventHint)

// SentryHook represents sentry-logrus hook instance.
type SentryHook struct {
	client    *sentry.Client
	r         *http.Request
	cfg       *config.ServiceConfig
	converter Converter
}

// NewSentryHook will create new logrus hook for sentry.
// TODO: ideally we would want to use Hub here so we dont instantiate new
// client every time
func NewSentryHook(client *sentry.Client, cfg *config.ServiceConfig, conv Converter) *SentryHook {
	if conv == nil {
		conv = DefaultConverter
	}
	return &SentryHook{
		converter: conv,
		client:    client,
		cfg:       cfg,
	}
}

// Levels will return valid levels for the sentry hook.
func (h *SentryHook) Levels() []logrus.Level {
	return h.cfg.GetLogLevels(false)
}

// Fire will send new sentry event.
func (hook *SentryHook) Fire(entry *logrus.Entry) error {
	event := sentry.NewEvent()
	event.Level = levelMap[entry.Level]

	hook.converter(entry, event, nil)
	hook.client.CaptureEvent(event, nil, nil)
	return nil
}

// DefaultConverter is default processor for generating sentry events to be sent.
func DefaultConverter(entry *logrus.Entry, event *sentry.Event, eHint *sentry.EventHint) {
	event.Message = entry.Message

	for k, v := range entry.Data {
		event.Extra[k] = v
	}
	if err, ok := entry.Data[logrus.ErrorKey].(error); ok {
		stacktrace := sentry.ExtractStacktrace(err)
		if stacktrace == nil {
			stacktrace = sentry.NewStacktrace()
		}
		// unwrap the error if possible
		cause := err
		if ex, ok := err.(interface{ Cause() error }); ok {
			if c := ex.Cause(); c != nil {
				cause = c
			}
		}
		exception := sentry.Exception{
			Type:       reflect.TypeOf(cause).String(),
			Value:      cause.Error(),
			Stacktrace: stacktrace,
		}
		event.Exception = []sentry.Exception{exception}
		eHint.OriginalException = cause
	}
}
