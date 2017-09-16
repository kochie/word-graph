package main

import (
	"fmt"

	"github.com/kochie/word-graph/webGrab"
)

func main() {
	// word_graph := wordGraph.New()
	fmt.Println(webGrab.GetSynonyms("logic"))
}
