package filter

import (
	"fmt"
	"testing"
)

func TestReplace(t *testing.T) {
	err := Init("./words.txt")
	if err != nil {
		t.Errorf("open file error:%v\n", err)
		return
	}
	content := `习近平
TMD`
	ret, hit := Replace(content, "*")
	fmt.Printf("After replace: %s,hit:%v\n", ret, hit)
}
