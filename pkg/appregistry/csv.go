package appregistry

import (
	//"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	olm "github.com/operator-framework/operator-lifecycle-manager/pkg/api/apis/operators/v1alpha1"
)

const (
	// Name of the section under which the list of owned and required list of
	// CRD(s) is specified inside an operator manifest.
	customResourceDefinitions = "customresourcedefinitions"

	// The yaml attribute that points to the name of an older
	// ClusterServiceVersion object that the current ClusterServiceVersion
	// replaces.
	replaces = "replaces"
)

// ClusterServiceVersion is a structured representation of cluster service
// version object(s) specified inside the 'clusterServiceVersions' section of
// an operator manifest.
type ClusterServiceVersion struct {
	// Type metadata.
	metav1.TypeMeta `json:",inline"`

	// Object metadata.
	metav1.ObjectMeta `json:"metadata"`

	// Spec is the raw representation of the 'spec' element of
	// ClusterServiceVersion object. Since we are
	// not interested in the content of spec we are not parsing it.
	Spec olm.ClusterServiceVersionSpec `json:"spec"`
}

// GetReplaces returns the name of the older ClusterServiceVersion object that
// is replaced by this ClusterServiceVersion object.
//
// If not defined, the function returns an empty string.
func (csv *ClusterServiceVersion) GetReplaces() (string, error) {
	return csv.Spec.Replaces, nil
}

// GetCustomResourceDefintions returns a list of owned and required
// CustomResourceDefinition object(s) specified inside the
// 'customresourcedefinitions' section of a ClusterServiceVersion 'spec'.
//
// owned represents the list of CRD(s) managed by this ClusterServiceVersion
// object.
// required represents the list of CRD(s) that this ClusterServiceVersion
// object depends on.
//
// If owned or required is not defined in the spec then an empty list is
// returned respectively.
func (csv *ClusterServiceVersion) GetCustomResourceDefintions() (owned []*CRDKey, required []*CRDKey, err error) {
	ownedCRDKeys := make([]*CRDKey, 0)
	for _, crd := range csv.Spec.CustomResourceDefinitions.Owned {
		key := &CRDKey{
			Kind: crd.Kind,
			Name: crd.Name,
			Version: crd.Version,
		}
		ownedCRDKeys = append(ownedCRDKeys, key)
	}

	requiredCRDKeys := make([]*CRDKey, 0)
	for _, crd := range csv.Spec.CustomResourceDefinitions.Owned {
		key := &CRDKey{
			Kind: crd.Kind,
			Name: crd.Name,
			Version: crd.Version,
		}
		requiredCRDKeys = append(requiredCRDKeys, key)
	}

	owned = ownedCRDKeys
	required = requiredCRDKeys
	return
}
