package storage

import (
	"testing"

	data "github.com/goncalopereira/gocrawler/internal/data"
	"github.com/goncalopereira/gocrawler/internal/env"
	"github.com/stretchr/testify/assert"
)

func TestBasicInMemoryDBSaveTwoEntries(t *testing.T) {
	db := BasicInMemoryDB{DB: make(map[string]data.Request, env.MaxURLSStored())}
	req, _ := data.MakeRequestFromURL("https://www.goncalopereira.com/")

	ok := db.Save("key1", &req)
	assert.Equal(t, true, ok)

	nextOk := db.Save("key2", &req)
	assert.Equal(t, true, nextOk)
}

func TestBasicInMemoryDBNoSaveSecondEntry(t *testing.T) {
	db := BasicInMemoryDB{DB: make(map[string]data.Request, env.MaxURLSStored())}
	req, _ := data.MakeRequestFromURL("https://www.goncalopereira.com")

	ok := db.Save("key1", &req)
	assert.Equal(t, true, ok)

	nextOk := db.Save("key1", &req)
	assert.Equal(t, false, nextOk)
}
