package ui

import (
	"fmt"
	"image"
	"image/color"

	"gioui.org/f32"
	"gioui.org/io/pointer"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"kphy-engine/camera"
	"kphy-engine/renderer"
	"kphy-engine/shapes"
)

// AppState 应用状态
type AppState struct {
	Is3D          bool
	Btn2D         widget.Clickable
	Btn3D         widget.Clickable
	BtnClear      widget.Clickable
	BtnDeleteLast widget.Clickable
	ShowDropdown  bool
	BtnDropdown   widget.Clickable
	ShapeButtons  [7]widget.Clickable
	Objects       []shapes.SceneObject
	Zoom          float32
	Camera        camera.Camera3D
	PlaceCount    int
}

// NewAppState 创建新的应用状态
func NewAppState() *AppState {
	return &AppState{
		Is3D:    false,
		Zoom:    1.0,
		Objects: make([]shapes.SceneObject, 0),
		Camera:  camera.NewCamera3D(),
	}
}

// AvailableShapes 返回当前模式下可用的形状
func (s *AppState) AvailableShapes() []shapes.ShapeKind {
	if s.Is3D {
		return []shapes.ShapeKind{shapes.ShapeCube, shapes.ShapeSphere3D, shapes.ShapeCylinder}
	}
	return []shapes.ShapeKind{shapes.ShapeCircle, shapes.ShapeRect, shapes.ShapeTriangle, shapes.ShapePentagon}
}

// AddShape 添加形状
func (s *AppState) AddShape(kind shapes.ShapeKind) {
	if s.Is3D {
		col := s.PlaceCount % 3
		dep := (s.PlaceCount / 3) % 3
		row := (s.PlaceCount / 9) % 3

		obj := shapes.SceneObject{
			Kind:  kind,
			X:     float32(col)*1.2 - 1.2,
			Y:     float32(dep)*1.2 - 0.6, // Y轴现在是深度（向里）
			Z:     float32(row)*1.0 + 0.5, // Z轴现在是向上
			Size:  0.3,
			Color: shapes.ColorPalette[s.PlaceCount%len(shapes.ColorPalette)],
		}
		s.Objects = append(s.Objects, obj)
	} else {
		col := s.PlaceCount % 4
		row := s.PlaceCount / 4
		x := 0.15 + float32(col)*0.2
		y := 0.15 + float32(row)*0.2
		if y > 0.85 {
			y = 0.15
			s.PlaceCount = 0
		}

		obj := shapes.SceneObject{
			Kind:  kind,
			X:     x,
			Y:     y,
			Z:     0,
			Size:  0.08,
			Color: shapes.ColorPalette[s.PlaceCount%len(shapes.ColorPalette)],
		}
		s.Objects = append(s.Objects, obj)
	}
	s.PlaceCount++
}

// DeleteLastObject 删除最后一个对象
func (s *AppState) DeleteLastObject() {
	if len(s.Objects) > 0 {
		s.Objects = s.Objects[:len(s.Objects)-1]
		if s.PlaceCount > 0 {
			s.PlaceCount--
		}
	}
}

// ClearAllObjects 清除所有对象
func (s *AppState) ClearAllObjects() {
	s.Objects = s.Objects[:0]
	s.PlaceCount = 0
}

// DrawUI 绘制主UI
func DrawUI(gtx layout.Context, th *material.Theme, state *AppState) layout.Dimensions {
	paint.Fill(gtx.Ops, color.NRGBA{R: 0x1A, G: 0x1A, B: 0x2E, A: 0xFF})

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return DrawToolbar(gtx, th, state)
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return DrawCanvas(gtx, th, state)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return DrawStatusBar(gtx, th, state)
		}),
	)
}

