package physics2d

// Material 物理材质属性
type Material struct {
	Density         float32 // 密度：物体的质量密度（千克/平方米），影响质量计算
	Restitution     float32 // 弹性系数：碰撞后的能量保留比例（0-1），0=完全非弹性，1=完全弹性
	StaticFriction  float32 // 静摩擦系数：物体开始滑动所需的摩擦力（0-1）
	DynamicFriction float32 // 动摩擦系数：物体滑动时的摩擦力（0-1），通常小于静摩擦
}

// DefaultMaterial 默认材质
func DefaultMaterial() Material {
	return Material{
		Density:         1.0,
		Restitution:     0.5,
		StaticFriction:  0.6,
		DynamicFriction: 0.4,
	}
}

// Rubber 橡胶材质（高弹性）
func Rubber() Material {
	return Material{
		Density:         1.2,
		Restitution:     0.85,
		StaticFriction:  0.8,
		DynamicFriction: 0.6,
	}
}

// Steel 钢铁材质（低弹性，高摩擦）
func Steel() Material {
	return Material{
		Density:         7.8,
		Restitution:     0.3,
		StaticFriction:  0.5,
		DynamicFriction: 0.4,
	}
}

// Wood 木材材质
func Wood() Material {
	return Material{
		Density:         0.6,
		Restitution:     0.4,
		StaticFriction:  0.6,
		DynamicFriction: 0.5,
	}
}
