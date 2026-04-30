package registries

var DimensionType = map[string]any{
	"overworld": struct {
		AmbientLight float32 `nbt:"ambient_light"`
		Attributes any `nbt:"attributes"`
		CoordinateScale float32 `nbt:"coordinate_scale"`
		HasCeiling bool `nbt:"has_ceiling"`
		HasSkylight bool `nbt:"has_skylight"`
		Height int32 `nbt:"height"`
		Infiniburn string `nbt:"infiniburn"`
		LogicalHeight int32 `nbt:"logical_height"`
		MinY int32 `nbt:"min_y"`
		MonsterSpawnBlockLightLimit int32 `nbt:"monster_spawn_block_light_limit"`
		MonsterSpawnLightLevel any `nbt:"monster_spawn_light_level"`
		Timelines string `nbt:"timelines"`
	}{
		AmbientLight: float32(0.0),
		Attributes: map[string]any{
			"minecraft:audio/ambient_sounds": map[string]any{
				"mood": map[string]any{
					"block_search_extent": int32(8),
					"offset": float32(2.0),
					"sound": "minecraft:ambient.cave",
					"tick_delay": int32(6000),
				},
			},
			"minecraft:audio/background_music": map[string]any{
				"creative": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.creative",
				},
				"default": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.game",
				},
			},
			"minecraft:gameplay/bed_rule": map[string]any{
				"can_set_spawn": "always",
				"can_sleep": "when_dark",
				"error_message": map[string]any{
					"translate": "block.minecraft.bed.no_sleep",
				},
			},
			"minecraft:gameplay/nether_portal_spawns_piglin": true,
			"minecraft:gameplay/respawn_anchor_works": false,
			"minecraft:visual/cloud_color": "#ccffffff",
			"minecraft:visual/cloud_height": float32(192.33),
			"minecraft:visual/fog_color": "#c0d8ff",
			"minecraft:visual/sky_color": "#78a7ff",
		},
		CoordinateScale: float32(1.0),
		HasCeiling: false,
		HasSkylight: true,
		Height: int32(384),
		Infiniburn: "#minecraft:infiniburn_overworld",
		LogicalHeight: int32(384),
		MinY: int32(-64),
		MonsterSpawnBlockLightLimit: int32(0),
		MonsterSpawnLightLevel: map[string]any{
			"max_inclusive": int32(7),
			"min_inclusive": int32(0),
			"type": "minecraft:uniform",
		},
		Timelines: "#minecraft:in_overworld",
	},
	"overworld_caves": struct {
		AmbientLight float32 `nbt:"ambient_light"`
		Attributes any `nbt:"attributes"`
		CoordinateScale float32 `nbt:"coordinate_scale"`
		HasCeiling bool `nbt:"has_ceiling"`
		HasSkylight bool `nbt:"has_skylight"`
		Height int32 `nbt:"height"`
		Infiniburn string `nbt:"infiniburn"`
		LogicalHeight int32 `nbt:"logical_height"`
		MinY int32 `nbt:"min_y"`
		MonsterSpawnBlockLightLimit int32 `nbt:"monster_spawn_block_light_limit"`
		MonsterSpawnLightLevel any `nbt:"monster_spawn_light_level"`
		Timelines string `nbt:"timelines"`
	}{
		AmbientLight: float32(0.0),
		Attributes: map[string]any{
			"minecraft:audio/ambient_sounds": map[string]any{
				"mood": map[string]any{
					"block_search_extent": int32(8),
					"offset": float32(2.0),
					"sound": "minecraft:ambient.cave",
					"tick_delay": int32(6000),
				},
			},
			"minecraft:audio/background_music": map[string]any{
				"creative": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.creative",
				},
				"default": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.game",
				},
			},
			"minecraft:gameplay/bed_rule": map[string]any{
				"can_set_spawn": "always",
				"can_sleep": "when_dark",
				"error_message": map[string]any{
					"translate": "block.minecraft.bed.no_sleep",
				},
			},
			"minecraft:gameplay/nether_portal_spawns_piglin": true,
			"minecraft:gameplay/respawn_anchor_works": false,
			"minecraft:visual/cloud_color": "#ccffffff",
			"minecraft:visual/cloud_height": float32(192.33),
			"minecraft:visual/fog_color": "#c0d8ff",
			"minecraft:visual/sky_color": "#78a7ff",
		},
		CoordinateScale: float32(1.0),
		HasCeiling: true,
		HasSkylight: true,
		Height: int32(384),
		Infiniburn: "#minecraft:infiniburn_overworld",
		LogicalHeight: int32(384),
		MinY: int32(-64),
		MonsterSpawnBlockLightLimit: int32(0),
		MonsterSpawnLightLevel: map[string]any{
			"max_inclusive": int32(7),
			"min_inclusive": int32(0),
			"type": "minecraft:uniform",
		},
		Timelines: "#minecraft:in_overworld",
	},
	"the_end": struct {
		AmbientLight float32 `nbt:"ambient_light"`
		Attributes any `nbt:"attributes"`
		CoordinateScale float32 `nbt:"coordinate_scale"`
		HasCeiling bool `nbt:"has_ceiling"`
		HasFixedTime bool `nbt:"has_fixed_time"`
		HasSkylight bool `nbt:"has_skylight"`
		Height int32 `nbt:"height"`
		Infiniburn string `nbt:"infiniburn"`
		LogicalHeight int32 `nbt:"logical_height"`
		MinY int32 `nbt:"min_y"`
		MonsterSpawnBlockLightLimit int32 `nbt:"monster_spawn_block_light_limit"`
		MonsterSpawnLightLevel any `nbt:"monster_spawn_light_level"`
		Skybox string `nbt:"skybox"`
		Timelines string `nbt:"timelines"`
	}{
		AmbientLight: float32(0.25),
		Attributes: map[string]any{
			"minecraft:audio/ambient_sounds": map[string]any{
				"mood": map[string]any{
					"block_search_extent": int32(8),
					"offset": float32(2.0),
					"sound": "minecraft:ambient.cave",
					"tick_delay": int32(6000),
				},
			},
			"minecraft:audio/background_music": map[string]any{
				"default": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(6000),
					"replace_current_music": true,
					"sound": "minecraft:music.end",
				},
			},
			"minecraft:gameplay/bed_rule": map[string]any{
				"can_set_spawn": "never",
				"can_sleep": "never",
				"explodes": true,
			},
			"minecraft:gameplay/respawn_anchor_works": false,
			"minecraft:visual/fog_color": "#181318",
			"minecraft:visual/sky_color": "#000000",
			"minecraft:visual/sky_light_color": "#e580ff",
			"minecraft:visual/sky_light_factor": float32(0.0),
		},
		CoordinateScale: float32(1.0),
		HasCeiling: false,
		HasFixedTime: true,
		HasSkylight: true,
		Height: int32(256),
		Infiniburn: "#minecraft:infiniburn_end",
		LogicalHeight: int32(256),
		MinY: int32(0),
		MonsterSpawnBlockLightLimit: int32(0),
		MonsterSpawnLightLevel: int32(15),
		Skybox: "end",
		Timelines: "#minecraft:in_end",
	},
	"the_nether": struct {
		AmbientLight float32 `nbt:"ambient_light"`
		Attributes any `nbt:"attributes"`
		CardinalLight string `nbt:"cardinal_light"`
		CoordinateScale float32 `nbt:"coordinate_scale"`
		HasCeiling bool `nbt:"has_ceiling"`
		HasFixedTime bool `nbt:"has_fixed_time"`
		HasSkylight bool `nbt:"has_skylight"`
		Height int32 `nbt:"height"`
		Infiniburn string `nbt:"infiniburn"`
		LogicalHeight int32 `nbt:"logical_height"`
		MinY int32 `nbt:"min_y"`
		MonsterSpawnBlockLightLimit int32 `nbt:"monster_spawn_block_light_limit"`
		MonsterSpawnLightLevel any `nbt:"monster_spawn_light_level"`
		Skybox string `nbt:"skybox"`
		Timelines string `nbt:"timelines"`
	}{
		AmbientLight: float32(0.1),
		Attributes: map[string]any{
			"minecraft:gameplay/bed_rule": map[string]any{
				"can_set_spawn": "never",
				"can_sleep": "never",
				"explodes": true,
			},
			"minecraft:gameplay/can_start_raid": false,
			"minecraft:gameplay/fast_lava": true,
			"minecraft:gameplay/piglins_zombify": false,
			"minecraft:gameplay/respawn_anchor_works": true,
			"minecraft:gameplay/sky_light_level": float32(4.0),
			"minecraft:gameplay/snow_golem_melts": true,
			"minecraft:gameplay/water_evaporates": true,
			"minecraft:visual/default_dripstone_particle": map[string]any{
				"type": "minecraft:dripping_dripstone_lava",
			},
			"minecraft:visual/fog_end_distance": float32(96.0),
			"minecraft:visual/fog_start_distance": float32(10.0),
			"minecraft:visual/sky_light_color": "#7a7aff",
			"minecraft:visual/sky_light_factor": float32(0.0),
		},
		CardinalLight: "nether",
		CoordinateScale: float32(8.0),
		HasCeiling: true,
		HasFixedTime: true,
		HasSkylight: false,
		Height: int32(256),
		Infiniburn: "#minecraft:infiniburn_nether",
		LogicalHeight: int32(128),
		MinY: int32(0),
		MonsterSpawnBlockLightLimit: int32(15),
		MonsterSpawnLightLevel: int32(7),
		Skybox: "none",
		Timelines: "#minecraft:in_nether",
	},
}
