package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
	ttf "github.com/veandco/go-sdl2/sdl_ttf"
)

const (
	windowTitle  = "Tank & Gun"
	windowWidth  = 800
	windowHeight = 600
)

func main() {
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		log.Fatalf("Could not initialize SDL: %v\n", err)
	}
	defer sdl.Quit()

	if err := ttf.Init(); err != nil {
		log.Fatalf("Could not initialize SDL_ttf: %v\n", err)
	}

	window, err := sdl.CreateWindow(windowTitle,
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		windowWidth, windowHeight,
		sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatalf("Could not create window: %v\n", err)
	}
	defer window.Destroy()

	font, err := ttf.OpenFont("res/ptsansb.ttf", 32)
	if err != nil {
		log.Fatalf("Could not open font: %v\n", err)
	}
	defer font.Close()

	solid, err := font.RenderUTF8_Solid(windowTitle, sdl.Color{A: 255, R: 0, G: 244, B: 255})
	if err != nil {
		log.Fatalf("Could not render text: %v\n", err)
	}
	defer solid.Free()

	surface, err := window.GetSurface()
	if err != nil {
		log.Fatalf("Could not get window surface: %v\n", err)
	}
	defer surface.Free()

	if err := solid.Blit(nil, surface, nil); err != nil {
		log.Fatalf("Could not put text onto surface: %v\n", err)
	}
	window.UpdateSurface()

	sdl.Delay(3000)
}
