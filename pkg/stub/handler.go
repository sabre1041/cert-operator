package stub

import (
	"context"

	"github.com/redhat-cop/cert-operator/pkg/apis/cache/v1alpha1"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
//	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	v1 "github.com/openshift/api/route/v1"
)

func NewHandler() sdk.Handler {
	return &Handler{}
}

type Handler struct {
	// Fill me
}

func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {
	switch o := event.Object.(type) {
	case *v1.Route:
		route := o
		if route.ObjectMeta.Annotations == nil || route.ObjectMeta.Annotations["openshift.io/managed.cert"] == "" {
			return nil
		}

		logrus.Infof("Found a route waiting for a cert : %v", route)

		//err := sdk.Create(newbusyBoxPod(o))
		//if err != nil && !errors.IsAlreadyExists(err) {
		//	logrus.Errorf("Failed to create busybox pod : %v", err)
		//	return err
		//}
	}
	return nil
}

// newbusyBoxPod demonstrates how to create a busybox pod
func newbusyBoxPod(cr *v1alpha1.Memcached) *corev1.Pod {
	labels := map[string]string{
		"app": "busy-box",
	}
	return &corev1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "busy-box",
			Namespace: cr.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(cr, schema.GroupVersionKind{
					Group:   v1alpha1.SchemeGroupVersion.Group,
					Version: v1alpha1.SchemeGroupVersion.Version,
					Kind:    "Memcached",
				}),
			},
			Labels: labels,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "busybox",
					Image:   "busybox",
					Command: []string{"sleep", "3600"},
				},
			},
		},
	}
}
