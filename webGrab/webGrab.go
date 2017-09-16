package webGrab

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// Synonyms struct has the associated words relating to the current word.
type Synonyms struct {
	firstTier  []string
	secondTier []string
	thirdTier  []string
}

type dataCategory struct {
	name  string
	color string
}

const baseURL string = "http://www.thesaurus.com/browse/"

func addSynonym(synonyms *Synonyms, categoryValue string, word string) {
	var data dataCategory
	fmt.Println(categoryValue)
	err := json.Unmarshal([]byte(categoryValue), &data)
	checkError(err)

	fmt.Println(data, word)

	switch {
	case data.name == "relevant-1":
		synonyms.thirdTier = append(synonyms.thirdTier, word)
		break
	case data.name == "relevant-2":
		synonyms.secondTier = append(synonyms.secondTier, word)
		break
	case data.name == "relevant-3":
		synonyms.firstTier = append(synonyms.firstTier, word)
		break
	default:
		fmt.Println("This default case statement should be unreachable!")
		break
	}
}

func readWordAnchor(t html.Token, synonyms *Synonyms) {
	for _, a := range t.Attr {
		if a.Key == "data-category" {
			doc, err := html.Parse(strings.NewReader(t.Data))
			checkError(err)
			addSynonym(synonyms, a.Val, doc.FirstChild.Data)
			break
		}
	}
}

// GetSynonyms will return a struct containing all the synonyms of a word.
func GetSynonyms(word string) (synonyms Synonyms) {
	response, err := http.Get(baseURL + word)
	checkError(err)

	z := html.NewTokenizer(response.Body)

	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return
		case tt == html.StartTagToken:
			t := z.Token()
			isAnchor := t.Data == "a"
			if isAnchor {
				readWordAnchor(t, &synonyms)
			}
		}
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
