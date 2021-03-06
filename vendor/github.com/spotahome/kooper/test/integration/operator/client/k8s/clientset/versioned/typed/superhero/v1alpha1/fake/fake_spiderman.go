/*
Copyright The Kubernetes Authors.

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
	v1alpha1 "github.com/spotahome/kooper/test/integration/operator/apis/superhero/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSpidermans implements SpidermanInterface
type FakeSpidermans struct {
	Fake *FakeSuperheroV1alpha1
	ns   string
}

var spidermansResource = schema.GroupVersionResource{Group: "superhero.comic.kooper", Version: "v1alpha1", Resource: "spidermans"}

var spidermansKind = schema.GroupVersionKind{Group: "superhero.comic.kooper", Version: "v1alpha1", Kind: "Spiderman"}

// Get takes name of the spiderman, and returns the corresponding spiderman object, and an error if there is any.
func (c *FakeSpidermans) Get(name string, options v1.GetOptions) (result *v1alpha1.Spiderman, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(spidermansResource, c.ns, name), &v1alpha1.Spiderman{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Spiderman), err
}

// List takes label and field selectors, and returns the list of Spidermans that match those selectors.
func (c *FakeSpidermans) List(opts v1.ListOptions) (result *v1alpha1.SpidermanList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(spidermansResource, spidermansKind, c.ns, opts), &v1alpha1.SpidermanList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SpidermanList{ListMeta: obj.(*v1alpha1.SpidermanList).ListMeta}
	for _, item := range obj.(*v1alpha1.SpidermanList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested spidermans.
func (c *FakeSpidermans) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(spidermansResource, c.ns, opts))

}

// Create takes the representation of a spiderman and creates it.  Returns the server's representation of the spiderman, and an error, if there is any.
func (c *FakeSpidermans) Create(spiderman *v1alpha1.Spiderman) (result *v1alpha1.Spiderman, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(spidermansResource, c.ns, spiderman), &v1alpha1.Spiderman{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Spiderman), err
}

// Update takes the representation of a spiderman and updates it. Returns the server's representation of the spiderman, and an error, if there is any.
func (c *FakeSpidermans) Update(spiderman *v1alpha1.Spiderman) (result *v1alpha1.Spiderman, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(spidermansResource, c.ns, spiderman), &v1alpha1.Spiderman{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Spiderman), err
}

// Delete takes name of the spiderman and deletes it. Returns an error if one occurs.
func (c *FakeSpidermans) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(spidermansResource, c.ns, name), &v1alpha1.Spiderman{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSpidermans) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(spidermansResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SpidermanList{})
	return err
}

// Patch applies the patch and returns the patched spiderman.
func (c *FakeSpidermans) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Spiderman, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(spidermansResource, c.ns, name, pt, data, subresources...), &v1alpha1.Spiderman{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Spiderman), err
}
