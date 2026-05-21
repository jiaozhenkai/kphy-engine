package physics3d

import "kphy-engine/math"

// ShapeType3D 3D形状类型
type ShapeType3D int

const (
	ShapeSphere3D   ShapeType3D = iota // 球体
	ShapeBox3D                         // 立方体
	ShapeCapsule3D                     // 胶囊体
	ShapeCylinder3D                    // 圆柱体
	ShapeMesh3D                        // 网格体
)

// Shape3D 3D形状接口
type Shape3D interface {
	GetType() ShapeType3D
	ComputeMass(density float32) float32
	ComputeInertiaTensor(density float32) [3][3]float32
	GetBoundingBox() math.AABB // （占位，后面可能需要3D版本）
}

// SphereShape3D 3D球体形状
type SphereShape3D struct {
	Radius float32     // 半径：球体的半径
	Offset math.Vector3 // 偏移：形状相对于刚体中心的偏移
}

// BoxShape3D 3D立方体形状
type BoxShape3D struct {
	HalfExtents math.Vector3 // 半长宽高：立方体的半宽、半高、半深
	Offset      math.Vector3 // 偏移：形状相对于刚体中心的偏移
}

// CapsuleShape3D 3D胶囊体形状
type CapsuleShape3D struct {
	Radius float32     // 半径：两端半球的半径
	Height float32     // 高度：胶囊体的总高度
	Offset math.Vector3 // 偏移：形状相对于刚体中心的偏移
}

// CylinderShape3D 3D圆柱体形状
type CylinderShape3D struct {
	Radius float32     // 半径：圆柱体底面半径
	Height float32     // 高度：圆柱体的高度
	Offset math.Vector3 // 偏移：形状相对于刚体中心的偏移
}

// MeshShape3D 3D网格形状
type MeshShape3D struct {
	Vertices []math.Vector3 // 顶点：网格的顶点列表
	Triangles [][3]int      // 三角形：三角形索引
}
