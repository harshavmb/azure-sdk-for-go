//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.
// DO NOT EDIT.

package armconnectedvmware

// ClustersClientCreateResponse contains the response from method ClustersClient.Create.
type ClustersClientCreateResponse struct {
	Cluster
}

// ClustersClientDeleteResponse contains the response from method ClustersClient.Delete.
type ClustersClientDeleteResponse struct {
	// placeholder for future response values
}

// ClustersClientGetResponse contains the response from method ClustersClient.Get.
type ClustersClientGetResponse struct {
	Cluster
}

// ClustersClientListByResourceGroupResponse contains the response from method ClustersClient.ListByResourceGroup.
type ClustersClientListByResourceGroupResponse struct {
	ClustersList
}

// ClustersClientListResponse contains the response from method ClustersClient.List.
type ClustersClientListResponse struct {
	ClustersList
}

// ClustersClientUpdateResponse contains the response from method ClustersClient.Update.
type ClustersClientUpdateResponse struct {
	Cluster
}

// DatastoresClientCreateResponse contains the response from method DatastoresClient.Create.
type DatastoresClientCreateResponse struct {
	Datastore
}

// DatastoresClientDeleteResponse contains the response from method DatastoresClient.Delete.
type DatastoresClientDeleteResponse struct {
	// placeholder for future response values
}

// DatastoresClientGetResponse contains the response from method DatastoresClient.Get.
type DatastoresClientGetResponse struct {
	Datastore
}

// DatastoresClientListByResourceGroupResponse contains the response from method DatastoresClient.ListByResourceGroup.
type DatastoresClientListByResourceGroupResponse struct {
	DatastoresList
}

// DatastoresClientListResponse contains the response from method DatastoresClient.List.
type DatastoresClientListResponse struct {
	DatastoresList
}

// DatastoresClientUpdateResponse contains the response from method DatastoresClient.Update.
type DatastoresClientUpdateResponse struct {
	Datastore
}

// GuestAgentsClientCreateResponse contains the response from method GuestAgentsClient.Create.
type GuestAgentsClientCreateResponse struct {
	GuestAgent
}

// GuestAgentsClientDeleteResponse contains the response from method GuestAgentsClient.Delete.
type GuestAgentsClientDeleteResponse struct {
	// placeholder for future response values
}

// GuestAgentsClientGetResponse contains the response from method GuestAgentsClient.Get.
type GuestAgentsClientGetResponse struct {
	GuestAgent
}

// GuestAgentsClientListByVMResponse contains the response from method GuestAgentsClient.ListByVM.
type GuestAgentsClientListByVMResponse struct {
	GuestAgentList
}

// HostsClientCreateResponse contains the response from method HostsClient.Create.
type HostsClientCreateResponse struct {
	Host
}

// HostsClientDeleteResponse contains the response from method HostsClient.Delete.
type HostsClientDeleteResponse struct {
	// placeholder for future response values
}

// HostsClientGetResponse contains the response from method HostsClient.Get.
type HostsClientGetResponse struct {
	Host
}

// HostsClientListByResourceGroupResponse contains the response from method HostsClient.ListByResourceGroup.
type HostsClientListByResourceGroupResponse struct {
	HostsList
}

// HostsClientListResponse contains the response from method HostsClient.List.
type HostsClientListResponse struct {
	HostsList
}

// HostsClientUpdateResponse contains the response from method HostsClient.Update.
type HostsClientUpdateResponse struct {
	Host
}

// HybridIdentityMetadataClientCreateResponse contains the response from method HybridIdentityMetadataClient.Create.
type HybridIdentityMetadataClientCreateResponse struct {
	HybridIdentityMetadata
}

// HybridIdentityMetadataClientDeleteResponse contains the response from method HybridIdentityMetadataClient.Delete.
type HybridIdentityMetadataClientDeleteResponse struct {
	// placeholder for future response values
}

// HybridIdentityMetadataClientGetResponse contains the response from method HybridIdentityMetadataClient.Get.
type HybridIdentityMetadataClientGetResponse struct {
	HybridIdentityMetadata
}

