package inverted

func BuildDocMap(token []string) map[string]bool {
	docMap := make(map[string]bool)

	for _, word := range token {
		if _, value := docMap[word]; !value {
			docMap[word] = true
		}
	}

	return docMap
}

func NewInvertedIndex(DocList []string) Inverted {
	globalDocMap := make([]map[string]bool, 0)

	for _, Doc := range DocList {
		token := Tokenize(Doc)
		docMap := BuildDocMap(token)
		globalDocMap = append(globalDocMap, docMap)
	}

	invertedIndex := BuildInvertedIndex()

	for DocMapIndex, DocMap := range globalDocMap {
		for DocEntry := range DocMap {
			invertedIndex.AddItem2HashMap(DocEntry, DocMapIndex)
		}
	}
	return *invertedIndex
}
