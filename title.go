package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
	ttf "github.com/veandco/go-sdl2/sdl_ttf"
)

// TitleScene is the initial splash screen
type TitleScene struct {
	title *sdl.Texture
}

const (
	fontFile = "res/Go-Bold.ttf"
	fontSize = 120
)

var (
	menu = [...]menuItem{{"Play", "play"}, {"Quit", "quit"}}
)

type menuItem struct {
	label  string
	action string
}

var (
	fgColor = sdl.Color{A: 255, R: 123, G: 123, B: 255}
)

// Init initializes resources
func (ts *TitleScene) Init(renderer *sdl.Renderer) error {
	for _, i := range menu {
		log.Printf("%s -> %s", i.label, i.action)
	}
	return nil
}

// HandleEvent handles events
func (ts *TitleScene) HandleEvent(event *sdl.Event) {
	switch evt := (*event).(type) {
	case *sdl.KeyDownEvent:
		if evt.Keysym.Sym == sdl.K_SPACE {
			log.Println("Space pressed")
			ts.title.Destroy()
			ts.title = nil
		}
	}
}

// Render renders the scene
func (ts *TitleScene) Render(renderer *sdl.Renderer) error {
	var err error
	if ts.title == nil {
		if ts.title, err = createTitle(renderer); err != nil {
			return fmt.Errorf("Could not create title: %v", err)
		}
	}
	if err := renderer.Clear(); err != nil {
		return fmt.Errorf("Could not clear target: %v", err)
	}

	if err := renderer.Copy(ts.title, nil, &sdl.Rect{X: 10, Y: int32(windowHeight) / 4,
		W: int32(windowWidth) - 20, H: int32(windowHeight) / 2}); err != nil {
		return fmt.Errorf("Could not copy texture: %v", err)
	}

	return nil
}

func createTitle(renderer *sdl.Renderer) (*sdl.Texture, error) {
	font, err := ttf.OpenFont(fontFile, fontSize)
	if err != nil {
		return nil, fmt.Errorf("Could not open font %s: %v", fontFile, err)
	}
	defer font.Close()

	r, g, b := randomColor(), randomColor(), randomColor()
	r1, g1, b1 := 255-r, 255-g, 255-b

	if err = renderer.SetDrawColor(r, g, b, fgColor.A); err != nil {
		return nil, fmt.Errorf("Could not set draw color: %v", err)
	}

	surface, err := font.RenderUTF8_Blended("Tank & Gun", sdl.Color{A: 255, R: r1, G: g1, B: b1})
	if err != nil {
		return nil, fmt.Errorf("Could not render font: %v", err)
	}

	title, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		return nil, fmt.Errorf("Could not create texture from surface: %v", err)
	}

	return title, nil
}

func randomColor() uint8 {
	return uint8(rand.Uint32() % 256)
}

// Destroy relases allocated resources
func (ts *TitleScene) Destroy() {
	ts.title.Destroy()
}
