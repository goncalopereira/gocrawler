package rules

import (
	"testing"

	"github.com/goncalopereira/gocrawler/internal/data"
	"github.com/stretchr/testify/assert"
)

func ValidURL() data.Request {
	req, _ := data.MakeRequestFromURL("https://www.goncalopereira.com")
	return req
}

func TestValidURLBadScheme(t *testing.T) {
	req := ValidURL()
	req.Depth = 10

	rules := []FollowRule{ValidURLRule{}}
	_, valid := ExecuteFollowRules(&req, "monzo://card", rules...)

	assert.Equal(t, false, valid)
}
func TestValidURL(t *testing.T) {
	req := ValidURL()
	req.Depth = 10

	rules := []FollowRule{ValidURLRule{}}
	_, valid := ExecuteFollowRules(&req, "https://www.goncalopereira.com", rules...)

	assert.Equal(t, true, valid)
}

func TestValidURLBadPath(t *testing.T) {
	req := ValidURL()
	req.Depth = 10

	rules := []FollowRule{ValidURLRule{}}
	_, valid := ExecuteFollowRules(&req, "https://www.goncalopereira.com/x.pdf", rules...)

	assert.Equal(t, false, valid)
}

func TestValidURLNoScheme(t *testing.T) {
	req := ValidURL()
	req.Depth = 10

	rules := []FollowRule{ValidURLRule{}}
	_, valid := ExecuteFollowRules(&req, ":///cdn-cgi/l/email-protection", rules...)

	assert.Equal(t, false, valid)
}

func TestFollowInvalidDepthLink(t *testing.T) {
	req := ValidURL()
	req.Depth = 10

	rules := []FollowRule{MaxDepthRule{MaxDepth: 1}}
	_, valid := ExecuteFollowRules(&req, "/", rules...)

	assert.Equal(t, false, valid)
}

func TestFollowValidDepthLink(t *testing.T) {
	req := ValidURL()
	req.Depth = 1

	rules := []FollowRule{MaxDepthRule{MaxDepth: 2}}
	_, valid := ExecuteFollowRules(&req, "/", rules...)

	assert.Equal(t, true, valid)
}

func TestFollowDifferentDomainSubdomainRuleLink(t *testing.T) {
	req := ValidURL()

	rules := []FollowRule{SubdomainRule{}}
	_, valid := ExecuteFollowRules(&req, "http://www.blah.com", rules...)

	assert.Equal(t, false, valid)
}

func TestFollowDifferentDomainNakedToWWW(t *testing.T) {
	req, _ := data.MakeRequestFromURL("https://goncalopereira.com")

	rules := []FollowRule{DifferentDomainRule{}}
	_, valid := ExecuteFollowRules(&req, "https://www.goncalopereira.com", rules...)

	assert.Equal(t, true, valid)
}

func TestFollowDifferentDomainWWWToNaked(t *testing.T) {
	req, _ := data.MakeRequestFromURL("https://www.goncalopereira.com")

	rules := []FollowRule{DifferentDomainRule{}}
	_, valid := ExecuteFollowRules(&req, "https://goncalopereira.com", rules...)

	assert.Equal(t, true, valid)
}

func TestFollowDifferentDomainSameDomain(t *testing.T) {
	req, _ := data.MakeRequestFromURL("https://www.goncalopereira.com")

	rules := []FollowRule{DifferentDomainRule{}}
	_, valid := ExecuteFollowRules(&req, "https://www.goncalopereira.com", rules...)

	assert.Equal(t, true, valid)
}

func TestFollowSubdomainLink(t *testing.T) {
	req := ValidURL()

	rules := []FollowRule{SubdomainRule{}}
	nextReq, valid := ExecuteFollowRules(&req, "https://subdomain.goncalopereira.com", rules...)

	assert.Equal(t, true, valid)
	assert.Equal(t, "subdomain.goncalopereira.com", nextReq.Host)
}

func TestFollowNoSubdomainLink(t *testing.T) {
	req := ValidURL()

	rules := []FollowRule{SubdomainRule{}}
	nextReq, valid := ExecuteFollowRules(&req, "https://goncalopereira.com", rules...)

	assert.Equal(t, true, valid)
	assert.Equal(t, "goncalopereira.com", nextReq.Host)
}

