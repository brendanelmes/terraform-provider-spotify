# terraform-provider-spotify

This is a terraform provider for managing spotify playlists.

> I am not affiliated with Hashicorp or Terraform.

## About

Although it is a fork of [conradludgate/terraform-provider-spotify](https://github.com/conradludgate/terraform-provider-spotify),
there are some distinct differences:

1. This provider assumes that there is no auth server, accepting only an API Key
2. There are no album or library resources, as this is a stripped back version initially for use in a teaching session

> There is a known issue with the `public` option that should be fixed following a migration to the `terraform-plugin-framework`

## Example

```tf
resource "spotify_playlist" "playlist" {
  name        = "Super Hot Fire Playlist"
  description = "My playlist is super hot fire"
  public      = false

  tracks = flatten([
    data.spotify_track.overkill.id,
    data.spotify_track.blackwater.id,
    data.spotify_track.overkill.id,
    data.spotify_search_track.search.tracks[*].id,
  ])
}

data "spotify_track" "overkill" {
  url = "https://open.spotify.com/track/4XdaaDFE881SlIaz31pTAG"
}
data "spotify_track" "blackwater" {
  spotify_id = "4lE6N1E0L8CssgKEUCgdbA"
}

data "spotify_search_track" "search" {
  name   = "Somebody Told Me"
  artist = "The Killers"
  album  = "Hot Fuss"
}

output "test" {
  value = data.spotify_search_track.search.tracks
}
```

## Installation

Add the following to your terraform configuration

```tf
terraform {
  required_providers {
    spotify = {
      source  = "brendanelmes/spotify"
      version = "~> 0.1.0"
    }
  }
}
```

## How to use

### Provider configuration

Configure the terraform provider like so:

```tf
provider "spotify" {
  api_key     = var.spotify_api_key
}

variable "spotify_api_key" {
  type = string
}
```
