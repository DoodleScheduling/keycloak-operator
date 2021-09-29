package keycloakrealm

import (
	"fmt"

	kc "github.com/keycloak/keycloak-operator/pkg/apis/keycloak/v1alpha1"
	"github.com/keycloak/keycloak-operator/pkg/common"
	"github.com/keycloak/keycloak-operator/pkg/model"
)

type Reconciler interface {
	Reconcile(cr *kc.KeycloakRealm) error
}

type KeycloakRealmReconciler struct { // nolint
	Keycloak kc.Keycloak
}

func NewKeycloakRealmReconciler(keycloak kc.Keycloak) *KeycloakRealmReconciler {
	return &KeycloakRealmReconciler{
		Keycloak: keycloak,
	}
}

func (i *KeycloakRealmReconciler) Reconcile(state *common.RealmState, cr *kc.KeycloakRealm) common.DesiredClusterState {
	switch {
	case cr.DeletionTimestamp != nil:
		return i.ReconcileRealmDelete(state, cr)
	case state.Realm != nil && cr.Spec.ApplyUpdates == true:
		return i.ReconcileRealmUpdate(state, cr)
	default:
		return i.ReconcileRealmCreate(state, cr)
	}
}

func (i *KeycloakRealmReconciler) ReconcileRealmUpdate(state *common.RealmState, cr *kc.KeycloakRealm) common.DesiredClusterState {
	desired := i.ReconcileRealmCreate(state, cr)
	desired.AddActions(i.getDesiredRealmClientScopeState(state, cr))

	return desired
}

func (i *KeycloakRealmReconciler) ReconcileRealmCreate(state *common.RealmState, cr *kc.KeycloakRealm) common.DesiredClusterState {
	desired := common.DesiredClusterState{}

	desired.AddAction(i.getKeycloakDesiredState())
	desired.AddAction(i.getDesiredRealmState(state, cr))

	for _, user := range cr.Spec.Realm.Users {
		desired.AddAction(i.getDesiredUserState(state, cr, user))
	}

	desired.AddAction(i.getBrowserRedirectorDesiredState(state, cr))

	return desired
}

func (i *KeycloakRealmReconciler) ReconcileRealmDelete(state *common.RealmState, cr *kc.KeycloakRealm) common.DesiredClusterState {
	desired := common.DesiredClusterState{}
	desired.AddAction(i.getKeycloakDesiredState())
	desired.AddAction(i.getDesiredRealmState(state, cr))
	return desired
}

// Always make sure keycloak is able to respond
func (i *KeycloakRealmReconciler) getKeycloakDesiredState() common.ClusterAction {
	return &common.PingAction{
		Msg: "check if keycloak is available",
	}
}

// Configure the browser redirector if provided by the user
func (i *KeycloakRealmReconciler) getBrowserRedirectorDesiredState(state *common.RealmState, cr *kc.KeycloakRealm) common.ClusterAction {
	if len(cr.Spec.RealmOverrides) == 0 {
		return nil
	}

	// Never update the realm configuration, leave it up to the users
	if state.Realm != nil {
		return nil
	}

	return &common.ConfigureRealmAction{
		Ref: cr,
		Msg: "configure browser redirector",
	}
}

func (i *KeycloakRealmReconciler) getDesiredRealmState(state *common.RealmState, cr *kc.KeycloakRealm) common.ClusterAction {
	switch {
	case cr.DeletionTimestamp != nil:
		return &common.DeleteRealmAction{
			Ref: cr,
			Msg: fmt.Sprintf("removing realm %v/%v", cr.Namespace, cr.Spec.Realm.Realm),
		}
	case state.Realm != nil && cr.Spec.ApplyUpdates == true:
		return &common.UpdateRealmAction{
			Ref: cr,
			Msg: fmt.Sprintf("update realm %v/%v", cr.Namespace, cr.Spec.Realm.Realm),
		}
	case state.Realm == nil:
		return &common.CreateRealmAction{
			Ref: cr,
			Msg: fmt.Sprintf("create realm %v/%v", cr.Namespace, cr.Spec.Realm.Realm),
		}
	default:
		return nil
	}
}

func (i *KeycloakRealmReconciler) getDesiredUserState(state *common.RealmState, cr *kc.KeycloakRealm, user *kc.KeycloakAPIUser) common.ClusterAction {
	val, ok := state.RealmUserSecrets[user.UserName]
	if !ok || val == nil {
		return &common.GenericCreateAction{
			Ref: model.RealmCredentialSecret(cr, user, &i.Keycloak),
			Msg: fmt.Sprintf("create credential secret for user %v in realm %v/%v", user.UserName, cr.Namespace, cr.Spec.Realm.Realm),
		}
	}

	return nil
}

func (i *KeycloakRealmReconciler) getDesiredRealmClientScopeState(state *common.RealmState, cr *kc.KeycloakRealm) []common.ClusterAction {
	var actions []common.ClusterAction
	var w []string

OUTER_SPEC:
	for _, specV := range cr.Spec.Realm.ClientScopes {
		//store local copy to reference
		p := specV
		for _, stateV := range state.Realm.Spec.Realm.ClientScopes {
			if stateV.Name != specV.Name {
				continue
			}

			p.ID = stateV.ID
			actions = append(actions, &common.UpdateClientScopeAction{
				ClientScope: &p,
				Realm:       cr.Spec.Realm.Realm,
				Msg:         fmt.Sprintf("update client scope %v in realm %v/%v", specV.Name, cr.Namespace, cr.Spec.Realm.Realm),
			})

			w = append(w, specV.Name)

			// If we have an existing resource update and go to the next in the specs
			continue OUTER_SPEC
		}

		// resource does not exist yet, create it
		actions = append(actions, &common.CreateClientScopeAction{
			ClientScope: &p,
			Realm:       cr.Spec.Realm.Realm,
			Msg:         fmt.Sprintf("create client scope %v in realm %v/%v", specV.Name, cr.Namespace, cr.Spec.Realm.Realm),
		})

		w = append(w, specV.Name)
	}

OUTER_STATE:
	for _, stateV := range state.Realm.Spec.Realm.ClientScopes {
		//store local copy to reference
		p := stateV
		for _, name := range w {
			if stateV.Name == name {
				continue OUTER_STATE
			}
		}

		// No create or update action found, removing resource
		actions = append(actions, &common.DeleteClientScopeAction{
			ClientScope: &p,
			Realm:       cr.Spec.Realm.Realm,
			Msg:         fmt.Sprintf("delete client scope %v in realm %v/%v", stateV.Name, cr.Namespace, cr.Spec.Realm.Realm),
		})
	}

	return actions
}
