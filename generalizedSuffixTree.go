package generalizedSuffixTree

type gst struct {
	elementsList []string
	root         *node
}

func (newGST *gst) Search(searchString string) (resultSet []string) {
	indexes := search(newGST, searchString)
	resultSet = make([]string, len(indexes))
	for i := 0; i < len(indexes); i++ {
		resultSet[i] = newGST.elementsList[indexes[i]]
	}

	return
}

func (newGST *gst) Add(newString string, index int) *gst {
	add(newString, index)
	return newGST
}

func createGST() *gst {
	return &gst{}
}
