package azureclient

import (
	"context"
	"log"

	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
)

// CreateClient represents logic for initialization of the Azure Devops Client
func CreateClient(organizationURL string, personalAccessToken string) (core.Client, context.Context) {

	// Create a connection to your organization
	connection := azuredevops.NewPatConnection(organizationURL, personalAccessToken)

	ctx := context.Background()

	// Create a client to interact with the Core area
	coreClient, err := core.NewClient(ctx, connection)
	if err != nil {
		log.Fatal(err)
	}

	return coreClient, ctx
}
