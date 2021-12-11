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

package v1beta1

import (
	"fmt"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var mydaemonsetlog = logf.Log.WithName("mydaemonset-resource")

func (r *MyDaemonset) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-apps-cncamp-io-v1beta1-mydaemonset,mutating=true,failurePolicy=fail,sideEffects=None,groups=apps.cncamp.io,resources=mydaemonsets,verbs=create;update,versions=v1beta1,name=mmydaemonset.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &MyDaemonset{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *MyDaemonset) Default() {
	mydaemonsetlog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-apps-cncamp-io-v1beta1-mydaemonset,mutating=false,failurePolicy=fail,sideEffects=None,groups=apps.cncamp.io,resources=mydaemonsets,verbs=create;update,versions=v1beta1,name=vmydaemonset.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &MyDaemonset{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *MyDaemonset) ValidateCreate() error {
	mydaemonsetlog.Info("validate create", "name", r.Name)
	if r.Spec.Image == "" {
		return fmt.Errorf("image is required")
	}
	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *MyDaemonset) ValidateUpdate(old runtime.Object) error {
	mydaemonsetlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *MyDaemonset) ValidateDelete() error {
	mydaemonsetlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
