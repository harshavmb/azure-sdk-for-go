//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package service_test

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/internal/generated"
	"log"
	"os"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/sas"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/service"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func Example_service_Client_NewClient() {
	accountName, ok := os.LookupEnv("AZURE_STORAGE_ACCOUNT_NAME")
	if !ok {
		panic("AZURE_STORAGE_ACCOUNT_NAME could not be found")
	}
	serviceURL := fmt.Sprintf("https://%s.blob.core.windows.net/", accountName)

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)
	serviceClient, err := service.NewClient(serviceURL, cred, nil)
	handleError(err)
	fmt.Println(serviceClient.URL())
}

func Example_service_Client_NewClientWithSharedKeyCredential() {
	accountName, ok := os.LookupEnv("AZURE_STORAGE_ACCOUNT_NAME")
	if !ok {
		panic("AZURE_STORAGE_ACCOUNT_NAME could not be found")
	}
	accountKey, ok := os.LookupEnv("AZURE_STORAGE_ACCOUNT_KEY")
	if !ok {
		panic("AZURE_STORAGE_ACCOUNT_KEY could not be found")
	}
	serviceURL := fmt.Sprintf("https://%s.blob.core.windows.net/", accountName)

	cred, err := service.NewSharedKeyCredential(accountName, accountKey)
	handleError(err)
	serviceClient, err := service.NewClientWithSharedKeyCredential(serviceURL, cred, nil)
	handleError(err)
	fmt.Println(serviceClient.URL())
}

func Example_service_Client_NewClientWithNoCredential() {
	accountName, ok := os.LookupEnv("AZURE_STORAGE_ACCOUNT_NAME")
	if !ok {
		panic("AZURE_STORAGE_ACCOUNT_NAME could not be found")
	}
	sharedAccessSignature, ok := os.LookupEnv("AZURE_STORAGE_SHARED_ACCESS_SIGNATURE")
	if !ok {
		panic("AZURE_STORAGE_SHARED_ACCESS_SIGNATURE could not be found")
	}
	serviceURL := fmt.Sprintf("https://%s.blob.core.windows.net/?%s", accountName, sharedAccessSignature)

	serviceClient, err := service.NewClientWithNoCredential(serviceURL, nil)
	handleError(err)
	fmt.Println(serviceClient.URL())
}

func Example_service_Client_NewClientFromConnectionString() {
	// Your connection string can be obtained from the Azure Portal.
	connectionString, ok := os.LookupEnv("AZURE_STORAGE_CONNECTION_STRING")
	if !ok {
		log.Fatal("the environment variable 'AZURE_STORAGE_CONNECTION_STRING' could not be found")
	}

	serviceClient, err := service.NewClientFromConnectionString(connectionString, nil)
	handleError(err)
	fmt.Println(serviceClient.URL())
}

func Example_service_Client_CreateContainer() {
	accountName, ok := os.LookupEnv("AZURE_STORAGE_ACCOUNT_NAME")
	if !ok {
		panic("AZURE_STORAGE_ACCOUNT_NAME could not be found")
	}
	serviceURL := fmt.Sprintf("https://%s.blob.core.windows.net/", accountName)

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)

	serviceClient, err := service.NewClient(serviceURL, cred, nil)
	handleError(err)

	_, err = serviceClient.CreateContainer(context.TODO(), "testcontainer", nil)
	handleError(err)

	// ======== 2. Delete a container ========
	defer func(serviceClient1 *service.Client, ctx context.Context, containerName string, options *container.DeleteOptions) {
		_, err = serviceClient1.DeleteContainer(ctx, containerName, options)
		if err != nil {
			log.Fatal(err)
		}
	}(serviceClient, context.TODO(), "testcontainer", nil)
}

func Example_service_Client_DeleteContainer() {
	accountName, ok := os.LookupEnv("AZURE_STORAGE_ACCOUNT_NAME")
	if !ok {
		panic("AZURE_STORAGE_ACCOUNT_NAME could not be found")
	}
	serviceURL := fmt.Sprintf("https://%s.blob.core.windows.net/", accountName)

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)
	serviceClient, err := service.NewClient(serviceURL, cred, nil)
	handleError(err)

	_, err = serviceClient.DeleteContainer(context.TODO(), "testcontainer", nil)
	handleError(err)
}

func Example_service_Client_RestoreContainer() {
	accountName, ok := os.LookupEnv("AZURE_STORAGE_ACCOUNT_NAME")
	if !ok {
		panic("AZURE_STORAGE_ACCOUNT_NAME could not be found")
	}
	serviceURL := fmt.Sprintf("https://%s.blob.core.windows.net/", accountName)

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)
	serviceClient, err := service.NewClient(serviceURL, cred, nil)
	handleError(err)

	listOptions := service.ListContainersOptions{
		Include: service.ListContainersInclude{
			Metadata: true, // Include Metadata
			Deleted:  true, // Include deleted containers in the result as well
		},
	}
	pager := serviceClient.NewListContainersPager(&listOptions)

	for pager.More() {
		resp, err := pager.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		for _, cont := range resp.ContainerItems {
			if *cont.Deleted {
				_, err = serviceClient.RestoreContainer(context.TODO(), *cont.Name, *cont.Version, nil)
				handleError(err)
			}
		}
	}
}

