package direktiv

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/sisatech/toml"
)

const (

	// DirektivDebug enables debug on dirtektiv applications
	DirektivDebug = "DIREKTIV_DEBUG"

	direktivWorkflowNamespace = "DIREKTIV_WFNS"

	// flowConfig
	flowBind     = "DIREKTIV_FLOW_BIND"
	flowEndpoint = "DIREKTIV_FLOW_ENDPOINT"
	flowProtocol = "DIREKTIV_FLOW_PROTOCOL"
	flowExchange = "DIREKTIV_FLOW_EXCHANGE"
	flowSidecar  = "DIREKTIV_FLOW_SIDECAR"
	flowMaxScale = "DIREKTIV_FLOW_MAX_SCALE"

	ingressBind     = "DIREKTIV_INGRESS_BIND"
	ingressEndpoint = "DIREKTIV_INGRESS_ENDPOINT"

	isolateEndpoint = "DIREKTIV_ISOLATE_ENDPOINT"

	// DBConn database connection
	DBConn = "DIREKTIV_DB"

	// instance logging
	instanceLoggingDriver = "DIREKTIV_INSTANCE_LOGGING_DRIVER"

	isolatePodClean = "DIREKTIV_ISOLATE_PODCLEAN"
)

// Config is the configuration for workflow and runner server
type Config struct {
	FlowAPI struct {
		Bind            string
		Endpoint        string
		IsolateEndpoint string
		Exchange        string
		Sidecar         string
		Protocol        string
		MaxScale        int
	} `toml:"flowAPI"`

	IngressAPI struct {
		Bind     string
		Endpoint string
	} `toml:"ingressAPI"`

	Database struct {
		DB string
	}

	InstanceLogging struct {
		Driver string
	}

	VariablesStorage struct {
		Driver string
	}

	Isolates struct {
		CleanupPods int
	} `toml:"isolates"`
}

func setIP(config *Config, env string, value *net.IP) error {

	v := os.Getenv(env)

	if len(v) > 0 {
		*value = net.ParseIP(v)
		if len(*value) == 0 {
			return fmt.Errorf("can not parse IP %s", v)
		}
		log.Debugf("setting %s to %s", env, value.String())
	}

	return nil

}

func setInt(config *Config, env string, value *int) error {

	v := os.Getenv(env)
	if len(v) > 0 {
		i, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*value = i
		log.Debugf("setting %s to %d", env, i)
	}

	return nil

}

func setString(config *Config, env string, value *string) error {

	v := os.Getenv(env)
	if len(v) > 0 {
		*value = v
		log.Debugf("setting %s via env", env)
	}

	return nil

}

// ReadConfig reads the configuration file and overwrites with environment variables if set
func ReadConfig(file string) (*Config, error) {

	c := new(Config)

	localIP := net.ParseIP("127.0.0.1")

	// defaults
	c.FlowAPI.Bind = fmt.Sprintf("%s:7777", localIP)
	c.FlowAPI.Endpoint = c.FlowAPI.Bind
	c.FlowAPI.Sidecar = "vorteil/sidecar"
	c.FlowAPI.Protocol = "http"

	c.IngressAPI.Bind = fmt.Sprintf("%s:6666", localIP)
	c.IngressAPI.Endpoint = c.IngressAPI.Bind

	c.Isolates.CleanupPods = 1

	// read config file if exists
	if len(file) > 0 {

		log.Printf("read config file %s", file)

		/* #nosec */
		data, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}

		err = toml.Unmarshal(data, c)
		if err != nil {
			return nil, err
		}

	}

	// overwrite with envs
	ints := []struct {
		name  string
		value *int
	}{
		{isolatePodClean, &c.Isolates.CleanupPods},
	}

	for _, i := range ints {
		err := setInt(c, i.name, i.value)
		if err != nil {
			return nil, err
		}
	}

	strings := []struct {
		name  string
		value *string
	}{
		{DBConn, &c.Database.DB},
		{instanceLoggingDriver, &c.InstanceLogging.Driver},
		{flowBind, &c.FlowAPI.Bind},
		{flowEndpoint, &c.FlowAPI.Endpoint},
		{ingressBind, &c.IngressAPI.Bind},
		{ingressEndpoint, &c.FlowAPI.Endpoint},
		{isolateEndpoint, &c.FlowAPI.IsolateEndpoint},
		{flowExchange, &c.FlowAPI.Exchange},
		{flowSidecar, &c.FlowAPI.Sidecar},
		{flowProtocol, &c.FlowAPI.Protocol},
	}

	for _, i := range strings {
		err := setString(c, i.name, i.value)
		if err != nil {
			return nil, err
		}
	}

	err := setInt(c, flowMaxScale, &c.FlowAPI.MaxScale)
	if err != nil {
		return nil, err
	}

	if c.FlowAPI.MaxScale < 0 {
		c.FlowAPI.MaxScale = 0
		log.Debugf("setting %s out of bounds, setting to 0", flowMaxScale)
	}

	// test database is set
	if len(c.Database.DB) == 0 {
		return nil, fmt.Errorf("no database configured")
	}

	return c, nil

}
