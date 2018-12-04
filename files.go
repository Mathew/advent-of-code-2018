package files

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func LoadFile(filePath string) string {
	absPath, _ := filepath.Abs(filePath)
	dat, err := ioutil.ReadFile(absPath)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(dat)
}
