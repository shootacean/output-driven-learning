package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {
	args := readArgs()
	fmt.Println(ls(args[0]))
}

// コマンドライン引数を取得する
func readArgs() []string {
	flag.Parse()
	return flag.Args()
}

// lsコマンドの再発明
func ls(dir string) []string {
	fileNames := getFileNames(dir)
	return fileNames
}

// 指定ディレクトリ配下のファイル名を取得する
func getFileNames(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	var paths []string
	for _, file := range files {
		paths = append(paths, filepath.Join(dir, file.Name()))
	}
	return paths
}
