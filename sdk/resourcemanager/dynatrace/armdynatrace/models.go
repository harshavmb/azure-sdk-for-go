//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdynatrace

import "time"

// AccountInfo - Dynatrace Account Information
type AccountInfo struct {
	// Account Id of the account this environment is linked to
	AccountID *string `json:"accountId,omitempty"`

	// Region in which the account is created
	RegionID *string `json:"regionId,omitempty"`
}

// AccountInfoSecure - Dynatrace account API Key
type AccountInfoSecure struct {
	// READ-ONLY; API Key of the user account
	APIKey *string `json:"apiKey,omitempty" azure:"ro"`

	// READ-ONLY; Account Id of the account this environment is linked to
	AccountID *string `json:"accountId,omitempty" azure:"ro"`

	// READ-ONLY; Region in which the account is created
	RegionID *string `json:"regionId,omitempty" azure:"ro"`
}

// AppServiceInfo - Details of App Services having Dynatrace OneAgent installed
type AppServiceInfo struct {
	// Update settings of OneAgent.
	AutoUpdateSetting *AutoUpdateSetting `json:"autoUpdateSetting,omitempty"`

	// The availability state of OneAgent.
	AvailabilityState *AvailabilityState `json:"availabilityState,omitempty"`

	// The name of the host group
	HostGroup *string `json:"hostGroup,omitempty"`

	// The name of the host
	HostName *string `json:"hostName,omitempty"`

	// Tells whether log modules are enabled or not
	LogModule *LogModule `json:"logModule,omitempty"`

	// The monitoring mode of OneAgent
	MonitoringType *MonitoringType `json:"monitoringType,omitempty"`

	// App service resource ID
	ResourceID *string `json:"resourceId,omitempty"`

	// The current update status of OneAgent.
	UpdateStatus *UpdateStatus `json:"updateStatus,omitempty"`

	// Version of the Dynatrace agent installed on the App Service.
	Version *string `json:"version,omitempty"`
}

// AppServiceListResponse - Response of a list App Services Operation.
type AppServiceListResponse struct {
	// The link to the next page of items
	NextLink *string `json:"nextLink,omitempty"`

	// The items on this page
	Value []*AppServiceInfo `json:"value,omitempty"`
}

// EnvironmentInfo - Dynatrace Environment Information
type EnvironmentInfo struct {
	// Id of the environment created
	EnvironmentID *string `json:"environmentId,omitempty"`

	// Ingestion key of the environment
	IngestionKey *string `json:"ingestionKey,omitempty"`

	// Landing URL for Dynatrace environment
	LandingURL *string `json:"landingURL,omitempty"`

	// Ingestion endpoint used for sending logs
	LogsIngestionEndpoint *string `json:"logsIngestionEndpoint,omitempty"`
}

// EnvironmentProperties - Properties of the Dynatrace environment.
type EnvironmentProperties struct {
	// Dynatrace Account Information
	AccountInfo *AccountInfo `json:"accountInfo,omitempty"`

	// Dynatrace Environment Information
	EnvironmentInfo *EnvironmentInfo `json:"environmentInfo,omitempty"`

	// The details of a Dynatrace single sign-on.
	SingleSignOnProperties *SingleSignOnProperties `json:"singleSignOnProperties,omitempty"`

	// User id
	UserID *string `json:"userId,omitempty"`
}

// FilteringTag - The definition of a filtering tag. Filtering tags are used for capturing resources and include/exclude them
// from being monitored.
type FilteringTag struct {
	// Valid actions for a filtering tag. Exclusion takes priority over inclusion.
	Action *TagAction `json:"action,omitempty"`

	// The name (also known as the key) of the tag.
	Name *string `json:"name,omitempty"`

	// The value of the tag.
	Value *string `json:"value,omitempty"`
}

// IdentityProperties - The properties of the managed service identities assigned to this resource.
type IdentityProperties struct {
	// REQUIRED; The type of managed identity assigned to this resource.
	Type *ManagedIdentityType `json:"type,omitempty"`

	// The identities assigned to this resource by the user.
	UserAssignedIdentities map[string]*UserAssignedIdentity `json:"userAssignedIdentities,omitempty"`

	// READ-ONLY; The active directory identifier of this principal.
	PrincipalID *string `json:"principalId,omitempty" azure:"ro"`

	// READ-ONLY; The Active Directory tenant id of the principal.
	TenantID *string `json:"tenantId,omitempty" azure:"ro"`
}