// DrawToolbar 绘制工具栏
func DrawToolbar(gtx layout.Context, th *material.Theme, state *AppState) layout.Dimensions {
	return layout.Background{}.Layout(gtx,
		func(gtx layout.Context) layout.Dimensions {
			size := image.Pt(gtx.Constraints.Max.X, gtx.Dp(unit.Dp(52)))
			defer clip.Rect{Max: size}.Push(gtx.Ops).Pop()
			paint.Fill(gtx.Ops, color.NRGBA{R: 0x2A, G: 0x2A, B: 0x42, A: 0xFF})
			return layout.Dimensions{Size: size}
		},
		func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{
				Left: unit.Dp(10), Right: unit.Dp(10),
				Top: unit.Dp(8), Bottom: unit.Dp(8),
			}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{
					Axis:      layout.Horizontal,
					Alignment: layout.Middle,
				}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(th, &state.Btn2D, "2D")
						if !state.Is3D {
							btn.Background = color.NRGBA{R: 0x42, G: 0xA5, B: 0xF5, A: 0xFF}
						} else {
							btn.Background = color.NRGBA{R: 0x44, G: 0x44, B: 0x66, A: 0xFF}
						}
						if state.Btn2D.Clicked(gtx) {
							state.Is3D = false
							state.ShowDropdown = false
						}
						return layout.Inset{Right: unit.Dp(4)}.Layout(gtx, btn.Layout)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(th, &state.Btn3D, "3D")
						if state.Is3D {
							btn.Background = color.NRGBA{R: 0xAB, G: 0x47, B: 0xBC, A: 0xFF}
						} else {
							btn.Background = color.NRGBA{R: 0x44, G: 0x44, B: 0x66, A: 0xFF}
						}
						if state.Btn3D.Clicked(gtx) {
							state.Is3D = true
							state.ShowDropdown = false
						}
						return layout.Inset{Right: unit.Dp(12)}.Layout(gtx, btn.Layout)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return DrawSeparator(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						label := "Add Shape ▼"
						btn := material.Button(th, &state.BtnDropdown, label)
						btn.Background = color.NRGBA{R: 0x43, G: 0xA0, B: 0x47, A: 0xFF}
						if state.BtnDropdown.Clicked(gtx) {
							state.ShowDropdown = !state.ShowDropdown
						}
						return layout.Inset{Right: unit.Dp(12)}.Layout(gtx, btn.Layout)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return DrawSeparator(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(th, &state.BtnDeleteLast, "Delete Last")
						btn.Background = color.NRGBA{R: 0xE5, G: 0x39, B: 0x35, A: 0xFF}
						if state.BtnDeleteLast.Clicked(gtx) {
							state.DeleteLastObject()
						}
						return layout.Inset{Right: unit.Dp(6)}.Layout(gtx, btn.Layout)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						btn := material.Button(th, &state.BtnClear, "Clear All")
						btn.Background = color.NRGBA{R: 0xBF, G: 0x36, B: 0x0C, A: 0xFF}
						if state.BtnClear.Clicked(gtx) {
							state.ClearAllObjects()
						}
						return layout.Inset{Right: unit.Dp(12)}.Layout(gtx, btn.Layout)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return DrawSeparator(gtx)
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						txt := fmt.Sprintf("Zoom: %.0f%%", state.Zoom*100)
						lbl := material.Body2(th, txt)
						lbl.Color = color.NRGBA{R: 0xBB, G: 0xBB, B: 0xDD, A: 0xFF}
						return layout.Inset{Left: unit.Dp(4)}.Layout(gtx, lbl.Layout)
					}),
				)
			})
		},
	)
}

// DrawSeparator 绘制分隔符
func DrawSeparator(gtx layout.Context) layout.Dimensions {
	size := image.Pt(gtx.Dp(unit.Dp(1)), gtx.Dp(unit.Dp(28)))
	defer clip.Rect{Max: size}.Push(gtx.Ops).Pop()
	paint.Fill(gtx.Ops, color.NRGBA{R: 0x55, G: 0x55, B: 0x77, A: 0xFF})
	return layout.Inset{Left: unit.Dp(6), Right: unit.Dp(6)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Dimensions{Size: size}
	})
}

// DrawCanvas 绘制画布
func DrawCanvas(gtx layout.Context, th *material.Theme, state *AppState) layout.Dimensions {
	canvasW := gtx.Constraints.Max.X
	canvasH := gtx.Constraints.Max.Y

	defer clip.Rect{Max: image.Pt(canvasW, canvasH)}.Push(gtx.Ops).Pop()

	r := renderer.NewRenderer(gtx)

	if state.Is3D {
		Draw3DScene(gtx, th, state, canvasW, canvasH, r)
	} else {
		paint.Fill(gtx.Ops, color.NRGBA{R: 0x10, G: 0x10, B: 0x20, A: 0xFF})
		DrawGrid(gtx, state, canvasW, canvasH)
		for _, obj := range state.Objects {
			Draw2DObject(r, state, obj, canvasW, canvasH)
		}
	}

	if state.ShowDropdown {
		DrawDropdown(gtx, th, state)
	}

	return layout.Dimensions{Size: image.Pt(canvasW, canvasH)}
}

