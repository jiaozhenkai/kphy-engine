package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/event"
	"gioui.org/io/key"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"

	"kphy-engine/ui"
)

func main() {
	go func() {
		window := new(app.Window)
		window.Option(
			app.Title("Physics Engine - Gio UI"),
			app.Size(unit.Dp(1100), unit.Dp(750)),
			app.MinSize(unit.Dp(500), unit.Dp(400)),
		)
		if err := run(window); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(window *app.Window) error {
	th := material.NewTheme()
	th.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))
	state := ui.NewAppState()
	var ops op.Ops

	for {
		switch e := window.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			handleKeys(gtx, state)
			handleScroll(gtx, state)
			ui.DrawUI(gtx, th, state)
			e.Frame(gtx.Ops)
		}
	}
}

func handleScroll(gtx layout.Context, state *ui.AppState) {
	// 注册事件目标
	event.Op(gtx.Ops, state)

	// 处理滚动事件
	for {
		ev, ok := gtx.Event(pointer.Filter{
			Target:  state,
			Kinds:   pointer.Scroll,
			ScrollX: pointer.ScrollRange{Min: -10, Max: 10},
			ScrollY: pointer.ScrollRange{Min: -10, Max: 10},
		})
		if !ok {
			break
		}
		if se, ok := ev.(pointer.Event); ok && se.Kind == pointer.Scroll {
			if state.Is3D {
				// 3D模式：控制相机距离
				if se.Scroll.Y < 0 {
					state.Camera.Dist *= 0.92
					if state.Camera.Dist < 3.0 {
						state.Camera.Dist = 3.0
					}
				} else if se.Scroll.Y > 0 {
					state.Camera.Dist *= 1.08
					if state.Camera.Dist > 20.0 {
						state.Camera.Dist = 20.0
					}
				}
			} else {
				// 2D模式：控制缩放
				if se.Scroll.Y < 0 {
					state.Zoom *= 1.08
					if state.Zoom > 4.0 {
						state.Zoom = 4.0
					}
				} else if se.Scroll.Y > 0 {
					state.Zoom /= 1.08
					if state.Zoom < 0.2 {
						state.Zoom = 0.2
					}
				}
			}
		}
	}
}

func handleKeys(gtx layout.Context, state *ui.AppState) {
	for {
		ev, ok := gtx.Event(
			key.Filter{Name: key.Name("+")},
			key.Filter{Name: key.Name("=")},
			key.Filter{Name: key.Name("-")},
			key.Filter{Name: key.Name(key.NameDeleteBackward)},
			key.Filter{Name: key.Name("Q")},
			key.Filter{Name: key.Name("E")},
			key.Filter{Name: key.Name("W")},
			key.Filter{Name: key.Name("S")},
		)
		if !ok {
			break
		}
		if ke, ok := ev.(key.Event); ok && ke.State == key.Press {
			switch ke.Name {
			case "+", "=":
				state.Zoom *= 1.1
				if state.Zoom > 3.0 {
					state.Zoom = 3.0
				}
			case "-":
				state.Zoom /= 1.1
				if state.Zoom < 0.3 {
					state.Zoom = 0.3
				}
			case key.NameDeleteBackward:
				state.DeleteLastObject()
			case "Q":
				state.Camera.Yaw -= 0.1
			case "E":
				state.Camera.Yaw += 0.1
			case "W":
				state.Camera.Pitch -= 0.05
				if state.Camera.Pitch < -0.2 {
					state.Camera.Pitch = -0.2
				}
			case "S":
				if state.Is3D {
					state.Camera.Pitch += 0.05
					if state.Camera.Pitch > 1.4 {
						state.Camera.Pitch = 1.4
					}
				}
			}
		}
	}
}
