package main

import (
	"log"
	"os"

	"gioui.org/app" // app contains Window handling.
	// gofont is used for loading the default font.
	// key is used for keyboard events.
	// system is used for system events (e.g. closing the window).
	// layout is used for layouting widgets.
	// op is used for recording different operations.
	"gioui.org/unit" // unit is used to define pixel-independent sizes
	// widget contains state handling for widgets.
	// material contains material design widgets.
)

func main() {
	// The ui loop is separated from the application window creation
	// such that it can be used for testing.
	ui := NewUI()

	// This creates a new application window and starts the UI.
	go func() {
		w := app.NewWindow(
			app.Title("Goi UI MDI test"),
			app.Size(unit.Dp(800), unit.Dp(600)),
		)
		if err := ui.Run(w); err != nil {
			log.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}()

	// This starts Gio main.
	app.Main()
}
