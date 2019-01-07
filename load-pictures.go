package main

import (
	"image"
	"os"

	"github.com/faiface/pixel"
)

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func loadSprite(path string) (pixel.Sprite, error) {
	pic, err := loadPicture(path)
	if err != nil {
		return *pixel.NewSprite(pic, pic.Bounds()), err
	}
	sprite := pixel.NewSprite(pic, pic.Bounds())
	return *sprite, nil
}

// func loadTTF(path string, size float64, origin pixel.Vec) *text.Text {
// 	file, err := os.Open(path)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	bytes, err := ioutil.ReadAll(file)
// 	if err != nil {
// 		panic(err)
// 	}

// 	font, err := truetype.Parse(bytes)
// 	if err != nil {
// 		panic(err)
// 	}

// 	face := truetype.NewFace(font, &truetype.Options{
// 		Size:              size,
// 		GlyphCacheEntries: 1,
// 	})

// 	atlas := text.NewAtlas(face, text.ASCII)

// 	txt := text.New(origin, atlas)

// 	return txt

// }
