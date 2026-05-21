package solver2d

import (
	"kphy-engine/math"
	"kphy-engine/physics2d"
)

// SolverConfig 解算器配置
type SolverConfig struct {
	PositionIterations int  // 位置解算迭代次数
	VelocityIterations int  // 速度解算迭代次数
	EnableSleep        bool // 是否允许物体休眠
	EnableWarmStart    bool // 是否启用热启动（冲量缓存）
}

// Solver 约束解算器
type Solver struct {
	Config SolverConfig // 解算器配置
}

// Constraint 约束接口
type Constraint interface {
	Prepare(dt float32)       // 准备约束
	ApplyImpulse()            // 应用速度冲量
	ApplyPositionCorrection() // 应用位置修正
}

// ContactConstraint 接触约束
// 处理碰撞产生的接触响应
type ContactConstraint struct {
	BodyA       *physics2d.RigidBody // 碰撞物体A
	BodyB       *physics2d.RigidBody // 碰撞物体B
	Contact     *physics2d.Contact   // 接触信息
	NormalMass  float32              // 法向质量矩阵
	TangentMass float32              // 切向质量矩阵
	Bias        float32              // 位置修正偏置
}

// DistanceConstraint 距离约束
// 保持两个物体之间的固定距离
type DistanceConstraint struct {
	BodyA     *physics2d.RigidBody // 物体A
	BodyB     *physics2d.RigidBody // 物体B
	LocalA    math.Vector2         // A上的锚点（局部坐标）
	LocalB    math.Vector2         // B上的锚点（局部坐标）
	Distance  float32              // 目标距离
	Stiffness float32              // 刚度系数
}

// HingeConstraint 铰链约束
// 约束两个物体围绕同一点旋转
type HingeConstraint struct {
	BodyA  *physics2d.RigidBody // 物体A
	BodyB  *physics2d.RigidBody // 物体B
	Anchor math.Vector2         // 铰链锚点
}
