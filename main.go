package main

import (
	"fmt"
	"github.com/mholt/archiver/v4"
	"io/fs"
)

func main() {
	fsys, err := archiver.FileSystem(`test.tar.gz`)
	if err != nil {
		panic(err)
	}

	err = fs.WalkDir(fsys, ".", func(path string, f fs.DirEntry, err error) error {
		if f.IsDir() {
			return nil
		}

		fmt.Printf(`trying to open %q`+"\n", path)
		fh, err := fsys.Open(path)
		if err != nil {
			return err
		}
		defer fh.Close()

		fmt.Printf(`file: %v`+"\n", path)

		return nil
	})
	if err != nil {
		panic(err)
	}
}