// LinkableEnvironmentListResponse - Response for getting all the linkable environments
type LinkableEnvironmentListResponse struct {
	// Link to the next set of results, if any.
	NextLink *string `json:"nextLink,omitempty"`

	// List of environments for which user is an admin
	Value []*LinkableEnvironmentResponse `json:"value,omitempty"`
}

// LinkableEnvironmentRequest - Request for getting all the linkable environments for a user
type LinkableEnvironmentRequest struct {
	// Azure region in which we want to link the environment
	Region *string `json:"region,omitempty"`

	// Tenant Id of the user in which they want to link the environment
	TenantID *string `json:"tenantId,omitempty"`

	// user principal id of the user
	UserPrincipal *string `json:"userPrincipal,omitempty"`
}

// LinkableEnvironmentResponse - Response for getting all the linkable environments
type LinkableEnvironmentResponse struct {
	// environment id for which user is an admin
	EnvironmentID *string `json:"environmentId,omitempty"`

	// Name of the environment
	EnvironmentName *string `json:"environmentName,omitempty"`

	// Billing plan information.
	PlanData *PlanData `json:"planData,omitempty"`
}

// LogRules - Set of rules for sending logs for the Monitor resource.
type LogRules struct {
	// List of filtering tags to be used for capturing logs. This only takes effect if SendActivityLogs flag is enabled. If empty,
	// all resources will be captured. If only Exclude action is specified, the
	// rules will apply to the list of all available resources. If Include actions are specified, the rules will only include
	// resources with the associated tags.
	FilteringTags []*FilteringTag `json:"filteringTags,omitempty"`

	// Flag specifying if AAD logs should be sent for the Monitor resource.
	SendAADLogs *SendAADLogsStatus `json:"sendAadLogs,omitempty"`

	// Flag specifying if activity logs from Azure resources should be sent for the Monitor resource.
	SendActivityLogs *SendActivityLogsStatus `json:"sendActivityLogs,omitempty"`

	// Flag specifying if subscription logs should be sent for the Monitor resource.
	SendSubscriptionLogs *SendSubscriptionLogsStatus `json:"sendSubscriptionLogs,omitempty"`
}

// MetricRules - Set of rules for sending metrics for the Monitor resource.
type MetricRules struct {
	// List of filtering tags to be used for capturing metrics. If empty, all resources will be captured. If only Exclude action
	// is specified, the rules will apply to the list of all available resources. If
	// Include actions are specified, the rules will only include resources with the associated tags.
	FilteringTags []*FilteringTag `json:"filteringTags,omitempty"`
}

// MonitorProperties - Properties specific to the monitor resource.
type MonitorProperties struct {
	// Properties of the Dynatrace environment.
	DynatraceEnvironmentProperties *EnvironmentProperties `json:"dynatraceEnvironmentProperties,omitempty"`

	// Marketplace subscription status.
	MarketplaceSubscriptionStatus *MarketplaceSubscriptionStatus `json:"marketplaceSubscriptionStatus,omitempty"`

	// Status of the monitor.
	MonitoringStatus *MonitoringStatus `json:"monitoringStatus,omitempty"`

	// Billing plan information.
	PlanData *PlanData `json:"planData,omitempty"`

	// User info.
	UserInfo *UserInfo `json:"userInfo,omitempty"`

	// READ-ONLY; Liftr Resource category.
	LiftrResourceCategory *LiftrResourceCategories `json:"liftrResourceCategory,omitempty" azure:"ro"`

	// READ-ONLY; The priority of the resource.
	LiftrResourcePreference *int32 `json:"liftrResourcePreference,omitempty" azure:"ro"`

	// READ-ONLY; Provisioning state of the resource.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// MonitorResource - Dynatrace Monitor Resource
type MonitorResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// REQUIRED; The resource-specific properties for this resource.
	Properties *MonitorProperties `json:"properties,omitempty"`

	// The managed service identities assigned to this resource.
	Identity *IdentityProperties `json:"identity,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; System metadata for this resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// MonitorResourceListResult - The response of a MonitorResource list operation.
type MonitorResourceListResult struct {
	// REQUIRED; The items on this page
	Value []*MonitorResource `json:"value,omitempty"`

	// The link to the next page of items
	NextLink *string `json:"nextLink,omitempty"`
}

// MonitorResourceUpdate - The updatable properties of the MonitorResource.
type MonitorResourceUpdate struct {
	// Properties of the Dynatrace environment.
	DynatraceEnvironmentProperties *EnvironmentProperties `json:"dynatraceEnvironmentProperties,omitempty"`

	// Marketplace subscription status.
	MarketplaceSubscriptionStatus *MarketplaceSubscriptionStatus `json:"marketplaceSubscriptionStatus,omitempty"`

	// Status of the monitor.
	MonitoringStatus *MonitoringStatus `json:"monitoringStatus,omitempty"`

	// Billing plan information.
	PlanData *PlanData `json:"planData,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// User info.
	UserInfo *UserInfo `json:"userInfo,omitempty"`
}

