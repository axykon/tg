package main

import (
	"fmt"
	"log"

	"github.com/veandco/go-sdl2/sdl"
	ttf "github.com/veandco/go-sdl2/sdl_ttf"
)

// TitleScene is the initial splash screen
type TitleScene struct {
}

const (
	fontFile = "res/Go-Bold.ttf"
	fontSize = 120
)

var (
	menu         = []menuItem{{"Play", "play"}, {"Settings", "settings"}, {"Quit", "quit"}}
	font         *ttf.Font
	bgColor      = sdl.Color{R: 30, G: 30, B: 30, A: 255}
	fgColor      = sdl.Color{R: 120, G: 120, B: 120, A: 255}
	selColor     = sdl.Color{R: 255, G: 255, B: 255, A: 255}
	itemSelected = 0
)

type menuItem struct {
	label  string
	action string
}

// Init initializes resources
func (ts *TitleScene) Init(renderer *sdl.Renderer) error {
	var err error
	font, err = ttf.OpenFont(fontFile, fontSize)
	if err != nil {
		return fmt.Errorf("Could not open font %s: %v", fontFile, err)
	}
	return nil
}

// HandleEvent handles events
func (ts *TitleScene) HandleEvent(event *sdl.Event) {
	switch evt := (*event).(type) {
	case *sdl.KeyDownEvent:
		if evt.Keysym.Sym == sdl.K_DOWN && itemSelected < len(menu)-1 {
			itemSelected++
		} else if evt.Keysym.Sym == sdl.K_UP && itemSelected > 0 {
			itemSelected--
		} else if evt.Keysym.Sym == sdl.K_RETURN {
			log.Printf("Action selected: %s", menu[itemSelected].action)
		}
	}
}

// Render renders the scene
func (ts *TitleScene) Render(renderer *sdl.Renderer) error {
	var err error
	if err = renderer.SetDrawColor(bgColor.R, bgColor.G, bgColor.B, bgColor.A); err != nil {
		return fmt.Errorf("Could not set draw color: %v", err)
	}
	if err = renderer.Clear(); err != nil {
		return fmt.Errorf("Could not clear target: %v", err)
	}

	y := 50

	for i, item := range menu {
		var itemColor sdl.Color
		if i == itemSelected {
			itemColor = selColor
		} else {
			itemColor = fgColor
		}

		surface, err := font.RenderUTF8_Blended(item.label, itemColor)
		if err != nil {
			return fmt.Errorf("Could not render font: %v", err)
		}
		defer surface.Free()

		w, h := surface.W, surface.H

		texture, err := renderer.CreateTextureFromSurface(surface)
		if err != nil {
			return fmt.Errorf("Could not create texture: %v", err)
		}
		defer texture.Destroy()

		renderer.Copy(texture, nil, &sdl.Rect{X: (int32(windowWidth) - w) / 2, Y: int32(y), W: w, H: h})

		y = y + int(h) + 20

	}

	return nil
}

// Destroy relases allocated resources
func (ts *TitleScene) Destroy() {
	font.Close()
}
