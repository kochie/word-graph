package wordGraph

import (
	"github.com/kochie/word-graph/word"
)

// WordGraph data structure.
type WordGraph struct {
	graph map[string]word.Word
}

// New constructs a new wordGraph.
func New() WordGraph {
	return WordGraph{
		graph: make(map[string]word.Word),
	}
}

// Add will include the Word w to the graph.
func (wdgh WordGraph) Add(w word.Word) bool {
	if wdgh.Exists(w.GetValue()) {
		return false
	}
	wdgh.graph[w.GetValue()] = w
	return true
}

// Remove will delete a value from the WordGraph if it exists.
func (wdgh WordGraph) Remove(word string) bool {
	if wdgh.Exists(word) {
		delete(wdgh.graph, word)
		return true
	}
	return false
}

// Exists checks if a value exists in the WordGraph
func (wdgh WordGraph) Exists(word string) bool {
	if _, ok := wdgh.graph[word]; ok {
		return true
	}
	return false
}

// Link will connect two words by their similarWords list.
func (wdgh WordGraph) Link(wordA string, wordB string) bool {
	if wdgh.Exists(wordA) && wdgh.Exists(wordB) {
		if wdgh.graph[wordA].Link(wdgh.graph[wordB]) {
			return true
		}
	}
	return false

}