// DrawGrid 绘制网格
func DrawGrid(gtx layout.Context, state *AppState, w, h int) {
	gridColor := color.NRGBA{R: 0x28, G: 0x28, B: 0x40, A: 0x88}
	spacing := int(50.0 * state.Zoom)
	if spacing < 10 {
		spacing = 10
	}

	for x := 0; x < w; x += spacing {
		rect := clip.Rect{Min: image.Pt(x, 0), Max: image.Pt(x+1, h)}
		stack := rect.Push(gtx.Ops)
		paint.Fill(gtx.Ops, gridColor)
		stack.Pop()
	}
	for y := 0; y < h; y += spacing {
		rect := clip.Rect{Min: image.Pt(0, y), Max: image.Pt(w, y+1)}
		stack := rect.Push(gtx.Ops)
		paint.Fill(gtx.Ops, gridColor)
		stack.Pop()
	}
}

// Draw3DScene 绘制3D场景
func Draw3DScene(gtx layout.Context, th *material.Theme, state *AppState, cw, ch int, r *renderer.Renderer) {
	Draw3DBackground(gtx, cw, ch)
	Draw3DGrid(gtx, state, cw, ch)
	Draw3DAxis(gtx, state, cw, ch)

	type projectedObj struct {
		obj   shapes.SceneObject
		sx    float32
		sy    float32
		depth float32
		scale float32
	}

	projected := make([]projectedObj, 0, len(state.Objects))
	for _, obj := range state.Objects {
		sx, sy, depth := state.Camera.Project3D(obj.X, obj.Y, obj.Z, cw, ch)
		scale := state.Camera.Dist / depth * state.Zoom
		if scale < 0.1 {
			scale = 0.1
		}
		if scale > 3.0 {
			scale = 3.0
		}
		projected = append(projected, projectedObj{obj: obj, sx: sx, sy: sy, depth: depth, scale: scale})
	}

	// 排序：远处的先绘制
	// 注意：为避免循环依赖，这里暂不排序
	// 实际项目中可以考虑将排序逻辑移到独立包

	for _, po := range projected {
		Draw3DObjectProjected(r, po.obj, po.sx, po.sy, po.scale, float32(cw))
	}

	DrawCameraHUD(gtx, th, state, cw, ch)
}

// Draw3DBackground 绘制3D背景
func Draw3DBackground(gtx layout.Context, cw, ch int) {
	horizonY := ch * 45 / 100
	func() {
		sky := clip.Rect{Max: image.Pt(cw, horizonY)}
		defer sky.Push(gtx.Ops).Pop()
		paint.Fill(gtx.Ops, color.NRGBA{R: 0xA0, G: 0xA0, B: 0xA8, A: 0xFF})
	}()
	bandH := ch * 6 / 100
	func() {
		band := clip.Rect{Min: image.Pt(0, horizonY-bandH), Max: image.Pt(cw, horizonY)}
		defer band.Push(gtx.Ops).Pop()
		paint.Fill(gtx.Ops, color.NRGBA{R: 0xBB, G: 0xBB, B: 0xBF, A: 0xFF})
	}()
	func() {
		ground := clip.Rect{Min: image.Pt(0, horizonY), Max: image.Pt(cw, ch)}
		defer ground.Push(gtx.Ops).Pop()
		paint.Fill(gtx.Ops, color.NRGBA{R: 0x4A, G: 0x3F, B: 0x35, A: 0xFF})
	}()
}

// Draw3DGrid 绘制3D网格
func Draw3DGrid(gtx layout.Context, state *AppState, cw, ch int) {
	gridColor := color.NRGBA{R: 0x88, G: 0x88, B: 0x88, A: 0x66}
	gridLines := 21
	extent := float32(5.0)
	half := extent / 2

	// 现在网格在X-Y平面（Y向里，X向右），Z轴向上
	for i := 0; i < gridLines; i++ {
		t := float32(i)/float32(gridLines-1)*extent - half

		x1, y1, _ := state.Camera.Project3D(-half, t, 0, cw, ch)
		x2, y2, _ := state.Camera.Project3D(half, t, 0, cw, ch)
		StrokeLine(gtx, x1, y1, x2, y2, 1, gridColor)

		x1, y1, _ = state.Camera.Project3D(t, -half, 0, cw, ch)
		x2, y2, _ = state.Camera.Project3D(t, half, 0, cw, ch)
		StrokeLine(gtx, x1, y1, x2, y2, 1, gridColor)
	}
}

