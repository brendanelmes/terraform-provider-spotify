package spotify

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/zmb3/spotify/v2"
	"golang.org/x/oauth2"
)

// ClientConfigurer for spotify API access
func ClientConfigurer(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	apiKey := d.Get("api_key").(string)
	if apiKey == "" {
		return nil, diag.Errorf("api_key must be provided")
	}

	// Create the HTTP client that injects the token for each request
	token := &oauth2.Token{AccessToken: apiKey}
	httpClient := oauth2.NewClient(ctx, oauth2.StaticTokenSource(token))

	client := spotify.New(httpClient)

	// optionally test the client
	user, err := client.CurrentUser(ctx)
	if err != nil {
		return nil, diag.Errorf("unable to authenticate with Spotify using provided api_key: %s", err)
	}
	fmt.Printf("Connected to Spotify user: %s\n", user.ID)

	return client, diags
}
