/*
Copyright 2018 The Kubernetes Authors.

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
package fake

import (
	v1alpha1 "github.com/marun/fnord/pkg/apis/federation/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFederatedNamespaces implements FederatedNamespaceInterface
type FakeFederatedNamespaces struct {
	Fake *FakeFederationV1alpha1
}

var federatednamespacesResource = schema.GroupVersionResource{Group: "federation.k8s.io", Version: "v1alpha1", Resource: "federatednamespaces"}

var federatednamespacesKind = schema.GroupVersionKind{Group: "federation.k8s.io", Version: "v1alpha1", Kind: "FederatedNamespace"}

// Get takes name of the federatedNamespace, and returns the corresponding federatedNamespace object, and an error if there is any.
func (c *FakeFederatedNamespaces) Get(name string, options v1.GetOptions) (result *v1alpha1.FederatedNamespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(federatednamespacesResource, name), &v1alpha1.FederatedNamespace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FederatedNamespace), err
}

// List takes label and field selectors, and returns the list of FederatedNamespaces that match those selectors.
func (c *FakeFederatedNamespaces) List(opts v1.ListOptions) (result *v1alpha1.FederatedNamespaceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(federatednamespacesResource, federatednamespacesKind, opts), &v1alpha1.FederatedNamespaceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FederatedNamespaceList{}
	for _, item := range obj.(*v1alpha1.FederatedNamespaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested federatedNamespaces.
func (c *FakeFederatedNamespaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(federatednamespacesResource, opts))
}

// Create takes the representation of a federatedNamespace and creates it.  Returns the server's representation of the federatedNamespace, and an error, if there is any.
func (c *FakeFederatedNamespaces) Create(federatedNamespace *v1alpha1.FederatedNamespace) (result *v1alpha1.FederatedNamespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(federatednamespacesResource, federatedNamespace), &v1alpha1.FederatedNamespace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FederatedNamespace), err
}

// Update takes the representation of a federatedNamespace and updates it. Returns the server's representation of the federatedNamespace, and an error, if there is any.
func (c *FakeFederatedNamespaces) Update(federatedNamespace *v1alpha1.FederatedNamespace) (result *v1alpha1.FederatedNamespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(federatednamespacesResource, federatedNamespace), &v1alpha1.FederatedNamespace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FederatedNamespace), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFederatedNamespaces) UpdateStatus(federatedNamespace *v1alpha1.FederatedNamespace) (*v1alpha1.FederatedNamespace, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(federatednamespacesResource, "status", federatedNamespace), &v1alpha1.FederatedNamespace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FederatedNamespace), err
}

// Delete takes name of the federatedNamespace and deletes it. Returns an error if one occurs.
func (c *FakeFederatedNamespaces) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(federatednamespacesResource, name), &v1alpha1.FederatedNamespace{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFederatedNamespaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(federatednamespacesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.FederatedNamespaceList{})
	return err
}

// Patch applies the patch and returns the patched federatedNamespace.
func (c *FakeFederatedNamespaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FederatedNamespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(federatednamespacesResource, name, data, subresources...), &v1alpha1.FederatedNamespace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FederatedNamespace), err
}
