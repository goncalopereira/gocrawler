package rules

import (
	data "github.com/goncalocool/coolcoolcool/internal/data"
	storage "github.com/goncalocool/coolcoolcool/internal/storage"
)

//WriteURLRule Check if URL was already seen (might be multiple requests to come)
type WriteURLRule struct {
	CurrentStorage storage.Storage
}

//Follow act on WriteURLRule
func (r WriteURLRule) Follow(req *data.Request, followReq *data.Request) (ok bool) {

	nodeKey := followReq.Key

	return r.CurrentStorage.Save(nodeKey, followReq)
}
