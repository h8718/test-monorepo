// nolint
package function

import (
	"fmt"
	"strconv"

	"github.com/mitchellh/hashstructure/v2"
)

type Config struct {
	Namespace string
	Name      string

	ServicePath  string
	WorkflowPath string

	Image string
	CMD   string
	Size  string
	Scale int

	Error error
}

func (c *Config) getId() string {
	str := fmt.Sprintf("%s-%s-%s-%s", c.Namespace, c.Name, c.ServicePath, c.WorkflowPath)
	v, err := hashstructure.Hash(str, hashstructure.FormatV2, nil)
	if err != nil {
		panic("unexpected hashstructure.Hash error: " + err.Error())
	}

	return fmt.Sprintf("obj-%d-obj", v)
}

func (c *Config) getValueHash() string {
	str := fmt.Sprintf("%s-%s-%s-%d", c.Image, c.CMD, c.Size, c.Scale)
	v, err := hashstructure.Hash(str, hashstructure.FormatV2, nil)
	if err != nil {
		panic("unexpected hashstructure.Hash error: " + err.Error())
	}

	return strconv.Itoa(int(v))
}

type Status interface {
	checks() any
	getId() string
	getValueHash() string
}

type ConfigStatus struct {
	Config *Config `json:"config"`
	Checks any     `json:"checks"`
}
