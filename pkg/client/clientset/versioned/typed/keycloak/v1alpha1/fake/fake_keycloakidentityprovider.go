// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/keycloak/keycloak-operator/pkg/apis/keycloak/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKeycloakIdentityProviders implements KeycloakIdentityProviderInterface
type FakeKeycloakIdentityProviders struct {
	Fake *FakeKeycloakV1alpha1
	ns   string
}

var keycloakidentityprovidersResource = schema.GroupVersionResource{Group: "keycloak.org", Version: "v1alpha1", Resource: "keycloakidentityproviders"}

var keycloakidentityprovidersKind = schema.GroupVersionKind{Group: "keycloak.org", Version: "v1alpha1", Kind: "KeycloakIdentityProvider"}

// Get takes name of the keycloakIdentityProvider, and returns the corresponding keycloakIdentityProvider object, and an error if there is any.
func (c *FakeKeycloakIdentityProviders) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.KeycloakIdentityProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(keycloakidentityprovidersResource, c.ns, name), &v1alpha1.KeycloakIdentityProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KeycloakIdentityProvider), err
}

// List takes label and field selectors, and returns the list of KeycloakIdentityProviders that match those selectors.
func (c *FakeKeycloakIdentityProviders) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.KeycloakIdentityProviderList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(keycloakidentityprovidersResource, keycloakidentityprovidersKind, c.ns, opts), &v1alpha1.KeycloakIdentityProviderList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KeycloakIdentityProviderList{ListMeta: obj.(*v1alpha1.KeycloakIdentityProviderList).ListMeta}
	for _, item := range obj.(*v1alpha1.KeycloakIdentityProviderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested keycloakIdentityProviders.
func (c *FakeKeycloakIdentityProviders) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(keycloakidentityprovidersResource, c.ns, opts))

}

// Create takes the representation of a keycloakIdentityProvider and creates it.  Returns the server's representation of the keycloakIdentityProvider, and an error, if there is any.
func (c *FakeKeycloakIdentityProviders) Create(ctx context.Context, keycloakIdentityProvider *v1alpha1.KeycloakIdentityProvider, opts v1.CreateOptions) (result *v1alpha1.KeycloakIdentityProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(keycloakidentityprovidersResource, c.ns, keycloakIdentityProvider), &v1alpha1.KeycloakIdentityProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KeycloakIdentityProvider), err
}

// Update takes the representation of a keycloakIdentityProvider and updates it. Returns the server's representation of the keycloakIdentityProvider, and an error, if there is any.
func (c *FakeKeycloakIdentityProviders) Update(ctx context.Context, keycloakIdentityProvider *v1alpha1.KeycloakIdentityProvider, opts v1.UpdateOptions) (result *v1alpha1.KeycloakIdentityProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(keycloakidentityprovidersResource, c.ns, keycloakIdentityProvider), &v1alpha1.KeycloakIdentityProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KeycloakIdentityProvider), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKeycloakIdentityProviders) UpdateStatus(ctx context.Context, keycloakIdentityProvider *v1alpha1.KeycloakIdentityProvider, opts v1.UpdateOptions) (*v1alpha1.KeycloakIdentityProvider, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(keycloakidentityprovidersResource, "status", c.ns, keycloakIdentityProvider), &v1alpha1.KeycloakIdentityProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KeycloakIdentityProvider), err
}

// Delete takes name of the keycloakIdentityProvider and deletes it. Returns an error if one occurs.
func (c *FakeKeycloakIdentityProviders) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(keycloakidentityprovidersResource, c.ns, name), &v1alpha1.KeycloakIdentityProvider{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKeycloakIdentityProviders) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(keycloakidentityprovidersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.KeycloakIdentityProviderList{})
	return err
}

// Patch applies the patch and returns the patched keycloakIdentityProvider.
func (c *FakeKeycloakIdentityProviders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KeycloakIdentityProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(keycloakidentityprovidersResource, c.ns, name, pt, data, subresources...), &v1alpha1.KeycloakIdentityProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KeycloakIdentityProvider), err
}
