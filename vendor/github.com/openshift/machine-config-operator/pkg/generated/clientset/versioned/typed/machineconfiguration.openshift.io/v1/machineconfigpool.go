// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	scheme "github.com/openshift/machine-config-operator/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MachineConfigPoolsGetter has a method to return a MachineConfigPoolInterface.
// A group's client should implement this interface.
type MachineConfigPoolsGetter interface {
	MachineConfigPools() MachineConfigPoolInterface
}

// MachineConfigPoolInterface has methods to work with MachineConfigPool resources.
type MachineConfigPoolInterface interface {
	Create(*v1.MachineConfigPool) (*v1.MachineConfigPool, error)
	Update(*v1.MachineConfigPool) (*v1.MachineConfigPool, error)
	UpdateStatus(*v1.MachineConfigPool) (*v1.MachineConfigPool, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.MachineConfigPool, error)
	List(opts metav1.ListOptions) (*v1.MachineConfigPoolList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.MachineConfigPool, err error)
	MachineConfigPoolExpansion
}

// machineConfigPools implements MachineConfigPoolInterface
type machineConfigPools struct {
	client rest.Interface
}

// newMachineConfigPools returns a MachineConfigPools
func newMachineConfigPools(c *MachineconfigurationV1Client) *machineConfigPools {
	return &machineConfigPools{
		client: c.RESTClient(),
	}
}

// Get takes name of the machineConfigPool, and returns the corresponding machineConfigPool object, and an error if there is any.
func (c *machineConfigPools) Get(name string, options metav1.GetOptions) (result *v1.MachineConfigPool, err error) {
	result = &v1.MachineConfigPool{}
	err = c.client.Get().
		Resource("machineconfigpools").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MachineConfigPools that match those selectors.
func (c *machineConfigPools) List(opts metav1.ListOptions) (result *v1.MachineConfigPoolList, err error) {
	result = &v1.MachineConfigPoolList{}
	err = c.client.Get().
		Resource("machineconfigpools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested machineConfigPools.
func (c *machineConfigPools) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("machineconfigpools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a machineConfigPool and creates it.  Returns the server's representation of the machineConfigPool, and an error, if there is any.
func (c *machineConfigPools) Create(machineConfigPool *v1.MachineConfigPool) (result *v1.MachineConfigPool, err error) {
	result = &v1.MachineConfigPool{}
	err = c.client.Post().
		Resource("machineconfigpools").
		Body(machineConfigPool).
		Do().
		Into(result)
	return
}

// Update takes the representation of a machineConfigPool and updates it. Returns the server's representation of the machineConfigPool, and an error, if there is any.
func (c *machineConfigPools) Update(machineConfigPool *v1.MachineConfigPool) (result *v1.MachineConfigPool, err error) {
	result = &v1.MachineConfigPool{}
	err = c.client.Put().
		Resource("machineconfigpools").
		Name(machineConfigPool.Name).
		Body(machineConfigPool).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *machineConfigPools) UpdateStatus(machineConfigPool *v1.MachineConfigPool) (result *v1.MachineConfigPool, err error) {
	result = &v1.MachineConfigPool{}
	err = c.client.Put().
		Resource("machineconfigpools").
		Name(machineConfigPool.Name).
		SubResource("status").
		Body(machineConfigPool).
		Do().
		Into(result)
	return
}

// Delete takes name of the machineConfigPool and deletes it. Returns an error if one occurs.
func (c *machineConfigPools) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("machineconfigpools").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *machineConfigPools) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	return c.client.Delete().
		Resource("machineconfigpools").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched machineConfigPool.
func (c *machineConfigPools) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.MachineConfigPool, err error) {
	result = &v1.MachineConfigPool{}
	err = c.client.Patch(pt).
		Resource("machineconfigpools").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}