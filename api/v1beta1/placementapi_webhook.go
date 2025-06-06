/*
Copyright 2022.

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

// Generated by:
//
// operator-sdk create webhook --group placement --version v1beta1 --kind PlacementAPI --programmatic-validation --defaulting
//

package v1beta1

import (
	"fmt"

	"github.com/openstack-k8s-operators/lib-common/modules/common/service"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// PlacementAPIDefaults -
type PlacementAPIDefaults struct {
	ContainerImageURL string
	APITimeout        int
}

var placementAPIDefaults PlacementAPIDefaults

// log is for logging in this package.
var placementapilog = logf.Log.WithName("placementapi-resource")

// SetupPlacementAPIDefaults - initialize PlacementAPI spec defaults for use with either internal or external webhooks
func SetupPlacementAPIDefaults(defaults PlacementAPIDefaults) {
	placementAPIDefaults = defaults
	placementapilog.Info("PlacementAPI defaults initialized", "defaults", defaults)
}

// SetupWebhookWithManager sets up the webhook with the Manager
func (r *PlacementAPI) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-placement-openstack-org-v1beta1-placementapi,mutating=true,failurePolicy=fail,sideEffects=None,groups=placement.openstack.org,resources=placementapis,verbs=create;update,versions=v1beta1,name=mplacementapi.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &PlacementAPI{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *PlacementAPI) Default() {
	placementapilog.Info("default", "name", r.Name)

	r.Spec.Default()
}

// Default - set defaults for this PlacementAPI spec
func (spec *PlacementAPISpec) Default() {
	if spec.ContainerImage == "" {
		spec.ContainerImage = placementAPIDefaults.ContainerImageURL
	}
	if spec.APITimeout == 0 {
		spec.APITimeout = placementAPIDefaults.APITimeout
	}

}

// Default - set defaults for this PlacementAPI core spec (this version is used by the OpenStackControlplane webhook)
func (spec *PlacementAPISpecCore) Default() {
	// nothing here yet
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-placement-openstack-org-v1beta1-placementapi,mutating=false,failurePolicy=fail,sideEffects=None,groups=placement.openstack.org,resources=placementapis,verbs=create;update,versions=v1beta1,name=vplacementapi.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &PlacementAPI{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *PlacementAPI) ValidateCreate() (admission.Warnings, error) {
	placementapilog.Info("validate create", "name", r.Name)

	errors := r.Spec.ValidateCreate(field.NewPath("spec"), r.Namespace)
	if len(errors) != 0 {
		placementapilog.Info("validation failed", "name", r.Name)
		return nil, apierrors.NewInvalid(
			schema.GroupKind{Group: "placement.openstack.org", Kind: "PlacementAPI"},
			r.Name, errors)
	}
	return nil, nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *PlacementAPI) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	placementapilog.Info("validate update", "name", r.Name)
	oldPlacement, ok := old.(*PlacementAPI)
	if !ok || oldPlacement == nil {
		return nil, apierrors.NewInternalError(fmt.Errorf("unable to convert existing object"))
	}

	errors := r.Spec.ValidateUpdate(oldPlacement.Spec, field.NewPath("spec"), r.Namespace)
	if len(errors) != 0 {
		placementapilog.Info("validation failed", "name", r.Name)
		return nil, apierrors.NewInvalid(
			schema.GroupKind{Group: "placement.openstack.org", Kind: "PlacementAPI"},
			r.Name, errors)
	}
	return nil, nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *PlacementAPI) ValidateDelete() (admission.Warnings, error) {
	placementapilog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil, nil
}

func (r PlacementAPISpec) ValidateCreate(basePath *field.Path, namespace string) field.ErrorList {
	return r.PlacementAPISpecCore.ValidateCreate(basePath, namespace)
}

func (r PlacementAPISpec) ValidateUpdate(old PlacementAPISpec, basePath *field.Path, namespace string) field.ErrorList {
	return r.PlacementAPISpecCore.ValidateCreate(basePath, namespace)
}

func (r PlacementAPISpecCore) ValidateCreate(basePath *field.Path, namespace string) field.ErrorList {
	var allErrs field.ErrorList

	// validate the service override key is valid
	allErrs = append(allErrs, service.ValidateRoutedOverrides(basePath.Child("override").Child("service"), r.Override.Service)...)

	allErrs = append(allErrs, ValidateDefaultConfigOverwrite(basePath, r.DefaultConfigOverwrite)...)

	// When a TopologyRef CR is referenced, fail if a different Namespace is
	// referenced because is not supported
	allErrs = append(allErrs, r.ValidateTopology(basePath, namespace)...)

	return allErrs
}

func (r PlacementAPISpecCore) ValidateUpdate(old PlacementAPISpecCore, basePath *field.Path, namespace string) field.ErrorList {
	var allErrs field.ErrorList

	// validate the service override key is valid
	allErrs = append(allErrs, service.ValidateRoutedOverrides(basePath.Child("override").Child("service"), r.Override.Service)...)

	allErrs = append(allErrs, ValidateDefaultConfigOverwrite(basePath, r.DefaultConfigOverwrite)...)

	// When a TopologyRef CR is referenced, fail if a different Namespace is
	// referenced because is not supported
	allErrs = append(allErrs, r.ValidateTopology(basePath, namespace)...)

	return allErrs
}

func ValidateDefaultConfigOverwrite(
	basePath *field.Path,
	validateConfigOverwrite map[string]string,
) field.ErrorList {
	var errors field.ErrorList
	for requested := range validateConfigOverwrite {
		if requested != "policy.yaml" {
			errors = append(
				errors,
				field.Invalid(
					basePath.Child("defaultConfigOverwrite"),
					requested,
					"Only the following keys are valid: policy.yaml",
				),
			)
		}
	}
	return errors
}

// SetDefaultRouteAnnotations sets HAProxy timeout values of the route
func (spec *PlacementAPISpecCore) SetDefaultRouteAnnotations(annotations map[string]string) {
	const haProxyAnno = "haproxy.router.openshift.io/timeout"
	// Use a custom annotation to flag when the operator has set the default HAProxy timeout
	// With the annotation func determines when to overwrite existing HAProxy timeout with the APITimeout
	const placementAnno = "api.placement.openstack.org/timeout"
	valPlacementAPI, okPlacementAPI := annotations[placementAnno]
	valHAProxy, okHAProxy := annotations[haProxyAnno]
	// Human operator set the HAProxy timeout manually
	if !okPlacementAPI && okHAProxy {
		return
	}
	// Human operator modified the HAProxy timeout manually without removing the Placemen flag
	if okPlacementAPI && okHAProxy && valPlacementAPI != valHAProxy {
		delete(annotations, placementAnno)
		placementapilog.Info("Human operator modified the HAProxy timeout manually without removing the Placement flag. Deleting the Placement flag to ensure proper configuration.")
		return
	}
	timeout := fmt.Sprintf("%ds", spec.APITimeout)
	annotations[placementAnno] = timeout
	annotations[haProxyAnno] = timeout
}
