package game

import (
	"fmt"
	"sync"
	"time"

	"github.com/veandco/go-sdl2/sdl"
	image "github.com/veandco/go-sdl2/sdl_image"
)

// Game is the main scene
type Game struct {
	renderer  *sdl.Renderer
	width     int
	height    int
	mutex     *sync.RWMutex
	nextScene string
	grass     *sdl.Texture
	grassX    int
}

// New creates the new game
func New(renderer *sdl.Renderer, width int, height int) *Game {
	g := &Game{renderer: renderer, width: width, height: height, mutex: &sync.RWMutex{}}
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
	if mask := image.Init(image.INIT_PNG); mask == 0 {
		return fmt.Errorf("PNG format is not supported")
	}
	defer image.Quit()

	var err error
	if g.grass, err = image.LoadTexture(g.renderer, "res/grass.png"); err != nil {
		return fmt.Errorf("Could not load grass resource: %v", err)
	}

	return nil
}

// Render renders the game
func (g *Game) Render() error {
	g.renderer.SetDrawColor(0, 0, 255, 128)
	g.renderer.Clear()

	const grassHeight = 100 / 2
	const grassWidth = 480 / 2

	for x := g.grassX; x < g.width; x += grassWidth {
		g.renderer.Copy(g.grass, nil, &sdl.Rect{X: int32(x), Y: int32(g.height) - grassHeight, W: 322, H: grassHeight})
	}
	return nil
}

// HandleEvent handles game events
func (g *Game) HandleEvent(event *sdl.Event) {

}

// Update updates the game logic
func (g *Game) Update() string {
	g.grassX--
	if g.grassX < -g.width {
		g.grassX = 0
	}

	g.mutex.RLock()
	defer g.mutex.RUnlock()
	return g.nextScene
}

// Destroy releases resorces
func (g *Game) Destroy() {
	g.grass.Destroy()
}
