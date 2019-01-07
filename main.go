package main

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	_ "image/png"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/grafov/kiwi"
	"golang.org/x/image/colornames"
)

var (
	log        = kiwi.Fork()
	oldW, oldH float64
	changed    bool
	smoothMode bool
)

func run() {
	if len(os.Args) > 1 {
		paths = os.Args[1:]
	} else {
		paths = []string{"."}
	}
	pictures := loadFiles(pathsToFiles(paths))

	cfg := pixelgl.WindowConfig{
		Title:  fmt.Sprintf("viewga: %v", paths),
		Bounds: pixel.R(0, 0, 1024, 768),
	}
	mainWin, err := pixelgl.NewWindow(cfg)
	if err != nil {
		log.Log("err", err)
		panic(err)
	}
	mainWin.Clear(colornames.Darkgray)

	for !mainWin.Closed() {
		switch {
		case mainWin.JustPressed(pixelgl.KeyS):
			smoothMode = !smoothMode
			changed = true
			mainWin.SetSmooth(smoothMode)
		case mainWin.JustPressed(pixelgl.KeyQ):
			mainWin.SetClosed(true)
			shutdown()
		}

		if !isMainWinChanged(mainWin) {
			if len(pictures) == 1 {
				showFullPicture(mainWin, pictures[0])
			} else {
				showPreviewBar(mainWin, pictures)
			}
		}

		mouse := mainWin.MousePosition()
		mainWin.SetTitle(fmt.Sprintf("viewga: %v", mouse))
		mainWin.Update()
	}
}

var paths []string

func main() {
	//	kiwi.SinkTo(os.Stderr, kiwi.AsLogfmt()).WithKey("err").Start()
	kiwi.SinkTo(os.Stdout, kiwi.AsLogfmt()).Start()
	image.RegisterFormat("png", "PNG", png.Decode, png.DecodeConfig)
	image.RegisterFormat("gif", "GIF", gif.Decode, gif.DecodeConfig)
	image.RegisterFormat("jpeg", "JPEG", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("jpg", "JPEG", jpeg.Decode, jpeg.DecodeConfig)

	pixelgl.Run(run)
	shutdown()
}

func shutdown() {
	kiwi.FlushAll()
	os.Exit(0)
}

// Check when it should be redrawed?
func isMainWinChanged(win *pixelgl.Window) bool {
	var (
		mainH = win.Bounds().H()
		mainW = win.Bounds().W()
	)
	if mainH == oldH || mainW == oldW {
		return false
	}
	oldH = mainH
	oldW = mainW
	if changed {
		changed = false
	}
	win.Clear(colornames.Darkgray)
	return true
}
