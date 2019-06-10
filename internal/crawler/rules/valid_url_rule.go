package rules

import (
	"strings"

	data "github.com/goncalocool/coolcoolcool/internal/data"
)

//ValidURLRule check if URL is valid
type ValidURLRule struct {
}

//Follow act on ValidURLRule
func (r ValidURLRule) Follow(req *data.Request, followReq *data.Request) (ok bool) {
	if followReq.Scheme != "http" && followReq.Scheme != "https" {
		return false
	}

	badPaths := []string{".pdf", ".png", " ", ".."}
	for _, bp := range badPaths {
		if strings.Contains(followReq.Path, bp) {
			return false
		}
	}

	return true
}
