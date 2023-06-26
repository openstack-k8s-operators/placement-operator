/*
Copyright 2023.

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

package functional_test

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

var _ = Describe("Placement webhook", func() {

	var placementAPI TestPlacementAPI

	BeforeEach(func() {
		// lib-common uses OPERATOR_TEMPLATES env var to locate the "templates"
		// directory of the operator. We need to set them othervise lib-common
		// will fail to generate the ConfigMap as it does not find common.sh
		err := os.Setenv("OPERATOR_TEMPLATES", "../../templates")
		Expect(err).NotTo(HaveOccurred())

		placementAPI = NewTestPlacementAPI(TestNamespace)
		placementAPI.Create()
	})

	AfterEach(func() {
		placementAPI.Delete()
	})

	When("databaseInstance is being updated", func() {
		It("should be blocked by the webhook and fail", func() {
			instance := placementAPI.Instance
			_, err := controllerutil.CreateOrPatch(
				th.Ctx, th.K8sClient, instance, func() error {
					instance.Spec.DatabaseInstance = "changed"
					return nil
				})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).Should(ContainSubstring("spec.databaseInstance: Forbidden: Value is immutable"))
		})
	})
})
