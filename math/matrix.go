package math

// Matrix2x2 2x2旋转矩阵
type Matrix2x2 struct {
	Data [2][2]float32
}

// Matrix3x3 3x3变换矩阵
type Matrix3x3 struct {
	Data [3][3]float32
}

// Matrix4x4 4x4变换矩阵（为3D预留）
type Matrix4x4 struct {
	Data [4][4]float32
}

// 矩阵操作方法声明（待实现）
func Mat2(angle float32) Matrix2x2                                           { return Matrix2x2{} }
func (m Matrix2x2) MulVec(v Vector2) Vector2                                 { return Vector2{} }
func (m Matrix2x2) Transpose() Matrix2x2                                     { return Matrix2x2{} }
func Mat3Identity() Matrix3x3                                                { return Matrix3x3{} }
func Mat3Transform(position Vector2, angle float32, scale Vector2) Matrix3x3 { return Matrix3x3{} }
func Mat4Identity() Matrix4x4                                                { return Matrix4x4{} }
