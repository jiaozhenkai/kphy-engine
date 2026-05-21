package collision2d

import (
	"kphy-engine/math"
	"kphy-engine/physics2d"
)

// CollisionDetector 碰撞检测器实现
type CollisionDetector struct {
	BroadPhase BroadPhase // 宽相位检测
}

// BroadPhase 宽相位检测接口
// 宽相位：快速找出可能碰撞的物体对，减少窄相位检测的工作量
type BroadPhase interface {
	InsertBody(body *physics2d.RigidBody)    // 插入物体
	RemoveBody(body *physics2d.RigidBody)    // 移除物体
	QueryOverlaps(aabb math.AABB) []*physics2d.RigidBody // 查询与AABB重叠的物体
	Update()                                  // 更新空间结构
}

// NarrowPhase 窄相位检测接口
// 窄相位：精确检测两个物体是否碰撞，计算接触信息
type NarrowPhase interface {
	TestCollision(bodyA, bodyB *physics2d.RigidBody) *ContactManifold
}

// SATDetector SAT（分离轴定理）碰撞检测器
type SATDetector struct{}

// GJKDetector GJK（Gilbert-Johnson-Keerthi）碰撞检测器
type GJKDetector struct{}
