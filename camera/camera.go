package camera

import "math"

// Camera3D 定义了3D相机
type Camera3D struct {
	Yaw   float32 // 围绕Y轴旋转的角度（偏航角）
	Pitch float32 // 围绕X轴旋转的角度（俯仰角）
	Dist  float32 // 到观察点的距离
}

// NewCamera3D 创建一个新的3D相机
func NewCamera3D() Camera3D {
	return Camera3D{
		Yaw:   0.4,
		Pitch: 0.45,
		Dist:  8.0,
	}
}

// Project3D 将3D世界坐标投影到2D屏幕坐标
// 现在Z轴向上，Y轴向里（代替原来的Z轴）
func (c *Camera3D) Project3D(x3d, y3d, z3d float32, canvasW, canvasH int) (float32, float32, float32) {
	// 首先绕Z轴旋转（偏航，因为现在Z向上）
	cosY := float32(math.Cos(float64(c.Yaw)))
	sinY := float32(math.Sin(float64(c.Yaw)))
	rx := x3d*cosY + y3d*sinY  // 注意这里现在使用y3d（原Z轴）
	ry := -x3d*sinY + y3d*cosY // 注意这里现在使用y3d（原Z轴）
	rz := z3d

	// 然后绕X轴旋转（俯仰）
	cosP := float32(math.Cos(float64(c.Pitch)))
	sinP := float32(math.Sin(float64(c.Pitch)))
	ry2 := ry*cosP - rz*sinP
	rz2 := ry*sinP + rz*cosP

	// 计算视深度
	viewDepth := ry2 + c.Dist // 现在ry2作为深度轴（代替原来的rz2）
	if viewDepth < 0.3 {
		viewDepth = 0.3
	}

	// 透视投影
	fov := float32(2.0)
	projX := rx * fov / viewDepth
	projY := rz2 * fov / viewDepth // 现在rz2作为Y轴（代替原来的ry2）

	// 映射到屏幕像素（画布中心）
	screenX := float32(canvasW)*0.5 + projX*float32(canvasW)*0.4
	screenY := float32(canvasH)*0.5 - projY*float32(canvasH)*0.4

	return screenX, screenY, viewDepth
}
