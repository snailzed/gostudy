package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	debug bool
	cmd   string
	help  bool
)

func main() {
	flag.BoolVar(&debug, "d", false, "if debug")
	flag.BoolVar(&help, "h", false, "show this help")
	flag.StringVar(&cmd, "cb", "default", "command")
	flag.Parse()
	if help {
		flag.PrintDefaults()
	}
	os.Stdout.WriteString(fmt.Sprintf("debug=%v, help=%v, cmd=%s\n", debug, help, cmd))
}

func testosArgs() {
	args := os.Args
	fmt.Printf("%T,%v", args, args)
}
