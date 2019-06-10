package crawler

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/goncalocool/coolcoolcool/internal/crawler/rules"
	data "github.com/goncalocool/coolcoolcool/internal/data"
)

type MockCommandPasses struct {
}

func (c *MockCommandPasses) Command(req *data.Request) (ok bool) {
	return true
}

func TestUnpackResponse(t *testing.T) {
	filter := make(chan data.Request, 10)
	responses := make(chan data.Response, 10)
	processedResponses := make(chan data.ProcessedResponse, 10)
	commands := []RequestCommand{
		new(MockCommandPasses)}

	c := Crawler{
		FilterRequests:     filter,
		WebResponses:       responses,
		ProcessedResponses: processedResponses,
		Commands:           commands,
		FollowRules:        []rules.FollowRule{},
		MaxLinks:           10,
		Shutdown:           make(chan bool, 1)}

	r1 := data.Request{Key: "1"}
	r2 := data.Request{Key: "2"}

	reqs := make(chan data.Request, 10)
	pr := data.ProcessedResponse{Requests: reqs}
	reqs <- r1
	reqs <- r2
	close(reqs)
	processedResponses <- pr

	c.DoUnpackResponse()

	assert.Equal(t, 0, len(processedResponses))
	assert.Equal(t, 0, len(responses))
	assert.Equal(t, 2, len(filter))
}

func TestProcessResponse(t *testing.T) {
	filter := make(chan data.Request, 10)
	responses := make(chan data.Response, 10)
	processedResponses := make(chan data.ProcessedResponse, 10)
	commands := []RequestCommand{
		new(MockCommandPasses)}

	c := Crawler{
		FilterRequests:     filter,
		WebResponses:       responses,
		ProcessedResponses: processedResponses,
		Commands:           commands,
		FollowRules:        []rules.FollowRule{},
		MaxLinks:           10,
		Shutdown:           make(chan bool, 1)}

	links := make(chan string, 10)
	links <- "/links1"
	links <- "/links2"
	close(links)
	resp := data.Response{Request: data.Request{Key: "4"}, Links: links}
	c.ProcessResponse(resp)

	assert.Equal(t, 1, len(processedResponses))

	pr := <-processedResponses

	assert.Equal(t, 2, len(pr.Requests))
	req := <-pr.Requests
	assert.Equal(t, req.Path, "/links1")
	req2 := <-pr.Requests
	assert.Equal(t, req2.Path, "/links2")

	assert.Equal(t, 0, len(responses))
	assert.Equal(t, 0, len(filter))
}

func TestCrawlerShutdown(t *testing.T) {
	filter := make(chan data.Request, 10)
	responses := make(chan data.Response, 10)
	processedResponses := make(chan data.ProcessedResponse, 10)
	commands := []RequestCommand{
		new(MockCommandPasses)}

	shutdown := make(chan bool, 1)
	c := Crawler{
		FilterRequests:     filter,
		WebResponses:       responses,
		ProcessedResponses: processedResponses,
		Commands:           commands,
		FollowRules:        []rules.FollowRule{},
		MaxLinks:           10,
		Shutdown:           shutdown}

	links := make(chan string, 10)
	links <- "/links1"
	links <- "/links2"
	close(links)
	responses <- data.Response{Request: data.Request{Key: "4"}, Links: links}
	go func() {
		time.Sleep(2 * time.Second)
		shutdown <- true
	}()
	c.Do()

	assert.Equal(t, 0, len(processedResponses))
	assert.Equal(t, 0, len(responses))
	assert.Equal(t, 0, len(filter))
}
