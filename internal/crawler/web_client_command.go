package crawler

import (
	data "github.com/goncalopereira/gocrawler/internal/data"
)

//WebClientCommand request filter that pushes valid requests to WebClient
type WebClientCommand struct {
	Requests chan data.Request
}

//Command act on WebClientCommand
func (wcc WebClientCommand) Command(req *data.Request) (ok bool) {
	copyReq := *req
	wcc.Requests <- copyReq
	return true
}
