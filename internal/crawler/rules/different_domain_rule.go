package rules

import (
	data "github.com/goncalopereira/gocrawler/internal/data"
)

//DifferentDomainRule check if its domain of original
type DifferentDomainRule struct {
}

//Follow act on DifferentDomainRule
func (r DifferentDomainRule) Follow(req *data.Request, followReq *data.Request) (ok bool) {

	//WWW to Naked
	if ("www." + req.Host) == followReq.Host {
		return true
	}

	//Naked to WWW
	if (req.Host) == "www."+followReq.Host {
		return true
	}

	if req.Host != followReq.Host {
		return false
	}

	return true
}
