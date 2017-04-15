package main

import (
	"fmt"
	"log"

	"github.com/veandco/go-sdl2/sdl"
	ttf "github.com/veandco/go-sdl2/sdl_ttf"
)

// MenuScene is the initial splash screen
type MenuScene struct {
	menu          []menuItem
	labels        []*sdl.Texture
	selected      int
	marginsHeight int
	labelHeight   int
}

const (
	fontFile     = "res/Go-Bold.ttf"
	fontSize     = 45
	labelSpacing = 40
)

var (
	bgColor  = sdl.Color{R: 30, G: 30, B: 30, A: 255}
	fgColor  = sdl.Color{R: 120, G: 120, B: 120, A: 255}
	selColor = sdl.Color{R: 255, G: 255, B: 255, A: 255}
)

type menuItem struct {
	label  string
	action string
}

// Init initializes resources
func (ts *MenuScene) Init() (err error) {
	renderer, err := window.GetRenderer()
	if err != nil {
		return fmt.Errorf("Could not get renderer: %v", err)
	}

	origTarget := renderer.GetRenderTarget()
	defer func() {
		err = renderer.SetRenderTarget(origTarget)
	}()

	ts.menu = []menuItem{{"Play", "play"}, {"Options", "options"}, {"Scores", "scores"}, {"Quit", "quit"}}
	ts.labels = make([]*sdl.Texture, len(ts.menu))

	font, err := ttf.OpenFont(fontFile, fontSize)
	if err != nil {
		return fmt.Errorf("Could not open font %s: %v", fontFile, err)
	}
	defer font.Close()

	ts.labelHeight = font.Height()
	_, windowHeight := window.GetSize()
	ts.marginsHeight = (windowHeight - (ts.labelHeight*len(ts.menu) + labelSpacing*(len(ts.menu)-1))) / 2

	log.Printf("window: %d, label: %d, spacing: %d, margin: %d", windowHeight, ts.labelHeight, labelSpacing, ts.marginsHeight)

	for i, item := range ts.menu {

		surface, err := font.RenderUTF8_Blended(item.label, fgColor)
		if err != nil {
			return fmt.Errorf("Could not render font: %v", err)
		}
		defer surface.Free()

		log.Printf("Label %s height is %d", item.label, surface.H)

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

		ts.labels[i], err = renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_TARGET,
			int(surface.W), int(surface.H+selSurface.H))
		renderer.SetRenderTarget(ts.labels[i])
		renderer.SetDrawColor(bgColor.R, bgColor.G, bgColor.B, bgColor.A)
		renderer.Clear()
		renderer.Copy(texture, nil, &sdl.Rect{X: 0, Y: 0, W: surface.W, H: surface.H})
		renderer.Copy(selTexture, nil, &sdl.Rect{X: 0, Y: surface.H, W: selSurface.W, H: selSurface.H})

	}

	return nil
}

// HandleEvent handles events
func (ts *MenuScene) HandleEvent(event *sdl.Event) {
	switch evt := (*event).(type) {
	case *sdl.KeyDownEvent:
		if evt.Keysym.Sym == sdl.K_DOWN && ts.selected < len(ts.menu)-1 {
			ts.selected++
		} else if evt.Keysym.Sym == sdl.K_UP && ts.selected > 0 {
			ts.selected--
		} else if evt.Keysym.Sym == sdl.K_RETURN {
			log.Printf("Action selected: %s", ts.menu[ts.selected].action)
		}
	}
}

// Render renders the scene
func (ts *MenuScene) Render() error {
	renderer, err := window.GetRenderer()
	windowWidth, _ := window.GetSize()
	if err != nil {
		return fmt.Errorf("Could not get renderer: %v", err)
	}

	if err = renderer.SetDrawColor(bgColor.R, bgColor.G, bgColor.B, bgColor.A); err != nil {
		return fmt.Errorf("Could not set draw color: %v", err)
	}

	if err = renderer.Clear(); err != nil {
		return fmt.Errorf("Could not clear target: %v", err)
	}

	y := ts.marginsHeight - labelSpacing

	for i, t := range ts.labels {
		y += labelSpacing
		_, _, w, h, _ := t.Query()

		var srcY int32
		if i == ts.selected {
			srcY = h / 2
		}

		renderer.Copy(t,
			&sdl.Rect{X: 0, Y: srcY, W: w, H: int32(ts.labelHeight)},
			&sdl.Rect{X: (int32(windowWidth) - w) / 2, Y: int32(y), W: w, H: int32(ts.labelHeight)})

		y += ts.labelHeight
	}

	return nil
}

// Destroy releases allocated resources
func (ts *MenuScene) Destroy() {
	for _, t := range ts.labels {
		t.Destroy()
	}
}
