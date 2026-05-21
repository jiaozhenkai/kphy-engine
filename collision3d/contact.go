// Package collision3d 实现 3D 碰撞检测类型。
//
// contact.go 存储 3D 碰撞检测的结果数据，是窄相位碰撞检测和物理约束解算之间的数据桥梁。
//
// 核心概念:
//   - Contact3D: 单个碰撞接触点，存储位置、法线、穿透深度等信息
//   - ContactManifold3D: 接触流形，接触点的集合
//
// 这些接触数据结构由窄相位碰撞检测填充，然后传递给解算器计算冲量、更新速度并修复重叠。
package collision3d

import (
	"kphy-engine/math"
	"kphy-engine/physics3d"
)

// Contact3D 3D碰撞接触点信息
type Contact3D struct {
	BodyA          *physics3d.RigidBody3D // 碰撞物体A
	BodyB          *physics3d.RigidBody3D // 碰撞物体B
	Point          math.Vector3           // 接触点：碰撞发生的位置
	Normal         math.Vector3           // 法线：碰撞法线（从A指向B）
	Depth          float32                // 穿透深度：物体相互穿透的距离
	Impulse        float32                // 总冲量：应用的总碰撞冲量
	NormalImpulse  float32                // 法向冲量：垂直于碰撞面的冲量（用于分离物体）
	TangentImpulse math.Vector3           // 切向冲量：平行于碰撞面的冲量（用于处理摩擦）
	MixFriction    float32                // 混合摩擦系数：两物体的摩擦系数混合值
	MixRestitution float32                // 混合弹性系数：两物体的弹性系数混合值
	IsTouching     bool                   // 是否接触：true表示物体正在接触
}

// ContactManifold3D 3D碰撞接触流形
type ContactManifold3D struct {
	BodyA    *physics3d.RigidBody3D // 碰撞物体A
	BodyB    *physics3d.RigidBody3D // 碰撞物体B
	Contacts []Contact3D            // 接触点列表
	Normal   math.Vector3           // 碰撞法线
}

// RayCastHit3D 3D射线击中信息
type RayCastHit3D struct {
	Body   *physics3d.RigidBody3D // 被击中的物体
	Point  math.Vector3           // 击中位置
	Normal math.Vector3           // 击中表面法线
	Depth  float32                // 穿透深度
	T      float32                // 射线参数：从起点到击中位置的距离比例(0-1)
}
