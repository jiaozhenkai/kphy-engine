# Kphy-Engine 物理引擎架构说明

## 整体架构

```
kphy-engine/
├── math/                 # 数学基础库
│   ├── vector2.go       # 2D向量
│   ├── vector3.go       # 3D向量 & 4D向量
│   ├── matrix.go        # 矩阵运算
│   └── util.go          # 数学工具和AABB
│
├── physics2d/           # ✨ 2D物理核心（优先实现）
│   ├── shape.go         # 2D形状定义
│   ├── material.go      # 2D物理材质
│   ├── body.go          # 2D刚体
│   └── world.go         # 2D物理世界
│
├── collision2d/         # ✨ 2D碰撞检测
│   ├── contact.go       # 2D接触点/流形
│   ├── collision.go     # 2D窄相位检测
│   └── broadphase.go    # 2D宽相位检测
│
├── solver2d/            # ✨ 2D约束解算器
│   └── solver.go        # 2D速度/位置解算
│
├── physics3d/           # 🚀 3D物理预留
│   ├── shape.go         # 3D形状定义
│   ├── body.go          # 3D刚体
│   └── world.go         # 3D物理世界
│
├── collision3d/         # 🚀 3D碰撞检测预留
│   ├── contact.go       # 3D接触点/流形
│   ├── collision.go     # 3D窄相位检测
│   └── broadphase.go    # 3D宽相位检测
│
├── solver3d/            # 🚀 3D约束解算器预留
│   └── solver.go        # 3D速度/位置解算
│
├── shapes/              # UI形状
├── renderer/            # 渲染
├── camera/              # 相机
├── ui/                  # 用户界面
└── main.go
```

## 核心模块说明

### 1. math/ - 数学基础库
- **Vector2**: 2D向量运算
- **Vector3**: 3D向量（预留）
- **Vector4**: 4D向量（四元数/齐次坐标）
- **Matrix2x2/3x3/4x4**: 矩阵变换
- **AABB**: 轴对齐包围盒

### 2. physics2d/ - 2D物理核心
- **Shape**: 2D形状接口（圆、多边形、矩形、胶囊）
- **Material**: 2D物理材质（密度、弹性、摩擦）
- **RigidBody**: 2D刚体（位置、速度、质量、外力）
- **World**: 2D物理世界管理

### 3. collision2d/ - 2D碰撞检测
- **BroadPhase**: 宽相位检测（SAP、四叉树、动态AABB树）
- **NarrowPhase**: 窄相位检测（SAT、GJK算法）
- **Contact/ContactManifold**: 碰撞接触点和流形

### 4. solver2d/ - 2D约束解算器
- **Solver**: 速度解算 + 位置解算
- **Constraint**: 约束接口
- **ContactConstraint/DistanceConstraint/HingeConstraint**: 各类约束

### 5. physics3d/ - 3D预留
- **Shape3D**: 3D形状接口（球体、立方体、胶囊、圆柱、网格）
- **Material3D**: 3D物理材质
- **RigidBody3D**: 3D刚体
- **World3D**: 3D物理世界

### 6. collision3d/ - 3D碰撞检测预留
- **BroadPhase3D**: 3D宽相位检测（SAP、八叉树、动态AABB树）
- **NarrowPhase3D**: 3D窄相位检测（GJK、SAT）
- **Contact3D/ContactManifold3D**: 3D碰撞接触点和流形

### 7. solver3d/ - 3D约束解算器预留
- **Solver3D**: 3D速度解算 + 位置解算
- **Constraint3D**: 3D约束接口
- **ContactConstraint3D/DistanceConstraint3D/BallSocketConstraint3D/HingeConstraint3D**: 各类3D约束

## 关键字段注释说明

### RigidBody 核心字段
| 字段              | 类型     | 含义                                                                    |
| ----------------- | -------- | ----------------------------------------------------------------------- |
| `Position`        | Vector2  | 物体在世界坐标系中的位置                                                |
| `Angle`           | float32  | 物体旋转角度（弧度）                                                    |
| `LinearVelocity`  | Vector2  | 物体的线速度（移动速度）                                                |
| `AngularVelocity` | float32  | 物体的角速度（旋转速度，弧度/秒）                                       |
| `Mass`            | float32  | 物体质量（千克）                                                        |
| **`InvMass`**     | float32  | **质量倒数**：`1/Mass`，0表示质量无穷大（静态物体），用于解算器避免除法 |
| `Inertia`         | float32  | 转动惯量：物体抗旋转的属性                                              |
| **`InvInertia`**  | float32  | **转动惯量倒数**：`1/Inertia`，0表示转动惯量无穷大                      |
| `Force`           | Vector2  | 当前帧作用在物体上的合力                                                |
| `Torque`          | float32  | 当前帧作用在物体上的合扭矩                                              |
| `Material`        | Material | 物体的物理材质                                                          |
| `GravityScale`    | float32  | 物体受重力影响的倍数（0不受重力）                                       |

### Contact 碰撞接触字段
| 字段             | 类型    | 含义                       |
| ---------------- | ------- | -------------------------- |
| `Normal`         | Vector2 | 碰撞法线（从A指向B的方向） |
| `Depth`          | float32 | 物体相互穿透的距离         |
| `NormalImpulse`  | float32 | 法向冲量（用于分离物体）   |
| `TangentImpulse` | float32 | 切向冲量（用于摩擦）       |
| `MixRestitution` | float32 | 混合后的弹性系数           |
| `MixFriction`    | float32 | 混合后的摩擦系数           |

### Material 材质字段
| 字段              | 类型    | 含义                                                   |
| ----------------- | ------- | ------------------------------------------------------ |
| `Density`         | float32 | 密度：用于计算质量（千克/平方米）                      |
| `Restitution`     | float32 | 弹性系数：碰撞后保留多少能量，0=完全非弹性，1=完全弹性 |
| `StaticFriction`  | float32 | 静摩擦系数：物体开始滑动所需的力                       |
| `DynamicFriction` | float32 | 动摩擦系数：物体滑动时的摩擦力                         |

## 数据流

```
输入(时间步长 dt)
    ↓
World.Step(dt)
    ↓
┌─────────────────────────────────────────┐
│  1. 宽相位 (BroadPhase)                 │
│     快速筛选潜在碰撞对                   │
├─────────────────────────────────────────┤
│  2. 窄相位 (NarrowPhase)                │
│     SAT/GJK → 计算接触流形              │
├─────────────────────────────────────────┤
│  3. 速度解算 (VelocitySolve)            │
│     应用冲量 → 更新速度                 │
├─────────────────────────────────────────┤
│  4. 位置解算 (PositionSolve)            │
│     修正穿透 → 更新位置                 │
├─────────────────────────────────────────┤
│  5. 积分 (Integrate)                    │
│     速度积分到位置                      │
└─────────────────────────────────────────┘
    ↓
输出(物体状态已更新)
```
