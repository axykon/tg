package menu

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	ttf "github.com/veandco/go-sdl2/sdl_ttf"
)

// Menu is the initial splash screen
type Menu struct {
	renderer      *sdl.Renderer
	width         int
	height        int
	font          *ttf.Font
	items         []item
	labels        []*sdl.Texture
	selected      int
	marginsHeight int
	labelHeight   int
	scene         string
}

const (
	labelSpacing = 20
)

var (
	bgColor  = sdl.Color{R: 30, G: 30, B: 30, A: 255}
	fgColor  = sdl.Color{R: 120, G: 120, B: 120, A: 255}
	selColor = sdl.Color{R: 255, G: 255, B: 255, A: 255}
)

// Item represents a menu item
type item struct {
	label string
	scene string
}

// New creates new menu with given items
func New(renderer *sdl.Renderer, font *ttf.Font, width int, height int) *Menu {
	return &Menu{renderer: renderer, font: font, width: width, height: height}
}

//Add adds new menu item
func (m *Menu) Add(label string, scene string) {
	m.items = append(m.items, item{label: label, scene: scene})
}

// Init initializes resources
func (m *Menu) Init() error {
	m.labels = make([]*sdl.Texture, len(m.items))
	m.labelHeight = m.font.Height()
	m.marginsHeight = (m.height - (m.labelHeight*len(m.items) + labelSpacing*(len(m.items)-1))) / 2

	for i := range m.items {
		if err := m.createLabel(i, m.font, m.renderer); err != nil {
			return fmt.Errorf("Could not create menu item: %v", err)
		}
	}

	return nil
}

// HandleEvent handles events
func (m *Menu) HandleEvent(event *sdl.Event) {
	switch evt := (*event).(type) {
	case *sdl.KeyDownEvent:
		if evt.Keysym.Sym == sdl.K_DOWN && m.selected < len(m.items)-1 {
			m.selected++
		} else if evt.Keysym.Sym == sdl.K_UP && m.selected > 0 {
			m.selected--
		} else if evt.Keysym.Sym == sdl.K_RETURN {
			m.scene = m.items[m.selected].scene
		}
	}
}

//Update implements scene's logic
func (m *Menu) Update() string {
	return m.scene
}

// Render renders the scene
func (m *Menu) Render() error {
	var err error
	if err = m.renderer.SetDrawColor(bgColor.R, bgColor.G, bgColor.B, bgColor.A); err != nil {
		return fmt.Errorf("Could not set draw color: %v", err)
	}

	if err = m.renderer.Clear(); err != nil {
		return fmt.Errorf("Could not clear target: %v", err)
	}

	y := m.marginsHeight - labelSpacing

	for i, t := range m.labels {
		y += labelSpacing
		_, _, w, h, _ := t.Query()

		var srcY int32
		if i == m.selected {
			srcY = h / 2
		}

		m.renderer.Copy(t,
			&sdl.Rect{X: 0, Y: srcY, W: w, H: int32(m.labelHeight)},
			&sdl.Rect{X: (int32(m.width) - w) / 2, Y: int32(y), W: w, H: int32(m.labelHeight)})

		y += m.labelHeight
	}

	return nil
}

func (m *Menu) createLabel(i int, font *ttf.Font, renderer *sdl.Renderer) (result error) {
	origTarget := renderer.GetRenderTarget()
	defer func() {
		result = renderer.SetRenderTarget(origTarget)
	}()

	item := m.items[i]

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

	m.labels[i], err = renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_TARGET,
		int(surface.W), m.labelHeight*2)
	renderer.SetRenderTarget(m.labels[i])
	renderer.SetDrawColor(bgColor.R, bgColor.G, bgColor.B, bgColor.A)
	renderer.Clear()
	renderer.Copy(texture, nil, &sdl.Rect{X: 0, Y: 0, W: surface.W, H: surface.H})
	renderer.Copy(selTexture, nil, &sdl.Rect{X: 0, Y: surface.H, W: selSurface.W, H: selSurface.H})

	return nil
}

// Destroy releases allocated resources
func (m *Menu) Destroy() {
	for _, t := range m.labels {
		t.Destroy()
	}
}
