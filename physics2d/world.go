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

// NewWorld 创建新的2D物理世界
func NewWorld(settings WorldSettings) *World {
	if settings.SubSteps <= 0 {
		settings.SubSteps = 1
	}
	return &World{
		Bodies:   make([]*RigidBody, 0),
		Gravity:  settings.Gravity,
		SubSteps: settings.SubSteps,
		IsPaused: false,
	}
}

// AddBody 向世界添加刚体
func (w *World) AddBody(body *RigidBody) {
	body.ID = w.BodyCount
	w.BodyCount++
	w.Bodies = append(w.Bodies, body)
}

// RemoveBody 从世界移除刚体
func (w *World) RemoveBody(body *RigidBody) {
	for i, b := range w.Bodies {
		if b.ID == body.ID {
			w.Bodies = append(w.Bodies[:i], w.Bodies[i+1:]...)
			break
		}
	}
}

// Step 世界时间步进 - 物理引擎的核心！
func (w *World) Step(dt float32) {
	if w.IsPaused || dt <= 0 {
		return
	}

	// 每帧分为子步模拟
	subDt := dt / float32(w.SubSteps)
	for i := 0; i < w.SubSteps; i++ {
		w.stepInternal(subDt)
	}
}

// stepInternal 单个子步的模拟
func (w *World) stepInternal(dt float32) {
	// 1. 积分外力和重力（更新速度）
	for _, body := range w.Bodies {
		if body.Type == BodyStatic {
			continue // 静态物体不移动
		}

		// 应用重力
		if body.GravityScale != 0 {
			gravityForce := w.Gravity.Mul(body.Mass).Mul(body.GravityScale)
			body.Force = body.Force.Add(gravityForce)
		}

		// 2. 积分速度和角度
		body.IntegrateLinear(dt)
		body.IntegrateAngular(dt)

		// 3. 积分位置和角度
		body.IntegratePosition(dt)
		body.IntegrateRotation(dt)

		// 4. 清除外力和力矩，等待下一帧
		body.Force = math.Vec2(0, 0)
		body.Torque = 0
	}

	// TODO: 这里后面会加碰撞检测和解算器
	// contacts := w.Collision.DetectCollisions(w.Bodies)
	// w.Solver.Solve(w.Bodies, contacts, dt)
}
