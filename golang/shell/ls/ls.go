package ls

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

// lsコマンドの再発明
func Exec(dir string) string {
	fileNames := getFileNames(dir)
	return strings.Join(fileNames, "\t")
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
