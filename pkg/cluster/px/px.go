package px

import (
	"github.com/harsh-px/px-operator/pkg/cluster"
	"github.com/sirupsen/logrus"
)

type pxCluster struct {
}

const pxClusterProviderName = "px"

func (p *pxCluster) Create(obj interface{}) error {
	logrus.Infof("creating a new portworx cluster")
	return nil
}

func (p *pxCluster) Upgrade(old interface{}, new interface{}) error {
	logrus.Infof("upgrading px cluster")
	return nil
}

func (p *pxCluster) Destroy(obj interface{}) error {
	return nil
}

// NewPXClusterProvider creates a new PX cluster
func NewPXClusterProvider(conf interface{}) (cluster.Cluster, error) {
	return &pxCluster{}, nil
}

func init() {
	cluster.Register(pxClusterProviderName, NewPXClusterProvider)
}
