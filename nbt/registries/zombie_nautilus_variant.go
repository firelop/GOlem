package registries

var ZombieNautilusVariant = map[string]any{
	"temperate": struct {
		AssetId string `nbt:"asset_id"`
		SpawnConditions any `nbt:"spawn_conditions"`
	}{
		AssetId: "minecraft:entity/nautilus/zombie_nautilus",
		SpawnConditions: []any{
			map[string]any{
				"priority": int32(0),
			},
		},
	},
	"warm": struct {
		AssetId string `nbt:"asset_id"`
		Model string `nbt:"model"`
		SpawnConditions any `nbt:"spawn_conditions"`
	}{
		AssetId: "minecraft:entity/nautilus/zombie_nautilus_coral",
		Model: "warm",
		SpawnConditions: []any{
			map[string]any{
				"condition": map[string]any{
					"biomes": "#minecraft:spawns_coral_variant_zombie_nautilus",
					"type": "minecraft:biome",
				},
				"priority": int32(1),
			},
		},
	},
}
