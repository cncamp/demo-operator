/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	appsv1beta1 "github.com/cncamp/demo-operator/api/v1beta1"
)

// MyDaemonsetReconciler reconciles a MyDaemonset object
type MyDaemonsetReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=apps.cncamp.io,resources=mydaemonsets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=apps.cncamp.io,resources=mydaemonsets/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=apps.cncamp.io,resources=mydaemonsets/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the MyDaemonset object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *MyDaemonsetReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)
	myds := &appsv1beta1.MyDaemonset{}
	if err := r.Client.Get(ctx, req.NamespacedName, myds); err != nil {
		fmt.Println(err)
	}
	nl := &v1.NodeList{}
	if myds.Spec.Image != "" {
		if err := r.Client.List(ctx, nl); err != nil {
			fmt.Println(err)
		}
		for _, n := range nl.Items {
			p := v1.Pod{
				TypeMeta: v12.TypeMeta{
					APIVersion: "v1",
					Kind:       "Pod",
				},
				ObjectMeta: v12.ObjectMeta{
					GenerateName: fmt.Sprintf("%s-", n.Name),
					Namespace:    myds.Namespace,
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Image: myds.Spec.Image,
							Name:  "container",
						},
					},
					NodeName: n.Name,
				},
			}
			if err := r.Client.Create(ctx, &p); err != nil {
				fmt.Println(err)
			}
		}
	}
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MyDaemonsetReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appsv1beta1.MyDaemonset{}).
		Complete(r)
}
