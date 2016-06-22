package generalizedSuffixTree

func findAllSuffixes(newString string) (suffixList []string) {
	newStringLength := len(newString)

	suffixList = make([]string, newStringLength)

	for i := 0; i < newStringLength; i++ {
		suffixList[i] = newString[i:newStringLength]
	}

	return
}
