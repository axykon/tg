package game

import "github.com/veandco/go-sdl2/sdl"

// Game is the main scene
type Game struct {
	renderer *sdl.Renderer
	width    int
	height   int
}

// New creates the new game
func New(renderer *sdl.Renderer, width int, height int) *Game {
	return &Game{renderer: renderer, width: width, height: height}
}

// Init initializes the game
func (g *Game) Init() error {
	return nil
}

// Renders the game
func (g *Game) Render() error {
	g.renderer.SetDrawColor(255, 255, 0, 128)
	g.renderer.Clear()
	return nil
}

// HandleEvent handles game events
func (g *Game) HandleEvent(event *sdl.Event) {
}

// Destroy releases resorces
func (g *Game) Destroy() {
}
