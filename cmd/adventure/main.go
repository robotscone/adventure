package main

import (
	"fmt"
	"runtime"

	"github.com/robotscone/adventure/internal/input"
	"github.com/veandco/go-sdl2/sdl"
)

func init() {
	// Ensure that the main function runs on the main thread
	// This will prevent any crashes where certain SDL2 functions expect to be
	// on the main thread
	runtime.LockOSThread()
}

func main() {
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Adventure", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)

	rect := sdl.Rect{X: 0, Y: 0, W: 200, H: 200}
	surface.FillRect(&rect, 0xFFFFFFFF)

	window.UpdateSurface()

	var quit bool
	for !quit {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				quit = true
			}
		}

		input.Update()
	}
}
