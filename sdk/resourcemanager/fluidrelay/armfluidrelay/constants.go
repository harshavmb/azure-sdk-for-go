//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armfluidrelay

const (
	moduleName    = "armfluidrelay"
	moduleVersion = "v1.0.0"
)

// CmkIdentityType - Values can be SystemAssigned or UserAssigned
type CmkIdentityType string

const (
	CmkIdentityTypeSystemAssigned CmkIdentityType = "SystemAssigned"
	CmkIdentityTypeUserAssigned   CmkIdentityType = "UserAssigned"
)

// PossibleCmkIdentityTypeValues returns the possible values for the CmkIdentityType const type.
func PossibleCmkIdentityTypeValues() []CmkIdentityType {
	return []CmkIdentityType{
		CmkIdentityTypeSystemAssigned,
		CmkIdentityTypeUserAssigned,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// KeyName - The key to regenerate.
type KeyName string

const (
	KeyNameKey1 KeyName = "key1"
	KeyNameKey2 KeyName = "key2"
)

// PossibleKeyNameValues returns the possible values for the KeyName const type.
func PossibleKeyNameValues() []KeyName {
	return []KeyName{
		KeyNameKey1,
		KeyNameKey2,
	}
}

// ProvisioningState - Provision states for FluidRelay RP
type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCanceled,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
	}
}

// ResourceIdentityType - The identity type.
type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned             ResourceIdentityType = "SystemAssigned"
	ResourceIdentityTypeUserAssigned               ResourceIdentityType = "UserAssigned"
	ResourceIdentityTypeSystemAssignedUserAssigned ResourceIdentityType = "SystemAssigned, UserAssigned"
	ResourceIdentityTypeNone                       ResourceIdentityType = "None"
)

// PossibleResourceIdentityTypeValues returns the possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{
		ResourceIdentityTypeSystemAssigned,
		ResourceIdentityTypeUserAssigned,
		ResourceIdentityTypeSystemAssignedUserAssigned,
		ResourceIdentityTypeNone,
	}
}

// StorageSKU - Sku of the storage associated with the resource
type StorageSKU string

const (
	StorageSKUBasic    StorageSKU = "basic"
	StorageSKUStandard StorageSKU = "standard"
)

// PossibleStorageSKUValues returns the possible values for the StorageSKU const type.
func PossibleStorageSKUValues() []StorageSKU {
	return []StorageSKU{
		StorageSKUBasic,
		StorageSKUStandard,
	}
}
