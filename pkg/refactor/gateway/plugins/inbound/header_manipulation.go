package inbound

import (
	"net/http"

	"github.com/direktiv/direktiv/pkg/refactor/core"
	"github.com/direktiv/direktiv/pkg/refactor/gateway/plugins"
)

const (
	HeaderManipulation = "header-manipulation"
)

type HeaderManipulationConfig struct {
	HeadersToAdd    map[string]string `json:"headers_to_add"    mapstructure:"headers_to_add" yaml:"headers_to_add"`
	HeadersToModify map[string]string `json:"headers_to_modify" yaml:"headers_to_modify"`
	HeadersToRemove []string          `json:"headers_to_remove" yaml:"headers_to_remove"`
}

type HeaderManipulationPlugin struct {
	configuration *HeaderManipulationConfig
}

func ConfigureHeaderManipulation(config interface{}, _ string) (core.PluginInstance, error) {
	headerManipulationConfig := &HeaderManipulationConfig{}

	err := plugins.ConvertConfig(config, headerManipulationConfig)
	if err != nil {
		return nil, err
	}

	return &HeaderManipulationPlugin{
		configuration: headerManipulationConfig,
	}, nil
}

func (hp *HeaderManipulationPlugin) Config() interface{} {
	return hp.configuration
}

func (hp *HeaderManipulationPlugin) ExecutePlugin(_ *core.ConsumerFile,
	_ http.ResponseWriter, r *http.Request,
) bool {
	for key, value := range hp.configuration.HeadersToAdd {
		r.Header.Add(key, value)
	}

	for key, value := range hp.configuration.HeadersToModify {
		r.Header.Set(key, value)
	}

	for _, key := range hp.configuration.HeadersToRemove {
		r.Header.Del(key)
	}

	return true
}

func (hp *HeaderManipulationPlugin) Type() string {
	return HeaderManipulation
}

//nolint:gochecknoinits
func init() {
	plugins.AddPluginToRegistry(plugins.NewPluginBase(
		HeaderManipulation,
		plugins.InboundPluginType,
		ConfigureHeaderManipulation))
}
