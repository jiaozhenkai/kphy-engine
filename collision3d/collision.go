package collision3d

import (
	"kphy-engine/math"
	"kphy-engine/physics3d"
)

// CollisionDetector3D 3D碰撞检测器
type CollisionDetector3D struct {
	BroadPhase BroadPhase3D // 宽相位检测
}

// BroadPhase3D 3D宽相位检测接口
type BroadPhase3D interface {
	InsertBody(body *physics3d.RigidBody3D)
	RemoveBody(body *physics3d.RigidBody3D)
	QueryOverlaps(bounds math.Vector3) []*physics3d.RigidBody3D // 占位
	Update()
}

// NarrowPhase3D 3D窄相位检测接口
type NarrowPhase3D interface {
	TestCollision(bodyA, bodyB *physics3d.RigidBody3D) *ContactManifold3D
}

// GJKDetector3D 3D GJK碰撞检测器
type GJKDetector3D struct{}

// SATDetector3D 3D SAT碰撞检测器（凸多面体）
type SATDetector3D struct{}
