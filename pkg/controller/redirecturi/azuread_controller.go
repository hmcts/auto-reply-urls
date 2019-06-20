package redirecturi

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/go-autorest/autorest"
	"os"
)

func getApplicationsClient() graphrbac.ApplicationsClient {
	appClient := graphrbac.NewApplicationsClient(TenantID())
	a, _ := GetGraphAuthorizer()
	appClient.Authorizer = a
	appClient.AddToUserAgent(UserAgent())

	return appClient
}

func GetADApplication(ctx context.Context, applicationObjectID string) (graphrbac.Application, error) {
	appClient := getApplicationsClient()
	return appClient.Get(ctx, applicationObjectID)
}

func UpdateADApplication(ctx context.Context, applicationObjectID string, replyUrls []string) (autorest.Response, error) {
	appClient := getApplicationsClient()

	fmt.Printf("Modifying %s to have the following redirect uris: %s\n", applicationObjectID, replyUrls)
	return appClient.Patch(ctx,
		applicationObjectID,
		graphrbac.ApplicationUpdateParameters{
			ReplyUrls: &replyUrls,
		})
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func AddRedirectUri(uri string) {
	ctx := context.Background()

	application, err := GetADApplication(ctx, ApplicationID())
	fmt.Println("Retrieved application")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	replyUrls := *application.ReplyUrls

	if contains(replyUrls, uri) == false {
		replyUrls = append(replyUrls, uri)
		fmt.Println("about to update app")

		_, err = UpdateADApplication(ctx, ApplicationID(), replyUrls)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(2)
		}
		fmt.Println("ad app updated")
	} else {
		fmt.Printf("Skipping appending url %s, as it already exists\n", uri)
	}
}
