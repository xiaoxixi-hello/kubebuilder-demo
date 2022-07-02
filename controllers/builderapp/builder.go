package builderapp

import (
	v1 "k8s.io/api/apps/v1"
	"reflect"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

type ResourceAppChangedPredicate struct {
	predicate.Funcs
}

type Predicate interface {
	// Create returns true if the Create event should be processed
	Create(event.CreateEvent) bool

	// Delete returns true if the Delete event should be processed
	Delete(event.DeleteEvent) bool

	// Update returns true if the Update event should be processed
	Update(event.UpdateEvent) bool

	// Generic returns true if the Generic event should be processed
	Generic(event.GenericEvent) bool
}

func (rl *ResourceAppChangedPredicate) Create(e event.CreateEvent) bool {
	return true
}

func (rl *ResourceAppChangedPredicate) Update(e event.UpdateEvent) bool {
	d1, ok1 := e.ObjectOld.(*v1.Deployment)
	d2, ok2 := e.ObjectNew.(*v1.Deployment)
	if ok1 && ok2 {
		if reflect.DeepEqual(d1.Spec, d2.Spec) {
			return false
		}
	}
	return true
}
func (rl *ResourceAppChangedPredicate) Delete(e event.DeleteEvent) bool {
	return true
}
func (rl *ResourceAppChangedPredicate) Generic(e event.GenericEvent) bool {
	return true
}
