package datatypes

type World struct {
	Name          string
	DimensionType string
	Seed          int64
	IsFlat        bool

	SpawnPosition *WorldPosition
}

func NewWorld(name string, dimensionType string, isFlat bool, seed int64, spawnPosition Position) *World {
	world := &World{
		Name:          name,
		DimensionType: dimensionType,
		Seed:          seed,
		IsFlat:        isFlat,
	}
	world.SpawnPosition = spawnPosition.ToWorldPosition(world)
	return world
}
