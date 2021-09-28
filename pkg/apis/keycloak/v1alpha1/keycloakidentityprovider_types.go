package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KeycloakIdentityProviderSpec defines the desired state of KeycloakIdentityProvider.
// +k8s:openapi-gen=true
type KeycloakIdentityProviderSpec struct {
	// Selector for looking up KeycloakRealm Custom Resources.
	// +kubebuilder:validation:Required
	RealmSelector *metav1.LabelSelector `json:"realmSelector"`
	// Keycloak Client REST object.
	// +kubebuilder:validation:Required
	IdentityProvider *KeycloakAPIIdentityProvider `json:"identityProvider"`
}

type KeycloakAPIIdentityProvider struct {
	// Identity Provider Alias.
	// +optional
	Alias string `json:"alias,omitempty"`
	// Identity Provider Display Name.
	// +optional
	DisplayName string `json:"displayName,omitempty"`
	// Identity Provider Internal ID.
	// +optional
	InternalID string `json:"internalId,omitempty"`
	// Identity Provider ID.
	// +optional
	ProviderID string `json:"providerId,omitempty"`
	// Identity Provider enabled flag.
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// Identity Provider Trust Email.
	// +optional
	TrustEmail bool `json:"trustEmail,omitempty"`
	// Identity Provider Store to Token.
	// +optional
	StoreToken bool `json:"storeToken,omitempty"`
	// Adds Read Token role when creating this Identity Provider.
	// +optional
	AddReadTokenRoleOnCreate bool `json:"addReadTokenRoleOnCreate,omitempty"`
	// Identity Provider First Broker Login Flow Alias.
	// +optional
	FirstBrokerLoginFlowAlias string `json:"firstBrokerLoginFlowAlias,omitempty"`
	// Identity Provider Post Broker Login Flow Alias.
	// +optional
	PostBrokerLoginFlowAlias string `json:"postBrokerLoginFlowAlias,omitempty"`
	// Identity Provider Link Only setting.
	// +optional
	LinkOnly bool `json:"linkOnly,omitempty"`
	// Identity Provider config.
	// +optional
	Config map[string]string `json:"config,omitempty"`
}

// KeycloakIdentityProviderStatus defines the observed state of KeycloakIdentityProvider
// +k8s:openapi-gen=true
type KeycloakIdentityProviderStatus struct {
	// Current phase of the operator.
	Phase StatusPhase `json:"phase"`
	// Human-readable message indicating details about current operator phase or error.
	Message string `json:"message"`
	// True if all resources are in a ready state and all work is done.
	Ready bool `json:"ready"`
	// A map of all the secondary resources types and names created for this CR. e.g "Deployment": [ "DeploymentName1", "DeploymentName2" ]
	SecondaryResources map[string][]string `json:"secondaryResources,omitempty"`
}

// KeycloakIdentityProvider is the Schema for the KeycloakIdentityProviders API.
// +genclient
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type KeycloakIdentityProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KeycloakIdentityProviderSpec   `json:"spec,omitempty"`
	Status KeycloakIdentityProviderStatus `json:"status,omitempty"`
}

// KeycloakIdentityProviderList contains a list of KeycloakIdentityProvider.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type KeycloakIdentityProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeycloakIdentityProvider `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KeycloakIdentityProvider{}, &KeycloakIdentityProviderList{})
}

func (i *KeycloakIdentityProvider) UpdateStatusSecondaryResources(kind string, resourceName string) {
	i.Status.SecondaryResources = UpdateStatusSecondaryResources(i.Status.SecondaryResources, kind, resourceName)
}
