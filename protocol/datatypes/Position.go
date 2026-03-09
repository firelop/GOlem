package datatypes

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
