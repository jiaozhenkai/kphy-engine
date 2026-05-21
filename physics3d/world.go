package physics3d

import "kphy-engine/math"

// Contact3D 3D碰撞接触点信息（占位定义）
type Contact3D struct {
	// 暂时为空，后面完善
}

// World3D 3D物理世界（预留）
type World3D struct {
	Bodies      []*RigidBody3D      // 物体列表：3D世界中所有的刚体
	Gravity     math.Vector3       // 重力：3D重力加速度向量
	Solver      Solver3D           // 解算器：3D物理约束解算器
	Collision   CollisionDetector3D // 碰撞检测器：3D碰撞检测
	BodyCount   int                // 物体计数：已创建物体数量
	SubSteps    int                // 子步数：每帧分为几个子步模拟
	IsPaused    bool               // 暂停状态：true表示世界暂停
	OnCollision func(bodyA, bodyB *RigidBody3D, contact *Contact3D) // 碰撞回调
}

// WorldSettings3D 3D物理世界配置
type WorldSettings3D struct {
	Gravity      math.Vector3 // 重力：3D重力向量
	SubSteps     int         // 子步数：每帧子步数，越多越精确但越慢
	PositionIter int         // 位置解算迭代次数
	VelocityIter int         // 速度解算迭代次数
}

// Solver3D 3D解算器接口
type Solver3D interface {
	Solve(bodies []*RigidBody3D, contacts []*Contact3D, dt float32)
}

// CollisionDetector3D 3D碰撞检测接口
type CollisionDetector3D interface {
	DetectCollisions(bodies []*RigidBody3D) []*Contact3D
	TestOverlap(bodyA, bodyB *RigidBody3D) bool
}
