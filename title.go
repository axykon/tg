package main

import (
	"fmt"
	"log"

	"github.com/veandco/go-sdl2/sdl"
	ttf "github.com/veandco/go-sdl2/sdl_ttf"
)

// TitleScene is the initial splash screen
type TitleScene struct {
	menu         []menuItem
	menuTextures []*sdl.Texture
}

const (
	fontFile = "res/Go-Bold.ttf"
)

var (
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
	origTarget := renderer.GetRenderTarget()

	var (
		err      error
		fontSize int
	)

	ts.menu = []menuItem{{"Play", "play"}, {"Quit", "quit"}, {"Options", "options"}}
	ts.menuTextures = make([]*sdl.Texture, len(ts.menu))

	fontSize = windowHeight / (len(ts.menu) + 1)

	font, err := ttf.OpenFont(fontFile, fontSize)
	if err != nil {
		return fmt.Errorf("Could not open font %s: %v", fontFile, err)
	}
	defer font.Close()

	for i, item := range ts.menu {

		surface, err := font.RenderUTF8_Blended(item.label, fgColor)
		if err != nil {
			return fmt.Errorf("Could not render font: %v", err)
		}
		defer surface.Free()

		//TODO: extract into a function
		texture, err := renderer.CreateTextureFromSurface(surface)
		if err != nil {
			return fmt.Errorf("Could not create texture: %v", err)
		}
		defer texture.Destroy()

		selSurface, err := font.RenderUTF8_Blended(item.label, selColor)
		if err != nil {
			return fmt.Errorf("Could not render font: %v", err)
		}
		defer selSurface.Free()

		selTexture, err := renderer.CreateTextureFromSurface(selSurface)
		if err != nil {
			return fmt.Errorf("Could not create texture: %v", err)
		}
		defer selTexture.Destroy()

		ts.menuTextures[i], err = renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_TARGET,
			int(surface.W), int(surface.H+selSurface.H))
		renderer.SetRenderTarget(ts.menuTextures[i])
		renderer.SetDrawColor(bgColor.R, bgColor.G, bgColor.B, bgColor.A)
		renderer.Clear()
		renderer.Copy(texture, nil, &sdl.Rect{X: 0, Y: 0, W: surface.W, H: surface.H})
		renderer.Copy(selTexture, nil, &sdl.Rect{X: 0, Y: surface.H, W: selSurface.W, H: selSurface.H})

	}

	renderer.SetRenderTarget(origTarget)

	return nil

}

// HandleEvent handles events
func (ts *TitleScene) HandleEvent(event *sdl.Event) {
	switch evt := (*event).(type) {
	case *sdl.KeyDownEvent:
		if evt.Keysym.Sym == sdl.K_DOWN && itemSelected < len(ts.menu)-1 {
			itemSelected++
		} else if evt.Keysym.Sym == sdl.K_UP && itemSelected > 0 {
			itemSelected--
		} else if evt.Keysym.Sym == sdl.K_RETURN {
			log.Printf("Action selected: %s", ts.menu[itemSelected].action)
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

	y := 0

	for i, t := range ts.menuTextures {
		_, _, w, h, _ := t.Query()

		var srcY int32
		if i == itemSelected {
			srcY = h / 2
		}

		renderer.Copy(t, &sdl.Rect{X: 0, Y: srcY, W: w, H: h / 2}, &sdl.Rect{X: (int32(windowWidth) - w) / 2, Y: int32(y), W: w, H: h / 2})

		y = y + int(h/2) // + 20

	}

	return nil
}

// Destroy releases allocated resources
func (ts *TitleScene) Destroy() {
	for _, t := range ts.menuTextures {
		t.Destroy()
	}
}
