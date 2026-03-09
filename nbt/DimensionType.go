package nbt

type DimensionType struct {
	CoordinateScale             float64 `nbt:"coordinate_scale"`
	HasSkylight                 bool    `nbt:"has_skylight"`
	HasCeiling                  bool    `nbt:"has_ceiling"`
	AmbientLight                float32 `nbt:"ambient_light"`
	HasFixedTime                bool    `nbt:"has_fixed_time,omitempty"`
	MonsterSpawnBlockLightLimit int32   `nbt:"monster_spawn_block_light_limit"`
	MonsterSpawnLightLevel      int32   `nbt:"monster_spawn_light_level"`
	LogicalHeight               int32   `nbt:"logical_height"`
	MinY                        int32   `nbt:"min_y"`
	Height                      int32   `nbt:"height"`
	Infiniburn                  string  `nbt:"infiniburn"`
	Skybox                      string  `nbt:"skybox,omitempty"`
	CardinalLight               string  `nbt:"cardinal_light,omitempty"`
}

func NewOverworldDimensionType() *DimensionType {
	return &DimensionType{
		CoordinateScale:             1.0,
		HasSkylight:                 true,
		HasCeiling:                  false,
		AmbientLight:                0.0,
		HasFixedTime:                false,
		MonsterSpawnBlockLightLimit: 0,
		MonsterSpawnLightLevel:      0,
		LogicalHeight:               384,
		MinY:                        -64,
		Height:                      384,
		Infiniburn:                  "#infiniburn_overworld",
		Skybox:                      "overworld",
		CardinalLight:               "",
	}
}
