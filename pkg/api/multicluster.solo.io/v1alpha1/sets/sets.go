// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./sets.go -destination mocks/sets.go

package v1alpha1sets

import (
	multicluster_solo_io_v1alpha1 "github.com/solo-io/skv2/pkg/api/multicluster.solo.io/v1alpha1"

	sksets "github.com/solo-io/skv2/contrib/pkg/sets"
	"github.com/solo-io/skv2/pkg/ezkube"
	"k8s.io/apimachinery/pkg/util/sets"
)

type KubernetesClusterSet interface {
	Keys() sets.String
	List() []*multicluster_solo_io_v1alpha1.KubernetesCluster
	Map() map[string]*multicluster_solo_io_v1alpha1.KubernetesCluster
	Insert(kubernetesCluster ...*multicluster_solo_io_v1alpha1.KubernetesCluster)
	Equal(kubernetesClusterSet KubernetesClusterSet) bool
	Has(kubernetesCluster *multicluster_solo_io_v1alpha1.KubernetesCluster) bool
	Delete(kubernetesCluster *multicluster_solo_io_v1alpha1.KubernetesCluster)
	Union(set KubernetesClusterSet) KubernetesClusterSet
	Difference(set KubernetesClusterSet) KubernetesClusterSet
	Intersection(set KubernetesClusterSet) KubernetesClusterSet
	Find(id ezkube.ResourceId) (*multicluster_solo_io_v1alpha1.KubernetesCluster, error)
	Length() int
}

func makeGenericKubernetesClusterSet(kubernetesClusterList []*multicluster_solo_io_v1alpha1.KubernetesCluster) sksets.ResourceSet {
	var genericResources []ezkube.ResourceId
	for _, obj := range kubernetesClusterList {
		genericResources = append(genericResources, obj)
	}
	return sksets.NewResourceSet(genericResources...)
}

type kubernetesClusterSet struct {
	set sksets.ResourceSet
}

func NewKubernetesClusterSet(kubernetesClusterList ...*multicluster_solo_io_v1alpha1.KubernetesCluster) KubernetesClusterSet {
	return &kubernetesClusterSet{set: makeGenericKubernetesClusterSet(kubernetesClusterList)}
}

func NewKubernetesClusterSetFromList(kubernetesClusterList *multicluster_solo_io_v1alpha1.KubernetesClusterList) KubernetesClusterSet {
	list := make([]*multicluster_solo_io_v1alpha1.KubernetesCluster, 0, len(kubernetesClusterList.Items))
	for idx := range kubernetesClusterList.Items {
		list = append(list, &kubernetesClusterList.Items[idx])
	}
	return &kubernetesClusterSet{set: makeGenericKubernetesClusterSet(list)}
}

func (s *kubernetesClusterSet) Keys() sets.String {
	return s.set.Keys()
}

func (s *kubernetesClusterSet) List() []*multicluster_solo_io_v1alpha1.KubernetesCluster {
	var kubernetesClusterList []*multicluster_solo_io_v1alpha1.KubernetesCluster
	for _, obj := range s.set.List() {
		kubernetesClusterList = append(kubernetesClusterList, obj.(*multicluster_solo_io_v1alpha1.KubernetesCluster))
	}
	return kubernetesClusterList
}

func (s *kubernetesClusterSet) Map() map[string]*multicluster_solo_io_v1alpha1.KubernetesCluster {
	newMap := map[string]*multicluster_solo_io_v1alpha1.KubernetesCluster{}
	for k, v := range s.set.Map() {
		newMap[k] = v.(*multicluster_solo_io_v1alpha1.KubernetesCluster)
	}
	return newMap
}

func (s *kubernetesClusterSet) Insert(
	kubernetesClusterList ...*multicluster_solo_io_v1alpha1.KubernetesCluster,
) {
	for _, obj := range kubernetesClusterList {
		s.set.Insert(obj)
	}
}

func (s *kubernetesClusterSet) Has(kubernetesCluster *multicluster_solo_io_v1alpha1.KubernetesCluster) bool {
	return s.set.Has(kubernetesCluster)
}

func (s *kubernetesClusterSet) Equal(
	kubernetesClusterSet KubernetesClusterSet,
) bool {
	return s.set.Equal(makeGenericKubernetesClusterSet(kubernetesClusterSet.List()))
}

func (s *kubernetesClusterSet) Delete(KubernetesCluster *multicluster_solo_io_v1alpha1.KubernetesCluster) {
	s.set.Delete(KubernetesCluster)
}

func (s *kubernetesClusterSet) Union(set KubernetesClusterSet) KubernetesClusterSet {
	return NewKubernetesClusterSet(append(s.List(), set.List()...)...)
}

func (s *kubernetesClusterSet) Difference(set KubernetesClusterSet) KubernetesClusterSet {
	newSet := s.set.Difference(makeGenericKubernetesClusterSet(set.List()))
	return &kubernetesClusterSet{set: newSet}
}

func (s *kubernetesClusterSet) Intersection(set KubernetesClusterSet) KubernetesClusterSet {
	newSet := s.set.Intersection(makeGenericKubernetesClusterSet(set.List()))
	var kubernetesClusterList []*multicluster_solo_io_v1alpha1.KubernetesCluster
	for _, obj := range newSet.List() {
		kubernetesClusterList = append(kubernetesClusterList, obj.(*multicluster_solo_io_v1alpha1.KubernetesCluster))
	}
	return NewKubernetesClusterSet(kubernetesClusterList...)
}

func (s *kubernetesClusterSet) Find(id ezkube.ResourceId) (*multicluster_solo_io_v1alpha1.KubernetesCluster, error) {
	obj, err := s.set.Find(&multicluster_solo_io_v1alpha1.KubernetesCluster{}, id)
	if err != nil {
		return nil, err
	}

	return obj.(*multicluster_solo_io_v1alpha1.KubernetesCluster), nil
}

func (s *kubernetesClusterSet) Length() int {
	return s.set.Length()
}
