package rules

import (
	"strings"

	data "github.com/goncalocool/coolcoolcool/internal/data"
)

//SubdomainRule check if its subdomain of original
type SubdomainRule struct {
}

//Follow act on SubdomainRule
func (r SubdomainRule) Follow(req *data.Request, followReq *data.Request) (ok bool) {
	originalDomain := strings.Split(req.Host, ".")
	nextDomain := strings.Split(followReq.Host, ".")

	originalBaseDomain := originalDomain[len(originalDomain)-2] + originalDomain[len(originalDomain)-1]
	nextBaseDomain := nextDomain[len(nextDomain)-2] + nextDomain[len(nextDomain)-1]
	if originalBaseDomain != nextBaseDomain {
		return false
	}

	return true
}
