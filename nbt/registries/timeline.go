package registries

var Timeline = map[string]any{
	"day": struct {
		PeriodTicks int32 `nbt:"period_ticks"`
		Tracks any `nbt:"tracks"`
	}{
		PeriodTicks: int32(24000),
		Tracks: map[string]any{
			"minecraft:audio/firefly_bush_sounds": map[string]any{
				"keyframes": []any{
					map[string]any{
						"ticks": int32(12600),
						"value": true,
					},
					map[string]any{
						"ticks": int32(23401),
						"value": false,
					},
				},
				"modifier": "or",
			},
			"minecraft:gameplay/bees_stay_in_hive": map[string]any{
				"keyframes": []any{
					map[string]any{
						"ticks": int32(12542),
						"value": true,
					},
					map[string]any{
						"ticks": int32(23460),
						"value": false,
					},
				},
				"modifier": "or",
			},
			"minecraft:gameplay/cat_waking_up_gift_chance": map[string]any{
				"ease": "constant",
				"keyframes": []any{
					map[string]any{
						"ticks": int32(362),
						"value": float32(0.0),
					},
					map[string]any{
						"ticks": int32(23667),
						"value": float32(0.7),
					},
				},
				"modifier": "maximum",
			},
			"minecraft:gameplay/creaking_active": map[string]any{
				"keyframes": []any{
					map[string]any{
						"ticks": int32(12600),
						"value": true,
					},
					map[string]any{
						"ticks": int32(23401),
						"value": false,
					},
				},
				"modifier": "or",
			},
			"minecraft:gameplay/eyeblossom_open": map[string]any{
				"keyframes": []any{
					map[string]any{
						"ticks": int32(12600),
						"value": true,
					},
					map[string]any{
						"ticks": int32(23401),
						"value": false,
					},
				},
			},
			"minecraft:gameplay/monsters_burn": map[string]any{
				"keyframes": []any{
					map[string]any{
						"ticks": int32(12542),
						"value": false,
					},
					map[string]any{
						"ticks": int32(23460),
						"value": true,
					},
				},
				"modifier": "or",
			},
			"minecraft:gameplay/sky_light_level": map[string]any{
				"keyframes": []any{
					map[string]any{
						"ticks": int32(133),
						"value": float32(1.0),
					},
					map[string]any{
						"ticks": int32(11867),
						"value": float32(1.0),
					},
					map[string]any{
						"ticks": int32(13670),
						"value": float32(0.26666668),
					},
					map[string]any{
						"ticks": int32(22330),
						"value": float32(0.26666668),
					},
				},
				"modifier": "multiply",
			},
			"minecraft:gameplay/turtle_egg_hatch_chance": map[string]any{
				"ease": "constant",
				"keyframes": []any{
					map[string]any{
						"ticks": int32(21062),
						"value": float32(1.0),
					},
					map[string]any{
						"ticks": int32(21905),
						"value": float32(0.002),
					},
				},
				"modifier": "maximum",
			},
			"minecraft:visual/cloud_color": map[string]any{
				"keyframes": []any{
					map[string]any{
						"ticks": int32(133),
						"value": int32(-1),
					},
					map[string]any{
						"ticks": int32(11867),
						"value": int32(-1),
					},
					map[string]any{
						"ticks": int32(13670),
						"value": int32(-15132378),
					},
					map[string]any{
						"ticks": int32(22330),
						"value": int32(-15132378),
					},
				},
				"modifier": "multiply",
			},
			"minecraft:visual/fog_color": map[string]any{
				"keyframes": []any{
					map[string]any{
						"ticks": int32(133),
						"value": "#ffffff",
					},
					map[string]any{
						"ticks": int32(11867),
						"value": "#ffffff",
					},
					map[string]any{
						"ticks": int32(13670),
						"value": "#0f0f16",
					},
					map[string]any{
						"ticks": int32(22330),
						"value": "#0f0f16",
					},
				},
				"modifier": "multiply",
			},
			"minecraft:visual/moon_angle": map[string]any{
				"ease": map[string]any{
					"cubic_bezier": []any{
						float32(0.362),
						float32(0.241),
						float32(0.638),
						float32(0.759),
					},
				},
				"keyframes": []any{
					map[string]any{
						"ticks": int32(6000),
						"value": float32(540.0),
					},
					map[string]any{
						"ticks": int32(6000),
						"value": float32(180.0),
					},
				},
			},
			"minecraft:visual/sky_color": map[string]any{
				"keyframes": []any{
					map[string]any{
						"ticks": int32(133),
						"value": "#ffffff",
					},
					map[string]any{
						"ticks": int32(11867),
						"value": "#ffffff",
					},
					map[string]any{
						"ticks": int32(13670),
						"value": "#000000",
					},
					map[string]any{
						"ticks": int32(22330),
						"value": "#000000",
					},
				},
				"modifier": "multiply",
			},
			"minecraft:visual/sky_light_color": map[string]any{
				"keyframes": []any{
					map[string]any{
						"ticks": int32(730),
						"value": "#ffffff",
					},
					map[string]any{
						"ticks": int32(11270),
						"value": "#ffffff",
					},
					map[string]any{
						"ticks": int32(13140),
						"value": "#7a7aff",
					},
					map[string]any{
						"ticks": int32(22860),
						"value": "#7a7aff",
					},
				},
				"modifier": "multiply",
			},
			"minecraft:visual/sky_light_factor": map[string]any{
				"keyframes": []any{
					map[string]any{
						"ticks": int32(730),
						"value": float32(1.0),
					},
					map[string]any{
						"ticks": int32(11270),
						"value": float32(1.0),
					},
					map[string]any{
						"ticks": int32(13140),
						"value": float32(0.24),
					},
					map[string]any{
						"ticks": int32(22860),
						"value": float32(0.24),
					},
				},
				"modifier": "multiply",
			},
			"minecraft:visual/star_angle": map[string]any{
				"ease": map[string]any{
					"cubic_bezier": []any{
						float32(0.362),
						float32(0.241),
						float32(0.638),
						float32(0.759),
					},
				},
				"keyframes": []any{
					map[string]any{
						"ticks": int32(6000),
						"value": float32(360.0),
					},
					map[string]any{
						"ticks": int32(6000),
						"value": float32(0.0),
					},
				},
			},
			"minecraft:visual/star_brightness": map[string]any{
				"keyframes": []any{
					map[string]any{
						"ticks": int32(92),
						"value": float32(0.037),
					},
					map[string]any{
						"ticks": int32(627),
						"value": float32(0.0),
					},
					map[string]any{
						"ticks": int32(11373),
						"value": float32(0.0),
					},
					map[string]any{
						"ticks": int32(11732),
						"value": float32(0.016),
					},
					map[string]any{
						"ticks": int32(11959),
						"value": float32(0.044),
					},
					map[string]any{
						"ticks": int32(12399),
						"value": float32(0.143),
					},
					map[string]any{
						"ticks": int32(12729),
						"value": float32(0.258),
					},
					map[string]any{
						"ticks": int32(13228),
						"value": float32(0.5),
					},
					map[string]any{
						"ticks": int32(22772),
						"value": float32(0.5),
					},
					map[string]any{
						"ticks": int32(23032),
						"value": float32(0.364),
					},
					map[string]any{
						"ticks": int32(23356),
						"value": float32(0.225),
					},
					map[string]any{
						"ticks": int32(23758),
						"value": float32(0.101),
					},
				},
				"modifier": "maximum",
			},
			"minecraft:visual/sun_angle": map[string]any{
				"ease": map[string]any{
					"cubic_bezier": []any{
						float32(0.362),
						float32(0.241),
						float32(0.638),
						float32(0.759),
					},
				},
				"keyframes": []any{
					map[string]any{
						"ticks": int32(6000),
						"value": float32(360.0),
					},
					map[string]any{
						"ticks": int32(6000),
						"value": float32(0.0),
					},
				},
			},
			"minecraft:visual/sunrise_sunset_color": map[string]any{
				"keyframes": []any{
					map[string]any{
						"ticks": int32(71),
						"value": "#5fefa333",
					},
					map[string]any{
						"ticks": int32(310),
						"value": "#29f5ba33",
					},
					map[string]any{
						"ticks": int32(565),
						"value": "#06fbd433",
					},
					map[string]any{
						"ticks": int32(730),
						"value": "#00ffe533",
					},
					map[string]any{
						"ticks": int32(11270),
						"value": "#00ffe533",
					},
					map[string]any{
						"ticks": int32(11397),
						"value": "#04fcd833",
					},
					map[string]any{
						"ticks": int32(11522),
						"value": "#0ff9cb33",
					},
					map[string]any{
						"ticks": int32(11690),
						"value": "#29f5ba33",
					},
					map[string]any{
						"ticks": int32(11929),
						"value": "#5fefa333",
					},
					map[string]any{
						"ticks": int32(12243),
						"value": "#b1e78733",
					},
					map[string]any{
						"ticks": int32(12358),
						"value": "#cce47e33",
					},
					map[string]any{
						"ticks": int32(12512),
						"value": "#e9e07233",
					},
					map[string]any{
						"ticks": int32(12613),
						"value": "#f6dd6b33",
					},
					map[string]any{
						"ticks": int32(12732),
						"value": "#feda6333",
					},
					map[string]any{
						"ticks": int32(12841),
						"value": "#fed75c33",
					},
					map[string]any{
						"ticks": int32(13035),
						"value": "#ecd25133",
					},
					map[string]any{
						"ticks": int32(13252),
						"value": "#c1cc4733",
					},
					map[string]any{
						"ticks": int32(13775),
						"value": "#36be3733",
					},
					map[string]any{
						"ticks": int32(13888),
						"value": "#1fbb3533",
					},
					map[string]any{
						"ticks": int32(14039),
						"value": "#09b73333",
					},
					map[string]any{
						"ticks": int32(14192),
						"value": "#00b33333",
					},
					map[string]any{
						"ticks": int32(21807),
						"value": "#00b23333",
					},
					map[string]any{
						"ticks": int32(21961),
						"value": "#09b73333",
					},
					map[string]any{
						"ticks": int32(22112),
						"value": "#1fbb3533",
					},
					map[string]any{
						"ticks": int32(22225),
						"value": "#36be3733",
					},
					map[string]any{
						"ticks": int32(22748),
						"value": "#c1cc4733",
					},
					map[string]any{
						"ticks": int32(22965),
						"value": "#ecd25133",
					},
					map[string]any{
						"ticks": int32(23159),
						"value": "#fed75c33",
					},
					map[string]any{
						"ticks": int32(23272),
						"value": "#feda6333",
					},
					map[string]any{
						"ticks": int32(23488),
						"value": "#e9e07233",
					},
					map[string]any{
						"ticks": int32(23642),
						"value": "#cce47e33",
					},
					map[string]any{
						"ticks": int32(23757),
						"value": "#b1e78733",
					},
				},
			},
		},
	},
	"early_game": struct {
		Tracks any `nbt:"tracks"`
	}{
		Tracks: map[string]any{
			"minecraft:gameplay/can_pillager_patrol_spawn": map[string]any{
				"keyframes": []any{
					map[string]any{
						"ticks": int32(0),
						"value": false,
					},
					map[string]any{
						"ticks": int32(120000),
						"value": true,
					},
				},
				"modifier": "and",
			},
		},
	},
	"moon": struct {
		PeriodTicks int32 `nbt:"period_ticks"`
		Tracks any `nbt:"tracks"`
	}{
		PeriodTicks: int32(192000),
		Tracks: map[string]any{
			"minecraft:gameplay/surface_slime_spawn_chance": map[string]any{
				"ease": "constant",
				"keyframes": []any{
					map[string]any{
						"ticks": int32(0),
						"value": float32(0.5),
					},
					map[string]any{
						"ticks": int32(24000),
						"value": float32(0.375),
					},
					map[string]any{
						"ticks": int32(48000),
						"value": float32(0.25),
					},
					map[string]any{
						"ticks": int32(72000),
						"value": float32(0.125),
					},
					map[string]any{
						"ticks": int32(96000),
						"value": float32(0.0),
					},
					map[string]any{
						"ticks": int32(120000),
						"value": float32(0.125),
					},
					map[string]any{
						"ticks": int32(144000),
						"value": float32(0.25),
					},
					map[string]any{
						"ticks": int32(168000),
						"value": float32(0.375),
					},
				},
				"modifier": "maximum",
			},
			"minecraft:visual/moon_phase": map[string]any{
				"keyframes": []any{
					map[string]any{
						"ticks": int32(0),
						"value": "full_moon",
					},
					map[string]any{
						"ticks": int32(24000),
						"value": "waning_gibbous",
					},
					map[string]any{
						"ticks": int32(48000),
						"value": "third_quarter",
					},
					map[string]any{
						"ticks": int32(72000),
						"value": "waning_crescent",
					},
					map[string]any{
						"ticks": int32(96000),
						"value": "new_moon",
					},
					map[string]any{
						"ticks": int32(120000),
						"value": "waxing_crescent",
					},
					map[string]any{
						"ticks": int32(144000),
						"value": "first_quarter",
					},
					map[string]any{
						"ticks": int32(168000),
						"value": "waxing_gibbous",
					},
				},
			},
		},
	},
	"villager_schedule": struct {
		PeriodTicks int32 `nbt:"period_ticks"`
		Tracks any `nbt:"tracks"`
	}{
		PeriodTicks: int32(24000),
		Tracks: map[string]any{
			"minecraft:gameplay/baby_villager_activity": map[string]any{
				"keyframes": []any{
					map[string]any{
						"ticks": int32(10),
						"value": "minecraft:idle",
					},
					map[string]any{
						"ticks": int32(3000),
						"value": "minecraft:play",
					},
					map[string]any{
						"ticks": int32(6000),
						"value": "minecraft:idle",
					},
					map[string]any{
						"ticks": int32(10000),
						"value": "minecraft:play",
					},
					map[string]any{
						"ticks": int32(12000),
						"value": "minecraft:rest",
					},
				},
			},
			"minecraft:gameplay/villager_activity": map[string]any{
				"keyframes": []any{
					map[string]any{
						"ticks": int32(10),
						"value": "minecraft:idle",
					},
					map[string]any{
						"ticks": int32(2000),
						"value": "minecraft:work",
					},
					map[string]any{
						"ticks": int32(9000),
						"value": "minecraft:meet",
					},
					map[string]any{
						"ticks": int32(11000),
						"value": "minecraft:idle",
					},
					map[string]any{
						"ticks": int32(12000),
						"value": "minecraft:rest",
					},
				},
			},
		},
	},
}
