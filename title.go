package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	ttf "github.com/veandco/go-sdl2/sdl_ttf"
)

// TitleScene is the initial splash screen
type TitleScene struct {
	title *sdl.Texture
}

const (
	fontFile = "res/Go-Bold.ttf"
	fontSize = 80
)

var (
	fgColor = sdl.Color{A: 255, R: 123, G: 123, B: 255}
)

// Init initializes resources
func (ts *TitleScene) Init(renderer *sdl.Renderer) error {
	font, err := ttf.OpenFont(fontFile, fontSize)
	if err != nil {
		return fmt.Errorf("Could not open font %s: %v", fontFile, err)
	}
	defer font.Close()

	if err = renderer.SetDrawColor(fgColor.R, fgColor.G, fgColor.B, fgColor.A); err != nil {
		return fmt.Errorf("Could not set draw color: %v", err)
	}

	surface, err := font.RenderUTF8_Blended("Tank & Gun", sdl.Color{A: 255, R: 255, G: 255, B: 255})
	if err != nil {
		return fmt.Errorf("Could not render font: %v", err)
	}

	ts.title, err = renderer.CreateTextureFromSurface(surface)
	if err != nil {
		return fmt.Errorf("Could not create texture from surface: %v", err)
	}

	return nil
}

// HandleEvent handles events
func HandleEvent(event *sdl.Event) {
}

// Render renders the scene
func (ts *TitleScene) Render(renderer *sdl.Renderer) error {
	if err := renderer.Clear(); err != nil {
		return fmt.Errorf("Could not clear target: %v", err)
	}

	if err := renderer.Copy(ts.title, nil, nil); err != nil {
		return fmt.Errorf("Could not copy texture: %v", err)
	}

	return nil
}

func (ts *TitleScene) initTitle() {
}

// Destroy relases allocated resources
func (ts *TitleScene) Destroy() {
	ts.title.Destroy()
}
