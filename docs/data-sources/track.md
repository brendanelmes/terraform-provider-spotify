---
page_title: "spotify_track Data Source - terraform-provider-spotify"
subcategory: ""
description: |-
  
---

# Data Source `spotify_track`

## Example Usage

```terraform
data "spotify_track" "beastly" {
  url = "https://open.spotify.com/track/5fl9iJbPf5ZPenmUJxH2k9"

  ## Computed
  # name = "Beastly"
  # artists = ["7pXu47GoqSYRajmBCjxdD6"]
  # album = "3IOq7Wxdsy7YfT22zMPlmX"
}

data "spotify_track" "new_beastly" {
  spotify_id = "1yQNHzLx4VFofmUNHT3FHa"

  ## Computed
  # name = "New Beastly"
  # artists = ["7pXu47GoqSYRajmBCjxdD6"]
  # album = "2yImgiwCG9KZnxzgulVthl"
}
```

## Schema

### Optional

- **id** (String) The ID of this resource.
- **spotify_id** (String) Spotify ID of the track
- **url** (String) Spotify URL of the track

### Read-only

- **album** (String) The spotify ID of the album
- **artists** (List of String) The spotify IDs of the artists
- **name** (String) The Name of the track


