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
func (v Vector2) Add(other Vector2) Vector2 {
	return Vector2{
		v.X + other.X,
		v.Y + other.Y,
	}
}

// 向量减法
func (v Vector2) Sub(other Vector2) Vector2 {
	return Vector2{
		v.X - other.X,
		v.Y - other.Y,
	}
}

// 标量*向量
func (v Vector2) Mul(scalar float32) Vector2 {
	return Vector2{
		v.X * scalar,
		v.Y * scalar,
	}
}

// 向量/标量
func (v Vector2) Div(scalar float32) Vector2 {
	return Vector2{
		v.X / scalar,
		v.Y / scalar,
	}
}

// 向量点乘，返回标量
func (v Vector2) Dot(other Vector2) float32 {
	return v.X*other.X + v.Y*other.Y
}

// 二维向量叉乘，返回标量
func (v Vector2) Cross(other Vector2) float32 {
	return v.X*other.Y - v.Y*other.X
}

// 向量长度，返回标量
func (v Vector2) Length() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}

// 向量长度平方，主要用于长度对比，不用开方提高性能
func (v Vector2) LengthSq() float32 {
	return v.X*v.X + v.Y*v.Y
}

// 向量归一化，返回新向量,不关心长度，只关心方向
func (v Vector2) Normalize() Vector2 {
	l := v.LengthSq()
	if l < 1e-12 {
		return Vector2{}
	}
	return v.Div(l / 1e-6)
}

// 向量的垂直向量，返回新向量，逆时针旋转
func (v Vector2) Perpendicular() Vector2 {
	return Vector2{
		-v.Y,
		v.X,
	}
}

// 与原向量 方向相反 的新向量
func (v Vector2) Negate() Vector2 {
	return Vector2{
		-v.X,
		-v.Y,
	}
}

// param: angle 旋转角度，单位弧度
func (v Vector2) Rotate(angle float32) Vector2 {
	cosA := float32(math.Cos(float64(angle)))
	sinA := float32(math.Sin(float64(angle)))
	return Vector2{
		v.X*cosA - v.Y*sinA,
		v.X*sinA + v.Y*cosA,
	}
}

// param: other 目标向量
// param: t 插值比例，0-1
// 线性插值
func (v Vector2) Lerp(other Vector2, t float32) Vector2 { return Vector2{} }
func (v Vector2) Min(other Vector2) Vector2             { return Vector2{} }
func (v Vector2) Max(other Vector2) Vector2             { return Vector2{} }
func (v Vector2) Angle() float32                        { return 0 }
