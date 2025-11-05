---
page_title: "spotify Provider"
subcategory: ""
description: |-

---

# spotify Provider

Spotify is a digital music, podcast, and video service that gives you access to millions of songs and other content from creators all over the world.

-> This provider does not include an auth server like `conradludgate/spotify`. It requires you to generate and provide your own api key.

## Example Usage

```terraform
provider "spotify" {
  api_key = var.spotify_api_key
}

variable "spotify_api_key" {
  type = string
}
```

## Schema

### Required

- **api_key** (String)
