package sitemap

import (
	"testing"

	"github.com/goncalocool/coolcoolcool/internal/data"
	"github.com/goncalocool/coolcoolcool/internal/storage"
)

func TestSiteMapOutput(t *testing.T) {
	r := data.Request{}
	linksDB := make(map[string]data.Request)
	linksDB["n1->n2"] = r
	linksDB["n3->n2"] = r
	linksStorage := storage.BasicInMemoryDB{DB: linksDB}
	urlDB := make(map[string]data.Request)
	urlDB["n1"] = r
	urlDB["n2"] = r
	urlDB["n3"] = r
	urlStorage := storage.BasicInMemoryDB{DB: urlDB}

	//Just checking if it doesnt blow up
	BasicSiteMap{LinksDB: linksStorage, URLDB: urlStorage}.Output()
}
