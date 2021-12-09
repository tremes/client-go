// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	templatev1 "github.com/openshift/api/template/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTemplateInstances implements TemplateInstanceInterface
type FakeTemplateInstances struct {
	Fake *FakeTemplateV1
	ns   string
}

var templateinstancesResource = schema.GroupVersionResource{Group: "template.openshift.io", Version: "v1", Resource: "templateinstances"}

var templateinstancesKind = schema.GroupVersionKind{Group: "template.openshift.io", Version: "v1", Kind: "TemplateInstance"}

// Get takes name of the templateInstance, and returns the corresponding templateInstance object, and an error if there is any.
func (c *FakeTemplateInstances) Get(ctx context.Context, name string, options v1.GetOptions) (result *templatev1.TemplateInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(templateinstancesResource, c.ns, name), &templatev1.TemplateInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*templatev1.TemplateInstance), err
}

// List takes label and field selectors, and returns the list of TemplateInstances that match those selectors.
func (c *FakeTemplateInstances) List(ctx context.Context, opts v1.ListOptions) (result *templatev1.TemplateInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(templateinstancesResource, templateinstancesKind, c.ns, opts), &templatev1.TemplateInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &templatev1.TemplateInstanceList{ListMeta: obj.(*templatev1.TemplateInstanceList).ListMeta}
	for _, item := range obj.(*templatev1.TemplateInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested templateInstances.
func (c *FakeTemplateInstances) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(templateinstancesResource, c.ns, opts))

}

// Create takes the representation of a templateInstance and creates it.  Returns the server's representation of the templateInstance, and an error, if there is any.
func (c *FakeTemplateInstances) Create(ctx context.Context, templateInstance *templatev1.TemplateInstance, opts v1.CreateOptions) (result *templatev1.TemplateInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(templateinstancesResource, c.ns, templateInstance), &templatev1.TemplateInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*templatev1.TemplateInstance), err
}

// Update takes the representation of a templateInstance and updates it. Returns the server's representation of the templateInstance, and an error, if there is any.
func (c *FakeTemplateInstances) Update(ctx context.Context, templateInstance *templatev1.TemplateInstance, opts v1.UpdateOptions) (result *templatev1.TemplateInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(templateinstancesResource, c.ns, templateInstance), &templatev1.TemplateInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*templatev1.TemplateInstance), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTemplateInstances) UpdateStatus(ctx context.Context, templateInstance *templatev1.TemplateInstance, opts v1.UpdateOptions) (*templatev1.TemplateInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(templateinstancesResource, "status", c.ns, templateInstance), &templatev1.TemplateInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*templatev1.TemplateInstance), err
}

// Delete takes name of the templateInstance and deletes it. Returns an error if one occurs.
func (c *FakeTemplateInstances) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(templateinstancesResource, c.ns, name, opts), &templatev1.TemplateInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTemplateInstances) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(templateinstancesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &templatev1.TemplateInstanceList{})
	return err
}

// Patch applies the patch and returns the patched templateInstance.
func (c *FakeTemplateInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *templatev1.TemplateInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(templateinstancesResource, c.ns, name, pt, data, subresources...), &templatev1.TemplateInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*templatev1.TemplateInstance), err
}