// HybridIdentityMetadataClientListByVMResponse contains the response from method HybridIdentityMetadataClient.ListByVM.
type HybridIdentityMetadataClientListByVMResponse struct {
	HybridIdentityMetadataList
}

// InventoryItemsClientCreateResponse contains the response from method InventoryItemsClient.Create.
type InventoryItemsClientCreateResponse struct {
	InventoryItem
}

// InventoryItemsClientDeleteResponse contains the response from method InventoryItemsClient.Delete.
type InventoryItemsClientDeleteResponse struct {
	// placeholder for future response values
}

// InventoryItemsClientGetResponse contains the response from method InventoryItemsClient.Get.
type InventoryItemsClientGetResponse struct {
	InventoryItem
}

// InventoryItemsClientListByVCenterResponse contains the response from method InventoryItemsClient.ListByVCenter.
type InventoryItemsClientListByVCenterResponse struct {
	InventoryItemsList
}

// MachineExtensionsClientCreateOrUpdateResponse contains the response from method MachineExtensionsClient.CreateOrUpdate.
type MachineExtensionsClientCreateOrUpdateResponse struct {
	MachineExtension
}

// MachineExtensionsClientDeleteResponse contains the response from method MachineExtensionsClient.Delete.
type MachineExtensionsClientDeleteResponse struct {
	// placeholder for future response values
}

// MachineExtensionsClientGetResponse contains the response from method MachineExtensionsClient.Get.
type MachineExtensionsClientGetResponse struct {
	MachineExtension
}

// MachineExtensionsClientListResponse contains the response from method MachineExtensionsClient.List.
type MachineExtensionsClientListResponse struct {
	MachineExtensionsListResult
}

// MachineExtensionsClientUpdateResponse contains the response from method MachineExtensionsClient.Update.
type MachineExtensionsClientUpdateResponse struct {
	MachineExtension
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationsList
}

// ResourcePoolsClientCreateResponse contains the response from method ResourcePoolsClient.Create.
type ResourcePoolsClientCreateResponse struct {
	ResourcePool
}

// ResourcePoolsClientDeleteResponse contains the response from method ResourcePoolsClient.Delete.
type ResourcePoolsClientDeleteResponse struct {
	// placeholder for future response values
}

// ResourcePoolsClientGetResponse contains the response from method ResourcePoolsClient.Get.
type ResourcePoolsClientGetResponse struct {
	ResourcePool
}

// ResourcePoolsClientListByResourceGroupResponse contains the response from method ResourcePoolsClient.ListByResourceGroup.
type ResourcePoolsClientListByResourceGroupResponse struct {
	ResourcePoolsList
}

// ResourcePoolsClientListResponse contains the response from method ResourcePoolsClient.List.
type ResourcePoolsClientListResponse struct {
	ResourcePoolsList
}

// ResourcePoolsClientUpdateResponse contains the response from method ResourcePoolsClient.Update.
type ResourcePoolsClientUpdateResponse struct {
	ResourcePool
}

// VCentersClientCreateResponse contains the response from method VCentersClient.Create.
type VCentersClientCreateResponse struct {
	VCenter
}

// VCentersClientDeleteResponse contains the response from method VCentersClient.Delete.
type VCentersClientDeleteResponse struct {
	// placeholder for future response values
}

// VCentersClientGetResponse contains the response from method VCentersClient.Get.
type VCentersClientGetResponse struct {
	VCenter
}

// VCentersClientListByResourceGroupResponse contains the response from method VCentersClient.ListByResourceGroup.
type VCentersClientListByResourceGroupResponse struct {
	VCentersList
}

// VCentersClientListResponse contains the response from method VCentersClient.List.
type VCentersClientListResponse struct {
	VCentersList
}

// VCentersClientUpdateResponse contains the response from method VCentersClient.Update.
type VCentersClientUpdateResponse struct {
	VCenter
}

// VirtualMachineTemplatesClientCreateResponse contains the response from method VirtualMachineTemplatesClient.Create.
type VirtualMachineTemplatesClientCreateResponse struct {
	VirtualMachineTemplate
}

// VirtualMachineTemplatesClientDeleteResponse contains the response from method VirtualMachineTemplatesClient.Delete.
type VirtualMachineTemplatesClientDeleteResponse struct {
	// placeholder for future response values
}