func TestFollowDifferentDomainLink(t *testing.T) {
	req := ValidURL()

	rules := []FollowRule{DifferentDomainRule{}}
	_, valid := ExecuteFollowRules(&req, "https://subdomain.goncalopereira.com", rules...)

	assert.Equal(t, false, valid)
}

func TestFollowRelativeLink(t *testing.T) {
	req := ValidURL()

	rules := []FollowRule{RelativeLinkRule{}}
	nextReq, valid := ExecuteFollowRules(&req, "/help", rules...)

	assert.Equal(t, true, valid)
	assert.Equal(t, "www.goncalopereira.com", nextReq.Host)
}

func TestFollowFragmentLink(t *testing.T) {
	req := ValidURL()

	rules := []FollowRule{RelativeLinkRule{}}
	nextReq, valid := ExecuteFollowRules(&req, "/cdn-cgi/l/email-protection#4c2429203c0c2123223623622f2321", rules...)

	assert.Equal(t, true, valid)
	assert.Equal(t, nextReq.Path, "/cdn-cgi/l/email-protection")
}

func TestFollowQueryStringLink(t *testing.T) {
	req := ValidURL()

	rules := []FollowRule{RelativeLinkRule{}}
	nextReq, valid := ExecuteFollowRules(&req, "/help?query=x", rules...)

	assert.Equal(t, true, valid)
	assert.Equal(t, "https://www.goncalopereira.com/help", nextReq.URL())

}

type MockSavePasses struct {
}

func (c MockSavePasses) Save(key string, req *data.Request) (ok bool) {
	return true
}

type MockSaveFails struct {
}

func (c MockSaveFails) Save(key string, req *data.Request) (ok bool) {
	return false
}

func TestWriteLinkRule(t *testing.T) {
	req := ValidURL()

	rules := []FollowRule{WriteLinkRule{CurrentStorage: MockSavePasses{}}}
	_, valid := ExecuteFollowRules(&req, "https://www.goncalopereira.com/help", rules...)

	assert.Equal(t, true, valid)
}

func TestWriteKeyRuleTrailingSlash(t *testing.T) {
	req := ValidURL()

	rules := []FollowRule{WriteKeyRule{}}
	followReq, valid := ExecuteFollowRules(&req, "https://www.goncalopereira.com/help/", rules...)

	assert.Equal(t, true, valid)
	assert.Equal(t, "www.goncalopereira.com/help", followReq.Key)
}

func TestWriteKeyRuleNakedDomain(t *testing.T) {
	req := ValidURL()

	rules := []FollowRule{WriteKeyRule{}}
	followReq, valid := ExecuteFollowRules(&req, "https://goncalopereira.com/help/", rules...)

	assert.Equal(t, true, valid)
	assert.Equal(t, "www.goncalopereira.com/help", followReq.Key)
}

func TestWriteKeyRuleStopSelfLink(t *testing.T) {
	req := ValidURL()
	req.Key = "www.goncalopereira.com"

	rules := []FollowRule{WriteKeyRule{}}
	_, valid := ExecuteFollowRules(&req, "https://www.goncalopereira.com", rules...)

	assert.Equal(t, false, valid)
}

func TestExistingURLWriteLinkRule(t *testing.T) {
	req := ValidURL()
	req.Depth = 10

	rules := []FollowRule{WriteLinkRule{CurrentStorage: MockSaveFails{}}}
	_, valid := ExecuteFollowRules(&req, "https://www.goncalopereira.com/help", rules...)

	assert.Equal(t, false, valid)
}

func TestWriteURLRule(t *testing.T) {
	req := ValidURL()
	req.Depth = 10

	rules := []FollowRule{WriteURLRule{CurrentStorage: MockSavePasses{}}}
	followReq, valid := ExecuteFollowRules(&req, "https://www.goncalopereira.com/help", rules...)
	followReq.Key = "x"
	assert.Equal(t, true, valid)
	assert.Equal(t, "x", followReq.Key)
}

func TestWriteURLRuleFails(t *testing.T) {
	req := ValidURL()
	req.Depth = 10

	rules := []FollowRule{WriteURLRule{CurrentStorage: MockSaveFails{}}}
	followReq, valid := ExecuteFollowRules(&req, "https://www.goncalopereira.com/help", rules...)
	followReq.Key = "x"
	assert.Equal(t, false, valid)
	assert.Equal(t, "x", followReq.Key)
}
