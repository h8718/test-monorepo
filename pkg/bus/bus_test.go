package bus

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestConfigNodefinderEmpty(t *testing.T) {

	config := &Config{}

	_, err := NewBus(config)
	if err == nil {
		t.Fail()
	}

}

func TestBasicBus(t *testing.T) {

	config := DefaultConfig()

	bus, err := NewBus(config)
	if err != nil {
		t.Fail()
	}
	defer bus.Stop()

	go bus.Start()

	ch := make(chan bool, 1)
	bus.WaitTillConnected(ch)
	<-ch

	if len(bus.Status().Nodes) != 1 {
		t.Fail()
	}

}

func printStatus(bus string, status *Status, t *testing.T) {

	t.Logf("bus: %s", bus)
	t.Logf("status: %s", status.Status)
	t.Logf("nodes: %d", len(status.Nodes))

}

func startBus(config *Config) (*Bus, error) {

	d, err := os.MkdirTemp(os.TempDir(), fmt.Sprintf("%v", time.Now().UnixMilli()))
	config.DataPath = d

	bus, err := NewBus(config)
	if err != nil {
		return nil, err
	}

	go bus.Start()
	ch := make(chan bool, 1)
	bus.WaitTillConnected(ch)
	ret := <-ch

	if !ret {
		return nil, fmt.Errorf("bus did not start in time")
	}

	return bus, nil
}

func TestClusterBus(t *testing.T) {

	bus1, err := startBus(DefaultConfig())
	if err != nil {
		t.Fail()
	}
	defer bus1.Stop()

	config2 := DefaultConfig()
	config2.NSQDListen = "0.0.0.0:5550"
	config2.LookupListen = "0.0.0.0:5551"
	config2.PREFIX = "[JENSJENS] "

	nfs := NewNodefinderStatic(
		[]string{
			// "127.0.0.1:5551",
			"127.0.0.1:4160",
		},
	)
	config2.Nodefinder = nfs
	bus2, err := startBus(config2)
	if err != nil {
		t.FailNow()
	}
	defer bus2.Stop()

	// config3 := DefaultConfig()
	// config3.NSQDListen = "0.0.0.0:4200"
	// config3.LookupListen = "0.0.0.0:4210"

	// nfs2 := NewNodefinderStatic(
	// 	[]string{
	// 		"127.0.0.1:4160",
	// 		"127.0.0.1:4210",
	// 	},
	// )
	// config3.Nodefinder = nfs2
	// bus3, err := startBus(config3)
	// if err != nil {
	// 	t.FailNow()
	// }
	// defer bus3.Stop()

	// config := nsq.NewConfig()

	// prod, err := nsq.NewProducer("127.0.0.1:5550", config)
	// prod.Publish("JENS", []byte("GERKE"))

	// cons, err := nsq.NewConsumer("JENS", "sjjs", config)

	// err = cons.ConnectToNSQLookupd("127.0.0.1:5551")

	// t.Logf("B!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!! %v %v %v", bus1.Status(), bus2.Status(), err)
	time.Sleep(3 * time.Second)
	printStatus("bus1", bus1.Status(), t)
	printStatus("bus2", bus2.Status(), t)
	// printStatus("bus3", bus3.Status(), t)

}
