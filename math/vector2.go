package math

// Vector2 2D向量
type Vector2 struct {
	X, Y float32
}

// Vec2 创建新的2D向量
func Vec2(x, y float32) Vector2 {
	return Vector2{X: x, Y: y}
}

// 以下是向量操作方法声明（待实现）
func (v Vector2) Add(other Vector2) Vector2 { return Vector2{} }
func (v Vector2) Sub(other Vector2) Vector2 { return Vector2{} }
func (v Vector2) Mul(scalar float32) Vector2 { return Vector2{} }
func (v Vector2) Div(scalar float32) Vector2 { return Vector2{} }
func (v Vector2) Dot(other Vector2) float32 { return 0 }
func (v Vector2) Cross(other Vector2) float32 { return 0 }
func (v Vector2) Length() float32 { return 0 }
func (v Vector2) LengthSq() float32 { return 0 }
func (v Vector2) Normalize() Vector2 { return Vector2{} }
func (v Vector2) Perpendicular() Vector2 { return Vector2{} }
func (v Vector2) Negate() Vector2 { return Vector2{} }
func (v Vector2) Rotate(angle float32) Vector2 { return Vector2{} }
func (v Vector2) Lerp(other Vector2, t float32) Vector2 { return Vector2{} }
func (v Vector2) Min(other Vector2) Vector2 { return Vector2{} }
func (v Vector2) Max(other Vector2) Vector2 { return Vector2{} }
func (v Vector2) Angle() float32 { return 0 }
