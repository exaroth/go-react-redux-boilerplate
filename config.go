package main

import (
	"fmt"
	"github.com/spf13/viper"
	"html/template"
	"io/ioutil"
	"strings"
)

const (
	appName            string = ""
	defaultTemplateDir        = "./templates"
	defaultStaticDir          = "./static/"
	templateExt               = ".tpl"
	defaultHostname           = "127.0.0.1"
	version                   = "0.0.1"
	lLevelWarning             = "warning"
	lLevelInfo                = "info"
	lLevelError               = "error"
)

const (
	defaultPort     int = 9000
	defaultLogLevel     = 1
)

// Env variable names
const (
	envPrefix      string = appName
	envDev                = "DEVELOPMENT"
	envHostname           = "APP_HOSTNAME"
	envPort               = "APP_PORT"
	envLogDisabled        = "LOGGING_DISABLED"
	envLogLevel           = "LOG_LEVEL"
	envLogRequests        = "LOG_REQUESTS"
)

type MainConfig struct {
	Port        int
	Hostname    string
	Templates   *template.Template
	DevEnv      bool
	AppName     string
	StaticDir   string
	TemplateDir string
	Version     string
	// Logging options
	LogDisabled    bool
	LogLevel       uint8
	LogRequestData bool
}

func (c *MainConfig) LoadTemplates() error {
	templates := []string{}
	files, err := ioutil.ReadDir(defaultTemplateDir)
	if err != nil {
		return err
	}
	for _, f := range files {
		fName := f.Name()
		if strings.HasSuffix(fName, templateExt) {
			templates = append(templates, fmt.Sprintf("%s/%s", defaultTemplateDir, fName))
		}
	}
	if len(templates) == 0 {
		return nil
	}
	c.Templates, err = template.ParseFiles(templates...)
	return err
}

func (c *MainConfig) GetTemplate(tplName string) *template.Template {
	tplName = fmt.Sprintf("%s%s", tplName, templateExt)
	return c.Templates.Lookup(tplName)
}

func getConfig() (cfg *MainConfig, err error) {
	viper.SetEnvPrefix(envPrefix)
	viper.AutomaticEnv()

	cfg = &MainConfig{}

	cfg.DevEnv = (viper.GetInt(envDev) == 1)

	cfg.Hostname = viper.GetString(envHostname)
	if cfg.Hostname == "" {
		cfg.Hostname = defaultHostname
	}
	cfg.Port = viper.GetInt(envPort)
	if cfg.Port == 0 {
		cfg.Port = defaultPort
	}
	cfg.LogDisabled = (viper.GetInt(envLogDisabled) == 1)
	cfg.LogRequestData = (viper.GetInt(envLogRequests) == 1)

	lLevelStr := viper.GetString(envLogLevel)
	switch lLevelStr {
	case lLevelError:
		cfg.LogLevel = 2
		break
	case lLevelWarning:
		cfg.LogLevel = 1
		break
	case lLevelInfo:
		cfg.LogLevel = 0
		break
	default:
		cfg.LogLevel = defaultLogLevel
	}

	cfg.AppName = appName
	cfg.Version = version
	cfg.StaticDir = defaultStaticDir

	err = func(funcs ...func() error) (er error) {
		for _, f := range funcs {
			if er = f(); er != nil {
				return
			}
		}
		return
	}(
		cfg.LoadTemplates,
		// Put config loader funcs here
	)
	return
}
