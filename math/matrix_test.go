package math

import (
	"math"
	"testing"
)

func TestMat2(t *testing.T) {
	// 测试 0 度旋转
	angle0 := float32(0.0)
	mat0 := Mat2(angle0)

	expected0 := Matrix2x2{
		Data: [2][2]float32{
			{1, 0},
			{0, 1},
		},
	}

	if !matrix2x2Equal(mat0, expected0) {
		t.Errorf("TestMat2 0度测试失败: 期望 %v, 得到 %v", expected0, mat0)
	}

	// 测试 90 度旋转（π/2 弧度）
	angle90 := float32(math.Pi / 2)
	mat90 := Mat2(angle90)

	expected90 := Matrix2x2{
		Data: [2][2]float32{
			{0, -1},
			{1, 0},
		},
	}

	// 考虑浮点数误差，分别比较元素
	if !float32Equal(mat90.Data[0][0], expected90.Data[0][0]) ||
		!float32Equal(mat90.Data[0][1], expected90.Data[0][1]) ||
		!float32Equal(mat90.Data[1][0], expected90.Data[1][0]) ||
		!float32Equal(mat90.Data[1][1], expected90.Data[1][1]) {
		t.Errorf("TestMat2 90度测试失败: 期望 %v, 得到 %v", expected90, mat90)
	}
}

func TestMatrix2x2_MulVec(t *testing.T) {
	// 测试 0 度旋转 (1,0) -> (1,0)
	mat0 := Mat2(0.0)
	vec := Vec2(1.0, 0.0)
	result0 := mat0.MulVec(vec)
	expected0 := Vec2(1.0, 0.0)
	if !vector2Equal(result0, expected0) {
		t.Errorf("TestMatrix2x2_MulVec 0度测试失败: 期望 %v, 得到 %v", expected0, result0)
	}

	// 测试 90 度旋转 (1,0) -> (0,1)
	angle90 := float32(math.Pi / 2)
	mat90 := Mat2(angle90)
	result90 := mat90.MulVec(vec)
	expected90 := Vec2(0.0, 1.0)
	if !vector2Equal(result90, expected90) {
		t.Errorf("TestMatrix2x2_MulVec 90度测试失败: 期望 %v, 得到 %v", expected90, result90)
	}
}

func TestMatrix2x2_Transpose(t *testing.T) {
	angle := float32(math.Pi / 4) // 45 度
	mat := Mat2(angle)
	result := mat.Transpose()

	// 对旋转矩阵来说，转置应该等于逆矩阵
	// 对于 45 度旋转，转置矩阵应该是 -45 度旋转
	negAngle := float32(-math.Pi / 4)
	expected := Mat2(negAngle)

	// 比较矩阵元素
	if !matrix2x2Equal(result, expected) {
		t.Errorf("TestMatrix2x2_Transpose 测试失败: 期望 %v, 得到 %v", expected, result)
	}
}

func TestMat3Identity(t *testing.T) {
	mat := Mat3Identity()

	expected := Matrix3x3{
		Data: [3][3]float32{
			{1, 0, 0},
			{0, 1, 0},
			{0, 0, 1},
		},
	}

	if !matrix3x3Equal(mat, expected) {
		t.Errorf("TestMat3Identity 测试失败: 期望 %v, 得到 %v", expected, mat)
	}
}

func TestMat3Transform(t *testing.T) {
	position := Vec2(10.0, 20.0)
	angle := float32(0.0)   // 不旋转
	scale := Vec2(2.0, 3.0) // x 方向缩放 2 倍，y 方向缩放 3 倍

	mat := Mat3Transform(position, angle, scale)

	// 对于 0 度旋转，变换矩阵应该是：
	// [2, 0, 10]
	// [0, 3, 20]
	// [0, 0, 1]
	expected := Matrix3x3{
		Data: [3][3]float32{
			{2, 0, 10},
			{0, 3, 20},
			{0, 0, 1},
		},
	}

	if !matrix3x3Equal(mat, expected) {
		t.Errorf("TestMat3Transform 测试失败: 期望 %v, 得到 %v", expected, mat)
	}
}

func TestMat4Identity(t *testing.T) {
	mat := Mat4Identity()

	expected := Matrix4x4{
		Data: [4][4]float32{
			{1, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 1},
		},
	}

	if !matrix4x4Equal(mat, expected) {
		t.Errorf("TestMat4Identity 测试失败: 期望 %v, 得到 %v", expected, mat)
	}
}

// 辅助函数：比较两个 float32 是否相等
func float32EqualMatrix(a, b float32) bool {
	return float32Equal(a, b)
}

// 辅助函数：比较两个 2x2 矩阵是否相等
func matrix2x2Equal(a, b Matrix2x2) bool {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if !float32Equal(a.Data[i][j], b.Data[i][j]) {
				return false
			}
		}
	}
	return true
}

// 辅助函数：比较两个 3x3 矩阵是否相等
func matrix3x3Equal(a, b Matrix3x3) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !float32Equal(a.Data[i][j], b.Data[i][j]) {
				return false
			}
		}
	}
	return true
}

// 辅助函数：比较两个 4x4 矩阵是否相等
func matrix4x4Equal(a, b Matrix4x4) bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if !float32Equal(a.Data[i][j], b.Data[i][j]) {
				return false
			}
		}
	}
	return true
}
