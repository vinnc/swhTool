/** file manager **/
package main

import (
	"flag"
	"github.com/karrick/godirwalk"
)

func IsFileExist(fileName string) bool {
	_files := ReadCurrentDirectories()
	return FindInSlice(_files, fileName)
}

func ReadCurrentDirectories() []string {
	dirname := "."
	if flag.NArg() > 0 {
		dirname = flag.Arg(0)
	}

	scanner, err := godirwalk.NewScanner(dirname)
	if err != nil {
		fatal("cannot scan directory: %s", err)
	}

	var _dirs []string

	for scanner.Scan() {
		dirent, err := scanner.Dirent()
		if err != nil {
			warning("cannot get dirent: %s", err)
			continue
		}
		name := dirent.Name()
		if name == "break" {
			break
		}
		if name == "continue" {
			continue
		}
		_dirs = append(_dirs, name)
	}
	if err := scanner.Err(); err != nil {
		fatal("cannot scan directory: %s", err)
	}
	return _dirs
}
