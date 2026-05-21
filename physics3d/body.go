package physics3d

import "kphy-engine/math"

// BodyType3D 3D刚体类型
type BodyType3D int

const (
	BodyStatic3D    BodyType3D = iota // 3D静态物体
	BodyDynamic3D                      // 3D动态物体
	BodyKinematic3D                    // 3D运动学物体
)

// RigidBody3D 3D刚体
type RigidBody3D struct {
	ID              int              // 刚体唯一标识ID
	Type            BodyType3D       // 刚体类型
	Position        math.Vector3     // 位置：3D世界坐标
	Orientation     math.Vector4     // 朝向：四元数（预留）
	LinearVelocity  math.Vector3     // 线速度：3D速度
	AngularVelocity math.Vector3     // 角速度：3D角速度
	Mass            float32          // 质量
	InvMass         float32          // 质量倒数：0表示无穷大
	InertiaTensor   [3][3]float32    // 转动惯量张量
	InvInertiaTensor [3][3]float32   // 转动惯量张量倒数
	Force           math.Vector3     // 外力：3D力
	Torque          math.Vector3     // 力矩：3D扭矩
	Material        Material3D       // 3D材质
	LinearDamping   float32          // 线阻尼
	AngularDamping  float32          // 角阻尼
	GravityScale    float32          // 重力缩放
	Shapes          []Shape3D        // 3D形状
}

// Material3D 3D物理材质
type Material3D struct {
	Density         float32 // 密度
	Restitution     float32 // 弹性系数
	StaticFriction  float32 // 静摩擦系数
	DynamicFriction float32 // 动摩擦系数
}
