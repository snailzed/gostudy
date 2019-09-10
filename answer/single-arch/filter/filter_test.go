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
	content := `回民吃猪肉
习近平
TMD
毛民进党hahahah
妹妹淫水 流
机吧
联国w3434
1989六四
性爱电影
李红智
梁光烈
巴黎市长
成人BT
免注册淫色电影
UltraSurf
鐵血三國志
性爱图库,3534543534,6666`
	ret := Replace(content, "*")
	fmt.Printf("After replace: %s\n\n", ret)
}
