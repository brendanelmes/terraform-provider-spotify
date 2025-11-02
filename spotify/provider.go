package spotify

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider for spotify
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("SPOTIFY_API_KEY", nil),
				Description: "Spotify API Key",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"spotify_playlist": resourcePlaylist(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"spotify_search_track": dataSourceSearchTrack(),
			"spotify_track":        dataSourceTrack(),
		},
		ConfigureContextFunc: ClientConfigurer,
	}
}
