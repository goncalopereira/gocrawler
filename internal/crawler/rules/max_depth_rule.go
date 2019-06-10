package rules

import (
	data "github.com/goncalocool/coolcoolcool/internal/data"
)

//MaxDepthRule rule about max depth usable
type MaxDepthRule struct {
	MaxDepth int
}

//Follow rule
func (r MaxDepthRule) Follow(req *data.Request, followReq *data.Request) (ok bool) {
	if followReq.Depth > r.MaxDepth {
		return false
	}

	return true
}
