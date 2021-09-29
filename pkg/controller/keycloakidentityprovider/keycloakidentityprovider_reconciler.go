package keycloakidentityprovider

import (
	"fmt"

	kc "github.com/keycloak/keycloak-operator/pkg/apis/keycloak/v1alpha1"
	"github.com/keycloak/keycloak-operator/pkg/common"
)

const (
	umaRoleName = "uma_protection"
)

type Reconciler interface {
	Reconcile(ipr *kc.KeycloakIdentityProvider) error
}

type KeycloakIdentityProviderReconciler struct { // nolint
	Keycloak kc.Keycloak
}

func NewKeycloakIdentityProviderReconciler(keycloak kc.Keycloak) *KeycloakIdentityProviderReconciler {
	return &KeycloakIdentityProviderReconciler{
		Keycloak: keycloak,
	}
}

func (i *KeycloakIdentityProviderReconciler) Reconcile(state *common.IdentityProviderState, ipr *kc.KeycloakIdentityProvider) common.DesiredClusterState {
	desired := common.DesiredClusterState{}

	desired.AddAction(i.pingKeycloak())
	if ipr.DeletionTimestamp != nil {
		desired.AddAction(i.getDeletedIdentityProviderState(state, ipr))
		return desired
	}

	if state.IdentityProvider == nil {
		desired.AddAction(i.getipreatedIdentityProviderState(state, ipr))
	} else {
		desired.AddAction(i.getUpdatedIdentityProviderState(state, ipr))
	}

	return desired
}

func (i *KeycloakIdentityProviderReconciler) pingKeycloak() common.ClusterAction {
	return common.PingAction{
		Msg: "check if keycloak is available",
	}
}

func (i *KeycloakIdentityProviderReconciler) getDeletedIdentityProviderState(state *common.IdentityProviderState, ipr *kc.KeycloakIdentityProvider) common.ClusterAction {
	return common.DeleteIdentityProviderAction{
		Ref:   ipr,
		Realm: state.Realm.Spec.Realm.Realm,
		Msg:   fmt.Sprintf("removing IdentityProvider %v/%v", ipr.Namespace, ipr.Spec.IdentityProvider.InternalID),
	}
}

func (i *KeycloakIdentityProviderReconciler) getipreatedIdentityProviderState(state *common.IdentityProviderState, ipr *kc.KeycloakIdentityProvider) common.ClusterAction {
	if state.Secret != nil {
		for k, v := range state.Secret.Data {
			ipr.Spec.IdentityProvider.Config[k] = string(v)
		}
	}

	return common.CreateIdentityProviderAction{
		Ref:   ipr,
		Realm: state.Realm.Spec.Realm.Realm,
		Msg:   fmt.Sprintf("ipreate IdentityProvider %v/%v", ipr.Namespace, ipr.Spec.IdentityProvider.InternalID),
	}
}

func (i *KeycloakIdentityProviderReconciler) getUpdatedIdentityProviderState(state *common.IdentityProviderState, ipr *kc.KeycloakIdentityProvider) common.ClusterAction {
	if state.Secret != nil {
		for k, v := range state.Secret.Data {
			ipr.Spec.IdentityProvider.Config[k] = string(v)
		}
	}

	return common.UpdateIdentityProviderAction{
		Ref:   ipr,
		Realm: state.Realm.Spec.Realm.Realm,
		Msg:   fmt.Sprintf("update IdentityProvider %v/%v", ipr.Namespace, ipr.Spec.IdentityProvider.InternalID),
	}
}
