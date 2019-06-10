package sitemap

import (
	"fmt"
)

//DisplayBFS shows nodes once Breadth First Search
func DisplayBFS(sorted []*Node) {
	var i = 0
	for _, kv := range sorted {
		i++
		fmt.Printf("%04d,%s%s%s %d\n", i, kv.VisitedFrom.Key, "->", kv.Key, kv.Value())
	}
}

//Display shows Nodes as strings by weight
func Display(sorted []*Node) {
	var i = 0
	for _, kv := range sorted {
		i++
		fmt.Printf("%04d,%s,%d\n", i, kv.Key, kv.Value())
	}
}
