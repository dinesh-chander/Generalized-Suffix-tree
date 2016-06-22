package generalizedSuffixTree

type gst struct {
	elementsList map[int]string
	root         *node
}

func (newGST *gst) Search(searchString string) (resultSet []string) {
	indexes := newGST.search(searchString)
	resultSet = make([]string, len(indexes))
	for i := 0; i < len(indexes); i++ {
		resultSet[i] = newGST.elementsList[indexes[i]]
	}
	return
}

func (newGST *gst) Add(newString string, index int) *gst {
	newGST.elementsList[index] = newString
	newGST.add(newString, index)
	return newGST
}

func createGST() *gst {
	return &gst{elementsList: make(map[int]string), root: createNewNode("", nil)}
}
