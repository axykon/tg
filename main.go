package main

import (
	"log"
	"os"
	"strings"

	"github.com/axykon/tg/game"
	"github.com/axykon/tg/menu"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	windowTitle = "Tank & Gun"
)

// Scene is an abstract scene which can be rendered on the screen
type Scene interface {
	Init() error
	Render() error
	HandleEvent(event *sdl.Event)
	Update() string
	Destroy()
}

var (
	window *sdl.Window
)

func main() {
	var err error

	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatalf("Could not initialize SDL: %v\n", err)
	}
	defer sdl.Quit()

	if err = ttf.Init(); err != nil {
		log.Fatalf("Could not initialize SDL_ttf: %v\n", err)
	}

	w, h, err := calcWindowSize()
	if err != nil {
		log.Fatalf("Could not get window size: %v", err)
	}

	window, err = sdl.CreateWindow(windowTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(w), int32(h), sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatalf("Could not create window: %v\n", err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		log.Fatalf("Could not create renderer: %v", err)
	}

	font, err := ttf.OpenFont("res/menu.ttf", 40)
	if err != nil {
		log.Fatalf("Could not open font: %v", err)
	}
	defer font.Close()

	var currentScene Scene
	var nextScene = "menu"

loop:
	for {
		event := sdl.PollEvent()
		switch event := event.(type) {
		case *sdl.QuitEvent:
			break loop
		case *sdl.KeyboardEvent:
			if event.Keysym.Sym == sdl.K_ESCAPE || event.Keysym.Sym == 'q' {
				break loop
			}
		}

		if currentScene != nil {
			currentScene.HandleEvent(&event)
			nextScene = currentScene.Update()
		}

		if nextScene != "" {
			log.Printf("Next scene is %s", nextScene)
			if currentScene != nil {
				currentScene.Destroy()
				currentScene = nil
			}

			switch nextScene {
			case "menu":
				m := menu.New(renderer, font, w, h)
				m.Add("Play", "game")
				m.Add("Quit", "exit")
				currentScene = m
			case "game":
				g := game.New(renderer, w, h)
				currentScene = g
			case "exit":
				break loop
			}

			if err = currentScene.Init(); err != nil {
				log.Fatalf("Could not init scene %s: %v", nextScene, err)
			}
			nextScene = ""

		}
		if err = currentScene.Render(); err != nil {
			log.Fatalf("Could not render scene: %v", err)
		}
		renderer.Present()
	}
}

func calcWindowSize() (w int, h int, err error) {
	if hostname, e := os.Hostname(); e != nil {
		err = e
	} else if strings.HasPrefix(hostname, "w") {
		w, h = 400, 250
	} else {
		w, h = 800, 600
	}
	return
}
