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

func (c *MainConfig) GetAddress() string {
	return fmt.Sprintf("%s:%d", c.Hostname, c.Port)
}

func (c *MainConfig) GetStaticDir() string {
	return c.StaticDir
}

func (c *MainConfig) GetTemplate(tplName string) *template.Template {
	tplName = fmt.Sprintf("%s%s", tplName, templateExt)
	return c.Templates.Lookup(tplName)
}

func getConfig() (c *MainConfig, err error) {
	viper.SetEnvPrefix(envPrefix)
	viper.AutomaticEnv()

	c = &MainConfig{}

	c.DevEnv = (viper.GetInt(envDev) == 1)

	c.Hostname = viper.GetString(envHostname)
	if c.Hostname == "" {
		c.Hostname = defaultHostname
	}
	c.Port = viper.GetInt(envPort)
	if c.Port == 0 {
		c.Port = defaultPort
	}
	c.LogDisabled = (viper.GetInt(envLogDisabled) == 1)
	c.LogRequestData = (viper.GetInt(envLogRequests) == 1)

	lLevelStr := viper.GetString(envLogLevel)
	switch lLevelStr {
	case lLevelError:
		c.LogLevel = 2
		break
	case lLevelWarning:
		c.LogLevel = 1
		break
	case lLevelInfo:
		c.LogLevel = 0
		break
	default:
		c.LogLevel = defaultLogLevel
	}

	c.AppName = appName
	c.Version = version
	c.StaticDir = defaultStaticDir

	err = func(funcs ...func() error) (er error) {
		for _, f := range funcs {
			if er = f(); er != nil {
				return
			}
		}
		return
	}(
		c.LoadTemplates,
		// Put config loader funcs here
	)
	return
}
