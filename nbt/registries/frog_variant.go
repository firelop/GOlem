package registries

var FrogVariant = map[string]any{
	"cold": struct {
		AssetId string `nbt:"asset_id"`
		SpawnConditions any `nbt:"spawn_conditions"`
	}{
		AssetId: "minecraft:entity/frog/cold_frog",
		SpawnConditions: []any{
			map[string]any{
				"condition": map[string]any{
					"biomes": "#minecraft:spawns_cold_variant_frogs",
					"type": "minecraft:biome",
				},
				"priority": int32(1),
			},
		},
	},
	"temperate": struct {
		AssetId string `nbt:"asset_id"`
		SpawnConditions any `nbt:"spawn_conditions"`
	}{
		AssetId: "minecraft:entity/frog/temperate_frog",
		SpawnConditions: []any{
			map[string]any{
				"priority": int32(0),
			},
		},
	},
	"warm": struct {
		AssetId string `nbt:"asset_id"`
		SpawnConditions any `nbt:"spawn_conditions"`
	}{
		AssetId: "minecraft:entity/frog/warm_frog",
		SpawnConditions: []any{
			map[string]any{
				"condition": map[string]any{
					"biomes": "#minecraft:spawns_warm_variant_frogs",
					"type": "minecraft:biome",
				},
				"priority": int32(1),
			},
		},
	},
}
