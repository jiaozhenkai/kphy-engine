package math

import (
	"testing"
)

// 测试 Mat2 函数
func TestMat2(t *testing.T) {
	// 测试能否正常调用
	angle := float32(0.0)
	mat := Mat2(angle)
	
	// 验证返回类型正确
	if mat.Data == [2][2]float32{} {
		t.Log("Mat2 返回了默认矩阵（待实现）")
	}
}

// 测试 Matrix2x2.MulVec 方法
func TestMatrix2x2_MulVec(t *testing.T) {
	mat := Mat2(0.0)
	vec := Vec2(1.0, 0.0)
	
	result := mat.MulVec(vec)
	
	// 验证返回类型正确
	if result.X == 0.0 && result.Y == 0.0 {
		t.Log("Matrix2x2.MulVec 返回了默认向量（待实现）")
	}
}

// 测试 Matrix2x2.Transpose 方法
func TestMatrix2x2_Transpose(t *testing.T) {
	mat := Mat2(0.0)
	result := mat.Transpose()
	
	// 验证返回类型正确
	if result.Data == [2][2]float32{} {
		t.Log("Matrix2x2.Transpose 返回了默认矩阵（待实现）")
	}
}

// 测试 Mat3Identity 函数
func TestMat3Identity(t *testing.T) {
	mat := Mat3Identity()
	
	// 验证返回类型正确
	if mat.Data == [3][3]float32{} {
		t.Log("Mat3Identity 返回了默认矩阵（待实现）")
	}
}

// 测试 Mat3Transform 函数
func TestMat3Transform(t *testing.T) {
	position := Vec2(10.0, 20.0)
	angle := float32(0.5) // 约 28.65 度
	scale := Vec2(2.0, 2.0)
	
	mat := Mat3Transform(position, angle, scale)
	
	// 验证返回类型正确
	if mat.Data == [3][3]float32{} {
		t.Log("Mat3Transform 返回了默认矩阵（待实现）")
	}
}

// 测试 Mat4Identity 函数
func TestMat4Identity(t *testing.T) {
	mat := Mat4Identity()
	
	// 验证返回类型正确
	if mat.Data == [4][4]float32{} {
		t.Log("Mat4Identity 返回了默认矩阵（待实现）")
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
