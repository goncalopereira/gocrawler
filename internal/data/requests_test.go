package data

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeRequestFromPOSTBodyURLonly(t *testing.T) {
	values := url.Values{}
	values["url"] = []string{"http://www.monzo.com"}
	req, ok := MakeRequestFromURL(values.Get("url"))

	assert.Equal(t, "www.monzo.com", req.Host)
	assert.Equal(t, true, ok)
	assert.Equal(t, "http://www.monzo.com", req.URL())
}

func TestMakeRequestFromPOSTBodyNoURL(t *testing.T) {
	values := url.Values{}
	req, ok := MakeRequestFromURL(values.Get("url"))

	assert.Equal(t, req, Request{})
	assert.Equal(t, false, ok)
}

func TestMakeRequestFromPOSTBodyNonURL(t *testing.T) {
	values := url.Values{}
	values["url"] = []string{"httxxzo.com"}
	req, ok := MakeRequestFromURL(values.Get("url"))

	assert.Equal(t, req, Request{})
	assert.Equal(t, false, ok)
}
