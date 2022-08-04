package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func handle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	const basedir = "../../downloaded"
	const extension = ".webp"
	const prefix = ""

	entries, err := os.ReadDir(basedir)
	handle(err)

	i := 1
	var files []fs.FileInfo
	for _, file := range entries {
		if file.Type().IsDir() {
			continue
		}
		if !strings.HasSuffix(file.Name(), extension) {
			continue
		}
		fileinfo, err := file.Info()
		handle(err)
		files = append(files, fileinfo)

	}
	sort.Slice(files, func(i, j int) bool {
		ifile := files[i]
		jfile := files[j]
		return ifile.Size() < jfile.Size()
	})
	for _, file := range files {
		handle(os.Rename(filepath.Join(basedir, file.Name()), filepath.Join(basedir, fmt.Sprintf("%s%04d%s", prefix, i, extension))))
		i++
	}

	for _, folder := range entries {
		if !folder.Type().IsDir() {
			continue
		}
		files, err := os.ReadDir(filepath.Join(basedir, folder.Name()))
		handle(err)
		for _, file := range files {
			if file.Type().IsDir() {
				continue
			}
			if !strings.HasSuffix(file.Name(), extension) {
				continue
			}
			handle(os.Rename(filepath.Join(basedir, folder.Name(), file.Name()), filepath.Join(basedir, fmt.Sprintf("%s%04d%s", prefix, i, extension))))
			i++
		}
		os.RemoveAll(filepath.Join(basedir, folder.Name()))
	}
}
