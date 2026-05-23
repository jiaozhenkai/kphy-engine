package math

import (
	"math"
	"testing"
)

// 浮点数比较的误差范围
const epsilon = 1e-6

func TestVec2(t *testing.T) {
	v := Vec2(3.0, 4.0)
	if v.X != 3.0 || v.Y != 4.0 {
		t.Errorf("Vec2 failed: expected (3,4), got (%v,%v)", v.X, v.Y)
	}
}

func TestVector2_Add(t *testing.T) {
	v1 := Vec2(1.0, 2.0)
	v2 := Vec2(3.0, 4.0)
	result := v1.Add(v2)
	expected := Vec2(4.0, 6.0)
	if !vector2Equal(result, expected) {
		t.Errorf("Add failed: expected %v, got %v", expected, result)
	}
}

func TestVector2_Sub(t *testing.T) {
	v1 := Vec2(5.0, 6.0)
	v2 := Vec2(3.0, 4.0)
	result := v1.Sub(v2)
	expected := Vec2(2.0, 2.0)
	if !vector2Equal(result, expected) {
		t.Errorf("Sub failed: expected %v, got %v", expected, result)
	}
}

func TestVector2_Mul(t *testing.T) {
	v := Vec2(2.0, 3.0)
	result := v.Mul(2.0)
	expected := Vec2(4.0, 6.0)
	if !vector2Equal(result, expected) {
		t.Errorf("Mul failed: expected %v, got %v", expected, result)
	}
}

func TestVector2_Div(t *testing.T) {
	v := Vec2(4.0, 6.0)
	result := v.Div(2.0)
	expected := Vec2(2.0, 3.0)
	if !vector2Equal(result, expected) {
		t.Errorf("Div failed: expected %v, got %v", expected, result)
	}
}

func TestVector2_Dot(t *testing.T) {
	v1 := Vec2(1.0, 2.0)
	v2 := Vec2(3.0, 4.0)
	result := v1.Dot(v2)
	expected := float32(11.0) // 1*3 + 2*4 = 11
	if !float32Equal(result, expected) {
		t.Errorf("Dot failed: expected %v, got %v", expected, result)
	}
}

func TestVector2_Cross(t *testing.T) {
	v1 := Vec2(1.0, 2.0)
	v2 := Vec2(3.0, 4.0)
	result := v1.Cross(v2)
	expected := float32(-2.0) // 1*4 - 2*3 = -2
	if !float32Equal(result, expected) {
		t.Errorf("Cross failed: expected %v, got %v", expected, result)
	}
}

func TestVector2_Length(t *testing.T) {
	v := Vec2(3.0, 4.0)
	result := v.Length()
	expected := float32(5.0)
	if !float32Equal(result, expected) {
		t.Errorf("Length failed: expected %v, got %v", expected, result)
	}
}

func TestVector2_LengthSq(t *testing.T) {
	v := Vec2(3.0, 4.0)
	result := v.LengthSq()
	expected := float32(25.0)
	if !float32Equal(result, expected) {
		t.Errorf("LengthSq failed: expected %v, got %v", expected, result)
	}
}

func TestVector2_Normalize(t *testing.T) {
	// 测试正常情况
	v := Vec2(3.0, 4.0)
	result := v.Normalize()
	// 归一化后应该是 (0.6, 0.8)
	expected := Vec2(0.6, 0.8)
	if !vector2Equal(result, expected) {
		t.Errorf("Normalize failed: expected %v, got %v", expected, result)
	}
	// 验证归一化后的长度应该接近 1
	length := result.Length()
	if !float32Equal(length, 1.0) {
		t.Errorf("Normalize length check failed: expected 1, got %v", length)
	}

	// 测试零向量
	zero := Vec2(0.0, 0.0)
	resultZero := zero.Normalize()
	// 零向量归一化应该返回零向量
	if !vector2Equal(resultZero, zero) {
		t.Errorf("Normalize zero vector failed: expected %v, got %v", zero, resultZero)
	}
}

func TestVector2_Perpendicular(t *testing.T) {
	v := Vec2(3.0, 4.0)
	result := v.Perpendicular()
	// 逆时针垂直应该是 (-4, 3)
	expected := Vec2(-4.0, 3.0)
	if !vector2Equal(result, expected) {
		t.Errorf("Perpendicular failed: expected %v, got %v", expected, result)
	}
	// 验证垂直：点积应该接近 0
	dot := v.Dot(result)
	if !float32Equal(dot, 0.0) {
		t.Errorf("Perpendicular dot check failed: expected 0, got %v", dot)
	}
}