// VirtualMachineTemplatesClientGetResponse contains the response from method VirtualMachineTemplatesClient.Get.
type VirtualMachineTemplatesClientGetResponse struct {
	VirtualMachineTemplate
}

// VirtualMachineTemplatesClientListByResourceGroupResponse contains the response from method VirtualMachineTemplatesClient.ListByResourceGroup.
type VirtualMachineTemplatesClientListByResourceGroupResponse struct {
	VirtualMachineTemplatesList
}

// VirtualMachineTemplatesClientListResponse contains the response from method VirtualMachineTemplatesClient.List.
type VirtualMachineTemplatesClientListResponse struct {
	VirtualMachineTemplatesList
}

// VirtualMachineTemplatesClientUpdateResponse contains the response from method VirtualMachineTemplatesClient.Update.
type VirtualMachineTemplatesClientUpdateResponse struct {
	VirtualMachineTemplate
}

// VirtualMachinesClientAssessPatchesResponse contains the response from method VirtualMachinesClient.AssessPatches.
type VirtualMachinesClientAssessPatchesResponse struct {
	VirtualMachineAssessPatchesResult
}

// VirtualMachinesClientCreateResponse contains the response from method VirtualMachinesClient.Create.
type VirtualMachinesClientCreateResponse struct {
	VirtualMachine
}

// VirtualMachinesClientDeleteResponse contains the response from method VirtualMachinesClient.Delete.
type VirtualMachinesClientDeleteResponse struct {
	// placeholder for future response values
}

// VirtualMachinesClientGetResponse contains the response from method VirtualMachinesClient.Get.
type VirtualMachinesClientGetResponse struct {
	VirtualMachine
}

// VirtualMachinesClientInstallPatchesResponse contains the response from method VirtualMachinesClient.InstallPatches.
type VirtualMachinesClientInstallPatchesResponse struct {
	VirtualMachineInstallPatchesResult
}

// VirtualMachinesClientListByResourceGroupResponse contains the response from method VirtualMachinesClient.ListByResourceGroup.
type VirtualMachinesClientListByResourceGroupResponse struct {
	VirtualMachinesList
}

// VirtualMachinesClientListResponse contains the response from method VirtualMachinesClient.List.
type VirtualMachinesClientListResponse struct {
	VirtualMachinesList
}

// VirtualMachinesClientRestartResponse contains the response from method VirtualMachinesClient.Restart.
type VirtualMachinesClientRestartResponse struct {
	// placeholder for future response values
}

// VirtualMachinesClientStartResponse contains the response from method VirtualMachinesClient.Start.
type VirtualMachinesClientStartResponse struct {
	// placeholder for future response values
}

// VirtualMachinesClientStopResponse contains the response from method VirtualMachinesClient.Stop.
type VirtualMachinesClientStopResponse struct {
	// placeholder for future response values
}

// VirtualMachinesClientUpdateResponse contains the response from method VirtualMachinesClient.Update.
type VirtualMachinesClientUpdateResponse struct {
	VirtualMachine
}

// VirtualNetworksClientCreateResponse contains the response from method VirtualNetworksClient.Create.
type VirtualNetworksClientCreateResponse struct {
	VirtualNetwork
}

// VirtualNetworksClientDeleteResponse contains the response from method VirtualNetworksClient.Delete.
type VirtualNetworksClientDeleteResponse struct {
	// placeholder for future response values
}

// VirtualNetworksClientGetResponse contains the response from method VirtualNetworksClient.Get.
type VirtualNetworksClientGetResponse struct {
	VirtualNetwork
}

// VirtualNetworksClientListByResourceGroupResponse contains the response from method VirtualNetworksClient.ListByResourceGroup.
type VirtualNetworksClientListByResourceGroupResponse struct {
	VirtualNetworksList
}

// VirtualNetworksClientListResponse contains the response from method VirtualNetworksClient.List.
type VirtualNetworksClientListResponse struct {
	VirtualNetworksList
}

// VirtualNetworksClientUpdateResponse contains the response from method VirtualNetworksClient.Update.
type VirtualNetworksClientUpdateResponse struct {
	VirtualNetwork
}
