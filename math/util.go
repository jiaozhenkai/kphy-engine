package math

const (
	Epsilon   = 1e-6
	PI        = 3.141592653589793
	DegToRad  = PI / 180.0
	RadToDeg  = 180.0 / PI
)

// 数学工具函数声明（待实现）
func Clamp(value, min, max float32) float32 { return 0 }
func Lerp(a, b, t float32) float32 { return 0 }
func Abs(x float32) float32 { return 0 }
func Sign(x float32) float32 { return 0 }
func Sqrt(x float32) float32 { return 0 }
func Sin(x float32) float32 { return 0 }
func Cos(x float32) float32 { return 0 }
func Atan2(y, x float32) float32 { return 0 }
func Min(a, b float32) float32 { return 0 }
func Max(a, b float32) float32 { return 0 }
func NearlyEqual(a, b float32) bool { return false }

// AABB 轴对齐包围盒
type AABB struct {
	Min, Max Vector2
}

// AABB 操作方法声明（待实现）
func NewAABB(min, max Vector2) AABB { return AABB{} }
func (a AABB) Center() Vector2 { return Vector2{} }
func (a AABB) Size() Vector2 { return Vector2{} }
func (a AABB) Expand(v Vector2) AABB { return AABB{} }
func (a AABB) Merge(other AABB) AABB { return AABB{} }
func (a AABB) Overlaps(other AABB) bool { return false }
func (a AABB) Contains(point Vector2) bool { return false }
