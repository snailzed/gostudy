package filter

func Replace(content string, replace string) (result string) {
	result = trie.FilterWords(content, replace)
	return
}
