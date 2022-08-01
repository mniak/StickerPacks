package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
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
	entries, err := os.ReadDir(basedir)
	handle(err)

	i := 1
	for _, file := range entries {
		if file.Type().IsDir() {
			continue
		}
		if !strings.HasSuffix(file.Name(), extension) {
			continue
		}
		handle(os.Rename(filepath.Join(basedir, file.Name()), filepath.Join(basedir, fmt.Sprintf("%04d"+extension, i))))
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
			handle(os.Rename(filepath.Join(basedir, folder.Name(), file.Name()), filepath.Join(basedir, fmt.Sprintf("%04d"+extension, i))))
			i++
		}
		os.RemoveAll(filepath.Join(basedir, folder.Name()))
	}
}
