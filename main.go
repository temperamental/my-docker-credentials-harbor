package main

import (
    "fmt"
    "github.com/docker/docker/api/types/credentials"
    "github.com/docker/docker/client"
    "os"
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

    // Create the credentials object.
    credentials := &credentials.Credentials{
        ServerAddress: "https://index.docker.io/v1/",
        Username:     username,
        Password:     password,
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

