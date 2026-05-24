package math

import "math"

// Matrix2x2 2x2旋转矩阵
type Matrix2x2 struct {
	Data [2][2]float32
}

// Matrix3x3 3x3变换矩阵
type Matrix3x3 struct {
	Data [3][3]float32
}

// Matrix4x4 4x4变换矩阵（为3D预留 TODO）
type Matrix4x4 struct {
	Data [4][4]float32
}

/*
Mat2 生成一个2D旋转矩阵:
cosA -sinA
sinA cosA
主要用途：对2D向量进行旋转变换
注意事项：
1. 参数 angle 的单位是弧度
2. 旋转方向为逆时针（数学坐标系）
*/

func Mat2(angle float32) Matrix2x2 {
	cosA := float32(math.Cos(float64(angle)))
	sinA := float32(math.Sin(float64(angle)))
	return Matrix2x2{
		Data: [2][2]float32{
			{cosA, -sinA},
			{sinA, cosA},
		},
	}
}

/*
MulVec 用矩阵乘以向量
主要用途：对向量进行旋转变换
注意事项：
1. 不修改原向量，返回新的变换后向量
2. 2x2矩阵仅用于旋转，不包含平移
*/

func (m Matrix2x2) MulVec(v Vector2) Vector2 {
	return Vector2{
		X: m.Data[0][0]*v.X + m.Data[0][1]*v.Y,
		Y: m.Data[1][0]*v.X + m.Data[1][1]*v.Y,
	}
}

/*
// Transpose 求矩阵的转置
主要用途：矩阵求逆的一部分、正交矩阵的转置等于逆
注意事项：
1. 对正交矩阵来说，转置等于逆矩阵
2. 此方法返回新矩阵，不修改原矩阵
*/
func (m Matrix2x2) Transpose() Matrix2x2 {
	return Matrix2x2{
		Data: [2][2]float32{
			{m.Data[0][0], m.Data[1][0]},
			{m.Data[0][1], m.Data[1][1]},
		},
	}
}

/*
Mat3Identity 生成3x3单位矩阵
主要用途：1. 矩阵乘法的单位元 2. 初始化变换矩阵
注意事项：
1. 任何矩阵乘以单位矩阵都等于自身
2. 单位矩阵的对角线都是1，其余为0
*/

func Mat3Identity() Matrix3x3 {
	return Matrix3x3{
		Data: [3][3]float32{
			{1, 0, 0},
			{0, 1, 0},
			{0, 0, 1},
		},
	}
}

/*
Mat3Transform 生成2D变换矩阵（包含平移、旋转、缩放）
主要用途：一次性组合多个2D变换（TRS: 平移-旋转-缩放）
注意事项：
1. 变换顺序为：先缩放，再旋转，最后平移（Scale -> Rotate -> Translate）
2. 参数 angle 的单位是弧度
3. 参数 scale 不能为0，否则会导致物体消失
*/

func Mat3Transform(position Vector2, angle float32, scale Vector2) Matrix3x3 {
	cosA := float32(math.Cos(float64(angle)))
	sinA := float32(math.Sin(float64(angle)))

	return Matrix3x3{
		Data: [3][3]float32{
			{scale.X * cosA, -scale.Y * sinA, position.X},
			{scale.X * sinA, scale.Y * cosA, position.Y},
			{0, 0, 1},
		},
	}
}

// Mat4Identity 生成4x4单位矩阵（3D）
// 主要用途：1. 3D矩阵乘法的单位元 2. 初始化3D变换矩阵
// 注意事项：
// 1. 任何4x4矩阵乘以单位矩阵都等于自身
// 2. 此为3D变换预留
func Mat4Identity() Matrix4x4 {
	return Matrix4x4{
		Data: [4][4]float32{
			{1, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 1},
		},
	}
}