func TestVector2_Negate(t *testing.T) {
	v := Vec2(3.0, 4.0)
	result := v.Negate()
	expected := Vec2(-3.0, -4.0)
	if !vector2Equal(result, expected) {
		t.Errorf("Negate failed: expected %v, got %v", expected, result)
	}
}

func TestVector2_Rotate(t *testing.T) {
	// 测试旋转 90 度（PI/2 弧度）
	v := Vec2(1.0, 0.0)
	angle := float32(math.Pi / 2.0)
	result := v.Rotate(angle)
	expected := Vec2(0.0, 1.0)
	if !vector2Equal(result, expected) {
		t.Errorf("Rotate 90deg failed: expected %v, got %v", expected, result)
	}

	// 测试旋转 180 度
	v2 := Vec2(1.0, 1.0)
	angle2 := float32(math.Pi)
	result2 := v2.Rotate(angle2)
	expected2 := Vec2(-1.0, -1.0)
	if !vector2Equal(result2, expected2) {
		t.Errorf("Rotate 180deg failed: expected %v, got %v", expected2, result2)
	}
}

func TestVector2_Lerp(t *testing.T) {
	v1 := Vec2(0.0, 0.0)
	v2 := Vec2(10.0, 10.0)

	// t=0
	result0 := v1.Lerp(v2, 0.0)
	if !vector2Equal(result0, v1) {
		t.Errorf("Lerp t=0 failed: expected %v, got %v", v1, result0)
	}

	// t=1
	result1 := v1.Lerp(v2, 1.0)
	if !vector2Equal(result1, v2) {
		t.Errorf("Lerp t=1 failed: expected %v, got %v", v2, result1)
	}

	// t=0.5
	result05 := v1.Lerp(v2, 0.5)
	expected05 := Vec2(5.0, 5.0)
	if !vector2Equal(result05, expected05) {
		t.Errorf("Lerp t=0.5 failed: expected %v, got %v", expected05, result05)
	}
}

func TestVector2_Min(t *testing.T) {
	v1 := Vec2(10.0, 5.0)
	v2 := Vec2(3.0, 8.0)
	result := v1.Min(v2)
	expected := Vec2(3.0, 5.0)
	if !vector2Equal(result, expected) {
		t.Errorf("Min failed: expected %v, got %v", expected, result)
	}
}

func TestVector2_Max(t *testing.T) {
	v1 := Vec2(10.0, 5.0)
	v2 := Vec2(3.0, 8.0)
	result := v1.Max(v2)
	expected := Vec2(10.0, 8.0)
	if !vector2Equal(result, expected) {
		t.Errorf("Max failed: expected %v, got %v", expected, result)
	}
}

func TestVector2_Angle(t *testing.T) {
	// 0 度
	v0 := Vec2(1.0, 0.0)
	angle0 := v0.Angle()
	if !float32Equal(angle0, 0.0) {
		t.Errorf("Angle 0deg failed: expected 0, got %v", angle0)
	}

	// 90 度 (PI/2)
	v90 := Vec2(0.0, 1.0)
	angle90 := v90.Angle()
	expected90 := float32(math.Pi / 2.0)
	if !float32Equal(angle90, expected90) {
		t.Errorf("Angle 90deg failed: expected %v, got %v", expected90, angle90)
	}

	// 180 度 (PI)
	v180 := Vec2(-1.0, 0.0)
	angle180 := v180.Angle()
	expected180 := float32(math.Pi)
	if !float32Equal(angle180, expected180) {
		t.Errorf("Angle 180deg failed: expected %v, got %v", expected180, angle180)
	}

	// -90 度 (-PI/2)
	v270 := Vec2(0.0, -1.0)
	angle270 := v270.Angle()
	expected270 := float32(-math.Pi / 2.0)
	if !float32Equal(angle270, expected270) {
		t.Errorf("Angle -90deg failed: expected %v, got %v", expected270, angle270)
	}
}

// 辅助函数：比较两个 Vector2 是否相等
func vector2Equal(a, b Vector2) bool {
	return float32Equal(a.X, b.X) && float32Equal(a.Y, b.Y)
}

// 辅助函数：比较两个 float32 是否相等
func float32Equal(a, b float32) bool {
	return math.Abs(float64(a-b)) < epsilon
}
