package physics2d

import "kphy-engine/math"

// Contact 碰撞接触点信息（占位定义）
type Contact struct {
	// 暂时为空，后面从collision2d导入
}

// World 2D物理世界
type World struct {
	Bodies      []*RigidBody                                    // 物体列表：世界中所有的刚体
	Gravity     math.Vector2                                    // 重力：世界的重力加速度向量
	Solver      Solver                                          // 解算器：物理约束解算器
	Collision   CollisionDetector                               // 碰撞检测器
	BodyCount   int                                             // 物体计数：已创建物体数量
	SubSteps    int                                             // 子步数：每帧分为几个子步模拟
	IsPaused    bool                                            // 暂停状态：true表示世界暂停
	OnCollision func(bodyA, bodyB *RigidBody, contact *Contact) // 碰撞回调
}

// WorldSettings 物理世界配置
type WorldSettings struct {
	Gravity      math.Vector2 // 重力：重力向量
	SubSteps     int          // 子步数：每帧子步数，越多越精确但越慢
	PositionIter int          // 位置解算迭代次数
	VelocityIter int          // 速度解算迭代次数
}

// Solver 解算器接口
type Solver interface {
	Solve(bodies []*RigidBody, contacts []*Contact, dt float32)
}

// CollisionDetector 碰撞检测接口
type CollisionDetector interface {
	DetectCollisions(bodies []*RigidBody) []*Contact
	TestOverlap(bodyA, bodyB *RigidBody) bool
}
