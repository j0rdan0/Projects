package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
)

func getClient(vaultURI string) (*azsecrets.Client, error) {
	credential, err := appAuthenticate()
	if err != nil {
		log.Panic(err)
	}
	client, err := azsecrets.NewClient(vaultURI, credential, nil)
	if err != nil {

		return nil, err
	} else {

		return client, nil
	}

}

func appAuthenticate() (*azidentity.ClientSecretCredential, error) {
	credential, err := azidentity.NewClientSecretCredential(config.TenantId, config.ClientId, config.Secret, nil)
	if err != nil {
		return nil, err
	} else {
		return credential, nil
	}
}

func getSecret(secretName string) {
	client, err := getClient(config.VaultURI)
	if err != nil {
		log.Panic(err.Error())
	}
	response, err := client.GetSecret(context.Background(), secretName, nil)

	if err != nil {
		log.Panic(err)
	} else {
		fmt.Println(*(response.Value))
	}
}
