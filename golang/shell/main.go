package main

import (
	"./ls"
	"flag"
	"fmt"
)

func main() {
	args := readArgs()
	result := ls.Exec(args[0])
	fmt.Println(result)
}

// コマンドライン引数を取得する
func readArgs() []string {
	flag.Parse()
	return flag.Args()
}
