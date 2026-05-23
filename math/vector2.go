package math

import "math"

// Vector2 2D向量
type Vector2 struct {
	X, Y float32
}

/*
 这里考虑的都是返回新向量，不修改原向量；主要考虑：
 1. 链式调用
 2. 避免修改原向量
 3. 线程安全，多个 goroutine 可以安全使用同一个向量
 4. 数学直觉上，a + b 不会改变 a 或 b，而是产生新值
*/

// Vec2 创建新的2D向量
func Vec2(x, y float32) Vector2 {
	return Vector2{X: x, Y: y}
}

// 向量加法
// usage：位移叠加、速度合成、力的合成
func (v Vector2) Add(other Vector2) Vector2 {
	return Vector2{
		v.X + other.X,
		v.Y + other.Y,
	}
}

// 向量减法
// usage：计算相对位移、求两向量之差
func (v Vector2) Sub(other Vector2) Vector2 {
	return Vector2{
		v.X - other.X,
		v.Y - other.Y,
	}
}

// 标量*向量
// usage：缩放向量长度、调整速度/力的大小
func (v Vector2) Mul(scalar float32) Vector2 {
	return Vector2{
		v.X * scalar,
		v.Y * scalar,
	}
}

// 向量/标量
// usage：按比例缩小向量
func (v Vector2) Div(scalar float32) Vector2 {
	return Vector2{
		v.X / scalar,
		v.Y / scalar,
	}
}

// 向量点乘，返回标量
// usage：1. 判断两向量夹角类型 2. 计算投影长度 3. 检测两物体是否朝向相同
func (v Vector2) Dot(other Vector2) float32 {
	return v.X*other.X + v.Y*other.Y
}

// 二维向量叉乘，返回标量
// usage：1. 判断旋转方向（左/右）2. 计算平行四边形面积 3. 判断点是否在三角形内
func (v Vector2) Cross(other Vector2) float32 {
	return v.X*other.Y - v.Y*other.X
}

// 向量长度，返回标量
// usage：计算距离、判断向量大小
func (v Vector2) Length() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}

// 向量长度平方，主要用于长度对比，不用开方提高性能
// usage：1. 比较向量长度大小（避免开方）2. 检测碰撞距离
func (v Vector2) LengthSq() float32 {
	return v.X*v.X + v.Y*v.Y
}

// 向量归一化，返回新向量,不关心长度，只关心方向
// usage：1. 获取方向向量（用于移动、朝向） 简化后续运算
func (v Vector2) Normalize() Vector2 {
	l := v.LengthSq()
	if l < 1e-12 {
		return Vector2{}
	}
	length := float32(math.Sqrt(float64(l))) // 直接对 l 开平方，避免重复计算
	return v.Div(length)
}

// 向量的垂直向量，返回新向量，逆时针旋转
// usage：1. 计算碰撞法线 2. 求切线方向（摩擦）3. 2D旋转操作
func (v Vector2) Perpendicular() Vector2 {
	return Vector2{
		-v.Y,
		v.X,
	}
}

// 与原向量 方向相反 的新向量
// usage：1. 反向速度/力 2. 反向碰撞法线
func (v Vector2) Negate() Vector2 {
	return Vector2{
		-v.X,
		-v.Y,
	}
}

// param: angle 旋转角度，单位弧度
// 逆时针旋转
// usage：1. 旋转物体朝向 2. 旋转速度/力的方向
func (v Vector2) Rotate(angle float32) Vector2 {
	cosA := float32(math.Cos(float64(angle)))
	sinA := float32(math.Sin(float64(angle)))
	return Vector2{
		v.X*cosA - v.Y*sinA,
		v.X*sinA + v.Y*cosA,
	}
}

// param: t 插值比例，0-1
// 线性插值，在两个向量之间进行平滑过渡 。
// usage：1. 物体移动动画 2. 颜色渐变 3. 相机平滑过渡 4. 属性平滑过渡
func (v Vector2) Lerp(other Vector2, t float32) Vector2 {
	return Vector2{v.X + t*(other.X-v.X), v.Y + t*(other.Y-v.Y)}
}

// 按照分量区最小值
/*
usage:
1. AABB（轴对齐包围盒）计算
2. 碰撞检测
3. 空间分区
*/
func (v Vector2) Min(other Vector2) Vector2 {
	return Vector2{
		float32(math.Min(float64(v.X), float64(other.X))),
		float32(math.Min(float64(v.Y), float64(other.Y))),
	}
}

// 按照分量区最大值
// usage：1. 构建 AABB（轴对齐包围盒） 2. 碰撞检测 3. 空间分区
func (v Vector2) Max(other Vector2) Vector2 {
	return Vector2{
		float32(math.Max(float64(v.X), float64(other.X))),
		float32(math.Max(float64(v.Y), float64(other.Y))),
	}
}

// 返回该向量相对于 X 轴正方向 的角度（弧度制,范围：[-PI, PI]）
// usage：1. 计算物体朝向 2. 旋转到指定方向 3. 计算角度差
func (v Vector2) Angle() float32 {
	return float32(math.Atan2(float64(v.Y), float64(v.X)))
}
