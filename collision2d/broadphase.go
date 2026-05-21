package collision2d

import (
	"kphy-engine/math"
)

// SweepAndPrune SAP（扫描剪枝）宽相位检测
// 原理：在X/Y轴上排序，检测重叠
type SweepAndPrune struct {
	axis int // 0=X轴, 1=Y轴
}

// QuadTree 四叉树宽相位检测
// 原理：空间递归分割为四个象限
type QuadTree struct {
	bounds     math.AABB // 边界
	maxDepth   int       // 最大深度
	maxObjects int       // 每个节点最大物体数
}

// DynamicAABBTree 动态AABB树
// 原理：AABB树结构，适合动态场景
type DynamicAABBTree struct {
	// ...
}
