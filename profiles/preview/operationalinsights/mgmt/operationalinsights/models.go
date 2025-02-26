//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package operationalinsights

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/operationalinsights/mgmt/2020-10-01/operationalinsights"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BillingType = original.BillingType

const (
	BillingTypeCluster    BillingType = original.BillingTypeCluster
	BillingTypeWorkspaces BillingType = original.BillingTypeWorkspaces
)

type ClusterEntityStatus = original.ClusterEntityStatus

const (
	Canceled            ClusterEntityStatus = original.Canceled
	Creating            ClusterEntityStatus = original.Creating
	Deleting            ClusterEntityStatus = original.Deleting
	Failed              ClusterEntityStatus = original.Failed
	ProvisioningAccount ClusterEntityStatus = original.ProvisioningAccount
	Succeeded           ClusterEntityStatus = original.Succeeded
	Updating            ClusterEntityStatus = original.Updating
)

type ClusterSkuNameEnum = original.ClusterSkuNameEnum

const (
	CapacityReservation ClusterSkuNameEnum = original.CapacityReservation
)

type DataIngestionStatus = original.DataIngestionStatus

const (
	ApproachingQuota      DataIngestionStatus = original.ApproachingQuota
	ForceOff              DataIngestionStatus = original.ForceOff
	ForceOn               DataIngestionStatus = original.ForceOn
	OverQuota             DataIngestionStatus = original.OverQuota
	RespectQuota          DataIngestionStatus = original.RespectQuota
	SubscriptionSuspended DataIngestionStatus = original.SubscriptionSuspended
)

type IdentityType = original.IdentityType

const (
	None           IdentityType = original.None
	SystemAssigned IdentityType = original.SystemAssigned
	UserAssigned   IdentityType = original.UserAssigned
)

type PublicNetworkAccessType = original.PublicNetworkAccessType

const (
	Disabled PublicNetworkAccessType = original.Disabled
	Enabled  PublicNetworkAccessType = original.Enabled
)

type WorkspaceEntityStatus = original.WorkspaceEntityStatus

const (
	WorkspaceEntityStatusCanceled            WorkspaceEntityStatus = original.WorkspaceEntityStatusCanceled
	WorkspaceEntityStatusCreating            WorkspaceEntityStatus = original.WorkspaceEntityStatusCreating
	WorkspaceEntityStatusDeleting            WorkspaceEntityStatus = original.WorkspaceEntityStatusDeleting
	WorkspaceEntityStatusFailed              WorkspaceEntityStatus = original.WorkspaceEntityStatusFailed
	WorkspaceEntityStatusProvisioningAccount WorkspaceEntityStatus = original.WorkspaceEntityStatusProvisioningAccount
	WorkspaceEntityStatusSucceeded           WorkspaceEntityStatus = original.WorkspaceEntityStatusSucceeded
	WorkspaceEntityStatusUpdating            WorkspaceEntityStatus = original.WorkspaceEntityStatusUpdating
)

type WorkspaceSkuNameEnum = original.WorkspaceSkuNameEnum

const (
	WorkspaceSkuNameEnumCapacityReservation WorkspaceSkuNameEnum = original.WorkspaceSkuNameEnumCapacityReservation
	WorkspaceSkuNameEnumFree                WorkspaceSkuNameEnum = original.WorkspaceSkuNameEnumFree
	WorkspaceSkuNameEnumLACluster           WorkspaceSkuNameEnum = original.WorkspaceSkuNameEnumLACluster
	WorkspaceSkuNameEnumPerGB2018           WorkspaceSkuNameEnum = original.WorkspaceSkuNameEnumPerGB2018
	WorkspaceSkuNameEnumPerNode             WorkspaceSkuNameEnum = original.WorkspaceSkuNameEnumPerNode
	WorkspaceSkuNameEnumPremium             WorkspaceSkuNameEnum = original.WorkspaceSkuNameEnumPremium
	WorkspaceSkuNameEnumStandalone          WorkspaceSkuNameEnum = original.WorkspaceSkuNameEnumStandalone
	WorkspaceSkuNameEnumStandard            WorkspaceSkuNameEnum = original.WorkspaceSkuNameEnumStandard
)

