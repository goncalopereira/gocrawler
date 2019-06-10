package crawler

import (
	"github.com/goncalocool/coolcoolcool/internal/crawler/rules"
	data "github.com/goncalocool/coolcoolcool/internal/data"
)

//Crawler makes decisions on what to follow and how to order
type Crawler struct {
	FilterRequests     chan data.Request
	WebResponses       <-chan data.Response
	ProcessedResponses chan data.ProcessedResponse
	MaxDepth           int
	Commands           []RequestCommand
	FollowRules        []rules.FollowRule
	MaxLinks           int
	Shutdown           chan bool
}

//ProcessResponse filters new links from a response and adds them as next request
func (c *Crawler) ProcessResponse(resp data.Response) {
	reqs := make(chan data.Request, c.MaxLinks)
	defer close(reqs)

	processedResponse := data.ProcessedResponse{Requests: reqs}
	for {
		select {
		case link, ok := <-resp.Links:
			if ok {
				nextRequest, okFollow := rules.ExecuteFollowRules(&resp.Request, link, c.FollowRules...)

				if okFollow {
					processedResponse.Requests <- nextRequest
				}
			} else {
				c.ProcessedResponses <- processedResponse
				return
			}
		default:
			continue
		}
	}
}

//DoUnpackResponse Queues up more requests into web clients, waits for requests to drain
//ALSO caught if requests is full
//Breadth First Search IF single web client/single crawler
//nested select because select with two channels has random priority
//wait for all requests to prevent ordering issues with web client timings
//unpacking one at a time also prevent blocking filterRequests due to explosion in links
func (c *Crawler) DoUnpackResponse() {
	select {
	case resp, ok := <-c.ProcessedResponses:
		if ok {
			for req := range resp.Requests {
				c.FilterRequests <- req
			}

		}
	default:
		return
	}
}

//Do Crawler selects next action based on incoming responses and requests
func (c *Crawler) Do() {
	defer close(c.Shutdown)
	for {

		select {
		case req, ok := <-c.FilterRequests:
			//drain requests first
			if ok {
				ExecuteCommands(&req, c.Commands...)
			}
		case res, ok := <-c.WebResponses:
			//in the meantime process some future links as this will access DB
			if ok {
				c.ProcessResponse(res)
			}
		case <-c.Shutdown:
			return
		default:
			c.DoUnpackResponse()
		}
	}
}
