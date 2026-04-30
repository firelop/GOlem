package registries

var WolfVariant = map[string]any{
	"ashen": struct {
		Assets struct {
			Angry string `nbt:"angry"`
			Tame string `nbt:"tame"`
			Wild string `nbt:"wild"`
		} `nbt:"assets"`
		SpawnConditions any `nbt:"spawn_conditions"`
	}{
		Assets: struct {
			Angry string `nbt:"angry"`
			Tame string `nbt:"tame"`
			Wild string `nbt:"wild"`
		}{
			Angry: "minecraft:entity/wolf/wolf_ashen_angry",
			Tame: "minecraft:entity/wolf/wolf_ashen_tame",
			Wild: "minecraft:entity/wolf/wolf_ashen",
		},
		SpawnConditions: []any{
			map[string]any{
				"condition": map[string]any{
					"biomes": "minecraft:snowy_taiga",
					"type": "minecraft:biome",
				},
				"priority": int32(1),
			},
		},
	},
	"black": struct {
		Assets struct {
			Angry string `nbt:"angry"`
			Tame string `nbt:"tame"`
			Wild string `nbt:"wild"`
		} `nbt:"assets"`
		SpawnConditions any `nbt:"spawn_conditions"`
	}{
		Assets: struct {
			Angry string `nbt:"angry"`
			Tame string `nbt:"tame"`
			Wild string `nbt:"wild"`
		}{
			Angry: "minecraft:entity/wolf/wolf_black_angry",
			Tame: "minecraft:entity/wolf/wolf_black_tame",
			Wild: "minecraft:entity/wolf/wolf_black",
		},
		SpawnConditions: []any{
			map[string]any{
				"condition": map[string]any{
					"biomes": "minecraft:old_growth_pine_taiga",
					"type": "minecraft:biome",
				},
				"priority": int32(1),
			},
		},
	},
	"chestnut": struct {
		Assets struct {
			Angry string `nbt:"angry"`
			Tame string `nbt:"tame"`
			Wild string `nbt:"wild"`
		} `nbt:"assets"`
		SpawnConditions any `nbt:"spawn_conditions"`
	}{
		Assets: struct {
			Angry string `nbt:"angry"`
			Tame string `nbt:"tame"`
			Wild string `nbt:"wild"`
		}{
			Angry: "minecraft:entity/wolf/wolf_chestnut_angry",
			Tame: "minecraft:entity/wolf/wolf_chestnut_tame",
			Wild: "minecraft:entity/wolf/wolf_chestnut",
		},
		SpawnConditions: []any{
			map[string]any{
				"condition": map[string]any{
					"biomes": "minecraft:old_growth_spruce_taiga",
					"type": "minecraft:biome",
				},
				"priority": int32(1),
			},
		},
	},
	"pale": struct {
		Assets struct {
			Angry string `nbt:"angry"`
			Tame string `nbt:"tame"`
			Wild string `nbt:"wild"`
		} `nbt:"assets"`
		SpawnConditions any `nbt:"spawn_conditions"`
	}{
		Assets: struct {
			Angry string `nbt:"angry"`
			Tame string `nbt:"tame"`
			Wild string `nbt:"wild"`
		}{
			Angry: "minecraft:entity/wolf/wolf_angry",
			Tame: "minecraft:entity/wolf/wolf_tame",
			Wild: "minecraft:entity/wolf/wolf",
		},
		SpawnConditions: []any{
			map[string]any{
				"priority": int32(0),
			},
		},
	},
	"rusty": struct {
		Assets struct {
			Angry string `nbt:"angry"`
			Tame string `nbt:"tame"`
			Wild string `nbt:"wild"`
		} `nbt:"assets"`
		SpawnConditions any `nbt:"spawn_conditions"`
	}{
		Assets: struct {
			Angry string `nbt:"angry"`
			Tame string `nbt:"tame"`
			Wild string `nbt:"wild"`
		}{
			Angry: "minecraft:entity/wolf/wolf_rusty_angry",
			Tame: "minecraft:entity/wolf/wolf_rusty_tame",
			Wild: "minecraft:entity/wolf/wolf_rusty",
		},
		SpawnConditions: []any{
			map[string]any{
				"condition": map[string]any{
					"biomes": "#minecraft:is_jungle",
					"type": "minecraft:biome",
				},
				"priority": int32(1),
			},
		},
	},
	"snowy": struct {
		Assets struct {
			Angry string `nbt:"angry"`
			Tame string `nbt:"tame"`
			Wild string `nbt:"wild"`
		} `nbt:"assets"`
		SpawnConditions any `nbt:"spawn_conditions"`
	}{
		Assets: struct {
			Angry string `nbt:"angry"`
			Tame string `nbt:"tame"`
			Wild string `nbt:"wild"`
		}{
			Angry: "minecraft:entity/wolf/wolf_snowy_angry",
			Tame: "minecraft:entity/wolf/wolf_snowy_tame",
			Wild: "minecraft:entity/wolf/wolf_snowy",
		},
		SpawnConditions: []any{
			map[string]any{
				"condition": map[string]any{
					"biomes": "minecraft:grove",
					"type": "minecraft:biome",
				},
				"priority": int32(1),
			},
		},
	},
	"spotted": struct {
		Assets struct {
			Angry string `nbt:"angry"`
			Tame string `nbt:"tame"`
			Wild string `nbt:"wild"`
		} `nbt:"assets"`
		SpawnConditions any `nbt:"spawn_conditions"`
	}{
		Assets: struct {
			Angry string `nbt:"angry"`
			Tame string `nbt:"tame"`
			Wild string `nbt:"wild"`
		}{
			Angry: "minecraft:entity/wolf/wolf_spotted_angry",
			Tame: "minecraft:entity/wolf/wolf_spotted_tame",
			Wild: "minecraft:entity/wolf/wolf_spotted",
		},
		SpawnConditions: []any{
			map[string]any{
				"condition": map[string]any{
					"biomes": "#minecraft:is_savanna",
					"type": "minecraft:biome",
				},
				"priority": int32(1),
			},
		},
	},
	"striped": struct {
		Assets struct {
			Angry string `nbt:"angry"`
			Tame string `nbt:"tame"`
			Wild string `nbt:"wild"`
		} `nbt:"assets"`
		SpawnConditions any `nbt:"spawn_conditions"`
	}{
		Assets: struct {
			Angry string `nbt:"angry"`
			Tame string `nbt:"tame"`
			Wild string `nbt:"wild"`
		}{
			Angry: "minecraft:entity/wolf/wolf_striped_angry",
			Tame: "minecraft:entity/wolf/wolf_striped_tame",
			Wild: "minecraft:entity/wolf/wolf_striped",
		},
		SpawnConditions: []any{
			map[string]any{
				"condition": map[string]any{
					"biomes": "#minecraft:is_badlands",
					"type": "minecraft:biome",
				},
				"priority": int32(1),
			},
		},
	},
	"woods": struct {
		Assets struct {
			Angry string `nbt:"angry"`
			Tame string `nbt:"tame"`
			Wild string `nbt:"wild"`
		} `nbt:"assets"`
		SpawnConditions any `nbt:"spawn_conditions"`
	}{
		Assets: struct {
			Angry string `nbt:"angry"`
			Tame string `nbt:"tame"`
			Wild string `nbt:"wild"`
		}{
			Angry: "minecraft:entity/wolf/wolf_woods_angry",
			Tame: "minecraft:entity/wolf/wolf_woods_tame",
			Wild: "minecraft:entity/wolf/wolf_woods",
		},
		SpawnConditions: []any{
			map[string]any{
				"condition": map[string]any{
					"biomes": "minecraft:forest",
					"type": "minecraft:biome",
				},
				"priority": int32(1),
			},
		},
	},
}
