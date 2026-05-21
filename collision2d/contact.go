// Package collision2d 实现 2D 碰撞检测类型。
//
// contact.go 存储碰撞检测的结果数据，是窄相位碰撞检测和物理约束解算之间的数据桥梁。
//
// 核心概念:
//   - Contact: 单个碰撞接触点，存储位置、法线、穿透深度等信息
//   - ContactManifold: 接触流形，接触点的集合（例如矩形和矩形碰撞可能有2个接触点）
//
// 这些接触数据结构由窄相位碰撞检测填充，然后传递给解算器计算冲量、更新速度并修复重叠。
package collision2d

import (
	"kphy-engine/math"
	"kphy-engine/physics2d"
)

// Contact 碰撞接触点信息
type Contact struct {
	BodyA          *physics2d.RigidBody // 碰撞物体A
	BodyB          *physics2d.RigidBody // 碰撞物体B
	Point          math.Vector2         // 接触点：碰撞发生的位置
	Normal         math.Vector2         // 法线：碰撞法线（从A指向B）
	Depth          float32              // 穿透深度：物体相互穿透的距离
	Impulse        float32              // 总冲量：应用的总碰撞冲量
	NormalImpulse  float32              // 法向冲量：垂直于碰撞面的冲量（用于分离物体）
	TangentImpulse float32              // 切向冲量：平行于碰撞面的冲量（用于处理摩擦）
	MixFriction    float32              // 混合摩擦系数：两物体的摩擦系数混合值
	MixRestitution float32              // 混合弹性系数：两物体的弹性系数混合值
	IsTouching     bool                 // 是否接触：true表示物体正在接触
}

// ContactManifold 碰撞接触流形
type ContactManifold struct {
	BodyA    *physics2d.RigidBody // 碰撞物体A
	BodyB    *physics2d.RigidBody // 碰撞物体B
	Contacts []Contact            // 接触点列表
	Normal   math.Vector2         // 碰撞法线
}

// RayCastHit 射线击中信息
type RayCastHit struct {
	Body   *physics2d.RigidBody // 被击中的物体
	Point  math.Vector2         // 击中位置
	Normal math.Vector2         // 击中表面法线
	Depth  float32              // 穿透深度
	T      float32              // 射线参数：从起点到击中位置的距离比例(0-1)
}
