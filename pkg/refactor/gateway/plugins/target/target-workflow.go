package target

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/direktiv/direktiv/pkg/refactor/core"
	"github.com/direktiv/direktiv/pkg/refactor/gateway/plugins"
	"github.com/direktiv/direktiv/pkg/refactor/spec"
)

const (
	TargetFlowPluginName = "target-flow"
)

type TargetWorkflowConfig struct {
	Namespace   string `yaml:"namespace" mapstructure:"namespace"`
	Flow        string `yaml:"flow" mapstructure:"flow"`
	Async       bool   `yaml:"async" mapstructure:"async"`
	ContentType string `yaml:"content_type"  mapstructure:"content_type"`
}

// TargetFlowPlugin executes a flow in a configured namespace.
// Flows can be executed async and sync
type TargetFlowPlugin struct {
	config *TargetWorkflowConfig
}

func ConfigureTargetFlowPlugin(config interface{}, ns string) (plugins.PluginInstance, error) {
	targetflowConfig := &TargetWorkflowConfig{
		Async: true,
	}

	err := plugins.ConvertConfig(TargetNamespaceFilePluginName, config, targetflowConfig)
	if err != nil {
		return nil, err
	}

	// set default to gateway namespace
	if targetflowConfig.Namespace == "" {
		targetflowConfig.Namespace = ns
	}

	// throw error if non magic namespace targets different namespace
	if targetflowConfig.Namespace != ns && ns != core.MagicalGatewayNamespace {
		return nil, fmt.Errorf("plugin can not target different namespace")
	}

	return &TargetFlowPlugin{
		config: targetflowConfig,
	}, nil
}

func (tf TargetFlowPlugin) Config() interface{} {
	return tf.config
}

func (tf TargetFlowPlugin) ExecutePlugin(c *spec.ConsumerFile,
	w http.ResponseWriter, r *http.Request) bool {

	url, err := createURL(tf.config.Namespace, tf.config.Flow, tf.config.Async)
	if err != nil {
		plugins.ReportError(w, http.StatusInternalServerError,
			"can not create url", err)

		return false
	}

	client := http.Client{}

	req, err := http.NewRequestWithContext(r.Context(), http.MethodPost,
		url.String(), io.NopCloser(r.Body))
	if err != nil {
		plugins.ReportError(w, http.StatusInternalServerError,
			"can not create request", err)

		return false
	}
	defer r.Body.Close()

	resp, err := client.Do(req)
	if err != nil {
		plugins.ReportError(w, http.StatusInternalServerError,
			"can not execute flow", err)

		return false
	}

	// overwrite content type
	if tf.config.ContentType != "" {
		w.Header().Set("Content-Type", tf.config.ContentType)
	}

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		plugins.ReportError(w, http.StatusInternalServerError,
			"can not serve response", err)

		return false
	}
	resp.Body.Close()

	return true
}

//nolint:gochecknoinits
func init() {
	plugins.AddPluginToRegistry(plugins.NewPluginBase(
		TargetFlowPluginName,
		plugins.TargetPluginType,
		ConfigureTargetFlowPlugin))
}

func createURL(ns, flow string, wait bool) (*url.URL, error) {
	ex := "execute"
	if wait {
		ex = "wait"
	}

	// prefix with slash if set relative
	if !strings.HasPrefix(flow, "/") {
		flow = "/" + flow
	}

	urlString := fmt.Sprintf("http://localhost:%s/api/namespaces/%s/tree%s?op=%s&ref=latest",
		os.Getenv("DIREKTIV_API_V1_PORT"), ns, flow, ex)

	return url.Parse(urlString)
}
