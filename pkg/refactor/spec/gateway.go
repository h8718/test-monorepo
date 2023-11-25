package spec

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

type Plugins struct {
	Auth     []PluginConfig `json:"auth,omitempty"     yaml:"auth"`
	Inbound  []PluginConfig `json:"inbound,omitempty"  yaml:"inbound"`
	Target   PluginConfig   `json:"target,omitempty"   yaml:"target"`
	Outbound []PluginConfig `json:"outbound,omitempty" yaml:"outbound"`
}

type EndpointFile struct {
	DirektivAPI    string   `json:"direktiv_api,omitempty" yaml:"direktiv_api"`
	Methods        []string `json:"methods"                yaml:"methods"`
	PathExtension  string   `json:"path_extension"         yaml:"path_extension"`
	AllowAnonymous bool     `json:"allow_anonymous"        yaml:"allow_anonymous"`
	Plugins        Plugins  `json:"plugins"                yaml:"plugins"`
	Timeout        int      `json:"timeout"                yaml:"timeout"`
}

type ConsumerFile struct {
	DirektivAPI string   `json:"direktiv_api,omitempty" yaml:"direktiv_api"`
	Username    string   `json:"username"               yaml:"username"`
	Password    string   `json:"password"               yaml:"password"`
	APIKey      string   `json:"api_key"                yaml:"api_key"`
	Tags        []string `json:"tags"                   yaml:"tags"`
	Groups      []string `json:"groups"                 yaml:"groups"`
}

type PluginConfig struct {
	Type          string      `json:"type"          yaml:"type"`
	Configuration interface{} `json:"configuration" yaml:"configuration"`
}

func ParseConsumerFile(data []byte) (*ConsumerFile, error) {
	res := &ConsumerFile{}
	err := yaml.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	if !strings.HasPrefix(res.DirektivAPI, "consumer/v1") {
		return nil, fmt.Errorf("invalid consumer api version")
	}

	return res, nil
}

func ParseEndpointFile(data []byte) (*EndpointFile, error) {
	res := &EndpointFile{}
	err := yaml.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	if !strings.HasPrefix(res.DirektivAPI, "endpoint/v1") {
		return nil, fmt.Errorf("invalid endpoint api version")
	}

	return res, nil
}
