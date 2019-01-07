package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/faiface/pixel"
)

type picture struct {
	Sprite *pixel.Sprite
	W, H   float64
	Name   string
}

// pathsToFiles returns files list for the selected paths.
func pathsToFiles(paths []string) []string {
	var files []string
	for _, path := range paths {
		fd, err := os.Open(path)
		if err != nil {
			continue
		}
		f, err := fd.Stat()
		if err != nil {
			fd.Close()
			continue
		}
		if f.IsDir() && !strings.Contains(path, "*") {
			path = filepath.Join(path, "*")
		}
		names, err := filepath.Glob(path)
		if err != nil {
			fd.Close()
			continue
		}
		files = append(files, names...)
		fd.Close()
	}
	log.Log("files", files)
	return files
}

func loadFiles(files []string) []picture {
	var (
		pic      pixel.Picture
		pictures []picture
		err      error
	)
	for _, file := range files {
		pic, err = loadPicture(file)
		if err != nil {
			log.Log("err", err, "file", file, "desc", "can't load preview")
			continue
		}
		bounds := pic.Bounds()
		pictures = append(pictures, picture{
			Sprite: pixel.NewSprite(pic, bounds),
			W:      bounds.W(),
			H:      bounds.H(),
			Name:   file,
		})
		log.Log("sprite", file, "w", bounds.W(), "h", bounds.H())
	}
	return pictures
}
