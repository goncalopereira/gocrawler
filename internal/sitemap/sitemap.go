package sitemap

import (
	"github.com/goncalopereira/gocrawler/internal/storage"
)

//Sitemap is the interface for outputing SiteMaps
type Sitemap interface {
	Output()
}

//BasicSiteMap is a SiteMap that works with BasicInMemoryDB
type BasicSiteMap struct {
	LinksDB storage.BasicInMemoryDB
	URLDB   storage.BasicInMemoryDB
}

//Output converts a DB into output (strings)
func (sm BasicSiteMap) Output() {
	weighted := BuildNodes(sm.LinksDB.DB, sm.URLDB.DB)

	sortedWeight := SortWeightedList(weighted)
	Display(sortedWeight)

	sortedBFS := BFS(sortedWeight[0])
	DisplayBFS(sortedBFS)
}
