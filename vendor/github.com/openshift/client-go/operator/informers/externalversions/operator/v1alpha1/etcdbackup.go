// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	context "context"
	time "time"

	apioperatorv1alpha1 "github.com/openshift/api/operator/v1alpha1"
	versioned "github.com/openshift/client-go/operator/clientset/versioned"
	internalinterfaces "github.com/openshift/client-go/operator/informers/externalversions/internalinterfaces"
	operatorv1alpha1 "github.com/openshift/client-go/operator/listers/operator/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// EtcdBackupInformer provides access to a shared informer and lister for
// EtcdBackups.
type EtcdBackupInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() operatorv1alpha1.EtcdBackupLister
}

type etcdBackupInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewEtcdBackupInformer constructs a new informer for EtcdBackup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewEtcdBackupInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredEtcdBackupInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredEtcdBackupInformer constructs a new informer for EtcdBackup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredEtcdBackupInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1alpha1().EtcdBackups().List(context.Background(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1alpha1().EtcdBackups().Watch(context.Background(), options)
			},
			ListWithContextFunc: func(ctx context.Context, options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1alpha1().EtcdBackups().List(ctx, options)
			},
			WatchFuncWithContext: func(ctx context.Context, options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1alpha1().EtcdBackups().Watch(ctx, options)
			},
		},
		&apioperatorv1alpha1.EtcdBackup{},
		resyncPeriod,
		indexers,
	)
}

func (f *etcdBackupInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredEtcdBackupInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *etcdBackupInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apioperatorv1alpha1.EtcdBackup{}, f.defaultInformer)
}

func (f *etcdBackupInformer) Lister() operatorv1alpha1.EtcdBackupLister {
	return operatorv1alpha1.NewEtcdBackupLister(f.Informer().GetIndexer())
}
