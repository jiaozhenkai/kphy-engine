package physics2d

import "kphy-engine/math"

// BodyType 刚体类型
type BodyType int

const (
	BodyStatic    BodyType = iota // 静态物体：不受力，不动，质量无穷大
	BodyDynamic                   // 动态物体：受物理影响，有质量
	BodyKinematic                 // 运动学物体：不受力，但可以移动，质量无穷大
)

// RigidBody 2D刚体
type RigidBody struct {
	ID int // 刚体唯一标识ID

	Type BodyType // 刚体类型：静态/动态/运动学

	Position math.Vector2 // 位置：物体在世界坐标系中的坐标
	Angle    float32      // 角度：物体的旋转角度（弧度）

	LinearVelocity  math.Vector2 // 线速度：物体沿X/Y轴的移动速度
	AngularVelocity float32      // 角速度：物体的旋转速度（弧度/秒）

	Mass    float32 // 质量：物体的质量（千克）
	InvMass float32 // 质量倒数：1/Mass，0表示质量无穷大（静态物体）

	Inertia    float32 // 转动惯量：物体抗旋转的属性
	InvInertia float32 // 转动惯量倒数：1/Inertia，0表示转动惯量无穷大

	Force  math.Vector2 // 外力：作用在物体上的合力
	Torque float32      // 力矩：作用在物体上的合扭矩

	Material Material // 材质：物体的物理材质属性（密度、弹性、摩擦）

	LinearDamping  float32 // 线阻尼：物体的线性速度衰减（空气阻力）
	AngularDamping float32 // 角阻尼：物体的角速度衰减

	GravityScale float32 // 重力缩放：物体受重力影响的倍数（0不受重力，1正常）

	IsSensor       bool // 是否是传感器：true表示只检测碰撞，不产生物理响应
	CollisionGroup int  // 碰撞组：物体所属的碰撞分组
	CollisionMask  int  // 碰撞掩码：与哪些组的物体碰撞

	UserData interface{} // 用户数据：可存储任意自定义数据

	Shapes []Shape // 形状：物体拥有的碰撞形状
}
