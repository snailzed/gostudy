package base

import "fmt"

func main() {
	fmt.Println("Start!")
	goto End //直接跳转到End标签代码
	fmt.Println("Processing!")

End:

	fmt.Println("END!")
}
