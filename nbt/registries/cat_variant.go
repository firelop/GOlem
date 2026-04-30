package registries

var CatVariant = map[string]any{
	"all_black": struct {
		AssetId string `nbt:"asset_id"`
		SpawnConditions any `nbt:"spawn_conditions"`
	}{
		AssetId: "minecraft:entity/cat/all_black",
		SpawnConditions: []any{
			map[string]any{
				"condition": map[string]any{
					"structures": "#minecraft:cats_spawn_as_black",
					"type": "minecraft:structure",
				},
				"priority": int32(1),
			},
			map[string]any{
				"condition": map[string]any{
					"range": map[string]any{
						"min": float32(0.9),
					},
					"type": "minecraft:moon_brightness",
				},
				"priority": int32(0),
			},
		},
	},
	"black": struct {
		AssetId string `nbt:"asset_id"`
		SpawnConditions any `nbt:"spawn_conditions"`
	}{
		AssetId: "minecraft:entity/cat/black",
		SpawnConditions: []any{
			map[string]any{
				"priority": int32(0),
			},
		},
	},
	"british_shorthair": struct {
		AssetId string `nbt:"asset_id"`
		SpawnConditions any `nbt:"spawn_conditions"`
	}{
		AssetId: "minecraft:entity/cat/british_shorthair",
		SpawnConditions: []any{
			map[string]any{
				"priority": int32(0),
			},
		},
	},
	"calico": struct {
		AssetId string `nbt:"asset_id"`
		SpawnConditions any `nbt:"spawn_conditions"`
	}{
		AssetId: "minecraft:entity/cat/calico",
		SpawnConditions: []any{
			map[string]any{
				"priority": int32(0),
			},
		},
	},
	"jellie": struct {
		AssetId string `nbt:"asset_id"`
		SpawnConditions any `nbt:"spawn_conditions"`
	}{
		AssetId: "minecraft:entity/cat/jellie",
		SpawnConditions: []any{
			map[string]any{
				"priority": int32(0),
			},
		},
	},
	"persian": struct {
		AssetId string `nbt:"asset_id"`
		SpawnConditions any `nbt:"spawn_conditions"`
	}{
		AssetId: "minecraft:entity/cat/persian",
		SpawnConditions: []any{
			map[string]any{
				"priority": int32(0),
			},
		},
	},
	"ragdoll": struct {
		AssetId string `nbt:"asset_id"`
		SpawnConditions any `nbt:"spawn_conditions"`
	}{
		AssetId: "minecraft:entity/cat/ragdoll",
		SpawnConditions: []any{
			map[string]any{
				"priority": int32(0),
			},
		},
	},
	"red": struct {
		AssetId string `nbt:"asset_id"`
		SpawnConditions any `nbt:"spawn_conditions"`
	}{
		AssetId: "minecraft:entity/cat/red",
		SpawnConditions: []any{
			map[string]any{
				"priority": int32(0),
			},
		},
	},
	"siamese": struct {
		AssetId string `nbt:"asset_id"`
		SpawnConditions any `nbt:"spawn_conditions"`
	}{
		AssetId: "minecraft:entity/cat/siamese",
		SpawnConditions: []any{
			map[string]any{
				"priority": int32(0),
			},
		},
	},
	"tabby": struct {
		AssetId string `nbt:"asset_id"`
		SpawnConditions any `nbt:"spawn_conditions"`
	}{
		AssetId: "minecraft:entity/cat/tabby",
		SpawnConditions: []any{
			map[string]any{
				"priority": int32(0),
			},
		},
	},
	"white": struct {
		AssetId string `nbt:"asset_id"`
		SpawnConditions any `nbt:"spawn_conditions"`
	}{
		AssetId: "minecraft:entity/cat/white",
		SpawnConditions: []any{
			map[string]any{
				"priority": int32(0),
			},
		},
	},
}
