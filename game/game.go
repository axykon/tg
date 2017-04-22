package game

import (
	"sync"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

// Game is the main scene
type Game struct {
	renderer  *sdl.Renderer
	width     int
	height    int
	mutex     *sync.Mutex
	nextScene string
}

// New creates the new game
func New(renderer *sdl.Renderer, width int, height int) *Game {
	g := &Game{renderer: renderer, width: width, height: height, mutex: &sync.Mutex{}}
	timer := time.NewTimer(time.Second * 5)
	go func() {
		<-timer.C
		g.mutex.Lock()
		defer g.mutex.Unlock()
		g.nextScene = "menu"
	}()
	return g
}

// Init initializes the game
func (g *Game) Init() error {
	return nil
}

// Render renders the game
func (g *Game) Render() error {
	g.renderer.SetDrawColor(255, 255, 0, 128)
	g.renderer.Clear()

	g.renderer.SetDrawColor(0, 255, 0, 255)
	g.renderer.FillRect(&sdl.Rect{X: 20, Y: 20, W: int32(g.width) - 40, H: int32(g.height) - 40})
	return nil
}

// HandleEvent handles game events
func (g *Game) HandleEvent(event *sdl.Event) {

}

// Update updates the game logic
func (g *Game) Update() string {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	return g.nextScene
}

// Destroy releases resorces
func (g *Game) Destroy() {
}
