package http

import (
	"io"

	data "github.com/goncalocool/coolcoolcool/internal/data"
	"golang.org/x/net/html"
)

//BodyParser interface for new ways to parse body
type BodyParser interface {
	Parse(body io.Reader, request data.Request) data.Response
}

//BasicBodyParser basic implementation for a href
type BasicBodyParser struct {
	MaxLinks int
}

//Parse handling links from body reader out to Response
// moving ioReader to chan would create issues with pointer
// links are A tag, Link Tag and Nav tag but we check A
// no JS parsing
func (r BasicBodyParser) Parse(body io.Reader, request data.Request) data.Response {
	links := make(chan string, r.MaxLinks)
	defer close(links)

	response := data.Response{Request: request, Links: links}
	z := html.NewTokenizer(body)
	for {
		tt := z.Next()
		switch {
		case r.MaxLinks == len(response.Links):
			fallthrough //links number, prevent deadlock
		case tt == html.ErrorToken:
			//End of page
			return response
		case tt == html.StartTagToken:
			t := z.Token()
			if t.Data == "a" {
				for _, attr := range t.Attr {
					if attr.Key == "href" {
						response.Links <- attr.Val
					}
				}
			}
		}
	}
}