func Example_service_Client_ListContainers() {
	accountName, ok := os.LookupEnv("AZURE_STORAGE_ACCOUNT_NAME")
	if !ok {
		panic("AZURE_STORAGE_ACCOUNT_NAME could not be found")
	}
	serviceURL := fmt.Sprintf("https://%s.blob.core.windows.net/", accountName)

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)
	serviceClient, err := service.NewClient(serviceURL, cred, nil)
	handleError(err)

	listContainersOptions := service.ListContainersOptions{
		Include: service.ListContainersInclude{
			Metadata: true, // Include Metadata
			Deleted:  true, // Include deleted containers in the result as well
		},
	}
	pager := serviceClient.NewListContainersPager(&listContainersOptions)

	for pager.More() {
		resp, err := pager.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		for _, container := range resp.ContainerItems {
			fmt.Println(*container.Name)
		}
	}
}

func Example_service_Client_GetSASURL() {
	cred, err := azblob.NewSharedKeyCredential("myAccountName", "myAccountKey")
	handleError(err)
	serviceClient, err := service.NewClientWithSharedKeyCredential("https://<myAccountName>.blob.core.windows.net", cred, nil)
	handleError(err)

	resources := sas.AccountResourceTypes{Service: true}
	permission := sas.AccountPermissions{Read: true}
	start := time.Now()
	expiry := start.AddDate(1, 0, 0)
	sasURL, err := serviceClient.GetSASURL(resources, permission, sas.AccountServices{Blob: true}, start, expiry)
	handleError(err)

	serviceURL := fmt.Sprintf("https://<myAccountName>.blob.core.windows.net/?%s", sasURL)
	serviceClientWithSAS, err := service.NewClientWithNoCredential(serviceURL, nil)
	handleError(err)
	_ = serviceClientWithSAS
}

func Example_service_Client_SetProperties() {
	accountName, ok := os.LookupEnv("AZURE_STORAGE_ACCOUNT_NAME")
	if !ok {
		panic("AZURE_STORAGE_ACCOUNT_NAME could not be found")
	}
	serviceURL := fmt.Sprintf("https://%s.blob.core.windows.net/", accountName)

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)
	serviceClient, err := service.NewClient(serviceURL, cred, nil)
	handleError(err)

	enabled := true  // enabling retention period
	days := int32(5) // setting retention period to 5 days
	serviceSetPropertiesResponse, err := serviceClient.SetProperties(context.TODO(), &service.SetPropertiesOptions{
		DeleteRetentionPolicy: &service.RetentionPolicy{Enabled: &enabled, Days: &days},
	})

	handleError(err)
	fmt.Println(serviceSetPropertiesResponse)
}

func Example_service_Client_GetProperties() {
	accountName, ok := os.LookupEnv("AZURE_STORAGE_ACCOUNT_NAME")
	if !ok {
		panic("AZURE_STORAGE_ACCOUNT_NAME could not be found")
	}
	serviceURL := fmt.Sprintf("https://%s.blob.core.windows.net/", accountName)

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)

	serviceClient, err := service.NewClient(serviceURL, cred, nil)
	handleError(err)

	serviceGetPropertiesResponse, err := serviceClient.GetProperties(context.TODO(), nil)
	handleError(err)

	fmt.Println(serviceGetPropertiesResponse)
}

// This example shows how to create and use an Azure Storage account Shared Access Signature (SAS).
func Example_service_SASSignatureValues_Sign() {
	accountName, accountKey := os.Getenv("AZURE_STORAGE_ACCOUNT_NAME"), os.Getenv("AZURE_STORAGE_ACCOUNT_KEY")

	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	handleError(err)

	sasQueryParams, err := sas.AccountSignatureValues{
		Protocol:      sas.ProtocolHTTPS,
		ExpiryTime:    time.Now().UTC().Add(48 * time.Hour),
		Permissions:   to.Ptr(sas.AccountPermissions{Read: true, List: true}).String(),
		Services:      to.Ptr(sas.AccountServices{Blob: true}).String(),
		ResourceTypes: to.Ptr(sas.AccountResourceTypes{Container: true, Object: true}).String(),
	}.SignWithSharedKey(credential)
	handleError(err)

	sasURL := fmt.Sprintf("https://%s.blob.core.windows.net/?%s", accountName, sasQueryParams.Encode())

	// This URL can be used to authenticate requests now
	serviceClient, err := service.NewClientWithNoCredential(sasURL, nil)
	handleError(err)

	// You can also break a blob URL up into it's constituent parts
	blobURLParts, _ := blob.ParseURL(serviceClient.URL())
	fmt.Printf("SAS expiry time = %s\n", blobURLParts.SAS.ExpiryTime())
}

