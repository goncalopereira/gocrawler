package rules

import (
	"strings"

	data "github.com/goncalopereira/gocrawler/internal/data"
)

//WriteKeyRule Pick Key and validate
type WriteKeyRule struct {
}

//Follow act on WriteKeyRule
func (r WriteKeyRule) Follow(req *data.Request, followReq *data.Request) (ok bool) {

	MakeKey(followReq)

	if req.Key == "" { //original requests
		MakeKey(req)
	}

	if followReq.Key == req.Key { //self
		return false
	}

	return true
}

//MakeKey how to make new unique keys
func MakeKey(req *data.Request) {
	sHost := strings.Split(req.Host, ".")
	var h = req.Host
	if len(sHost) == 2 {
		//naked domain attach www.
		h = "www." + h
	}

	req.Key = h + strings.TrimRight(req.Path, "/") //deduplicate trailing '/
}
