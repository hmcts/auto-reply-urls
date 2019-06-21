package redirecturi

import (
	"fmt"
	"github.com/Azure/go-autorest/autorest/azure"
	"os"
)

var (
	// these are our *global* config settings, to be shared by all packages.
	// each has corresponding public accessors below.
	// if anything requires a `Set` accessor, that indicates it perhaps
	// shouldn't be set here, because mutable vars shouldn't be global.
	clientID      = os.Getenv("AZURE_CLIENT_ID")
	clientSecret  = os.Getenv("AZURE_CLIENT_SECRET")
	tenantID      = os.Getenv("AZURE_TENANT_ID")
	appId         = os.Getenv("OBJECT_ID")
	cloudName     = "AzurePublicCloud"
	userAgent     string
	environment   *azure.Environment
)

// ClientID is the OAuth client ID.
func ClientID() string {
	if len(clientID) == 0 {
		panic("\"AZURE_CLIENT_ID\" environment variable must be set")
	}
	return clientID
}

func ObjectId() string {
	if len(clientID) == 0 {
		panic("\"OBJECT_ID\" environment variable must be set")
	}
	return appId
}

// ClientSecret is the OAuth client secret.
func ClientSecret() string {
	if len(clientSecret) == 0 {
		panic("\"AZURE_CLIENT_SECRET\" environment variable must be set")
	}

	return clientSecret
}

// TenantID is the AAD tenant to which this client belongs.
func TenantID() string {
	if len(clientSecret) == 0 {
		panic("\"AZURE_TENANT_ID\" environment variable must be set")
	}

	return tenantID
}

// UserAgent() specifies a string to append to the agent identifier.
func UserAgent() string {
	if len(userAgent) > 0 {
		return userAgent
	}
	return "auto-redirect-uri"
}

// Environment() returns an `azure.Environment{...}` for the current cloud.
func Environment() *azure.Environment {
	if environment != nil {
		return environment
	}
	env, err := azure.EnvironmentFromName(cloudName)
	if err != nil {
		// TODO: move to initialization of var
		panic(fmt.Sprintf(
			"invalid cloud name '%s' specified, cannot continue\n", cloudName))
	}
	environment = &env
	return environment
}
