package cluster

import (
	"github.com/sirupsen/logrus"
)

// This package follows the factory pattern.
// Reference: http://matthewbrown.io/2016/01/23/factory-pattern-in-golang/

type ClusterInitFunction func(conf interface{}) (Cluster, error)

// Cluster an interface to manage a storage cluster
type Cluster interface {
	// Create creates the given cluster
	Create(obj interface{}) error
	// Upgrade upgrades the given cluster from old to new
	Upgrade(old interface{}, new interface{}) error
	// Destory destroys all components of the given cluster
	Destroy(obj interface{}) error
}

var clusterFactories = make(map[string]ClusterInitFunction)

func Register(name string, initFunc ClusterInitFunction) {
	if initFunc == nil {
		logrus.Fatalf("nil initFunc provided by cluster provider: %s", name)
	}

	_, registered := clusterFactories[name]
	if registered {
		logrus.Errorf("Cluster %s already registered. Ignoring.", name)
		return
	}

	clusterFactories[name] = initFunc
}
