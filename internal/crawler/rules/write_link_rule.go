package rules

import (
	data "github.com/goncalocool/coolcoolcool/internal/data"
	storage "github.com/goncalocool/coolcoolcool/internal/storage"
)

//WriteLinkRule check if Link is valid
type WriteLinkRule struct {
	CurrentStorage storage.Storage
}

//Follow act on WriteLinkRule
func (r WriteLinkRule) Follow(req *data.Request, followReq *data.Request) (ok bool) {

	//decision: write any links before sending to web client even if broken to prevent retries
	//split sync wait of web client with the response processing+DB
	//possible to check this on RequestFilter

	//distributed IDs are HARD, we generate a Key and use Referer as ParentKey for an unique link
	//save as simply as possible
	//prevent maps in maps

	nodeKey := req.Key + "->" + followReq.Key

	return r.CurrentStorage.Save(nodeKey, followReq)
}
