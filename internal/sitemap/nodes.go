package sitemap

import (
	"container/list"
	"sort"
	"strings"

	"github.com/goncalopereira/gocrawler/internal/data"
)

//Node represents a page with links coming in and out from other pages
type Node struct {
	Key         string
	LinksIn     []*Node
	LinksOut    []*Node
	VisitedFrom *Node
}

//Value represents the weight of a page, pseudo page-rank counts links coming in
func (n Node) Value() int {
	return len(n.LinksIn)
}

//GetNode gets or creates the node for a new page
func GetNode(m map[string]*Node, key string) *Node {
	current, ok := m[key]
	if ok {
		return current
	}

	c := &Node{Key: key, LinksIn: []*Node{}, LinksOut: []*Node{}}
	m[key] = c
	return c

}

//BuildNodes DB to page Nodes
func BuildNodes(linksDB map[string]data.Request, urlDB map[string]data.Request) map[string]*Node {
	weighted := make(map[string]*Node, len(urlDB))

	for k := range linksDB {
		keys := strings.Split(k, "->")
		fromKey := keys[0]
		toKey := keys[1]

		from := GetNode(weighted, fromKey)
		to := GetNode(weighted, toKey)

		to.LinksIn = append(to.LinksIn, from)
		from.LinksOut = append(from.LinksOut, to)
	}
	return weighted
}

//BFS does search for nodes
//https://en.wikipedia.org/wiki/Breadth-first_search
func BFS(startNode *Node) []*Node {
	var result []*Node

	q := list.New()
	q.PushBack(startNode)
	startNode.VisitedFrom = &Node{Key: ""}
	for q.Len() > 0 {
		v := q.Front()
		q.Remove(v)
		n, _ := v.Value.(*Node)
		result = append(result, n)
		for _, child := range n.LinksOut {
			if child.VisitedFrom == nil {
				child.VisitedFrom = n
				q.PushBack(child)
			}
		}
	}

	return result
}

//SortWeightedList pseudo page-rank, reverse link/number of pages linking in
func SortWeightedList(weighted map[string]*Node) []*Node {
	var result []*Node
	for _, v := range weighted {
		result = append(result, v)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Value() > result[j].Value()
	})

	return result
}
