package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
	ttf "github.com/veandco/go-sdl2/sdl_ttf"
)

const (
	windowTitle  = "TANK & GUN"
	windowWidth  = 800
	windowHeight = 600
)

var (
	window   *sdl.Window
	renderer *sdl.Renderer
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatalf("Could not initialize SDL: %v\n", err)
	}
	defer sdl.Quit()

	if err := ttf.Init(); err != nil {
		log.Fatalf("Could not initialize SDL_ttf: %v\n", err)
	}

	window, renderer, err := sdl.CreateWindowAndRenderer(windowWidth, windowHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatalf("Could not create window: %v\n", err)
	}
	// defer renderer.Destroy()
	defer window.Destroy()

	font, err := ttf.OpenFont("res/ptsansb.ttf", 40)
	if err != nil {
		log.Fatalf("Could not open font: %v\n", err)
	}
	//defer font.Close()
	defer ttf.Quit()

	message, err := font.RenderUTF8_Solid(windowTitle, sdl.Color{A: 255, R: 255, G: 255, B: 255})
	if err != nil {
		log.Fatalf("Could not render text: %v\n", err)
	}
	defer message.Free()

	surface, err := window.GetSurface()
	if err != nil {
		log.Fatalf("Could not get window surface: %v\n", err)
	}
	texture, err := renderer.CreateTextureFromSurface(message)
	if err != nil {
		log.Fatalf("Could not create texture: %v\n", err)
	}
	defer surface.Free()
	defer texture.Destroy()

	renderer.SetDrawColor(255, 125, 0, 255)

	if err := renderer.Clear(); err != nil {
		log.Fatalf("Could not clear rendering target: %v", err)
	}
	renderer.Copy(texture, nil, &sdl.Rect{X: 10, Y: windowHeight / 4, W: windowWidth - 20, H: windowHeight / 2})
	// renderer.CopyEx(texture, nil, nil, 0, &sdl.Point{100, 100}, sdl.FLIP_NONE)
	renderer.Present()

	// if err := window.UpdateSurface(); err != nil {
	// 	log.Fatalf("Could not update window surface: %v\n", err)
	// }

loop:
	for {
		switch event := sdl.WaitEvent().(type) {
		case *sdl.QuitEvent:
			break loop
		case *sdl.KeyDownEvent:
			if event.Keysym.Sym == sdl.K_ESCAPE {
				break loop
			}
		}
	}
}
