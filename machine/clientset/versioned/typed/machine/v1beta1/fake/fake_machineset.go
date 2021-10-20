// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/openshift/api/machine/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMachineSets implements MachineSetInterface
type FakeMachineSets struct {
	Fake *FakeMachineV1beta1
	ns   string
}

var machinesetsResource = schema.GroupVersionResource{Group: "machine.openshift.io", Version: "v1beta1", Resource: "machinesets"}

var machinesetsKind = schema.GroupVersionKind{Group: "machine.openshift.io", Version: "v1beta1", Kind: "MachineSet"}

// Get takes name of the machineSet, and returns the corresponding machineSet object, and an error if there is any.
func (c *FakeMachineSets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.MachineSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(machinesetsResource, c.ns, name), &v1beta1.MachineSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MachineSet), err
}

// List takes label and field selectors, and returns the list of MachineSets that match those selectors.
func (c *FakeMachineSets) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.MachineSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(machinesetsResource, machinesetsKind, c.ns, opts), &v1beta1.MachineSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.MachineSetList{ListMeta: obj.(*v1beta1.MachineSetList).ListMeta}
	for _, item := range obj.(*v1beta1.MachineSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested machineSets.
func (c *FakeMachineSets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(machinesetsResource, c.ns, opts))

}

// Create takes the representation of a machineSet and creates it.  Returns the server's representation of the machineSet, and an error, if there is any.
func (c *FakeMachineSets) Create(ctx context.Context, machineSet *v1beta1.MachineSet, opts v1.CreateOptions) (result *v1beta1.MachineSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(machinesetsResource, c.ns, machineSet), &v1beta1.MachineSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MachineSet), err
}

// Update takes the representation of a machineSet and updates it. Returns the server's representation of the machineSet, and an error, if there is any.
func (c *FakeMachineSets) Update(ctx context.Context, machineSet *v1beta1.MachineSet, opts v1.UpdateOptions) (result *v1beta1.MachineSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(machinesetsResource, c.ns, machineSet), &v1beta1.MachineSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MachineSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMachineSets) UpdateStatus(ctx context.Context, machineSet *v1beta1.MachineSet, opts v1.UpdateOptions) (*v1beta1.MachineSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(machinesetsResource, "status", c.ns, machineSet), &v1beta1.MachineSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MachineSet), err
}

// Delete takes name of the machineSet and deletes it. Returns an error if one occurs.
func (c *FakeMachineSets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(machinesetsResource, c.ns, name), &v1beta1.MachineSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMachineSets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(machinesetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.MachineSetList{})
	return err
}

// Patch applies the patch and returns the patched machineSet.
func (c *FakeMachineSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.MachineSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(machinesetsResource, c.ns, name, pt, data, subresources...), &v1beta1.MachineSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MachineSet), err
}
