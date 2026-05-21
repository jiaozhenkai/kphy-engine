package solver3d

import (
	"kphy-engine/math"
	"kphy-engine/physics3d"
)

// SolverConfig3D 3D解算器配置
type SolverConfig3D struct {
	PositionIterations int  // 位置解算迭代次数
	VelocityIterations int  // 速度解算迭代次数
	EnableSleep        bool // 是否允许物体休眠
	EnableWarmStart    bool // 是否启用热启动（冲量缓存）
}

// Solver3D 3D约束解算器
type Solver3D struct {
	Config SolverConfig3D // 解算器配置
}

// Constraint3D 3D约束接口
type Constraint3D interface {
	Prepare(dt float32)
	ApplyImpulse()
	ApplyPositionCorrection()
}

// ContactConstraint3D 3D接触约束
type ContactConstraint3D struct {
	BodyA        *physics3d.RigidBody3D
	BodyB        *physics3d.RigidBody3D
	Contact      *physics3d.Contact3D
	NormalMass   [3][3]float32 // 3D质量矩阵
	TangentMass  [3][3]float32 // 3D切向质量矩阵
	Bias         float32
}

// DistanceConstraint3D 3D距离约束
type DistanceConstraint3D struct {
	BodyA    *physics3d.RigidBody3D
	BodyB    *physics3d.RigidBody3D
	LocalA   math.Vector3
	LocalB   math.Vector3
	Distance float32
	Stiffness float32
}

// BallSocketConstraint 球铰约束（3D）
type BallSocketConstraint struct {
	BodyA  *physics3d.RigidBody3D
	BodyB  *physics3d.RigidBody3D
	Anchor math.Vector3
}

// HingeConstraint3D 3D铰链约束
type HingeConstraint3D struct {
	BodyA  *physics3d.RigidBody3D
	BodyB  *physics3d.RigidBody3D
	Anchor math.Vector3
	Axis   math.Vector3 // 旋转轴
}