// Draw3DAxis 绘制3D坐标轴
func Draw3DAxis(gtx layout.Context, state *AppState, cw, ch int) {
	ox, oy, _ := state.Camera.Project3D(0, 0, 0, cw, ch)
	axisLen := float32(2.5)

	// X轴：向右（红色）
	ex, ey, _ := state.Camera.Project3D(axisLen, 0, 0, cw, ch)
	StrokeLine(gtx, ox, oy, ex, ey, 2.5, color.NRGBA{R: 0xFF, G: 0x3C, B: 0x50, A: 0xFF})

	// Y轴：向里（绿色）
	ex, ey, _ = state.Camera.Project3D(0, axisLen, 0, cw, ch)
	StrokeLine(gtx, ox, oy, ex, ey, 2.5, color.NRGBA{R: 0x7C, G: 0xC3, B: 0x2E, A: 0xFF})

	// Z轴：向上（蓝色）
	ex, ey, _ = state.Camera.Project3D(0, 0, axisLen, cw, ch)
	StrokeLine(gtx, ox, oy, ex, ey, 2.5, color.NRGBA{R: 0x38, G: 0x8A, B: 0xFF, A: 0xFF})
}

// Draw3DObjectProjected 绘制投影的3D对象
func Draw3DObjectProjected(r *renderer.Renderer, obj shapes.SceneObject, sx, sy, scale, canvasW float32) {
	size := canvasW * 0.06 * scale

	foggedColor := obj.Color
	fogFactor := scale
	if fogFactor > 1.0 {
		fogFactor = 1.0
	}
	foggedColor.A = uint8(float32(obj.Color.A) * renderer.ClampF(fogFactor, 0.3, 1.0))

	switch obj.Kind {
	case shapes.ShapeCube:
		r.Draw3DCube(sx, sy, size, foggedColor)
	case shapes.ShapeSphere3D:
		r.Draw3DSphere(sx, sy, size/2, foggedColor)
	case shapes.ShapeCylinder:
		r.Draw3DCylinder(sx, sy, size, foggedColor)
	}

	shadowAlpha := uint8(float32(0x55) * renderer.ClampF(fogFactor, 0.15, 0.7))
	shadowColor := color.NRGBA{A: shadowAlpha}
	shadowW := int(size * 0.7)
	shadowH := int(size * 0.15)
	shadowY := int(sy + size*0.45)
	bounds := image.Rect(int(sx)-shadowW/2, shadowY, int(sx)+shadowW/2, shadowY+shadowH)
	func() {
		defer clip.Ellipse(bounds).Push(r.Context().Ops).Pop()
		paint.Fill(r.Context().Ops, shadowColor)
	}()
}

// Draw2DObject 绘制2D对象
func Draw2DObject(r *renderer.Renderer, state *AppState, obj shapes.SceneObject, cw, ch int) {
	cx := obj.X * float32(cw)
	cy := float32(ch) - obj.Y*float32(ch) // 原点在左下角，向上为Y正
	size := obj.Size * float32(cw) * state.Zoom

	switch obj.Kind {
	case shapes.ShapeCircle:
		r.DrawCircle(cx, cy, size/2, obj.Color)
	case shapes.ShapeRect:
		r.DrawRect(cx-size/2, cy-size/2, size, size, obj.Color)
	case shapes.ShapeTriangle:
		r.DrawTriangle(cx, cy, size, obj.Color)
	case shapes.ShapePentagon:
		r.DrawPolygon(cx, cy, size/2, 5, obj.Color)
	}
}

// StrokeLine 绘制线条
func StrokeLine(gtx layout.Context, x1, y1, x2, y2 float32, width float32, c color.NRGBA) {
	var p clip.Path
	p.Begin(gtx.Ops)
	p.MoveTo(f32.Pt(x1, y1))
	p.LineTo(f32.Pt(x2, y2))
	spec := clip.Stroke{Path: p.End(), Width: width}
	defer spec.Op().Push(gtx.Ops).Pop()
	paint.Fill(gtx.Ops, c)
}

