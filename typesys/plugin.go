package typesys

import "net/url"

type Plugin struct {
	Name   string
	Source url.URL
}
