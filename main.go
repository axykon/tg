package main

import (
	"log"
	"os"
	"strings"

	"github.com/veandco/go-sdl2/sdl"
	ttf "github.com/veandco/go-sdl2/sdl_ttf"
)

const (
	windowTitle = "Tank & Gun"
)

type Scene interface {
	Init(r *sdl.Renderer) error
	Render(r *sdl.Renderer) error
	HandleEvent(event *sdl.Event)
	Destroy()
}

var (
	window                    *sdl.Window
	renderer                  *sdl.Renderer
	windowWidth, windowHeight int
)

func main() {
	if hostname, err := os.Hostname(); err != nil {
		log.Fatalf("Could not get host name: %v", err)
	} else if strings.HasPrefix(hostname, "w") {
		windowWidth, windowHeight = 200, 140
	} else {
		windowWidth, windowHeight = 800, 600
	}

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatalf("Could not initialize SDL: %v\n", err)
	}
	defer sdl.Quit()

	if err := ttf.Init(); err != nil {
		log.Fatalf("Could not initialize SDL_ttf: %v\n", err)
	}

	window, err := sdl.CreateWindow(windowTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		windowWidth, windowHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatalf("Could not create window: %v\n", err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		log.Fatalf("Could not create renderer: %v", err)
	}

	var t Scene = &TitleScene{}
	defer t.Destroy()
	if err = t.Init(renderer); err != nil {
		log.Fatalf("Could not init title scene: %v", err)
	}

loop:
	for {
		event := sdl.PollEvent()
		switch event := event.(type) {
		case *sdl.QuitEvent:
			break loop
		case *sdl.KeyDownEvent:
			if event.Keysym.Sym == sdl.K_ESCAPE || event.Keysym.Sym == 'q' {
				break loop
			}
		}
		t.HandleEvent(&event)
		if err = t.Render(renderer); err != nil {
			log.Fatalf("Could not render scene: %v", err)
		}
		renderer.Present()
	}
}
