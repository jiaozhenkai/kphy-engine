package shapes

import (
	"gioui.org/layout"
	"image/color"
)

// DrawAPI 定义了绘图API接口，供物理引擎调用
type DrawAPI interface {
	// DrawCircle 绘制圆形
	DrawCircle(cx, cy, radius float32, c color.NRGBA)

	// DrawRect 绘制矩形
	DrawRect(x, y, w, h float32, c color.NRGBA)

	// DrawTriangle 绘制三角形
	DrawTriangle(cx, cy, size float32, c color.NRGBA)

	// DrawPolygon 绘制多边形
	DrawPolygon(cx, cy, radius float32, sides int, c color.NRGBA)

	// DrawLine 绘制线段
	DrawLine(x1, y1, x2, y2, width float32, c color.NRGBA)

	// Draw3DCube 绘制3D立方体
	Draw3DCube(cx, cy, size float32, c color.NRGBA)

	// Draw3DSphere 绘制3D球体
	Draw3DSphere(cx, cy, radius float32, c color.NRGBA)

	// Draw3DCylinder 绘制3D圆柱体
	Draw3DCylinder(cx, cy, size float32, c color.NRGBA)

	// Context 获取当前的layout.Context（内部使用）
	Context() layout.Context

	// SetContext 设置layout.Context
	SetContext(ctx layout.Context)
}
