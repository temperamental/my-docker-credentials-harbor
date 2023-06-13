package main

// For reference
// * https://github.com/awslabs/amazon-ecr-credential-helper/blob/main/ecr-login/ecr.go
// * https://github.com/Azure/acr-docker-credential-helper/

import (
	"fmt"
	"os"

	"github.com/moby/moby/api/types/credentials"
	"github.com/moby/moby/client"
)

func main() {
	// Get the Docker client.
	cli, err := client.NewEnvClient()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Get the credentials from the environment.
	username := os.Getenv("HARBOR_CLIENT_NAME")
	password := os.Getenv("HARBOR_CLIENT_SECRET")

	// TODO: This seems specific to authenticating to hub.docker.io
	// Need to update to be Harbor specific

	// Create the credentials object.
	credentials := &credentials.Credentials{
		ServerAddress: "https://index.docker.io/v1/",
		Username:      username,
		Password:      password,
	}

	// Store the credentials in the Docker daemon.
	err = cli.StoreCredentials(credentials)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Print a success message.
	fmt.Println("Credentials stored successfully!")
}
