package core

import (
	"errors"
	"fmt"
	"time"
)

var ErrNotFound = errors.New("ErrNotFound")

// nolint:revive,stylecheck
type Config struct {
	ApiV1Port int `env:"DIREKTIV_API_V1_PORT" envDefault:"6665"`
	ApiV2Port int `env:"DIREKTIV_API_V2_PORT" envDefault:"6667"`
	GrpcPort  int `env:"DIREKTIV_GRPC_PORT"   envDefault:"6666"`

	Prometheus     string `env:"DIREKTIV_PROMETHEUS_BACKEND"`
	OpenTelemetry  string `env:"DIREKTIV_OPEN_TELEMETRY_BACKEND"`
	EnableEventing bool   `env:"DIREKTIV_ENABLE_EVENTING"`

	SecretKey string `env:"DIREKTIV_SECRET_KEY" envDefault:"01234567890123456789012345678912"`

	DB string `env:"DIREKTIV_DB,notEmpty"`

	EnableDocker bool `env:"DIREKITV_ENABLE_DOCKER"`

	KnativeServiceAccount string `env:"DIREKTIV_KNATIVE_SERVICE_ACCOUNT"`
	KnativeNamespace      string `env:"DIREKTIV_KNATIVE_NAMESPACE"`
	KnativeIngressClass   string `env:"DIREKTIV_KNATIVE_INGRESS_CLASS"`
	KnativeSidecar        string `env:"DIREKTIV_KNATIVE_SIDECAR"`
	KnativeMaxScale       int    `env:"DIREKTIV_KNATIVE_MAX_SCALE"  envDefault:"5"`
	KnativeNetShape       string `env:"DIREKTIV_KNATIVE_NET_SHAPE"`

	KnativeSizeMemorySmall  int `env:"DIREKTIV_KNATIVE_SIZE_MEMORY_SMALL"      envDefault:"512"`
	KnativeSizeMemoryMedium int `env:"DIREKTIV_KNATIVE_SIZE_MEMORY_MEDIUM"     envDefault:"1024"`
	KnativeSizeMemoryLarge  int `env:"DIREKTIV_KNATIVE_SIZE_MEMORY_LARGE"      envDefault:"2048"`

	KnativeSizeCPUSmall  string `env:"DIREKTIV_KNATIVE_SIZE_CPU_SMALL"    envDefault:"250m"`
	KnativeSizeCPUMedium string `env:"DIREKTIV_KNATIVE_SIZE_CPU_MEDIUM"   envDefault:"500m"`
	KnativeSizeCPULarge  string `env:"DIREKTIV_KNATIVE_SIZE_CPU_LARGE"    envDefault:"1"`

	KnativeSizeDiskSmall  int `env:"DIREKTIV_KNATIVE_SIZE_DISK_SMALL"   envDefault:"256"`
	KnativeSizeDiskMedium int `env:"DIREKTIV_KNATIVE_SIZE_DISK_MEDIUM"  envDefault:"1024"`
	KnativeSizeDiskLarge  int `env:"DIREKTIV_KNATIVE_SIZE_DISK_LARGE"   envDefault:"4096"`

	FunctionsTimeout int    `env:"DIREKTIV_FUNCTIONS_TIMEOUT" envDefault:"7200"`
	LogFormat        string `env:"DIREKTIV_LOG_JSON"`
	LogDebug         bool   `env:"DIREKTIV_DEBUG"`
}

func (conf *Config) GetFunctionsTimeout() time.Duration {
	return time.Second * time.Duration(conf.FunctionsTimeout)
}

func (conf *Config) IsValid() error {
	err := conf.checkInvalidEmptyFields()
	if err != nil {
		return err
	}

	return nil
}

func (conf *Config) checkInvalidEmptyFields() error {
	var invalidEmptyFields []string

	// knative setting only required when docker mode is disabled.
	if !conf.EnableDocker {
		if conf.KnativeServiceAccount == "" {
			invalidEmptyFields = append(invalidEmptyFields, "DIREKTIV_KNATIVE_SERVICE_ACCOUNT")
		}
		if conf.KnativeNamespace == "" {
			invalidEmptyFields = append(invalidEmptyFields, "DIREKTIV_KNATIVE_NAMESPACE")
		}
		if conf.KnativeIngressClass == "" {
			invalidEmptyFields = append(invalidEmptyFields, "DIREKTIV_KNATIVE_INGRESS_CLASS")
		}
		if conf.KnativeSidecar == "" {
			invalidEmptyFields = append(invalidEmptyFields, "DIREKTIV_KNATIVE_SIDECAR")
		}
	}

	if len(invalidEmptyFields) == 0 {
		return nil
	}

	return fmt.Errorf("following fields are required but got emty strings: %v", invalidEmptyFields)
}

type Version struct {
	UnixTime int64 `json:"unix_time"`
}

type App struct {
	Version *Version
	Config  *Config

	ServiceManager  ServiceManager
	RegistryManager RegistryManager
	GatewayManager  GatewayManager
}
