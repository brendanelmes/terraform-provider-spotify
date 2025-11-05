---
page_title: "spotify_playlist Resource - terraform-provider-spotify"
subcategory: ""
description: |-
  Resource to manage a spotify playlist.
---

# Resource `spotify_playlist`

Resource to manage a spotify playlist.

## Example Usage

```terraform
resource "spotify_playlist" "playlist" {
  name        = "My playlist"
  description = "My playlist is so awesome"
  public      = false

  tracks = [
    data.spotify_track.never_gonna_give_you_up.id,
    data.spotify_track.careless_whisper.id,
    "5jr6pG3khBcBXZRm8NogSe",
  ]
}

data "spotify_track" "never_gonna_give_you_up" {
  url = "https://open.spotify.com/track/4PTG3Z6ehGkBFwjybzWkR8"
}
data "spotify_track" "careless_whisper" {
  spotify_id = "5WDLRQ3VCdVrKw0njWe5E5?si=44a369c142b144ad"
}
```

## Schema

### Required

- **name** (String) The name of the resulting playlist
- **tracks** (List of String) A set of tracks for the playlist to contain

### Optional

- **description** (String) The description of the resulting playlist
- **id** (String) The ID of this resource.
- **public** (Boolean) Whether the playlist can be accessed publicly

### Read-only

- **snapshot_id** (String)

## Import

Import is supported using the following syntax:

```shell
# Using the playlist ID
# https://open.spotify.com/playlist/37i9dQZF1DWVs8I62NcHks (a playlist share link)
#                                   ^^^^^^^^^^^^^^^^^^^^^^
terraform import spotify_playlist.example 37i9dQZF1DWVs8I62NcHks
```
