package inverted

import (
	"fmt"
	"strings"
)

type PostingsListsMap struct {
	Term            string
	Frequency       int
	DocumentListing []int
}

type Inverted struct {
	HashMap map[string]*PostingsListsMap
	Items   []*PostingsListsMap
}

func (invertedIndex *Inverted) FindItem(Term string) int {
	defer fmt.Errorf("%s", "Not Found")
	for index, item := range invertedIndex.Items {
		if item.Term == Term {
			return index
		}
	}
	return 0
}

func (invertedIndex *Inverted) Search(searchTerm string) {
	Term := strings.ToLower(searchTerm)
	result := invertedIndex.HashMap[Term]
	if result != nil {
		fmt.Println("Found:", searchTerm, "in documents:", result.DocumentListing)
	} else {
		fmt.Println("Not Found:", searchTerm)
	}
}

func (invertedIndex *Inverted) AddItem2HashMap(Term string, Document int) {
	if invertedIndex.HashMap[Term] != nil {

		FoundItemPosition := invertedIndex.FindItem(Term)

		invertedIndex.Items[FoundItemPosition].Frequency++
		invertedIndex.Items[FoundItemPosition].DocumentListing = append(invertedIndex.Items[FoundItemPosition].DocumentListing, Document)
	} else {

		InvertedIndexEntry := &PostingsListsMap{
			Term:            Term,
			Frequency:       1,
			DocumentListing: []int{Document},
		}

		invertedIndex.HashMap[Term] = InvertedIndexEntry
		invertedIndex.Items = append(invertedIndex.Items, InvertedIndexEntry)
	}
}

func BuildInvertedIndex() *Inverted {
	invertedIndex := &Inverted{
		HashMap: make(map[string]*PostingsListsMap),
		Items:   []*PostingsListsMap{},
	}
	return invertedIndex
}
