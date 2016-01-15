// Copyright 2016 Richard Hawkins
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Package app manages the main game loop.

package main

import (
	_ "image/png"
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/hurricanerix/transylvania/display"
	"github.com/hurricanerix/transylvania/events"
	"github.com/hurricanerix/transylvania/examples/splash/splash"
	"github.com/hurricanerix/transylvania/sprite"
)

const windowWidth = 800
const windowHeight = 600

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func main() {
	screen, err := display.SetMode(640, 480)
	if err != nil {
		log.Fatalln("failed to set display mode:", err)
	}

	splash.Main(screen)

	background, err := sprite.Load("face.png", 1)
	if err != nil {
		panic(err)
	}
	background.Bind(screen.Program)

	for running := true; running; {
		// TODO move this somewhere else (maybe a Clear method of display
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// TODO refector events to be cleaner
		if screen.Window.ShouldClose() {
			running = !screen.Window.ShouldClose()
		}

		for _, event := range events.Get() {
			if event.Action == glfw.Press && event.Key == glfw.KeyEscape {
				running = false
				event.Window.SetShouldClose(true)
			}
		}

		screen.Fill(200.0/256.0, 200/256.0, 200/256.0)
		background.Draw(0, 0)

		screen.Flip()

		// TODO refector events to be cleaner
		glfw.PollEvents()
	}

}
