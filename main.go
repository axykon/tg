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
	Init() error
	Render() error
	HandleEvent(event *sdl.Event)
	Destroy()
}

var (
	window *sdl.Window
)

func main() {
	var err error
	log.Print(bgColor)

	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatalf("Could not initialize SDL: %v\n", err)
	}
	defer sdl.Quit()

	if err = ttf.Init(); err != nil {
		log.Fatalf("Could not initialize SDL_ttf: %v\n", err)
	}

	w, h, err := getWindowSize()
	if err != nil {
		log.Fatalf("Could not get window size: %v", err)
	}

	window, err = sdl.CreateWindow(windowTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		w, h, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatalf("Could not create window: %v\n", err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		log.Fatalf("Could not create renderer: %v", err)
	}

	var t Scene = &MenuScene{}
	defer t.Destroy()
	if err = t.Init(); err != nil {
		log.Fatalf("Could not init menu scene: %v", err)
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
		if err = t.Render(); err != nil {
			log.Fatalf("Could not render scene: %v", err)
		}
		renderer.Present()
	}
}

func getWindowSize() (w int, h int, err error) {
	if hostname, e := os.Hostname(); e != nil {
		err = e
	} else if strings.HasPrefix(hostname, "w") {
		w, h = 200, 140
	} else {
		w, h = 800, 600
	}
	return
}
