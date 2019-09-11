//构造trie字典树：减少时间复杂度
package util

type Node struct {
	char   rune           //当前节点对应的字符
	childs map[rune]*Node //当前节点的所有子节点，map存放： 下一个节点--》节点对象
	depth  int            //节点深度
	isEnd  bool           //是否结尾词
}
type Trie struct {
	root *Node
	size int
}

//创建新的字典树
func NewTrie() *Trie {
	return &Trie{
		root: NewNode(),
	}
}

//创建新节点
func NewNode() *Node {
	return &Node{
		childs: make(map[rune]*Node, 32),
	}
}

//将字符串添加到字典树中
func (t *Trie) Add(key string) {
	runeArr := []rune(key)
	node := t.root
	for index, char := range runeArr {
		ret, exists := node.childs[char]
		//不存在则新增
		if !exists {
			ret = NewNode()
			ret.depth = index + 1
			ret.char = char
			node.childs[char] = ret
		}
		node = ret
	}
	node.isEnd = true
	return
}

//查找需要过滤的字符串，并进行替换:从根节点开始查找
func (t *Trie) FilterWords(words, replace string) (result string, hit bool) {
	wordsRunes := []rune(words)
	node := t.root
	if node == nil {
		return
	}
	//resultRune := make([]rune, len(wordsRunes))
	resultRune := []rune{}
	start := 0
	for index, char := range wordsRunes {
		ret, exists := node.childs[char]
		//如果不存在的话，则保存中间结果,不需要过滤
		if !exists {
			resultRune = append(resultRune, wordsRunes[start:index+1]...)
			start = index + 1
			node = t.root
			continue
		}
		//存在的话
		node = ret
		if node.isEnd {
			hit = true
			node = t.root
			resultRune = append(resultRune, ([]rune(replace))...)
			start = index + 1
			continue
		}
	}
	result = string(resultRune)
	return
}
