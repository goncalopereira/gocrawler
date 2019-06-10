package sitemap

import (
	"testing"

	"github.com/goncalocool/coolcoolcool/internal/data"
	"github.com/stretchr/testify/assert"
)

func TestNodeValue(t *testing.T) {
	n := Node{Key: "new", LinksIn: []*Node{}, LinksOut: []*Node{}}
	n2 := Node{Key: "new2", LinksIn: []*Node{}, LinksOut: []*Node{}}

	assert.Equal(t, 0, n.Value())

	n.LinksIn = append(n.LinksIn, &n2)

	assert.Equal(t, 1, n.Value())
}

func TestNewNode(t *testing.T) {
	m := make(map[string]*Node)

	GetNode(m, "new")

	assert.Equal(t, 1, len(m))
	assert.Equal(t, Node{Key: "new"}.Key, m["new"].Key)
}
func TestExistingNode(t *testing.T) {

	m := make(map[string]*Node)

	GetNode(m, "new")

	assert.Equal(t, 1, len(m))
	assert.Equal(t, Node{Key: "new"}.Key, m["new"].Key)

	GetNode(m, "new")
	assert.Equal(t, 1, len(m))
	assert.Equal(t, Node{Key: "new"}.Key, m["new"].Key)
}

func TestBuildNodes(t *testing.T) {
	r := data.Request{}
	linksDB := make(map[string]data.Request)
	linksDB["n1->n2"] = r
	linksDB["n3->n2"] = r
	urlDB := make(map[string]data.Request)
	urlDB["n1"] = r
	urlDB["n2"] = r
	urlDB["n3"] = r
	weighed := BuildNodes(linksDB, urlDB)

	assert.Equal(t, 3, len(weighed))
	assert.Equal(t, 1, len(weighed["n1"].LinksOut))
	assert.Equal(t, 0, len(weighed["n1"].LinksIn))

	assert.Equal(t, 0, len(weighed["n2"].LinksOut))
	assert.Equal(t, 2, len(weighed["n2"].LinksIn))
}

func TestSortNodesWeight(t *testing.T) {
	r := data.Request{}
	linksDB := make(map[string]data.Request)
	linksDB["n1->n2"] = r
	linksDB["n3->n2"] = r
	urlDB := make(map[string]data.Request)
	urlDB["n1"] = r
	urlDB["n2"] = r
	urlDB["n3"] = r
	weighed := BuildNodes(linksDB, urlDB)

	sorted := SortWeightedList(weighed)

	assert.Equal(t, "n2", sorted[0].Key)
	assert.Equal(t, 2, sorted[0].Value())
	assert.Equal(t, 3, len(sorted))
}

func TestSortNodesBFS(t *testing.T) {
	r := data.Request{}
	linksDB := make(map[string]data.Request)
	linksDB["n1->n2"] = r
	linksDB["n3->n2"] = r
	linksDB["n1->n4"] = r
	linksDB["n2->n5"] = r
	urlDB := make(map[string]data.Request)
	urlDB["n1"] = r
	urlDB["n2"] = r
	urlDB["n3"] = r
	urlDB["n4"] = r
	urlDB["n5"] = r
	weighed := BuildNodes(linksDB, urlDB)

	//n1,n2,n4,n5 n3 not visible
	sorted := BFS(weighed["n1"])

	assert.Equal(t, "n1", sorted[0].Key)

	assert.Equal(t, 4, len(sorted))

	//randomly sorted
	if sorted[1].Key == "n2" {
		assert.Equal(t, "n2", sorted[1].Key)
		assert.Equal(t, "n4", sorted[2].Key)
	} else {
		assert.Equal(t, "n4", sorted[1].Key)
		assert.Equal(t, "n2", sorted[2].Key)
	}

	assert.Equal(t, "n5", sorted[3].Key)
}