// MonitoredResource - Details of resource being monitored by Dynatrace monitor resource
type MonitoredResource struct {
	// The ARM id of the resource.
	ID *string `json:"id,omitempty"`

	// Reason for why the resource is sending logs (or why it is not sending).
	ReasonForLogsStatus *string `json:"reasonForLogsStatus,omitempty"`

	// Reason for why the resource is sending metrics (or why it is not sending).
	ReasonForMetricsStatus *string `json:"reasonForMetricsStatus,omitempty"`

	// Flag indicating if resource is sending logs to Dynatrace.
	SendingLogs *SendingLogsStatus `json:"sendingLogs,omitempty"`

	// Flag indicating if resource is sending metrics to Dynatrace.
	SendingMetrics *SendingMetricsStatus `json:"sendingMetrics,omitempty"`
}

// MonitoredResourceListResponse - List of all the resources being monitored by Dynatrace monitor resource
type MonitoredResourceListResponse struct {
	// The link to the next page of items
	NextLink *string `json:"nextLink,omitempty"`

	// The items on this page
	Value []*MonitoredResource `json:"value,omitempty"`
}

// MonitoringTagRulesProperties - Properties for the Tag rules resource of a Monitor account.
type MonitoringTagRulesProperties struct {
	// Set of rules for sending logs for the Monitor resource.
	LogRules *LogRules `json:"logRules,omitempty"`

	// Set of rules for sending metrics for the Monitor resource.
	MetricRules *MetricRules `json:"metricRules,omitempty"`

	// READ-ONLY; Provisioning state of the resource.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// MonitorsClientBeginCreateOrUpdateOptions contains the optional parameters for the MonitorsClient.BeginCreateOrUpdate method.
type MonitorsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// MonitorsClientBeginDeleteOptions contains the optional parameters for the MonitorsClient.BeginDelete method.
type MonitorsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// MonitorsClientGetAccountCredentialsOptions contains the optional parameters for the MonitorsClient.GetAccountCredentials
// method.
type MonitorsClientGetAccountCredentialsOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientGetOptions contains the optional parameters for the MonitorsClient.Get method.
type MonitorsClientGetOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientGetSSODetailsOptions contains the optional parameters for the MonitorsClient.GetSSODetails method.
type MonitorsClientGetSSODetailsOptions struct {
	// The details of the get sso details request.
	Request *SSODetailsRequest
}

// MonitorsClientGetVMHostPayloadOptions contains the optional parameters for the MonitorsClient.GetVMHostPayload method.
type MonitorsClientGetVMHostPayloadOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientListAppServicesOptions contains the optional parameters for the MonitorsClient.ListAppServices method.
type MonitorsClientListAppServicesOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientListByResourceGroupOptions contains the optional parameters for the MonitorsClient.ListByResourceGroup method.
type MonitorsClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientListBySubscriptionIDOptions contains the optional parameters for the MonitorsClient.ListBySubscriptionID
// method.
type MonitorsClientListBySubscriptionIDOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientListHostsOptions contains the optional parameters for the MonitorsClient.ListHosts method.
type MonitorsClientListHostsOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientListLinkableEnvironmentsOptions contains the optional parameters for the MonitorsClient.ListLinkableEnvironments
// method.
type MonitorsClientListLinkableEnvironmentsOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientListMonitoredResourcesOptions contains the optional parameters for the MonitorsClient.ListMonitoredResources
// method.
type MonitorsClientListMonitoredResourcesOptions struct {
	// placeholder for future optional parameters
}

// MonitorsClientUpdateOptions contains the optional parameters for the MonitorsClient.Update method.
type MonitorsClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType `json:"actionType,omitempty" azure:"ro"`

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool `json:"isDataAction,omitempty" azure:"ro"`

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin `json:"origin,omitempty" azure:"ro"`
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string `json:"operation,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string `json:"provider,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string `json:"resource,omitempty" azure:"ro"`
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation `json:"value,omitempty" azure:"ro"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// PlanData - Billing plan information.
type PlanData struct {
	// different billing cycles like MONTHLY/WEEKLY. this could be enum
	BillingCycle *string `json:"billingCycle,omitempty"`

	// date when plan was applied
	EffectiveDate *time.Time `json:"effectiveDate,omitempty"`

	// plan id as published by Dynatrace
	PlanDetails *string `json:"planDetails,omitempty"`

	// different usage type like PAYG/COMMITTED. this could be enum
	UsageType *string `json:"usageType,omitempty"`
}

// SSODetailsRequest - Request for getting sso details for a user
type SSODetailsRequest struct {
	// user principal id of the user
	UserPrincipal *string `json:"userPrincipal,omitempty"`
}

// SSODetailsResponse - SSO details from the Dynatrace partner
type SSODetailsResponse struct {
	// array of Aad(azure active directory) domains
	AADDomains []*string `json:"aadDomains,omitempty"`

	// Array of admin user emails.
	AdminUsers []*string `json:"adminUsers,omitempty"`

	// Whether the SSO is enabled for this resource or not.
	IsSsoEnabled *SSOStatus `json:"isSsoEnabled,omitempty"`

	// URL for Azure AD metadata
	MetadataURL *string `json:"metadataUrl,omitempty"`

	// The login URL specific to this Dynatrace Environment
	SingleSignOnURL *string `json:"singleSignOnUrl,omitempty"`
}

// SingleSignOnClientBeginCreateOrUpdateOptions contains the optional parameters for the SingleSignOnClient.BeginCreateOrUpdate
// method.
type SingleSignOnClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// SingleSignOnClientGetOptions contains the optional parameters for the SingleSignOnClient.Get method.
type SingleSignOnClientGetOptions struct {
	// placeholder for future optional parameters
}

// SingleSignOnClientListOptions contains the optional parameters for the SingleSignOnClient.List method.
type SingleSignOnClientListOptions struct {
	// placeholder for future optional parameters
}

// SingleSignOnProperties - The details of a Dynatrace single sign-on.
type SingleSignOnProperties struct {
	// array of Aad(azure active directory) domains
	AADDomains []*string `json:"aadDomains,omitempty"`

	// Version of the Dynatrace agent installed on the VM.
	EnterpriseAppID *string `json:"enterpriseAppId,omitempty"`

	// State of Single Sign On
	SingleSignOnState *SingleSignOnStates `json:"singleSignOnState,omitempty"`

	// The login URL specific to this Dynatrace Environment
	SingleSignOnURL *string `json:"singleSignOnUrl,omitempty"`

	// READ-ONLY; Provisioning state of the resource.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// SingleSignOnResource - Single sign-on configurations for a given monitor resource.
type SingleSignOnResource struct {
	// REQUIRED; The resource-specific properties for this resource.
	Properties *SingleSignOnProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; System metadata for this resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// SingleSignOnResourceListResult - The response of a DynatraceSingleSignOnResource list operation.
type SingleSignOnResourceListResult struct {
	// REQUIRED; The items on this page
	Value []*SingleSignOnResource `json:"value,omitempty"`

	// The link to the next page of items
	NextLink *string `json:"nextLink,omitempty"`
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}

// TagRule - Tag rules for a monitor resource
type TagRule struct {
	// REQUIRED; The resource-specific properties for this resource.
	Properties *MonitoringTagRulesProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; System metadata for this resource.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// TagRuleListResult - The response of a TagRule list operation.
type TagRuleListResult struct {
	// REQUIRED; The items on this page
	Value []*TagRule `json:"value,omitempty"`

	// The link to the next page of items
	NextLink *string `json:"nextLink,omitempty"`
}

// TagRuleUpdate - The updatable properties of the TagRule.
type TagRuleUpdate struct {
	// Set of rules for sending logs for the Monitor resource.
	LogRules *LogRules `json:"logRules,omitempty"`

	// Set of rules for sending metrics for the Monitor resource.
	MetricRules *MetricRules `json:"metricRules,omitempty"`
}

// TagRulesClientBeginCreateOrUpdateOptions contains the optional parameters for the TagRulesClient.BeginCreateOrUpdate method.
type TagRulesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// TagRulesClientBeginDeleteOptions contains the optional parameters for the TagRulesClient.BeginDelete method.
type TagRulesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// TagRulesClientGetOptions contains the optional parameters for the TagRulesClient.Get method.
type TagRulesClientGetOptions struct {
	// placeholder for future optional parameters
}

// TagRulesClientListOptions contains the optional parameters for the TagRulesClient.List method.
type TagRulesClientListOptions struct {
	// placeholder for future optional parameters
}

// TagRulesClientUpdateOptions contains the optional parameters for the TagRulesClient.Update method.
type TagRulesClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// UserAssignedIdentity - A managed identity assigned by the user.
type UserAssignedIdentity struct {
	// REQUIRED; The active directory client identifier for this principal.
	ClientID *string `json:"clientId,omitempty"`

	// REQUIRED; The active directory identifier for this principal.
	PrincipalID *string `json:"principalId,omitempty"`
}

// UserInfo - User info.
type UserInfo struct {
	// Country of the user
	Country *string `json:"country,omitempty"`

	// Email of the user used by Dynatrace for contacting them if needed
	EmailAddress *string `json:"emailAddress,omitempty"`

	// First Name of the user
	FirstName *string `json:"firstName,omitempty"`

	// Last Name of the user
	LastName *string `json:"lastName,omitempty"`

	// Phone number of the user used by Dynatrace for contacting them if needed
	PhoneNumber *string `json:"phoneNumber,omitempty"`
}

// VMExtensionPayload - Response of payload to be passed while installing VM agent.
type VMExtensionPayload struct {
	// Id of the environment created
	EnvironmentID *string `json:"environmentId,omitempty"`

	// Ingestion key of the environment
	IngestionKey *string `json:"ingestionKey,omitempty"`
}

// VMHostsListResponse - Response of a list VM Host Operation.
type VMHostsListResponse struct {
	// The link to the next page of items
	NextLink *string `json:"nextLink,omitempty"`

	// The items on this page
	Value []*VMInfo `json:"value,omitempty"`
}

// VMInfo - Details of VM Resource having Dynatrace OneAgent installed
type VMInfo struct {
	// Update settings of OneAgent.
	AutoUpdateSetting *AutoUpdateSetting `json:"autoUpdateSetting,omitempty"`

	// The availability state of OneAgent.
	AvailabilityState *AvailabilityState `json:"availabilityState,omitempty"`

	// The name of the host group
	HostGroup *string `json:"hostGroup,omitempty"`

	// The name of the host
	HostName *string `json:"hostName,omitempty"`

	// Tells whether log modules are enabled or not
	LogModule *LogModule `json:"logModule,omitempty"`

	// The monitoring mode of OneAgent
	MonitoringType *MonitoringType `json:"monitoringType,omitempty"`

	// Azure VM resource ID
	ResourceID *string `json:"resourceId,omitempty"`

	// The current update status of OneAgent.
	UpdateStatus *UpdateStatus `json:"updateStatus,omitempty"`

	// Version of the Dynatrace agent installed on the VM.
	Version *string `json:"version,omitempty"`
}
