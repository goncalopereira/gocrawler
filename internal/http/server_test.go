package http

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	data "github.com/goncalopereira/gocrawler/internal/data"
)

func TestPostToServer(t *testing.T) {
	filter := make(chan data.Request, 10)

	h := NewRequestHandler{FilterRequests: filter}

	data := url.Values{}
	data.Set("url", "http://www.goncalopereira.com")

	httpReq, _ := http.NewRequest("POST", "/", strings.NewReader(data.Encode()))
	httpReq.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") // This makes it work
	rr := httptest.NewRecorder()
	http.HandlerFunc(h.Post).ServeHTTP(rr, httpReq)

	assert.Equal(t, 1, len(h.FilterRequests))
	req := <-filter
	assert.Equal(t, "http://www.goncalopereira.com", req.URL())
}
