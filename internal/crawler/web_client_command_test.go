package crawler

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/goncalopereira/gocrawler/internal/data"
)

func TestSendCopyRequestToWebClient(t *testing.T) {

	request := data.Request{}
	web := make(chan data.Request, 10)

	c := WebClientCommand{Requests: web}

	c.Command(&request)

	assert.Equal(t, len(c.Requests), 1)

	copyReq := <-c.Requests
	copyReq.Key = "1"

	assert.NotEqual(t, copyReq.Key, request.Key)
}
