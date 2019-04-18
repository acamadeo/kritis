/*
Copyright 2018 Google LLC

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/grafeas/kritis/pkg/kritis/apis/kritis/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGenericAttestationPolicies implements GenericAttestationPolicyInterface
type FakeGenericAttestationPolicies struct {
	Fake *FakeKritisV1beta1
	ns   string
}

var genericattestationpoliciesResource = schema.GroupVersionResource{Group: "kritis", Version: "v1beta1", Resource: "genericattestationpolicies"}

var genericattestationpoliciesKind = schema.GroupVersionKind{Group: "kritis", Version: "v1beta1", Kind: "GenericAttestationPolicy"}

// Get takes name of the genericAttestationPolicy, and returns the corresponding genericAttestationPolicy object, and an error if there is any.
func (c *FakeGenericAttestationPolicies) Get(name string, options v1.GetOptions) (result *v1beta1.GenericAttestationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(genericattestationpoliciesResource, c.ns, name), &v1beta1.GenericAttestationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.GenericAttestationPolicy), err
}

// List takes label and field selectors, and returns the list of GenericAttestationPolicies that match those selectors.
func (c *FakeGenericAttestationPolicies) List(opts v1.ListOptions) (result *v1beta1.GenericAttestationPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(genericattestationpoliciesResource, genericattestationpoliciesKind, c.ns, opts), &v1beta1.GenericAttestationPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.GenericAttestationPolicyList{}
	for _, item := range obj.(*v1beta1.GenericAttestationPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested genericAttestationPolicies.
func (c *FakeGenericAttestationPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(genericattestationpoliciesResource, c.ns, opts))

}

// Create takes the representation of a genericAttestationPolicy and creates it.  Returns the server's representation of the genericAttestationPolicy, and an error, if there is any.
func (c *FakeGenericAttestationPolicies) Create(genericAttestationPolicy *v1beta1.GenericAttestationPolicy) (result *v1beta1.GenericAttestationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(genericattestationpoliciesResource, c.ns, genericAttestationPolicy), &v1beta1.GenericAttestationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.GenericAttestationPolicy), err
}

// Update takes the representation of a genericAttestationPolicy and updates it. Returns the server's representation of the genericAttestationPolicy, and an error, if there is any.
func (c *FakeGenericAttestationPolicies) Update(genericAttestationPolicy *v1beta1.GenericAttestationPolicy) (result *v1beta1.GenericAttestationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(genericattestationpoliciesResource, c.ns, genericAttestationPolicy), &v1beta1.GenericAttestationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.GenericAttestationPolicy), err
}

// Delete takes name of the genericAttestationPolicy and deletes it. Returns an error if one occurs.
func (c *FakeGenericAttestationPolicies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(genericattestationpoliciesResource, c.ns, name), &v1beta1.GenericAttestationPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGenericAttestationPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(genericattestationpoliciesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.GenericAttestationPolicyList{})
	return err
}

// Patch applies the patch and returns the patched genericAttestationPolicy.
func (c *FakeGenericAttestationPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.GenericAttestationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(genericattestationpoliciesResource, c.ns, name, data, subresources...), &v1beta1.GenericAttestationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.GenericAttestationPolicy), err
}
