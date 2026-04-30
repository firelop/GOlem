package datatypes

type DoubleVec3 struct {
	X float64
	Y float64
	Z float64
}

func NewDoubleVec3(x float64, y float64, z float64) *DoubleVec3 {
	return &DoubleVec3{
		X: x,
		Y: y,
		Z: z,
	}
}

func (d *DoubleVec3) ToPosition() *Position {
	return &Position{
		X: int32(d.X),
		Y: int16(d.Y),
		Z: int32(d.Z),
	}
}

type Position struct {
	X int32 // int 24 bits
	Y int16 // int 12 bits
	Z int32 // int 24 bits
}

func NewPosition(x int32, y int16, z int32) *Position {
	return &Position{
		X: x,
		Y: y,
		Z: z,
	}
}

func (p *Position) ToWorldPosition(world *World) *WorldPosition {
	return &WorldPosition{
		World:    world,
		Position: *p,
	}
}

type WorldPosition struct {
	World *World
	Position
}

type Angle struct {
	Yaw   float32
	Pitch float32
}

func NewAngle(yaw float32, pitch float32) *Angle {
	return &Angle{
		Yaw:   yaw,
		Pitch: pitch,
	}
}
