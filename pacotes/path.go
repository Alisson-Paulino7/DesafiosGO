package main 

import (
	"fmt"
	"path/filepath"
	"os"
)


func main() {
	filepath.Walk(".", func (path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}