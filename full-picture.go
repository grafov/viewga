package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func showFullPicture(win *pixelgl.Window, pic picture) {
	var (
		mainH = win.Bounds().H()
		mainW = win.Bounds().W()
	)
	var mat = pixel.IM
	mat = pixel.IM.Moved(pixel.V(pic.W/2., pic.H/2.))
	mat = mat.ScaledXY(pixel.ZV, pixel.V(mainW/pic.W, mainH/pic.H))
	pic.Sprite.Draw(win, mat)
}
