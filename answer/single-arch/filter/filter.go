package filter

func Replace(content string, replace string) (result string, hit bool) {
	result, hit = trie.FilterWords(content, replace)
	return
}
