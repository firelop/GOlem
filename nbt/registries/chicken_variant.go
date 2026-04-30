package registries

var ChickenVariant = map[string]any{
	"cold": struct {
		AssetId string `nbt:"asset_id"`
		Model string `nbt:"model"`
		SpawnConditions any `nbt:"spawn_conditions"`
	}{
		AssetId: "minecraft:entity/chicken/cold_chicken",
		Model: "cold",
		SpawnConditions: []any{
			map[string]any{
				"condition": map[string]any{
					"biomes": "#minecraft:spawns_cold_variant_farm_animals",
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
		AssetId: "minecraft:entity/chicken/temperate_chicken",
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
		AssetId: "minecraft:entity/chicken/warm_chicken",
		SpawnConditions: []any{
			map[string]any{
				"condition": map[string]any{
					"biomes": "#minecraft:spawns_warm_variant_farm_animals",
					"type": "minecraft:biome",
				},
				"priority": int32(1),
			},
		},
	},
}
