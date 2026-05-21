package collision3d

import (
	"kphy-engine/math"
)

// SweepAndPrune3D 3D SAP（扫描剪枝）宽相位检测
type SweepAndPrune3D struct {
	// 预留实现
}

// Octree 八叉树（3D版本的四叉树）
type Octree struct {
	bounds     math.Vector3 // 边界
	maxDepth   int          // 最大深度
	maxObjects int          // 每个节点最大物体数
}

// DynamicAABBTree3D 3D动态AABB树
type DynamicAABBTree3D struct {
	// 预留实现
}