type AssociatedWorkspace = original.AssociatedWorkspace
type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type CapacityReservationProperties = original.CapacityReservationProperties
type Cluster = original.Cluster
type ClusterListResult = original.ClusterListResult
type ClusterListResultIterator = original.ClusterListResultIterator
type ClusterListResultPage = original.ClusterListResultPage
type ClusterPatch = original.ClusterPatch
type ClusterPatchProperties = original.ClusterPatchProperties
type ClusterProperties = original.ClusterProperties
type ClusterSku = original.ClusterSku
type ClustersClient = original.ClustersClient
type ClustersCreateOrUpdateFuture = original.ClustersCreateOrUpdateFuture
type ClustersDeleteFuture = original.ClustersDeleteFuture
type DeletedWorkspacesClient = original.DeletedWorkspacesClient
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorDetail = original.ErrorDetail
type ErrorResponse = original.ErrorResponse
type Identity = original.Identity
type KeyVaultProperties = original.KeyVaultProperties
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type PrivateLinkScopedResource = original.PrivateLinkScopedResource
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type TrackedResource = original.TrackedResource
type UserIdentityProperties = original.UserIdentityProperties
type Workspace = original.Workspace
type WorkspaceCapping = original.WorkspaceCapping
type WorkspaceListResult = original.WorkspaceListResult
type WorkspacePatch = original.WorkspacePatch
type WorkspaceProperties = original.WorkspaceProperties
type WorkspaceSku = original.WorkspaceSku
type WorkspacesClient = original.WorkspacesClient
type WorkspacesCreateOrUpdateFuture = original.WorkspacesCreateOrUpdateFuture
type WorkspacesDeleteFuture = original.WorkspacesDeleteFuture

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewClusterListResultIterator(page ClusterListResultPage) ClusterListResultIterator {
	return original.NewClusterListResultIterator(page)
}
func NewClusterListResultPage(cur ClusterListResult, getNextPage func(context.Context, ClusterListResult) (ClusterListResult, error)) ClusterListResultPage {
	return original.NewClusterListResultPage(cur, getNextPage)
}
func NewClustersClient(subscriptionID string) ClustersClient {
	return original.NewClustersClient(subscriptionID)
}
func NewClustersClientWithBaseURI(baseURI string, subscriptionID string) ClustersClient {
	return original.NewClustersClientWithBaseURI(baseURI, subscriptionID)
}
func NewDeletedWorkspacesClient(subscriptionID string) DeletedWorkspacesClient {
	return original.NewDeletedWorkspacesClient(subscriptionID)
}
func NewDeletedWorkspacesClientWithBaseURI(baseURI string, subscriptionID string) DeletedWorkspacesClient {
	return original.NewDeletedWorkspacesClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewWorkspacesClient(subscriptionID string) WorkspacesClient {
	return original.NewWorkspacesClient(subscriptionID)
}
func NewWorkspacesClientWithBaseURI(baseURI string, subscriptionID string) WorkspacesClient {
	return original.NewWorkspacesClientWithBaseURI(baseURI, subscriptionID)
}
func PossibleBillingTypeValues() []BillingType {
	return original.PossibleBillingTypeValues()
}
func PossibleClusterEntityStatusValues() []ClusterEntityStatus {
	return original.PossibleClusterEntityStatusValues()
}
func PossibleClusterSkuNameEnumValues() []ClusterSkuNameEnum {
	return original.PossibleClusterSkuNameEnumValues()
}
func PossibleDataIngestionStatusValues() []DataIngestionStatus {
	return original.PossibleDataIngestionStatusValues()
}
func PossibleIdentityTypeValues() []IdentityType {
	return original.PossibleIdentityTypeValues()
}
func PossiblePublicNetworkAccessTypeValues() []PublicNetworkAccessType {
	return original.PossiblePublicNetworkAccessTypeValues()
}
func PossibleWorkspaceEntityStatusValues() []WorkspaceEntityStatus {
	return original.PossibleWorkspaceEntityStatusValues()
}
func PossibleWorkspaceSkuNameEnumValues() []WorkspaceSkuNameEnum {
	return original.PossibleWorkspaceSkuNameEnumValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
