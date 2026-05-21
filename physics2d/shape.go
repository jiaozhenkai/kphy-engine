package physics2d

import "kphy-engine/math"

// ShapeType 形状类型
type ShapeType int

const (
	ShapeCircle  ShapeType = iota // 圆形
	ShapePolygon                  // 多边形
	ShapeBox                      // 矩形（多边形的优化版本）
	ShapeCapsule                  // 胶囊体
)

// Shape 形状接口
type Shape interface {
	GetType() ShapeType
	ComputeMass(density float32) float32
	ComputeInertia(density float32) float32
	GetAABB(body *RigidBody) math.AABB
	GetRadius() float32
}

// CircleShape 圆形形状
type CircleShape struct {
	Radius float32     // 半径：圆形的半径
	Offset math.Vector2 // 偏移：形状相对于刚体中心的偏移
}

// PolygonShape 多边形形状
type PolygonShape struct {
	Vertices []math.Vector2 // 顶点：多边形的顶点列表
	Normals  []math.Vector2 // 法线：每条边的法线向量
}

// BoxShape 矩形形状（优化的多边形）
type BoxShape struct {
	HalfExtents math.Vector2 // 半长宽：矩形的半宽和半高
	Offset      math.Vector2 // 偏移：形状相对于刚体中心的偏移
}

// CapsuleShape 胶囊体形状
type CapsuleShape struct {
	Radius float32     // 半径：两端圆的半径
	Height float32     // 高度：胶囊体的总高度
	Offset math.Vector2 // 偏移：形状相对于刚体中心的偏移
}
