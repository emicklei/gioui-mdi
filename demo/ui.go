package main

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/event"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget/material"
	mdi "github.com/emicklei/gioui-mdi"
)

// UI holds all of the application state.
type UI struct {
	// Theme is used to hold the fonts used throughout the application.
	Theme *material.Theme

	stack *mdi.MDIWindowStack
}

// NewUI creates a new UI using the Go Fonts.
func NewUI() *UI {
	ui := &UI{}
	ui.Theme = material.NewTheme(gofont.Collection())
	ui.stack = mdi.NewWindowStack(ui.Theme)
	return ui
}

// Run handles window events and renders the application.
func (ui *UI) Run(w *app.Window) error {
	var ops op.Ops

	// listen for events happening on the window.
	for e := range w.Events() {
		// detect the type of the event.
		switch e := e.(type) {
		// this is sent when the application should re-render.
		case system.FrameEvent:
			// Reset the operations back to zero.
			ops.Reset()

			// gtx is used to pass around rendering and event information.
			gtx := layout.NewContext(&ops, e)

			ui.handleDragging(&ops, gtx.Queue)
			ui.stack.HandleDragging(&ops, gtx.Queue)

			// render and handle UI.
			ui.Layout(gtx)
			// render and handle the operations from the UI.
			e.Frame(gtx.Ops)

		// this is sent when the application is closed.
		case system.DestroyEvent:
			return e.Err
		}
	}

	return nil
}

// Layout displays the main program layout.
func (ui *UI) Layout(gtx layout.Context) layout.Dimensions {
	//return ui.test.Layout(ui.Theme, gtx)
	return ui.stack.Layout(ui.Theme, gtx)
}

func (ui *UI) handleDragging(ops *op.Ops, q event.Queue) {
	// Process events that arrived between the last frame and this one.
	ui.stack.HandleDragging(ops, q)
}
