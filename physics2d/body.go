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

// NewRigidBody 创建2D刚体
func NewRigidBody(typ BodyType, position math.Vector2, mass float32) *RigidBody {
	invMass := float32(0)
	invInertia := float32(0)

	if typ == BodyDynamic && mass > 0 {
		invMass = 1 / mass
		invInertia = float32(0.1) // 先给个默认转动惯量倒数
	}

	return &RigidBody{
		Type:            typ,
		Position:        position,
		Angle:           0,
		LinearVelocity:  math.Vec2(0, 0),
		AngularVelocity: 0,
		Mass:            mass,
		InvMass:         invMass,
		Inertia:         10,
		InvInertia:      invInertia,
		Force:           math.Vec2(0, 0),
		Torque:          0,
		Material:        DefaultMaterial(),
		LinearDamping:   0.01,
		AngularDamping:  0.01,
		GravityScale:    1.0,
		IsSensor:        false,
		Shapes:          make([]Shape, 0),
	}
}

// SetMass 设置质量
func (b *RigidBody) SetMass(mass float32) {
	if b.Type == BodyStatic {
		return
	}
	b.Mass = mass
	if mass > 0 {
		b.InvMass = 1 / mass
	} else {
		b.InvMass = 0
	}
}

// ApplyForce 施加力到质心
func (b *RigidBody) ApplyForce(force math.Vector2) {
	if b.Type != BodyDynamic {
		return
	}
	b.Force = b.Force.Add(force)
}

// ApplyForceAtPoint 在世界坐标的某个点施加力
func (b *RigidBody) ApplyForceAtPoint(force math.Vector2, point math.Vector2) {
	if b.Type != BodyDynamic {
		return
	}

	b.Force = b.Force.Add(force)

	// 计算从质心到点的向量
	arm := point.Sub(b.Position)
	// 计算力矩
	b.Torque += arm.Cross(force)
}

// ApplyLinearImpulse 施加冲量到质心（直接改变速度）
func (b *RigidBody) ApplyLinearImpulse(impulse math.Vector2) {
	if b.Type != BodyDynamic {
		return
	}
	b.LinearVelocity = b.LinearVelocity.Add(impulse.Mul(b.InvMass))
}

// ApplyTorque 施加力矩
func (b *RigidBody) ApplyTorque(torque float32) {
	if b.Type != BodyDynamic {
		return
	}
	b.Torque += torque
}

// SetLinearVelocity 设置线速度
func (b *RigidBody) SetLinearVelocity(vel math.Vector2) {
	b.LinearVelocity = vel
}

// SetAngularVelocity 设置角速度
func (b *RigidBody) SetAngularVelocity(vel float32) {
	b.AngularVelocity = vel
}

// IntegrateLinear 积分线性速度（F = ma）
func (b *RigidBody) IntegrateLinear(dt float32) {
	if b.Type != BodyDynamic || b.InvMass <= 0 {
		return
	}

	// v += (F/m) * dt
	acceleration := b.Force.Mul(b.InvMass)
	b.LinearVelocity = b.LinearVelocity.Add(acceleration.Mul(dt))

	// 应用阻尼
	b.LinearVelocity = b.LinearVelocity.Mul(1 - b.LinearDamping*dt)
}

// IntegratePosition 积分位置（x += v * dt）
func (b *RigidBody) IntegratePosition(dt float32) {
	if b.Type != BodyDynamic {
		return
	}

	b.Position = b.Position.Add(b.LinearVelocity.Mul(dt))
}

// IntegrateAngular 积分角速度（τ = Iα）
func (b *RigidBody) IntegrateAngular(dt float32) {
	if b.Type != BodyDynamic || b.InvInertia <= 0 {
		return
	}

	// ω += (τ/I) * dt
	angularAcc := b.Torque * b.InvInertia
	b.AngularVelocity += angularAcc * dt

	// 应用阻尼
	b.AngularVelocity *= (1 - b.AngularDamping*dt)
}

// IntegrateRotation 积分角度（θ += ω * dt）
func (b *RigidBody) IntegrateRotation(dt float32) {
	if b.Type != BodyDynamic {
		return
	}

	b.Angle += b.AngularVelocity * dt
}
