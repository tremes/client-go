package v1alpha1

import (
	configv1 "github.com/openshift/api/config/v1"
	operatorv1 "github.com/openshift/api/operator/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//
// DataGather provides data gather configuration options and status for the particular Insights data gathering.
//
// Compatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.
// +openshift:compatibility-gen:level=4
type DataGather struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// spec holds user settable values for configuration
	// +kubebuilder:validation:Required
	Spec DataGatherSpec `json:"spec"`
	// status holds observed values from the cluster. They may not be overridden.
	// +optional
	Status DataGatherStatus `json:"status"`
}

type DataGatherSpec struct {
	// dataPolicy allows user to enable additional global obfuscation of the IP addresses and base domain
	// in the Insights archive data. Valid values are "None" and "ObfuscateNetworking".
	// When set to None the data is not obfuscated.
	// When set to ObfuscateNetworking the IP addresses and the cluster domain name are obfuscated.
	// When omitted, this means no opinion and the platform is left to choose a reasonable default, which is subject to change over time.
	// The current default is None.
	// +optional
	DataPolicy DataPolicy `json:"dataPolicy"`
	// disabledGatherers is a list of gatherers to be excluded from the gathering. All the gatherers can be disabled by providing "all" value.
	// If all the gatherers are disabled, the Insights operator does not gather any data.
	// The particular gatherers IDs can be found at https://github.com/openshift/insights-operator/blob/master/docs/gathered-data.md.
	// Run the following command to get the names of last active gatherers:
	// "oc get insightsoperators.operator.openshift.io cluster -o json | jq '.status.gatherStatus.gatherers[].name'"
	// An example of disabling gatherers looks like this: `disabledGatherers: ["clusterconfig/machine_configs", "workloads/workload_info"]`
	// +optional
	DisabledGatherers []string `json:"disabledGatherers"`
}

const (
	// No data obfuscation
	NoPolicy DataPolicy = "None"
	// IP addresses and cluster domain name are obfuscated
	ObfuscateNetworking DataPolicy = "ObfuscateNetworking"
	// Data gathering is running
	Running State = "Running"
	// Data gathering is completed
	Completed State = "Completed"
	// Data gathering failed
	Failed State = "Failed"
	// Data gathering is pending
	Pending State = "Pending"
)

// dataPolicy declares valid data policy types
// +kubebuilder:validation:Enum="";None;ObfuscateNetworking
type DataPolicy string

// state declares valid gathering state types
// +kubebuilder:validation:Enum=Running;Completed;Failed
type State string

type DataGatherStatus struct {
	// state reflects the current state of the data gathering process.
	State State `json:"state,omitempty"`
	// gatherers is a list of active gatherers (and their statuses) in the last gathering.
	// +listType=atomic
	Gatherers []GathererStatus `json:"gatherers,omitempty"`
	// startTime is the time when Insights data gathering started.
	StartTime metav1.Time `json:"startTime,omitempty"`
	// finishTime is the time when Insights data gathering finished.
	FinishTime metav1.Time `json:"finishTime,omitempty"`
	// relatedObjects is a list of resources which are useful when debugging or inspecting the data
	// gathering Pod
	RelatedObjects []configv1.ObjectReference `json:"relatedObjects,omitempty"`
	// healthChecks provides basic information about active Insights health checks
	// in a cluster.
	// +listType=atomic
	// +optional
	HealthChecks []operatorv1.HealthCheck `json:"healthChecks,omitempty"`
	// conditions provide details on the status of the gatherer job.
	// +listType=atomic
	// +optional
	Conditions []metav1.Condition `json:"conditions"`
}

// gathererStatus represents information about a particular
// data gatherer.
type GathererStatus struct {
	// conditions provide details on the status of each gatherer.
	// +listType=atomic
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinItems=1
	Conditions []metav1.Condition `json:"conditions"`
	// name is the name of the gatherer.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MaxLength=256
	// +kubebuilder:validation:MinLength=5
	Name string `json:"name"`
	// lastGatherDuration represents the time spent gathering.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Type=string
	// +kubebuilder:validation:Pattern="^([1-9][0-9]*(\\.[0-9]+)?(ns|us|µs|ms|s|m|h))+$"
	LastGatherDuration metav1.Duration `json:"lastGatherDuration"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DataGatherList is a collection of items
//
// Compatibility level 4: No compatibility is provided, the API can change at any point for any reason. These capabilities should not be used by applications needing long term support.
// +openshift:compatibility-gen:level=4
type DataGatherList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []DataGather `json:"items"`
}
