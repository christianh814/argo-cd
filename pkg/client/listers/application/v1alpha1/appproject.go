// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/argoproj/argo-cd/v2/pkg/apis/application/v1alpha1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// AppProjectLister helps list AppProjects.
// All objects returned here must be treated as read-only.
type AppProjectLister interface {
	// List lists all AppProjects in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AppProject, err error)
	// AppProjects returns an object that can list and get AppProjects.
	AppProjects(namespace string) AppProjectNamespaceLister
	AppProjectListerExpansion
}

// appProjectLister implements the AppProjectLister interface.
type appProjectLister struct {
	listers.ResourceIndexer[*v1alpha1.AppProject]
}

// NewAppProjectLister returns a new AppProjectLister.
func NewAppProjectLister(indexer cache.Indexer) AppProjectLister {
	return &appProjectLister{listers.New[*v1alpha1.AppProject](indexer, v1alpha1.Resource("appproject"))}
}

// AppProjects returns an object that can list and get AppProjects.
func (s *appProjectLister) AppProjects(namespace string) AppProjectNamespaceLister {
	return appProjectNamespaceLister{listers.NewNamespaced[*v1alpha1.AppProject](s.ResourceIndexer, namespace)}
}

// AppProjectNamespaceLister helps list and get AppProjects.
// All objects returned here must be treated as read-only.
type AppProjectNamespaceLister interface {
	// List lists all AppProjects in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AppProject, err error)
	// Get retrieves the AppProject from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AppProject, error)
	AppProjectNamespaceListerExpansion
}

// appProjectNamespaceLister implements the AppProjectNamespaceLister
// interface.
type appProjectNamespaceLister struct {
	listers.ResourceIndexer[*v1alpha1.AppProject]
}