func Example_service_Client_NewClientWithUserDelegationCredential() {
	accountName, ok := os.LookupEnv("AZURE_STORAGE_ACCOUNT_NAME")
	if !ok {
		panic("AZURE_STORAGE_ACCOUNT_NAME could not be found")
	}
	// Create Managed Identity (OAuth) Credentials using Client ID
	clientOptions := azcore.ClientOptions{} // Fill clientOptions as needed
	optsClientID := azidentity.ManagedIdentityCredentialOptions{ClientOptions: clientOptions, ID: azidentity.ClientID("7cf7db0d-...")}
	cred, err := azidentity.NewManagedIdentityCredential(&optsClientID)
	handleError(err)
	clientOptionsService := service.ClientOptions{} // Same as azcore.ClientOptions using service instead

	svcClient, err := service.NewClient(fmt.Sprintf("https://%s.blob.core.windows.net/", accountName), cred, &clientOptionsService)
	handleError(err)

	// Set current and past time and create key
	currentTime := time.Now().UTC().Add(-10 * time.Second)
	pastTime := currentTime.Add(48 * time.Hour)
	info := generated.KeyInfo{
		Start:  to.Ptr(currentTime.UTC().Format(sas.TimeFormat)),
		Expiry: to.Ptr(pastTime.UTC().Format(sas.TimeFormat)),
	}

	udc, err := svcClient.GetUserDelegationCredential(context.Background(), info, nil)
	handleError(err)

	fmt.Println("User Delegation Key has been created for ", accountName)

	// Create Account Signature Values with desired permissions and sign with user delegation credential
	sasQueryParams, err := sas.AccountSignatureValues{
		Protocol:      sas.ProtocolHTTPS,
		ExpiryTime:    time.Now().UTC().Add(48 * time.Hour),
		Permissions:   to.Ptr(sas.AccountPermissions{Read: true, List: true}).String(),
		Services:      to.Ptr(sas.AccountServices{Blob: true}).String(),
		ResourceTypes: to.Ptr(sas.AccountResourceTypes{Container: true, Object: true}).String(),
	}.SignWithUserDelegation(udc)
	handleError(err)

	sasURL := fmt.Sprintf("https://%s.blob.core.windows.net/?%s", accountName, sasQueryParams.Encode())

	// This URL can be used to authenticate requests now
	serviceClient, err := service.NewClientWithNoCredential(sasURL, nil)
	handleError(err)

	// You can also break a blob URL up into it's constituent parts
	blobURLParts, _ := blob.ParseURL(serviceClient.URL())
	fmt.Printf("SAS expiry time = %s\n", blobURLParts.SAS.ExpiryTime())

	// Create Managed Identity (OAuth) Credentials using Resource ID
	optsResourceID := azidentity.ManagedIdentityCredentialOptions{ClientOptions: clientOptions, ID: azidentity.ResourceID("/subscriptions/...")}
	cred, err = azidentity.NewManagedIdentityCredential(&optsResourceID)
	handleError(err)

	svcClient, err = service.NewClient("svcURL", cred, &clientOptionsService)
	handleError(err)

	udc, err = svcClient.GetUserDelegationCredential(context.Background(), info, nil)
	handleError(err)
	fmt.Println("User Delegation Key has been created for ", accountName)

	// Create Account Signature Values with desired permissions and sign with user delegation credential
	sasQueryParams, err = sas.AccountSignatureValues{
		Protocol:      sas.ProtocolHTTPS,
		ExpiryTime:    time.Now().UTC().Add(48 * time.Hour),
		Permissions:   to.Ptr(sas.AccountPermissions{Read: true, List: true}).String(),
		Services:      to.Ptr(sas.AccountServices{Blob: true}).String(),
		ResourceTypes: to.Ptr(sas.AccountResourceTypes{Container: true, Object: true}).String(),
	}.SignWithUserDelegation(udc)
	handleError(err)

	sasURL = fmt.Sprintf("https://%s.blob.core.windows.net/?%s", accountName, sasQueryParams.Encode())

	// This URL can be used to authenticate requests now
	serviceClient, err = service.NewClientWithNoCredential(sasURL, nil)
	handleError(err)

	// You can also break a blob URL up into it's constituent parts
	blobURLParts, _ = blob.ParseURL(serviceClient.URL())
	fmt.Printf("SAS expiry time = %s\n", blobURLParts.SAS.ExpiryTime())

}
