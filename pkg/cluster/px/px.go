package px

import (
	portworx "github.com/harsh-px/px-operator/pkg/apis/portworx.com"
	"github.com/harsh-px/px-operator/pkg/cluster"
	"github.com/sirupsen/logrus"
)

type pxCluster struct {
}

func (p *pxCluster) Create(obj interface{}) error {
	logrus.Infof("creating a new portworx cluster")
	// TODO add gatekeeper check to ensure only one cluster is running
	return nil
}

func (p *pxCluster) Upgrade(new interface{}) error {
	logrus.Infof("upgrading px cluster")
	return nil
}

func (p *pxCluster) Destroy(obj interface{}) error {
	logrus.Infof("destroying px cluster")
	return nil
}

// NewPXClusterProvider creates a new PX cluster
func NewPXClusterProvider(conf interface{}) (cluster.Cluster, error) {
	return &pxCluster{}, nil
}

func init() {
	cluster.Register(portworx.GroupName, NewPXClusterProvider)
}
