package cluster

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/serf/serf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestNodeConfig(t *testing.T) {
	newPort := 3333
	newTimeout := time.Minute

	config := DefaultConfig()

	assert.Equal(t, config.SerfPort, defaultSerfPort)
	assert.Equal(t, config.SerfReapTimeout, defaultReapTimeout)

	assert.Equal(t, config.NSQDPort, defaultNSQDPort)
	assert.Equal(t, config.NSQLookupPort, defaultNSQLookupPort)
	assert.Equal(t, config.NSQDListenHTTPPort, defaultNSQDListenHTTPPort)
	assert.Equal(t, config.NSQLookupListenHTTPPort, defaultNSQLookupListenHTTPPort)

	config.SerfPort = newPort
	config.SerfReapTimeout = time.Minute

	assert.Equal(t, config.SerfPort, newPort)
	assert.Equal(t, config.SerfReapTimeout, newTimeout)
}

func TestNewNode(t *testing.T) {
	config := DefaultConfig()

	node, err := NewNode(context.TODO(), config, NewNodeFinderStatic(nil), 200*time.Millisecond, zap.NewNop().Sugar())
	require.NoError(t, err)
	defer node.Stop()

	nodes := node.serfServer.NumNodes()
	assert.Equal(t, nodes, 1)
}

func rightNumber(nodes []*Node) bool {
	for i := 0; i < len(nodes); i++ {
		mnodes := nodes[i].serfServer.Members()
		count := len(mnodes)
		for j := 0; j < len(mnodes); j++ {
			if mnodes[j].Status == serf.StatusLeft {
				count--
			}
		}

		if count != len(nodes) {
			return false
		}

		nn, _ := nodes[i].bus.nodes()

		if len(nn.Producers) != len(nodes) {
			return false
		}

	}

	return true
}

func createConfig(t *testing.T, topics []string, change bool) (Config, []randomPort) {
	config := DefaultConfig()

	if change {
		config.NSQInactiveTimeout = 10 * time.Second
		config.NSQTombstoneTimeout = 5 * time.Second

		config.SerfReapTimeout = 3 * time.Second
		config.SerfReapInterval = 1 * time.Second
		config.SerfTombstoneTimeout = 5 * time.Second
	}

	config.NSQTopics = topics

	ports := getPorts(t)
	config.SerfPort = ports[0].port
	config.NSQDPort = ports[1].port
	config.NSQDListenHTTPPort = ports[2].port
	config.NSQLookupPort = ports[3].port
	config.NSQLookupListenHTTPPort = ports[4].port

	return config, ports
}

func createCluster(t *testing.T, count int, topics []string, change bool) ([]*Node, error) {
	logger := zap.NewNop().Sugar()

	// Preallocate slices to avoid reallocations
	nfNodes := make([]string, count)
	configs := make([]Config, count)
	ports := make([][]randomPort, count)
	finalNodes := make([]*Node, 0, count)

	hn, _ := os.Hostname()

	for i := 0; i < count; i++ {
		config, ports1 := createConfig(t, topics, change)
		nfNodes[i] = fmt.Sprintf("%s:%d", hn, config.SerfPort)
		configs[i] = config
		ports[i] = ports1
	}

	for i := 0; i < count; i++ {
		closePorts(ports[i])
	}

	nf := NewNodeFinderStatic(nfNodes)

	for i := 0; i < count; i++ {
		c := configs[i]
		node, err := NewNode(context.TODO(), c, nf, 200*time.Millisecond, logger)
		if err != nil {
			return nil, err
		}
		finalNodes = append(finalNodes, node)
	}

	return finalNodes, nil
}

func TestClusterSubscribe(t *testing.T) {
	count := 3
	nodes, err := createCluster(t, count, []string{"topic1", "topic2"}, false)
	require.NoError(t, err)

	for i := 0; i < count; i++ {
		defer nodes[i].Stop()
	}
	// check three node cluster
	require.Eventually(t, func() bool {
		return rightNumber(nodes)
	}, 10*time.Second, time.Millisecond*100)

	// Test topic1 subscription
	testTopic1Subscription(t, nodes)

	// Test topic2 subscription
	testTopic2Subscription(t, nodes)
}

func testTopic1Subscription(t *testing.T, nodes []*Node) {
	err := nodes[0].Publish("topic1", []byte("msg"))
	require.NoError(t, err)

	counter1 := &counterHandler{}
	unsub1, err := nodes[0].Subscribe("topic1", nodes[0].busChannelName, counter1.counter)
	require.NoError(t, err)
	defer unsub1()

	require.Eventually(t, func() bool {
		return counter1.cc == 1
	}, 30*time.Second, time.Millisecond*100)

	counter2 := &counterHandler{}
	unsub2, err := nodes[1].Subscribe("topic1", nodes[1].busChannelName, counter2.counter)
	require.NoError(t, err)
	defer unsub2()

	counter3 := &counterHandler{}
	unsub3, err := nodes[2].Subscribe("topic1", nodes[2].busChannelName, counter3.counter)
	require.NoError(t, err)
	defer unsub3()

	require.Eventually(t, func() bool {
		return nodes[0].Publish("topic1", []byte("msg")) == nil
	}, time.Millisecond*200, time.Millisecond*20)

	require.Eventually(t, func() bool {
		return counter1.cc == 2 && counter2.cc == 1 && counter3.cc == 1
	}, 30*time.Second, time.Millisecond*100)

	add := 10
	for i := 0; i < add; i++ {
		err = nodes[0].Publish("topic1", []byte("msg1"))
		require.NoError(t, err)
	}

	require.Eventually(t, func() bool {
		return counter1.cc == add+2 && counter2.cc == add+1 && counter3.cc == add+1
	}, 30*time.Second, time.Millisecond*100)
}

func testTopic2Subscription(t *testing.T, nodes []*Node) {
	err := nodes[0].Publish("topic2", []byte("msg"))
	require.NoError(t, err)

	// test single subscriber
	counter1 := &counterHandler{}
	unsub1, err := nodes[0].Subscribe("topic2", "shared", counter1.counter)
	require.NoError(t, err)
	defer unsub1()

	counter2 := &counterHandler{}
	unsub2, err := nodes[1].Subscribe("topic2", "shared", counter2.counter)
	require.NoError(t, err)
	defer unsub2()

	counter3 := &counterHandler{}
	unsub3, err := nodes[2].Subscribe("topic2", "shared", counter3.counter)
	require.NoError(t, err)
	defer unsub3()

	c := 5
	for i := 0; i < c; i++ {
		require.Eventually(t, func() bool {
			return nodes[0].Publish("topic2", []byte("msg1")) == nil
		}, time.Millisecond*200, time.Millisecond*20)
	}

	require.Eventually(t, func() bool {
		return counter1.cc+counter2.cc+counter3.cc >= c+1
	}, 30*time.Second, time.Millisecond*30, "did not get received events")
}

type counterHandler struct {
	cc int
}

var j int

func (ch *counterHandler) counter(msg []byte) error {
	j++
	ch.cc++

	return nil
}
