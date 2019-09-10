package filter

import (
	"bufio"
	"gostudy/answer/single-arch/util"
	"io"
	"os"
	"strings"
)

var (
	trie *util.Trie
)

//初始化字典树
func Init(filename string) (err error) {
	trie = util.NewTrie()
	fd, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fd.Close()
	reader := bufio.NewReader(fd)
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return err
		}
		if err == io.EOF {
			break
		}
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		trie.Add(line)

	}
	return nil
}
