package main

import (
	"time"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
)

const (
	BlockSize   = 20
	FieldHeight = 20
	FieldWidth  = 10
	TetroSize   = 4

	// Window size
	W = BlockSize * FieldWidth
	H = BlockSize * FieldHeight

	TimerPeriod = 250 // milliseconds
)

var (
	x = 4
	y = 6
)

var window *glfw.Window

func main() {
	// Create the window
	var err error
	err = glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()
	window, err = glfw.CreateWindow(W, H, "squares", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	window.SetKeyCallback(keyPressed)

	if err := gl.Init(); err != nil {
		panic(err)
	}
	// Timer
	ticker := time.NewTicker(time.Millisecond * TimerPeriod)
	go func() {
		for range ticker.C {
			//fmt.Println("tick ", t)
			update()
		}
	}()
	// Init OpenGL
	gl.Ortho(0, W, H, 0, -1, 1)
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(255, 255, 255, 0)
	gl.LineWidth(1)
	gl.Color3f(1, 0, 0)
	for !window.ShouldClose() {
		draw()
		glfw.PollEvents()
	}
}

func square(i, j int) {
	j-- // center
	i--
	gl.Begin(gl.POLYGON)
	gl.Vertex2i(int32(j*BlockSize), int32(i*BlockSize))
	gl.Vertex2i(int32((j+1)*BlockSize-1), int32(i*BlockSize))
	gl.Vertex2i(int32((j+1)*BlockSize-1), int32((i+1)*BlockSize-1))
	gl.Vertex2i(int32(j*BlockSize), int32((i+1)*BlockSize-1))
	gl.End()
}

func update() (x int, y int) {
	return
}

func draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.Color3ub(192, 192, 192)
	// square(1, 6)
	// square(2, 6)
	// square(3, 6)
	square(x, y)
	// square(5, 6)
	// square(30, 6)
	window.SwapBuffers()
}

func keyPressed(w *glfw.Window, k glfw.Key, s int, act glfw.Action, mods glfw.ModifierKey) {
	if act != glfw.Press {
		return
	}
	// Rotation
	switch k {
	case glfw.KeyUp:
		x--
	case glfw.KeyDown:
		x++
	case glfw.KeyLeft:
		y--
	case glfw.KeyRight:
		y++
	}
}
