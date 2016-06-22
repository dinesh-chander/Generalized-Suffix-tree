//	package generalizedSuffixTree provides fast retrieval and insertion operations
package generalizedSuffixTree

//	add method is a private method for gst struct
//	It will add a new given string in the gst and returns a boolean value i.e the operation status
func (newGST *gst) add(newString string, index int) {
	suffixList := findAllSuffixes(newString)
	for _, suffix := range suffixList {
		newGST.insert(suffix, index)
	}
}

func (newGST *gst) insert(newSuffix string, index int) {
	targetNode, suffixToBeAppended := findSuitablePositionToInsertNewNode(newSuffix, newGST.root)
	appendNewNode(targetNode, suffixToBeAppended, index)
}

func createNewNode(newString string, parent *node) *node {
	return &node{edgeLabel: newString, child_nodes: make(map[byte]*node), indexData: make(map[int][]int), parent: parent}
}

// findSuitablePositionToInsertNewNode method will find the correct node to insert a new node
// This method will start from root node and then find substring in the child nodes
func findSuitablePositionToInsertNewNode(newString string, currentNode *node) (*node, string) {
	if len(newString) == 0 {
		return currentNode, ""
	}

	if currentNode.child_nodes[newString[0]] == nil {
		return currentNode, newString
	} else {
		return findSuitablePositionToInsertNewNode(newString[1:len(newString)], currentNode.child_nodes[newString[0]])
	}
}

func updateIndexes(targetNode *node, index int, newPosition int) {
	if targetNode.indexData[index] == nil {
		targetNode.indexData[index] = make([]int, 0)
	}

	targetNode.indexData[index] = append(targetNode.indexData[index], newPosition)
}

func appendNewNode(targetNode *node, suffixToBeAppended string, index int) {

	return
}
