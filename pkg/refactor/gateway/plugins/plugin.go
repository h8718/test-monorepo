package plugins

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/direktiv/direktiv/pkg/refactor/core"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
)

// nolint
const (
	ConsumerUserHeader   = "Direktiv-Consumer-User"
	ConsumerTagsHeader   = "Direktiv-Consumer-Tags"
	ConsumerGroupsHeader = "Direktiv-Consumer-Groups"
)

var registry = make(map[string]Plugin)

type PluginType string

var (
	AuthPluginType     PluginType = "auth"
	TargetPluginType   PluginType = "target"
	InboundPluginType  PluginType = "inbound"
	OutboundPluginType PluginType = "outbound"
)

type Plugin interface {
	Configure(config interface{}, namespace string) (core.PluginInstance, error)
	Name() string
	Type() PluginType
}

// PluginBase is a basic implementation of the Plugin interface.
type PluginBase struct {
	pname    string
	ptype    PluginType
	configFn func(interface{}, string) (core.PluginInstance, error)
}

func (p PluginBase) Name() string {
	return p.pname
}

func (p PluginBase) Type() PluginType {
	return p.ptype
}

func (p PluginBase) Configure(config interface{}, ns string) (core.PluginInstance, error) {
	return p.configFn(config, ns)
}

func NewPluginBase(pname string, ptype PluginType,
	configFn func(interface{}, string) (core.PluginInstance, error),
) Plugin {
	return &PluginBase{
		pname:    pname,
		ptype:    ptype,
		configFn: configFn,
	}
}

// ConvertConfig converts an interface into the config struct of the plugin.
// It is used in the `Configure` function of the Plugin.
func ConvertConfig(config interface{}, target interface{}) error {
	if config != nil {
		err := mapstructure.Decode(config, target)
		if err != nil {
			return errors.Wrap(err, "configuration invalid")
		}
	}

	return nil
}

func GetAllPlugins() map[string]Plugin {
	return registry
}

func AddPluginToRegistry(plugin Plugin) {
	slog.Info("adding plugin", slog.String("name", plugin.Name()))
	registry[plugin.Name()] = plugin
}

func GetPluginFromRegistry(plugin string) (Plugin, error) {
	p, ok := registry[plugin]
	if !ok {
		return nil, fmt.Errorf("plugin %s does not exist", plugin)
	}

	return p, nil
}

var (
	URLParamCtxKey       = &ContextKey{"URLParamContext"}
	ConsumersParamCtxKey = &ContextKey{"ConsumersParamCtxKey"}
)

type ContextKey struct {
	name string
}

func (k *ContextKey) String() string {
	return "plugin context value " + k.name
}

func ReportError(w http.ResponseWriter, status int, msg string, err error) {
	slog.Error("can not process plugin", slog.String("error", err.Error()))
	w.WriteHeader(status)
	errMsg := fmt.Sprintf("%s: %s", msg, err.Error())

	// nolint
	w.Write([]byte(errMsg))
}

func ReportNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	// nolint
	w.Write([]byte("not found"))
}

func IsJSON(str string) bool {
	var js json.RawMessage

	return json.Unmarshal([]byte(str), &js) == nil
}
