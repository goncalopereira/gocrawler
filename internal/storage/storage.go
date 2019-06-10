package storage

import (
	"github.com/goncalopereira/gocrawler/internal/data"
)

//THE DB!!!!

//Storage interface for database, in memory, etc etc
//Sync
//Save decision from implementation
type Storage interface {
	Save(key string, req *data.Request) (ok bool)
}

//BasicInMemoryDB just records the information
//decision: map is not goroutine safe
//We never read on webclients
//We have a single goroutine handling the saves
//More goroutines, sync.Map? Shard? Log and Clean Duplicates?
//we care about last write wins, is write safe?
type BasicInMemoryDB struct {
	DB map[string]data.Request
}

//Save returns if it was able to save or though it was an existing entry
//Sync
func (s BasicInMemoryDB) Save(key string, req *data.Request) (ok bool) {

	if _, ok := s.DB[key]; ok {
		return false
	}

	copyReq := *req
	s.DB[key] = copyReq

	return true
}
