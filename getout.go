/*************************************
*
* Access outlook 365 email using the
* microsoft graph api
*
* Getting going with sample code from
* https://github.com/yaegashi/msgraph.go/blob/master/cmd/msgraph-me/main.go
*
*************************************/
package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/yaegashi/msgraph.go/msauth"
	msgraph "github.com/yaegashi/msgraph.go/v1.0"
	"golang.org/x/oauth2"
	//msgraph "github.com/yaegashi/msgraph.go/v1.0"
)

const (
	defaultTentantID      = "common"
	defaultClientID       = "45c7f99c-0a94-42ff-a6d8-a8d657229e8c"
	defaultTokenCachePath = "token_cache.json"
)

var defaultScopes = []string{"offline_access", "User.Read", "Calendars.Read", "Files.Read", "Mail.ReadWrite"}

func main() {
	//fmt.Println("Go get outlook mail.")
	fmt.Println("Go get User info first.")
	var tenantID, clientID, tokenCachePath string
	flag.StringVar(&tenantID, "tenant-id", defaultTentantID, "Tenant ID")
	flag.StringVar(&clientID, "client-id", defaultClientID, "Client ID")
	flag.StringVar(&tokenCachePath, "token-cache-path", defaultTokenCachePath, "Token cache path")
	flag.Parse()

	ctx := context.Background()
	ms := msauth.NewManager()
	ms.LoadFile(tokenCachePath)
	ts, err := ms.DeviceAuthorizationGrant(ctx, tenantID, clientID, defaultScopes, nil)
	if err != nil {
		log.Fatal(err)
	}
	ms.SaveFile(tokenCachePath)
	httpClient := oauth2.NewClient(ctx, ts)
	graphClient := msgraph.NewClient(httpClient)
	req := graphClient.Me().Request()
	log.Printf("GET %s", req.URL())
	req = (*msgraph.UserRequest)(graphClient.Me().Outlook().Request())
	log.Printf("GET %s", req.URL())
	req.Top(10)
}
