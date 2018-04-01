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
)

const (
	defaultPort int = 9000
)

// Env variable names
const (
	envPrefix  string = appName
	envDev            = "DEVELOPMENT"
	envApiAddr        = "API_ADDR"
)

type MainConfig struct {
	Port        int
	Hostname    string
	Templates   *template.Template
	DevEnv      bool
	ApiAddr     string
	AppName     string
	StaticDir   string
	TemplateDir string
	Version     string
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
	return c.Templates.Lookup(tplName)
}

func getConfig() (cfg *MainConfig, err error) {
	viper.SetEnvPrefix(envPrefix)
	viper.AutomaticEnv()

	cfg = &MainConfig{}

	cfg.DevEnv = (viper.GetInt(envDev) == 1)
	cfg.ApiAddr = viper.GetString(envApiAddr)
	if cfg.ApiAddr == "" {
	}
	cfg.Hostname = viper.GetString("APP_HOSTNAME")
	if cfg.Hostname == "" {
		cfg.Hostname = defaultHostname
	}

	cfg.AppName = appName
	cfg.Version = version
	cfg.StaticDir = defaultStaticDir

	cfg.Port = viper.GetInt("APP_PORT")
	if cfg.Port == 0 {
		cfg.Port = defaultPort
	}
	err = func(funcs ...func() error) (er error) {
		for _, f := range funcs {
			if er = f(); er != nil {
				return
			}
		}
		return
	}(
		cfg.LoadTemplates,
		// Put loader funcs here
	)
	return
}
