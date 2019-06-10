package rules

import (
	data "github.com/goncalocool/coolcoolcool/internal/data"
)

//RelativeLinkRule fix relative links for web client
type RelativeLinkRule struct {
}

//Follow act on RelativeLinkRule
func (r RelativeLinkRule) Follow(req *data.Request, followReq *data.Request) (ok bool) {

	if followReq.Host == "" {
		followReq.Scheme = req.Scheme
		followReq.Host = req.Host
	}

	return true
}