// DrawDropdown 绘制下拉菜单
func DrawDropdown(gtx layout.Context, th *material.Theme, state *AppState) {
	shapesList := state.AvailableShapes()
	startX := gtx.Dp(unit.Dp(200))
	startY := gtx.Dp(unit.Dp(4))
	itemW := gtx.Dp(unit.Dp(140))
	itemH := gtx.Dp(unit.Dp(36))
	totalH := itemH * len(shapesList)

	func() {
		shadow := clip.Rect{
			Min: image.Pt(startX+3, startY+3),
			Max: image.Pt(startX+itemW+3, startY+totalH+3),
		}
		defer shadow.Push(gtx.Ops).Pop()
		paint.Fill(gtx.Ops, color.NRGBA{A: 0x55})
	}()

	func() {
		bg := clip.Rect{
			Min: image.Pt(startX, startY),
			Max: image.Pt(startX+itemW, startY+totalH),
		}
		defer bg.Push(gtx.Ops).Pop()
		paint.Fill(gtx.Ops, color.NRGBA{R: 0x38, G: 0x38, B: 0x5A, A: 0xF8})
	}()

	for i, shape := range shapesList {
		idx := int(shape)
		iy := startY + itemH*i

		func() {
			defer op.Offset(image.Pt(startX, iy)).Push(gtx.Ops).Pop()
			localGtx := gtx
			localGtx.Constraints = layout.Exact(image.Pt(itemW, itemH))

			btn := material.Button(th, &state.ShapeButtons[idx], shape.String())
			btn.Background = color.NRGBA{R: 0x38, G: 0x38, B: 0x5A, A: 0x00}
			btn.Color = color.NRGBA{R: 0xEE, G: 0xEE, B: 0xEE, A: 0xFF}

			if state.ShapeButtons[idx].Clicked(localGtx) {
				state.AddShape(shape)
				state.ShowDropdown = false
			}
			btn.Layout(localGtx)
		}()
	}
}

// DrawCameraHUD 绘制相机HUD
func DrawCameraHUD(gtx layout.Context, th *material.Theme, state *AppState, cw, ch int) {
	defer op.Offset(image.Pt(cw-200, 10)).Push(gtx.Ops).Pop()
	localGtx := gtx
	localGtx.Constraints = layout.Exact(image.Pt(190, 80))

	legend := fmt.Sprintf("X (red)  = right\nY (green) = forward\nZ (blue) = up\nQ/E: Orbit  W/S: Tilt\nScroll: Camera Distance")
	lbl := material.Caption(th, legend)
	lbl.Color = color.NRGBA{R: 0x99, G: 0xCC, B: 0xEE, A: 0xCC}
	lbl.Layout(localGtx)
}

// DrawStatusBar 绘制状态栏
func DrawStatusBar(gtx layout.Context, th *material.Theme, state *AppState) layout.Dimensions {
	return layout.Background{}.Layout(gtx,
		func(gtx layout.Context) layout.Dimensions {
			size := image.Pt(gtx.Constraints.Max.X, gtx.Dp(unit.Dp(26)))
			defer clip.Rect{Max: size}.Push(gtx.Ops).Pop()
			paint.Fill(gtx.Ops, color.NRGBA{R: 0x1E, G: 0x1E, B: 0x36, A: 0xFF})
			return layout.Dimensions{Size: size}
		},
		func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{Left: unit.Dp(10), Top: unit.Dp(4)}.Layout(gtx,
				func(gtx layout.Context) layout.Dimensions {
					mode := "2D"
					extra := "Scroll/+/- zoom"
					if state.Is3D {
						mode = "3D"
						extra = "Scroll: camera distance | Q/E orbit | W/S tilt"
					}
					txt := fmt.Sprintf("Mode: %s | Objects: %d | Zoom: %.0f%% | %s | Backspace: delete",
						mode, len(state.Objects), state.Zoom*100, extra)
					lbl := material.Caption(th, txt)
					lbl.Color = color.NRGBA{R: 0x88, G: 0x88, B: 0xAA, A: 0xFF}
					return lbl.Layout(gtx)
				},
			)
		},
	)
}

// HandleScroll 处理滚动事件
func HandleScroll(gtx layout.Context, state *AppState) {
	// 注册事件目标
	// 注意：这里我们需要处理事件注册
	// 由于简化处理，我们暂时只处理2D模式下的滚动
	// 在main.go中保留了更完整的处理

	// 处理指针事件
	for {
		ev, ok := gtx.Event(pointer.Filter{
			Target: state,
			Kinds:  pointer.Scroll,
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

// HandleKeys 处理键盘事件
func HandleKeys(gtx layout.Context, state *AppState) {
	// 注意：这里简化了处理，实际项目中可能需要更完整的实现
}
