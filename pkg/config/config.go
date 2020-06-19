package config

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var Config *ServiceConfig

func init() {
	var err error
	Config, err = NewConfig()
	if err != nil {
		panic(err)
	}
}

// ServiceEnv denotes type of environment app is running.
type ServiceEnv string

const (
	// ServiceEnvLocal is local environment.
	ServiceEnvLocal ServiceEnv = "local"
	// ServiceEnvTest is environment for tests.
	ServiceEnvTest = "test"
	// ServiceEnvStaging is environment used in staging.
	ServiceEnvStaging = "staging"
	// ServiceEnvProduction is production environment.
	ServiceEnvProduction = "prod"
)

func (se ServiceEnv) isValid() bool {
	switch se {
	case ServiceEnvLocal, ServiceEnvTest, ServiceEnvStaging, ServiceEnvProduction:
		return true
	}
	return false
}

// NewServiceEnv will return a ServiceEnv or error
// if invalid type was passed.
func NewServiceEnv(in string) (ServiceEnv, error) {
	if in == "" {
		return ServiceEnvProduction, nil
	}
	serviceEnv := ServiceEnv(in)
	if !serviceEnv.isValid() {
		return ServiceEnvProduction, fmt.Errorf("Invalid service env defined: %s", in)
	}
	return serviceEnv, nil
}

const (
	defaultAppPort         int = 9000
	defaultPrometheusPort      = 9102
	defaultReportingPeriod     = 10
)

const (
	defaultAppAddr         string = "0.0.0.0"
	defaultPrometheusAddr         = "0.0.0.0"
	defaultStaticDirName          = "static"
	defaultTemplateDirName        = "templates"
	templateExt                   = ".tpl"
)

// ServiceConfig holds all global app configuration
// options.
type ServiceConfig struct {

	// This denotes type of the service env app is running in
	ServiceEnv ServiceEnv

	// This denotes logging level for running app
	LogLevel log.Level

	// ServiceAddr is local address app is running at
	ServiceAddr string
	// ServicePort denotes port service is running at locally.
	ServicePort int

	// ServicePort denotes port service is running at locally.
	StaticDir string

	// PrometheusServerAddr is address under which to run Prometheus exporter
	PrometheusServerAddr string
	// PrometheusServerPort is port under which to run Prometheus exporter
	PrometheusServerPort int

	// Telemetry Config
	ReportingPeriod    time.Duration
	StackdriverEnabled bool
	StackdriverProject string
	TraceSampleRate    float64
	TraceAlways        bool

	// this denotes how much time we should wait for response from the remote
	// when making external requests.
	RequestTimeout time.Duration
}

func (cfg *ServiceConfig) loadPrometheusEnvironment() error {
	cfg.PrometheusServerAddr = viper.GetString("PROMETHEUS_EXPORTER_ADDR")
	if cfg.PrometheusServerAddr == "" {
		cfg.PrometheusServerAddr = defaultPrometheusAddr
	}
	cfg.PrometheusServerPort = viper.GetInt("PROMETHEUS_EXPORTER_PORT")
	if cfg.PrometheusServerPort == 0 {
		cfg.PrometheusServerPort = defaultPrometheusPort
	}
	return nil
}

func (cfg *ServiceConfig) loadTelemetryEnvironment() error {
	cfg.StackdriverEnabled = viper.GetBool("STACKDRIVER_ENABLED")
	cfg.TraceSampleRate = viper.GetFloat64("TRACE_SAMPLE_RATE")
	cfg.TraceAlways = viper.GetBool("TRACE_ALWAYS")

	period := viper.GetInt("REPORTING_PERIOD")
	if period == 0 {
		period = defaultReportingPeriod
	}
	cfg.ReportingPeriod = time.Duration(period) * time.Second
	return nil
}

func (cfg *ServiceConfig) loadServerEnvironment() error {
	var err error
	cfg.ServiceEnv, err = NewServiceEnv(viper.GetString("SERVICE_ENV"))
	if err != nil {
		return err
	}
	rawLogLevel := viper.GetString("LOG_LEVEL")
	if rawLogLevel == "" {
		rawLogLevel = "info"
	}
	cfg.LogLevel, err = log.ParseLevel(rawLogLevel)
	if err != nil {
		return err
	}
	serviceAddr := viper.GetString("SERVICE_ADDR")
	if serviceAddr == "" {
		serviceAddr = defaultAppAddr
	}
	servicePort := viper.GetInt("SERVICE_PORT")
	if servicePort == 0 {
		servicePort = defaultAppPort
	}

	cfg.ServiceAddr = serviceAddr
	cfg.ServicePort = servicePort

	staticDir := viper.GetString("STATIC_DIR")
	if staticDir == "" {
		staticDir = defaultStaticDirName
	}
	cfg.StaticDir = staticDir

	cfg.RequestTimeout = time.Duration(viper.GetInt("REQUEST_TIMEOUT")) * time.Second

	return nil
}

// This method will retrieve all applicable log levels to pass
// to logrus hooks based on the defined base log level and debug
// parameter.
func (cfg *ServiceConfig) GetLogLevels(debug bool) []log.Level {
	baseLLevel := cfg.LogLevel
	if debug {
		baseLLevel = log.DebugLevel
	}

	lLevels := []log.Level{}

	for i := 0; i <= int(baseLLevel); i++ {
		lLevels = append(lLevels, log.Level(i))
	}
	return lLevels
}

func (c *ServiceConfig) GetStaticDir() string {
	return fmt.Sprintf("./%s", defaultStaticDirName)
}

// GetServiceAddr will retrieve local address of the service.
func (cfg *ServiceConfig) GetServiceAddr() string {
	return fmt.Sprintf("%s:%d", cfg.ServiceAddr, cfg.ServicePort)
}

// GetServiceAddr will retrieve local address of the service.
func (cfg *ServiceConfig) GetPrometheusAddr() string {
	if cfg.ServiceEnv == ServiceEnvProduction || cfg.ServiceEnv == ServiceEnvStaging {
		return fmt.Sprintf("%s:%d", cfg.PrometheusServerAddr, cfg.PrometheusServerPort)
	}
	return ""
}

// NewConfig will create new config.
func NewConfig() (*ServiceConfig, error) {
	viper.AutomaticEnv()

	cfg := &ServiceConfig{}
	err := func(funcs ...func() error) error {
		for _, f := range funcs {
			err := f()
			if err != nil {
				return err
			}
		}
		return nil
	}(
		cfg.loadServerEnvironment,
		cfg.loadPrometheusEnvironment,
		cfg.loadTelemetryEnvironment,
	)

	return cfg, err
}
