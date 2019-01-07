package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func showPreviewBar(win *pixelgl.Window, pictures []picture) {
	// Draw preview bar
	const border = 2
	var (
		totalX = 10.
		totalY = 5.
	)
	if int(totalX) > len(pictures) {
		totalX = float64(len(pictures))
	}
	var (
		mainH = win.Bounds().H()
		mainW = win.Bounds().W()
		picW  = mainW / totalX
		picH  = mainH / totalY
		count = 0
	)
	for y := mainH - picH - picH/2; y > picH/2; y -= picH + border {
		for x := picW / 2; x < mainW-picW/2; x += picW + border {
			if count >= len(pictures) {
				break
			}
			pic := pictures[count]
			count++
			var mat = pixel.IM
			mat = pixel.IM.Moved(pixel.V(x, y))
			fraction := pixel.V(picW/pic.W, picH/pic.H)
			mat = mat.ScaledXY(pixel.V(x, y), fraction)
			log.Log("title", pic.Name, "x", int(x), "y", int(y), "w", int(pic.W), "h", int(pic.H))
			pic.Sprite.Draw(win, mat)
		}
	}
}
