package renderer

import (
	"image"
	"image/color"
	"math"

	"gioui.org/f32"
	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

// Renderer 实现了 DrawAPI 接口，提供具体的绘图功能
type Renderer struct {
	ctx layout.Context
}

// NewRenderer 创建一个新的 Renderer
func NewRenderer(ctx layout.Context) *Renderer {
	return &Renderer{ctx: ctx}
}

// Context 获取当前的 layout.Context
func (r *Renderer) Context() layout.Context {
	return r.ctx
}

// SetContext 设置 layout.Context
func (r *Renderer) SetContext(ctx layout.Context) {
	r.ctx = ctx
}

// DrawCircle 绘制圆形
func (r *Renderer) DrawCircle(cx, cy, radius float32, c color.NRGBA) {
	ri := int(radius)
	bounds := image.Rect(int(cx)-ri, int(cy)-ri, int(cx)+ri, int(cy)+ri)
	defer clip.Ellipse(bounds).Push(r.ctx.Ops).Pop()
	paint.Fill(r.ctx.Ops, c)
}

// DrawRect 绘制矩形
func (r *Renderer) DrawRect(x, y, w, h float32, c color.NRGBA) {
	rect := clip.Rect{
		Min: image.Pt(int(x), int(y)),
		Max: image.Pt(int(x+w), int(y+h)),
	}
	defer rect.Push(r.ctx.Ops).Pop()
	paint.Fill(r.ctx.Ops, c)
}

// DrawRotatedRect 绘制旋转矩形
func (r *Renderer) DrawRotatedRect(cx, cy, w, h, angle float32, c color.NRGBA) {
	halfW, halfH := w/2, h/2

	var p clip.Path
	p.Begin(r.ctx.Ops)

	// 计算四个角点
	cosA := float32(math.Cos(float64(angle)))
	sinA := float32(math.Sin(float64(angle)))

	// 四个角点相对于中心的位置
	points := []f32.Point{
		{X: -halfW, Y: -halfH},
		{X: halfW, Y: -halfH},
		{X: halfW, Y: halfH},
		{X: -halfW, Y: halfH},
	}

	// 旋转并移动到正确位置
	for i, point := range points {
		// 旋转
		rotatedX := point.X*cosA - point.Y*sinA
		rotatedY := point.X*sinA + point.Y*cosA

		// 移动到中心
		finalX := cx + rotatedX
		finalY := cy + rotatedY

		if i == 0 {
			p.MoveTo(f32.Pt(finalX, finalY))
		} else {
			p.LineTo(f32.Pt(finalX, finalY))
		}
	}
	p.Close()

	defer clip.Outline{Path: p.End()}.Op().Push(r.ctx.Ops).Pop()
	paint.Fill(r.ctx.Ops, c)
}

// DrawTriangle 绘制三角形
func (r *Renderer) DrawTriangle(cx, cy, size float32, c color.NRGBA) {
	half := size / 2
	var p clip.Path
	p.Begin(r.ctx.Ops)
	p.MoveTo(f32.Pt(cx, cy-half))
	p.LineTo(f32.Pt(cx+half, cy+half))
	p.LineTo(f32.Pt(cx-half, cy+half))
	p.Close()
	defer clip.Outline{Path: p.End()}.Op().Push(r.ctx.Ops).Pop()
	paint.Fill(r.ctx.Ops, c)
}

// DrawRotatedTriangle 绘制旋转三角形
func (r *Renderer) DrawRotatedTriangle(cx, cy, size float32, angle float32, c color.NRGBA) {
	half := size / 2

	var p clip.Path
	p.Begin(r.ctx.Ops)

	cosA := float32(math.Cos(float64(angle)))
	sinA := float32(math.Sin(float64(angle)))

	// 三个角点相对于中心的位置
	points := []f32.Point{
		{X: 0, Y: -half},
		{X: half, Y: half},
		{X: -half, Y: half},
	}

	for i, point := range points {
		// 旋转
		rotatedX := point.X*cosA - point.Y*sinA
		rotatedY := point.X*sinA + point.Y*cosA

		finalX := cx + rotatedX
		finalY := cy + rotatedY

		if i == 0 {
			p.MoveTo(f32.Pt(finalX, finalY))
		} else {
			p.LineTo(f32.Pt(finalX, finalY))
		}
	}
	p.Close()

	defer clip.Outline{Path: p.End()}.Op().Push(r.ctx.Ops).Pop()
	paint.Fill(r.ctx.Ops, c)
}

// DrawPolygon 绘制多边形
func (r *Renderer) DrawPolygon(cx, cy, radius float32, sides int, c color.NRGBA) {
	var p clip.Path
	p.Begin(r.ctx.Ops)
	for i := 0; i < sides; i++ {
		angle := float64(i)*2*math.Pi/float64(sides) - math.Pi/2
		px := cx + radius*float32(math.Cos(angle))
		py := cy + radius*float32(math.Sin(angle))
		if i == 0 {
			p.MoveTo(f32.Pt(px, py))
		} else {
			p.LineTo(f32.Pt(px, py))
		}
	}
	p.Close()
	defer clip.Outline{Path: p.End()}.Op().Push(r.ctx.Ops).Pop()
	paint.Fill(r.ctx.Ops, c)
}

// DrawRotatedPolygon 绘制旋转多边形
func (r *Renderer) DrawRotatedPolygon(cx, cy, radius float32, sides int, angle float32, c color.NRGBA) {
	var p clip.Path
	p.Begin(r.ctx.Ops)

	cosA := float32(math.Cos(float64(angle)))
	sinA := float32(math.Sin(float64(angle)))

	for i := 0; i < sides; i++ {
		// 多边形角点角度
		polyAngle := float64(i)*2*math.Pi/float64(sides) - math.Pi/2

		// 相对于中心的位置
		localX := radius * float32(math.Cos(polyAngle))
		localY := radius * float32(math.Sin(polyAngle))

		// 旋转
		rotatedX := localX*cosA - localY*sinA
		rotatedY := localX*sinA + localY*cosA

		// 移动到最终位置
		finalX := cx + rotatedX
		finalY := cy + rotatedY

		if i == 0 {
			p.MoveTo(f32.Pt(finalX, finalY))
		} else {
			p.LineTo(f32.Pt(finalX, finalY))
		}
	}
	p.Close()

	defer clip.Outline{Path: p.End()}.Op().Push(r.ctx.Ops).Pop()
	paint.Fill(r.ctx.Ops, c)
}

// DrawLine 绘制线段
func (r *Renderer) DrawLine(x1, y1, x2, y2, width float32, c color.NRGBA) {
	var p clip.Path
	p.Begin(r.ctx.Ops)
	p.MoveTo(f32.Pt(x1, y1))
	p.LineTo(f32.Pt(x2, y2))
	spec := clip.Stroke{Path: p.End(), Width: width}
	defer spec.Op().Push(r.ctx.Ops).Pop()
	paint.Fill(r.ctx.Ops, c)
}

// Draw3DCube 绘制带渐变效果的3D立方体
func (r *Renderer) Draw3DCube(cx, cy, size float32, c color.NRGBA) {
	half := size / 2
	off := size * 0.3

	// 前面添加柔和的垂直渐变
	r.drawSoftRectGradient(cx-half, cy-half, size, size, c)

	// 顶面（更亮，添加渐变）
	lighter := LerpColor(c, color.NRGBA{R: 255, G: 255, B: 255, A: c.A}, 0.4)
	var p clip.Path
	p.Begin(r.ctx.Ops)
	p.MoveTo(f32.Pt(cx-half, cy-half))
	p.LineTo(f32.Pt(cx-half+off, cy-half-off))
	p.LineTo(f32.Pt(cx+half+off, cy-half-off))
	p.LineTo(f32.Pt(cx+half, cy-half))
	p.Close()
	func() {
		defer clip.Outline{Path: p.End()}.Op().Push(r.ctx.Ops).Pop()
		paint.Fill(r.ctx.Ops, lighter)
	}()

	// 右面（更暗，添加渐变）
	darker := LerpColor(c, color.NRGBA{R: 0, G: 0, B: 0, A: c.A}, 0.5)
	var p2 clip.Path
	p2.Begin(r.ctx.Ops)
	p2.MoveTo(f32.Pt(cx+half, cy-half))
	p2.LineTo(f32.Pt(cx+half+off, cy-half-off))
	p2.LineTo(f32.Pt(cx+half+off, cy+half-off))
	p2.LineTo(f32.Pt(cx+half, cy+half))
	p2.Close()
	func() {
		defer clip.Outline{Path: p2.End()}.Op().Push(r.ctx.Ops).Pop()
		paint.Fill(r.ctx.Ops, darker)
	}()

	// 添加渐变边缘效果
	r.drawFadedEdge(cx, cy, half, off, c)
}

// 辅助函数：绘制带柔和渐变的矩形
func (r *Renderer) drawSoftRectGradient(x, y, w, h float32, base color.NRGBA) {
	steps := 20
	stepH := h / float32(steps)

	for i := 0; i < steps; i++ {
		t := float32(i) / float32(steps-1)
		// 平滑的非线性过渡
		tSmooth := t * t * (3 - 2*t)

		// 顶部略亮，底部略暗
		lightColor := LerpColor(base, color.NRGBA{R: 255, G: 255, B: 255, A: base.A}, 0.15)
		darkColor := LerpColor(base, color.NRGBA{R: 0, G: 0, B: 0, A: base.A}, 0.15)

		col := LerpColor(lightColor, darkColor, tSmooth)
		ry := y + stepH*float32(i)
		r.DrawRect(x, ry, w, stepH+1, col)
	}
}

// Draw3DSphere 绘制带渐变效果的3D球体
func (r *Renderer) Draw3DSphere(cx, cy, radius float32, c color.NRGBA) {
	// 使用多层渐变绘制球体，模拟平滑光照
	r.drawRadialGradient(cx, cy, radius, c)

	// 高光
	highlight := LerpColor(c, color.NRGBA{R: 255, G: 255, B: 255, A: 0xBB}, 0.7)
	hr := int(radius * 0.3)
	hx := int(cx - radius*0.3)
	hy := int(cy - radius*0.3)
	hbounds := image.Rect(hx-hr, hy-hr, hx+hr, hy+hr)
	func() {
		defer clip.Ellipse(hbounds).Push(r.ctx.Ops).Pop()
		paint.Fill(r.ctx.Ops, highlight)
	}()
}

// Draw3DCylinder 绘制带渐变效果的3D圆柱体
func (r *Renderer) Draw3DCylinder(cx, cy, size float32, c color.NRGBA) {
	w := size * 0.55
	h := size * 0.85

	// 主体（使用垂直渐变）
	r.drawVerticalGradient(cx-w/2, cy-h/2, w, h, c)

	// 顶部椭圆（更亮）
	lighter := LerpColor(c, color.NRGBA{R: 255, G: 255, B: 255, A: c.A}, 0.4)
	ellH := int(size * 0.15)
	topY := int(cy - h/2)
	topBounds := image.Rect(int(cx-w/2), topY-ellH, int(cx+w/2), topY+ellH)
	func() {
		defer clip.Ellipse(topBounds).Push(r.ctx.Ops).Pop()
		paint.Fill(r.ctx.Ops, lighter)
	}()

	// 底部椭圆（更暗）
	darker := LerpColor(c, color.NRGBA{R: 0, G: 0, B: 0, A: c.A}, 0.4)
	botY := int(cy + h/2)
	botBounds := image.Rect(int(cx-w/2), botY-ellH, int(cx+w/2), botY+ellH)
	func() {
		defer clip.Ellipse(botBounds).Push(r.ctx.Ops).Pop()
		paint.Fill(r.ctx.Ops, darker)
	}()
}

// 辅助函数：绘制径向渐变（模拟球体光照，更模糊的效果）
func (r *Renderer) drawRadialGradient(cx, cy, radius float32, base color.NRGBA) {
	steps := 40 // 增加层数，让渐变更平滑
	for i := steps; i >= 0; i-- {
		t := float32(i) / float32(steps)
		// 非线性分布，让边缘渐变更柔和
		tSmooth := t * t * (3 - 2*t)

		ri := int(radius * (0.3 + 0.7*tSmooth))
		if ri < 1 {
			continue
		}

		// 更复杂的光照模型，从边缘到中心：暗 -> 中等 -> 亮
		var col color.NRGBA
		if t < 0.5 {
			// 边缘到中间：从暗到基本色
			col = LerpColor(base, color.NRGBA{R: 0, G: 0, B: 0, A: base.A}, 0.7*(1-2*t))
		} else {
			// 中间到高亮区：从基本色到更亮的颜色
			highlightColor := color.NRGBA{
				R: ClampU8(int(base.R) + 60),
				G: ClampU8(int(base.G) + 60),
				B: ClampU8(int(base.B) + 60),
				A: base.A,
			}
			col = LerpColor(base, highlightColor, 0.5*(2*t-1))
		}

		// 调整alpha实现更好的层次感
		col.A = base.A

		bounds := image.Rect(int(cx)-ri, int(cy)-ri, int(cx)+ri, int(cy)+ri)
		func() {
			defer clip.Ellipse(bounds).Push(r.ctx.Ops).Pop()
			paint.Fill(r.ctx.Ops, col)
		}()
	}
}

// 辅助函数：绘制垂直渐变（更模糊的效果）
func (r *Renderer) drawVerticalGradient(x, y, w, h float32, base color.NRGBA) {
	steps := 30 // 增加层数
	stepH := h / float32(steps)
	for i := 0; i < steps; i++ {
		t := float32(i) / float32(steps-1)
		// 使用平滑的非线性过渡
		tSmooth := t * t * (3 - 2*t)

		// 顶部很亮，中间是基本色，底部很暗
		topColor := color.NRGBA{
			R: ClampU8(int(base.R) + 80),
			G: ClampU8(int(base.G) + 80),
			B: ClampU8(int(base.B) + 80),
			A: base.A,
		}
		bottomColor := color.NRGBA{
			R: ClampU8(int(float32(base.R) * 0.5)),
			G: ClampU8(int(float32(base.G) * 0.5)),
			B: ClampU8(int(float32(base.B) * 0.5)),
			A: base.A,
		}

		var col color.NRGBA
		if tSmooth < 0.5 {
			col = LerpColor(topColor, base, tSmooth*2)
		} else {
			col = LerpColor(base, bottomColor, (tSmooth-0.5)*2)
		}

		ry := y + stepH*float32(i)
		r.DrawRect(x, ry, w, stepH+2, col)
	}
}

// 辅助函数：绘制渐隐的边缘
func (r *Renderer) drawFadedEdge(cx, cy, half, off float32, base color.NRGBA) {
	// 这里可以添加更多的边缘渐变效果
	// 为了简化，暂时只做基础实现
}

// 辅助函数：颜色插值
func LerpColor(c1, c2 color.NRGBA, t float32) color.NRGBA {
	t = clampF(t, 0, 1)
	return color.NRGBA{
		R: clampU8(int(float32(c1.R)*(1-t) + float32(c2.R)*t)),
		G: clampU8(int(float32(c1.G)*(1-t) + float32(c2.G)*t)),
		B: clampU8(int(float32(c1.B)*(1-t) + float32(c2.B)*t)),
		A: clampU8(int(float32(c1.A)*(1-t) + float32(c2.A)*t)),
	}
}

// 辅助函数：限制uint8范围
func ClampU8(v int) uint8 {
	if v > 255 {
		return 255
	}
	if v < 0 {
		return 0
	}
	return uint8(v)
}

// 辅助函数：限制float32范围
func ClampF(v, minV, maxV float32) float32 {
	if v < minV {
		return minV
	}
	if v > maxV {
		return maxV
	}
	return v
}

// clampU8 保持小写版本用于内部使用
func clampU8(v int) uint8 {
	return ClampU8(v)
}

// clampF 保持小写版本用于内部使用
func clampF(v, minV, maxV float32) float32 {
	return ClampF(v, minV, maxV)
}
