package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

// Basic logger interface.
type logger interface {
	Warning(string, error, *http.Request)
	Error(string, error, *http.Request)
	Info(string, error, *http.Request)
}

// Represents request retrieved data.
type logReadyRequestInfo struct {
	URL    string                 `json:"url"`
	Method string                 `json:"method"`
	Params map[string]interface{} `json:"params,omitempty"`
}

// Logger module.
type Logger struct {
	cfg *MainConfig
}

// Struct representing unmarshalled output of log.
type loggerOutput struct {
	RequestInfo *logReadyRequestInfo `json:"request_info,omitempty"`
	Error       string               `json:"error,omitempty"`
}

// Base logger output fn parsing and  loading the data to be logged
func (l *Logger) getLogData(severity uint8, message string, err error, req *http.Request) *loggerOutput {
	if l.cfg.LogDisabled || l.cfg.LogLevel < severity {
		return nil
	}
	var errMess string
	if err != nil {
		errMess = err.Error()
	}
	return &loggerOutput{
		Error:       errMess,
		RequestInfo: l.parseRequest(req),
	}
}

// Scrape request information, url, params to be logged along with
// the other info.
func (l *Logger) parseRequest(req *http.Request) *logReadyRequestInfo {
	// TODO: for now we are not passing any params to the app,
	// so for now Params key is nil at all times. (kw)
	if req == nil || !l.cfg.LogRequestData {
		return nil
	}
	return &logReadyRequestInfo{
		URL:    fmt.Sprintf("%s%s", req.Host, req.URL.String()),
		Method: req.Method,
	}
}

// Log info.
func (l *Logger) Info(message string, err error, req *http.Request) {
	data := l.getLogData(2, message, err, req)
	if data != nil {
		log.WithFields(log.Fields{
			"data": data,
		}).Info(message)
	}
}

// Log warning.
func (l *Logger) Warning(message string, err error, req *http.Request) {
	data := l.getLogData(1, message, err, req)
	if data != nil {
		log.WithFields(log.Fields{
			"data": data,
		}).Warning(message)
	}
}

//  Log error.
func (l *Logger) Error(message string, err error, req *http.Request) {
	data := l.getLogData(0, message, err, req)
	if data != nil {
		log.WithFields(log.Fields{
			"data": data,
		}).Error(message)
	}
}

// Initialize logger module.
func NewLogger(c *MainConfig) logger {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	// Set level to info for logrus as we manually check
	// severity(kw).
	log.SetLevel(log.InfoLevel)
	return &Logger{
		cfg: c,
	}
}
