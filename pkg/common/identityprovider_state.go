package common

import (
	"context"

	kc "github.com/keycloak/keycloak-operator/pkg/apis/keycloak/v1alpha1"
	v1 "k8s.io/api/core/v1"
	apiErrors "k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type IdentityProviderState struct {
	IdentityProvider *kc.KeycloakAPIIdentityProvider
	Secret           *v1.Secret
	Context          context.Context
	Realm            *kc.KeycloakRealm
}

func NewIdentityProviderState(context context.Context, realm *kc.KeycloakRealm) *IdentityProviderState {
	return &IdentityProviderState{
		Context: context,
		Realm:   realm,
	}
}

func (i *IdentityProviderState) Read(context context.Context, ipr *kc.KeycloakIdentityProvider, realmClient KeycloakInterface, controllerClient client.Client) error {
	if ipr.Spec.IdentityProvider.InternalID == "" {
		return nil
	}

	IdentityProvider, err := realmClient.GetIdentityProvider(ipr.Spec.IdentityProvider.InternalID, i.Realm.Spec.Realm.Realm)

	if err != nil {
		return err
	}

	i.IdentityProvider = IdentityProvider

	if ipr.Spec.Secret != nil {
		err = i.readIdentityProviderSecret(context, ipr, controllerClient)
		if err != nil {
			return err
		}
	}

	return nil
}

func (i *IdentityProviderState) readIdentityProviderSecret(context context.Context, ipr *kc.KeycloakIdentityProvider, controllerClient client.Client) error {
	secret := &v1.Secret{}
	key := client.ObjectKey{
		Name:      ipr.Spec.Secret.Name,
		Namespace: ipr.Namespace,
	}

	err := controllerClient.Get(context, key, secret)
	if err != nil {
		if !apiErrors.IsNotFound(err) {
			return err
		}
	} else {
		i.Secret = secret.DeepCopy()
	}

	return nil
}
