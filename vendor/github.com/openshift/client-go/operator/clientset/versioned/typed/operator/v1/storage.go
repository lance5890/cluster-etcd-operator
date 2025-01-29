// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	context "context"

	operatorv1 "github.com/openshift/api/operator/v1"
	applyconfigurationsoperatorv1 "github.com/openshift/client-go/operator/applyconfigurations/operator/v1"
	scheme "github.com/openshift/client-go/operator/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// StoragesGetter has a method to return a StorageInterface.
// A group's client should implement this interface.
type StoragesGetter interface {
	Storages() StorageInterface
}

// StorageInterface has methods to work with Storage resources.
type StorageInterface interface {
	Create(ctx context.Context, storage *operatorv1.Storage, opts metav1.CreateOptions) (*operatorv1.Storage, error)
	Update(ctx context.Context, storage *operatorv1.Storage, opts metav1.UpdateOptions) (*operatorv1.Storage, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, storage *operatorv1.Storage, opts metav1.UpdateOptions) (*operatorv1.Storage, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*operatorv1.Storage, error)
	List(ctx context.Context, opts metav1.ListOptions) (*operatorv1.StorageList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *operatorv1.Storage, err error)
	Apply(ctx context.Context, storage *applyconfigurationsoperatorv1.StorageApplyConfiguration, opts metav1.ApplyOptions) (result *operatorv1.Storage, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, storage *applyconfigurationsoperatorv1.StorageApplyConfiguration, opts metav1.ApplyOptions) (result *operatorv1.Storage, err error)
	StorageExpansion
}

// storages implements StorageInterface
type storages struct {
	*gentype.ClientWithListAndApply[*operatorv1.Storage, *operatorv1.StorageList, *applyconfigurationsoperatorv1.StorageApplyConfiguration]
}

// newStorages returns a Storages
func newStorages(c *OperatorV1Client) *storages {
	return &storages{
		gentype.NewClientWithListAndApply[*operatorv1.Storage, *operatorv1.StorageList, *applyconfigurationsoperatorv1.StorageApplyConfiguration](
			"storages",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *operatorv1.Storage { return &operatorv1.Storage{} },
			func() *operatorv1.StorageList { return &operatorv1.StorageList{} },
		),
	}
}
