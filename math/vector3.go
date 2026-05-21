package math

// Vector3 3D向量（为3D物理预留）
type Vector3 struct {
	X, Y, Z float32
}

// Vector4 4D向量（四元数/齐次坐标用）
type Vector4 struct {
	X, Y, Z, W float32
}

// Vec3 创建新的3D向量
func Vec3(x, y, z float32) Vector3 {
	return Vector3{X: x, Y: y, Z: z}
}

// Vec4 创建新的4D向量
func Vec4(x, y, z, w float32) Vector4 {
	return Vector4{X: x, Y: y, Z: z, W: w}
}

// 3D向量操作方法声明（待实现）
func (v Vector3) Add(other Vector3) Vector3             { return Vector3{} }
func (v Vector3) Sub(other Vector3) Vector3             { return Vector3{} }
func (v Vector3) Mul(scalar float32) Vector3            { return Vector3{} }
func (v Vector3) Div(scalar float32) Vector3            { return Vector3{} }
func (v Vector3) Dot(other Vector3) float32             { return 0 }
func (v Vector3) Cross(other Vector3) Vector3           { return Vector3{} }
func (v Vector3) Length() float32                       { return 0 }
func (v Vector3) LengthSq() float32                     { return 0 }
func (v Vector3) Normalize() Vector3                    { return Vector3{} }
func (v Vector3) Negate() Vector3                       { return Vector3{} }
func (v Vector3) Lerp(other Vector3, t float32) Vector3 { return Vector3{} }
func (v Vector3) ToVec2() Vector2                       { return Vector2{} }
