package v1alpha1

import (
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Cluster describes a Portworx cluster
type Cluster struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`
	Spec            ClusterSpec `json:"spec,omitempty"`

	// Status represents the current status of the Portworx cluster
	// +optional
	Status ClusterStatus `json:"status,omitempty"`
}

// ClusterList is a list of Cluster objects in Kubernetes
type ClusterList struct {
	meta.TypeMeta `json:",inline"`
	meta.ListMeta `json:"metadata,omitempty"`

	Items []Cluster `json:"items"`
}

// Specification for a Cluster
type ClusterSpec struct {
	// Specific image to use on all nodes of the cluster.
	// +optional
	Image string `json:"image,omitempty"`

	// All nodes participating in this cluster
	Nodes []NodeSpec `json:"nodes,omitempty"`
}

type ClusterStatus struct {
	StatusInfo
	Conditions   []ClusterCondition `json:"conditions,omitempty"`
	NodeStatuses []NodeStatus       `json:"nodeStatuses,omitempty"`
}

// Node defines a single instance of available storage on a
// node and the appropriate options to apply to it to make it available
// to the cluster.
type Node struct {
	meta.TypeMeta   `json:",inline"`
	meta.ObjectMeta `json:"metadata,omitempty"`
	Spec            NodeSpec `json:"spec,omitempty"`

	// Status represents the current status of the storage node
	// +optional
	Status NodeStatus `json:"status,omitempty"`
}

// NodeList is a list of Node objects in Kubernetes.
type NodeList struct {
	meta.TypeMeta `json:",inline"`
	meta.ListMeta `json:"metadata,omitempty"`

	Items []Node `json:"items"`
}

// NodeSpec holds specification parameters for a Node.
type NodeSpec struct {
	// Request the storage node be scheduled on a specific node
	// Must have set either Node or NodeSelector
	// +optional
	Name string `json:"name,omitempty"`

	// Storage network if any
	Network *NodeNetwork `json:"network,omitempty"`

	// Raw block devices available on the Node to be used for storage.
	// Devices or Directories must be set and their use are specific to
	// the implementation
	// Must have set either Devices or Directories
	// +optional
	Devices []string `json:"devices,omitempty"`
}

// NodeNetwork specifies which network interfaces the Node should use for data
// and management transport
type NodeNetwork struct {
	Data string `json:"data"`
	Mgmt string `json:"mgmt"`
}

type StatusCondition struct {
	Time    meta.Time `json:"time,omitempty"`
	Message string    `json:"message,omitempty"`
	Reason  string    `json:"reason,omitempty"`
}

type StatusInfo struct {
	Ready bool `json:"ready"`

	// The following follow the same definition as PodStatus
	Message string `json:"message,omitempty"`
	Reason  string `json:"reason,omitempty"`
}

type ClusterConditionType string

const (
	ClusterConditionReady   ClusterConditionType = "Ready"
	ClusterConditionOffline ClusterConditionType = "Offline"
)

type ClusterCondition struct {
	StatusCondition
	Type ClusterConditionType `json:"type,omitempty"`
}

type NodeConditionType string

const (
	NodeConditionReady   NodeConditionType = "Ready"
	NodeConditionOffline NodeConditionType = "Offline"
)

type NodeCondition struct {
	StatusCondition
	Type NodeConditionType `json:"type,omitempty"`
}

type NodeStatus struct {
	StatusInfo
	Added      bool            `json:"added,omitempty"`
	Conditions []NodeCondition `json:"conditions,omitempty"`
	PodName    string          `json:"podName,omitempty"`
	Name       string          `json:"name,omitempty"`
}
