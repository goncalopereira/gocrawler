package rules

import (
	"net/url"

	data "github.com/goncalocool/coolcoolcool/internal/data"
)

//FollowRule modifies next request and returns if it should be followed
type FollowRule interface {
	Follow(req *data.Request, followReq *data.Request) (ok bool)
}

//ExecuteFollowRules checks all FollowRule for new link
func ExecuteFollowRules(req *data.Request, link string, rules ...FollowRule) (followRequest data.Request, valid bool) {
	nextURL, err := url.Parse(link)

	if err != nil {
		return data.Request{}, false
	}

	nextReq := data.Request{
		Path:   nextURL.Path,
		Host:   nextURL.Host,
		Scheme: nextURL.Scheme,
		Depth:  req.Depth + 1}

	for _, rule := range rules {
		ok := rule.Follow(req, &nextReq)

		if !ok {
			return nextReq, false
		}

	}

	return nextReq, true
}
