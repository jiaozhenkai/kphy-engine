package shapes

import "image/color"

// ShapeKind 定义了可用的形状类型
type ShapeKind int

const (
	ShapeCircle ShapeKind = iota
	ShapeRect
	ShapeTriangle
	ShapePentagon
	ShapeCube
	ShapeSphere3D
	ShapeCylinder
)

func (s ShapeKind) String() string {
	switch s {
	case ShapeCircle:
		return "Circle"
	case ShapeRect:
		return "Rectangle"
	case ShapeTriangle:
		return "Triangle"
	case ShapePentagon:
		return "Pentagon"
	case ShapeCube:
		return "Cube"
	case ShapeSphere3D:
		return "Sphere"
	case ShapeCylinder:
		return "Cylinder"
	}
	return "Unknown"
}

// SceneObject 表示场景中的一个对象
type SceneObject struct {
	Kind  ShapeKind
	X, Y  float32
	Z     float32
	Size  float32
	Color color.NRGBA
}

// ColorPalette 提供默认的颜色调色板
var ColorPalette = []color.NRGBA{
	{R: 0x42, G: 0xA5, B: 0xF5, A: 0xFF},
	{R: 0xEF, G: 0x53, B: 0x50, A: 0xFF},
	{R: 0x66, G: 0xBB, B: 0x6A, A: 0xFF},
	{R: 0xFF, G: 0xCA, B: 0x28, A: 0xFF},
	{R: 0xAB, G: 0x47, B: 0xBC, A: 0xFF},
	{R: 0xFF, G: 0x70, B: 0x43, A: 0xFF},
}
