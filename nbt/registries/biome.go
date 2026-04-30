package registries

var Biome = map[string]any{
	"badlands": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		CreatureSpawnProbability float32 `nbt:"creature_spawn_probability"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:audio/background_music": map[string]any{
				"default": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.overworld.badlands",
				},
			},
			"minecraft:gameplay/snow_golem_melts": true,
			"minecraft:visual/sky_color": "#6eb1ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		CreatureSpawnProbability: float32(0.03),
		Downfall: float32(0.0),
		Effects: map[string]any{
			"foliage_color": "#9e814d",
			"grass_color": "#90814d",
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:ore_gold_extra",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:patch_grass_badlands",
				"minecraft:patch_dry_grass_badlands",
				"minecraft:patch_dead_bush_badlands",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_sugar_cane_badlands",
				"minecraft:patch_pumpkin",
				"minecraft:patch_cactus_decorated",
				"minecraft:patch_firefly_bush_near_water",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: false,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:sheep",
					Weight: int32(12),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:pig",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:chicken",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:cow",
					Weight: int32(8),
				},
				{
					Maxcount: int32(2),
					Mincount: int32(1),
					Type: "minecraft:armadillo",
					Weight: int32(6),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(2.0),
	},
	"bamboo_jungle": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:audio/background_music": map[string]any{
				"default": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.overworld.bamboo_jungle",
				},
			},
			"minecraft:gameplay/increased_fire_burnout": true,
			"minecraft:visual/sky_color": "#77a8ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.9),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:bamboo",
				"minecraft:bamboo_vegetation",
				"minecraft:flower_warm",
				"minecraft:patch_grass_jungle",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
				"minecraft:vines",
				"minecraft:patch_melon",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:sheep",
					Weight: int32(12),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:pig",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:chicken",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:cow",
					Weight: int32(8),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:chicken",
					Weight: int32(10),
				},
				{
					Maxcount: int32(2),
					Mincount: int32(1),
					Type: "minecraft:parrot",
					Weight: int32(40),
				},
				{
					Maxcount: int32(2),
					Mincount: int32(1),
					Type: "minecraft:panda",
					Weight: int32(80),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:ocelot",
					Weight: int32(2),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(0.95),
	},
	"basalt_deltas": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:audio/ambient_sounds": map[string]any{
				"additions": map[string]any{
					"sound": "minecraft:ambient.basalt_deltas.additions",
					"tick_chance": float32(0.0111),
				},
				"loop": "minecraft:ambient.basalt_deltas.loop",
				"mood": map[string]any{
					"block_search_extent": int32(8),
					"offset": float32(2.0),
					"sound": "minecraft:ambient.basalt_deltas.mood",
					"tick_delay": int32(6000),
				},
			},
			"minecraft:audio/background_music": map[string]any{
				"default": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.nether.basalt_deltas",
				},
			},
			"minecraft:visual/ambient_particles": []any{
				map[string]any{
					"particle": map[string]any{
						"type": "minecraft:white_ash",
					},
					"probability": float32(0.118093334),
				},
			},
			"minecraft:visual/fog_color": "#685f70",
		},
		Carvers: "minecraft:nether_cave",
		Downfall: float32(0.0),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:delta",
				"minecraft:small_basalt_columns",
				"minecraft:large_basalt_columns",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:basalt_blobs",
				"minecraft:blackstone_blobs",
				"minecraft:spring_delta",
				"minecraft:patch_fire",
				"minecraft:patch_soul_fire",
				"minecraft:glowstone_extra",
				"minecraft:glowstone",
				"minecraft:brown_mushroom_nether",
				"minecraft:red_mushroom_nether",
				"minecraft:ore_magma",
				"minecraft:spring_closed_double",
				"minecraft:ore_gold_deltas",
				"minecraft:ore_quartz_deltas",
				"minecraft:ore_ancient_debris_large",
				"minecraft:ore_debris_small",
			},
		},
		HasPrecipitation: false,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(2),
					Mincount: int32(1),
					Type: "minecraft:strider",
					Weight: int32(60),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:ghast",
					Weight: int32(40),
				},
				{
					Maxcount: int32(5),
					Mincount: int32(2),
					Type: "minecraft:magma_cube",
					Weight: int32(100),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(2.0),
	},
	"beach": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:visual/sky_color": "#78a7ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.4),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:flower_default",
				"minecraft:patch_grass_badlands",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(5),
					Mincount: int32(2),
					Type: "minecraft:turtle",
					Weight: int32(5),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(0.8),
	},
	"birch_forest": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:audio/background_music": map[string]any{
				"default": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.overworld.forest",
				},
			},
			"minecraft:visual/sky_color": "#7aa5ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.6),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:forest_flowers",
				"minecraft:wildflowers_birch_forest",
				"minecraft:trees_birch",
				"minecraft:patch_bush",
				"minecraft:flower_default",
				"minecraft:patch_grass_forest",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:sheep",
					Weight: int32(12),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:pig",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:chicken",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:cow",
					Weight: int32(8),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(0.6),
	},
	"cherry_grove": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:audio/background_music": map[string]any{
				"default": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.overworld.cherry_grove",
				},
			},
			"minecraft:visual/sky_color": "#7ba4ff",
			"minecraft:visual/water_fog_color": "#5db7ef",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.8),
		Effects: map[string]any{
			"foliage_color": "#b6db61",
			"grass_color": "#b6db61",
			"water_color": "#5db7ef",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
				"minecraft:ore_emerald",
			},
			[]any{
				"minecraft:ore_infested",
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:patch_tall_grass_2",
				"minecraft:patch_grass_plain",
				"minecraft:flower_cherry",
				"minecraft:trees_cherry",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(2),
					Mincount: int32(1),
					Type: "minecraft:pig",
					Weight: int32(1),
				},
				{
					Maxcount: int32(6),
					Mincount: int32(2),
					Type: "minecraft:rabbit",
					Weight: int32(2),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(2),
					Type: "minecraft:sheep",
					Weight: int32(2),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(0.5),
	},
	"cold_ocean": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
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
				"underwater": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.under_water",
				},
			},
			"minecraft:visual/sky_color": "#7ba4ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.5),
		Effects: map[string]any{
			"water_color": "#3d57d6",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:trees_water",
				"minecraft:flower_default",
				"minecraft:patch_grass_badlands",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
				"minecraft:seagrass_cold",
				"minecraft:kelp_cold",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:drowned",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(3),
					Type: "minecraft:cod",
					Weight: int32(15),
				},
				{
					Maxcount: int32(5),
					Mincount: int32(1),
					Type: "minecraft:salmon",
					Weight: int32(15),
				},
			},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:squid",
					Weight: int32(3),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:nautilus",
					Weight: int32(2),
				},
			},
		},
		Temperature: float32(0.5),
	},
	"crimson_forest": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:audio/ambient_sounds": map[string]any{
				"additions": map[string]any{
					"sound": "minecraft:ambient.crimson_forest.additions",
					"tick_chance": float32(0.0111),
				},
				"loop": "minecraft:ambient.crimson_forest.loop",
				"mood": map[string]any{
					"block_search_extent": int32(8),
					"offset": float32(2.0),
					"sound": "minecraft:ambient.crimson_forest.mood",
					"tick_delay": int32(6000),
				},
			},
			"minecraft:audio/background_music": map[string]any{
				"default": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.nether.crimson_forest",
				},
			},
			"minecraft:visual/ambient_particles": []any{
				map[string]any{
					"particle": map[string]any{
						"type": "minecraft:crimson_spore",
					},
					"probability": float32(0.025),
				},
			},
			"minecraft:visual/fog_color": "#330303",
		},
		Carvers: "minecraft:nether_cave",
		Downfall: float32(0.0),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
			},
			[]any{
			},
			[]any{
			},
			[]any{
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:spring_open",
				"minecraft:patch_fire",
				"minecraft:glowstone_extra",
				"minecraft:glowstone",
				"minecraft:ore_magma",
				"minecraft:spring_closed",
				"minecraft:ore_gravel_nether",
				"minecraft:ore_blackstone",
				"minecraft:ore_gold_nether",
				"minecraft:ore_quartz_nether",
				"minecraft:ore_ancient_debris_large",
				"minecraft:ore_debris_small",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_lava",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:weeping_vines",
				"minecraft:crimson_fungi",
				"minecraft:crimson_forest_vegetation",
			},
		},
		HasPrecipitation: false,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(2),
					Mincount: int32(1),
					Type: "minecraft:strider",
					Weight: int32(60),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(2),
					Type: "minecraft:zombified_piglin",
					Weight: int32(1),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(3),
					Type: "minecraft:hoglin",
					Weight: int32(9),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(3),
					Type: "minecraft:piglin",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(2.0),
	},
	"dark_forest": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:audio/background_music": map[string]any{
				"default": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.overworld.forest",
				},
			},
			"minecraft:visual/sky_color": "#79a6ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.8),
		Effects: map[string]any{
			"dry_foliage_color": "#7b5334",
			"grass_color_modifier": "dark_forest",
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:dark_forest_vegetation",
				"minecraft:forest_flowers",
				"minecraft:flower_default",
				"minecraft:patch_grass_forest",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_leaf_litter",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:sheep",
					Weight: int32(12),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:pig",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:chicken",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:cow",
					Weight: int32(8),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(0.7),
	},
	"deep_cold_ocean": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
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
				"underwater": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.under_water",
				},
			},
			"minecraft:visual/sky_color": "#7ba4ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.5),
		Effects: map[string]any{
			"water_color": "#3d57d6",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:trees_water",
				"minecraft:flower_default",
				"minecraft:patch_grass_badlands",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
				"minecraft:seagrass_deep_cold",
				"minecraft:kelp_cold",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:drowned",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(3),
					Type: "minecraft:cod",
					Weight: int32(15),
				},
				{
					Maxcount: int32(5),
					Mincount: int32(1),
					Type: "minecraft:salmon",
					Weight: int32(15),
				},
			},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:squid",
					Weight: int32(3),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:nautilus",
					Weight: int32(2),
				},
			},
		},
		Temperature: float32(0.5),
	},
	"deep_dark": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:audio/background_music": map[string]any{
				"default": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.overworld.deep_dark",
				},
			},
			"minecraft:visual/sky_color": "#78a7ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.4),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
				"minecraft:sculk_vein",
				"minecraft:sculk_patch_deep_dark",
			},
			[]any{
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:patch_tall_grass_2",
				"minecraft:trees_plains",
				"minecraft:flower_plains",
				"minecraft:patch_grass_plain",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_pumpkin",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(0.8),
	},
	"deep_frozen_ocean": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
		TemperatureModifier string `nbt:"temperature_modifier"`
	}{
		Attributes: map[string]any{
			"minecraft:visual/sky_color": "#7ba4ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.5),
		Effects: map[string]any{
			"water_color": "#3938c9",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:iceberg_packed",
				"minecraft:iceberg_blue",
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
				"minecraft:blue_ice",
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:trees_water",
				"minecraft:flower_default",
				"minecraft:patch_grass_badlands",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(2),
					Mincount: int32(1),
					Type: "minecraft:polar_bear",
					Weight: int32(1),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:drowned",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(5),
					Mincount: int32(1),
					Type: "minecraft:salmon",
					Weight: int32(15),
				},
			},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:squid",
					Weight: int32(1),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:nautilus",
					Weight: int32(2),
				},
			},
		},
		Temperature: float32(0.5),
		TemperatureModifier: "frozen",
	},
	"deep_lukewarm_ocean": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
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
				"underwater": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.under_water",
				},
			},
			"minecraft:visual/sky_color": "#7ba4ff",
			"minecraft:visual/water_fog_color": "#041633",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.5),
		Effects: map[string]any{
			"water_color": "#45adf2",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:trees_water",
				"minecraft:flower_default",
				"minecraft:patch_grass_badlands",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
				"minecraft:seagrass_deep_warm",
				"minecraft:kelp_warm",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:drowned",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(3),
					Type: "minecraft:cod",
					Weight: int32(8),
				},
				{
					Maxcount: int32(3),
					Mincount: int32(1),
					Type: "minecraft:pufferfish",
					Weight: int32(5),
				},
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:tropical_fish",
					Weight: int32(25),
				},
			},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:squid",
					Weight: int32(8),
				},
				{
					Maxcount: int32(2),
					Mincount: int32(1),
					Type: "minecraft:dolphin",
					Weight: int32(2),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:nautilus",
					Weight: int32(5),
				},
			},
		},
		Temperature: float32(0.5),
	},
	"deep_ocean": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
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
				"underwater": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.under_water",
				},
			},
			"minecraft:visual/sky_color": "#7ba4ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.5),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:trees_water",
				"minecraft:flower_default",
				"minecraft:patch_grass_badlands",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
				"minecraft:seagrass_deep",
				"minecraft:kelp_cold",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:drowned",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(3),
					Type: "minecraft:cod",
					Weight: int32(10),
				},
			},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:squid",
					Weight: int32(1),
				},
				{
					Maxcount: int32(2),
					Mincount: int32(1),
					Type: "minecraft:dolphin",
					Weight: int32(1),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:nautilus",
					Weight: int32(5),
				},
			},
		},
		Temperature: float32(0.5),
	},
	"desert": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:audio/background_music": map[string]any{
				"default": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.overworld.desert",
				},
			},
			"minecraft:gameplay/snow_golem_melts": true,
			"minecraft:visual/sky_color": "#6eb1ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.0),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:fossil_upper",
				"minecraft:fossil_lower",
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
				"minecraft:desert_well",
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:flower_default",
				"minecraft:patch_grass_badlands",
				"minecraft:patch_dry_grass_desert",
				"minecraft:patch_dead_bush_2",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_sugar_cane_desert",
				"minecraft:patch_pumpkin",
				"minecraft:patch_cactus_desert",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: false,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(3),
					Mincount: int32(2),
					Type: "minecraft:rabbit",
					Weight: int32(12),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:camel",
					Weight: int32(1),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(19),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(1),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(50),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:husk",
					Weight: int32(80),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:parched",
					Weight: int32(50),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(2.0),
	},
	"dripstone_caves": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:audio/background_music": map[string]any{
				"default": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.overworld.dripstone_caves",
				},
			},
			"minecraft:visual/sky_color": "#78a7ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.4),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
				"minecraft:large_dripstone",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper_large",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
				"minecraft:dripstone_cluster",
				"minecraft:pointed_dripstone",
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:patch_tall_grass_2",
				"minecraft:trees_plains",
				"minecraft:flower_plains",
				"minecraft:patch_grass_plain",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_pumpkin",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:drowned",
					Weight: int32(95),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(0.8),
	},
	"end_barrens": struct {
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Carvers: []any{
		},
		Downfall: float32(0.5),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{},
		HasPrecipitation: false,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(0.5),
	},
	"end_highlands": struct {
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Carvers: []any{
		},
		Downfall: float32(0.5),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:end_gateway_return",
			},
			[]any{
			},
			[]any{
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:chorus_plant",
			},
		},
		HasPrecipitation: false,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(0.5),
	},
	"end_midlands": struct {
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Carvers: []any{
		},
		Downfall: float32(0.5),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{},
		HasPrecipitation: false,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(0.5),
	},
	"eroded_badlands": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		CreatureSpawnProbability float32 `nbt:"creature_spawn_probability"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:audio/background_music": map[string]any{
				"default": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.overworld.badlands",
				},
			},
			"minecraft:gameplay/snow_golem_melts": true,
			"minecraft:visual/sky_color": "#6eb1ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		CreatureSpawnProbability: float32(0.03),
		Downfall: float32(0.0),
		Effects: map[string]any{
			"foliage_color": "#9e814d",
			"grass_color": "#90814d",
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:ore_gold_extra",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:patch_grass_badlands",
				"minecraft:patch_dry_grass_badlands",
				"minecraft:patch_dead_bush_badlands",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_sugar_cane_badlands",
				"minecraft:patch_pumpkin",
				"minecraft:patch_cactus_decorated",
				"minecraft:patch_firefly_bush_near_water",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: false,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:sheep",
					Weight: int32(12),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:pig",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:chicken",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:cow",
					Weight: int32(8),
				},
				{
					Maxcount: int32(2),
					Mincount: int32(1),
					Type: "minecraft:armadillo",
					Weight: int32(6),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(2.0),
	},
	"flower_forest": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:audio/background_music": map[string]any{
				"default": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.overworld.flower_forest",
				},
			},
			"minecraft:visual/sky_color": "#79a6ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.8),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:flower_forest_flowers",
				"minecraft:trees_flower_forest",
				"minecraft:flower_flower_forest",
				"minecraft:patch_grass_badlands",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:sheep",
					Weight: int32(12),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:pig",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:chicken",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:cow",
					Weight: int32(8),
				},
				{
					Maxcount: int32(3),
					Mincount: int32(2),
					Type: "minecraft:rabbit",
					Weight: int32(4),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(0.7),
	},
	"forest": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:audio/background_music": map[string]any{
				"default": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.overworld.forest",
				},
			},
			"minecraft:visual/sky_color": "#79a6ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.8),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:forest_flowers",
				"minecraft:trees_birch_and_oak_leaf_litter",
				"minecraft:patch_bush",
				"minecraft:flower_default",
				"minecraft:patch_grass_forest",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:sheep",
					Weight: int32(12),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:pig",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:chicken",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:cow",
					Weight: int32(8),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:wolf",
					Weight: int32(5),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(0.7),
	},
	"frozen_ocean": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
		TemperatureModifier string `nbt:"temperature_modifier"`
	}{
		Attributes: map[string]any{
			"minecraft:visual/sky_color": "#7fa1ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.5),
		Effects: map[string]any{
			"water_color": "#3938c9",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:iceberg_packed",
				"minecraft:iceberg_blue",
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
				"minecraft:blue_ice",
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:trees_water",
				"minecraft:flower_default",
				"minecraft:patch_grass_badlands",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(2),
					Mincount: int32(1),
					Type: "minecraft:polar_bear",
					Weight: int32(1),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:drowned",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(5),
					Mincount: int32(1),
					Type: "minecraft:salmon",
					Weight: int32(15),
				},
			},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:squid",
					Weight: int32(1),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:nautilus",
					Weight: int32(2),
				},
			},
		},
		Temperature: float32(0.0),
		TemperatureModifier: "frozen",
	},
	"frozen_peaks": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:audio/background_music": map[string]any{
				"default": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.overworld.frozen_peaks",
				},
			},
			"minecraft:gameplay/increased_fire_burnout": true,
			"minecraft:visual/sky_color": "#859dff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.9),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
				"minecraft:ore_emerald",
			},
			[]any{
				"minecraft:ore_infested",
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
				"minecraft:spring_lava_frozen",
			},
			[]any{
				"minecraft:glow_lichen",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(3),
					Mincount: int32(1),
					Type: "minecraft:goat",
					Weight: int32(5),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(-0.7),
	},
	"frozen_river": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
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
				"underwater": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.under_water",
				},
			},
			"minecraft:visual/sky_color": "#7fa1ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.5),
		Effects: map[string]any{
			"water_color": "#3938c9",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:trees_water",
				"minecraft:patch_bush",
				"minecraft:flower_default",
				"minecraft:patch_grass_badlands",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:drowned",
					Weight: int32(1),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(5),
					Mincount: int32(1),
					Type: "minecraft:salmon",
					Weight: int32(5),
				},
			},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:squid",
					Weight: int32(2),
				},
			},
		},
		Temperature: float32(0.0),
	},
	"grove": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:audio/background_music": map[string]any{
				"default": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.overworld.grove",
				},
			},
			"minecraft:visual/sky_color": "#81a0ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.8),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
				"minecraft:ore_emerald",
			},
			[]any{
				"minecraft:ore_infested",
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
				"minecraft:spring_lava_frozen",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:trees_grove",
				"minecraft:patch_pumpkin",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:wolf",
					Weight: int32(1),
				},
				{
					Maxcount: int32(3),
					Mincount: int32(2),
					Type: "minecraft:rabbit",
					Weight: int32(8),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(2),
					Type: "minecraft:fox",
					Weight: int32(4),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(-0.2),
	},
	"ice_spikes": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		CreatureSpawnProbability float32 `nbt:"creature_spawn_probability"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:visual/sky_color": "#7fa1ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		CreatureSpawnProbability: float32(0.07),
		Downfall: float32(0.5),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
				"minecraft:ice_spike",
				"minecraft:ice_patch",
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:trees_snowy",
				"minecraft:flower_default",
				"minecraft:patch_grass_badlands",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(3),
					Mincount: int32(2),
					Type: "minecraft:rabbit",
					Weight: int32(10),
				},
				{
					Maxcount: int32(2),
					Mincount: int32(1),
					Type: "minecraft:polar_bear",
					Weight: int32(1),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(20),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:stray",
					Weight: int32(80),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(0.0),
	},
	"jagged_peaks": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:audio/background_music": map[string]any{
				"default": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.overworld.jagged_peaks",
				},
			},
			"minecraft:gameplay/increased_fire_burnout": true,
			"minecraft:visual/sky_color": "#859dff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.9),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
				"minecraft:ore_emerald",
			},
			[]any{
				"minecraft:ore_infested",
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
				"minecraft:spring_lava_frozen",
			},
			[]any{
				"minecraft:glow_lichen",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(3),
					Mincount: int32(1),
					Type: "minecraft:goat",
					Weight: int32(5),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(-0.7),
	},
	"jungle": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:audio/background_music": map[string]any{
				"default": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.overworld.jungle",
				},
			},
			"minecraft:gameplay/increased_fire_burnout": true,
			"minecraft:visual/sky_color": "#77a8ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.9),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:bamboo_light",
				"minecraft:trees_jungle",
				"minecraft:flower_warm",
				"minecraft:patch_grass_jungle",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
				"minecraft:vines",
				"minecraft:patch_melon",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:sheep",
					Weight: int32(12),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:pig",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:chicken",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:cow",
					Weight: int32(8),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:chicken",
					Weight: int32(10),
				},
				{
					Maxcount: int32(2),
					Mincount: int32(1),
					Type: "minecraft:parrot",
					Weight: int32(40),
				},
				{
					Maxcount: int32(2),
					Mincount: int32(1),
					Type: "minecraft:panda",
					Weight: int32(1),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
				{
					Maxcount: int32(3),
					Mincount: int32(1),
					Type: "minecraft:ocelot",
					Weight: int32(2),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(0.95),
	},
	"lukewarm_ocean": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
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
				"underwater": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.under_water",
				},
			},
			"minecraft:visual/sky_color": "#7ba4ff",
			"minecraft:visual/water_fog_color": "#041633",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.5),
		Effects: map[string]any{
			"water_color": "#45adf2",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:trees_water",
				"minecraft:flower_default",
				"minecraft:patch_grass_badlands",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
				"minecraft:seagrass_warm",
				"minecraft:kelp_warm",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:drowned",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(3),
					Type: "minecraft:cod",
					Weight: int32(15),
				},
				{
					Maxcount: int32(3),
					Mincount: int32(1),
					Type: "minecraft:pufferfish",
					Weight: int32(5),
				},
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:tropical_fish",
					Weight: int32(25),
				},
			},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(2),
					Mincount: int32(1),
					Type: "minecraft:squid",
					Weight: int32(10),
				},
				{
					Maxcount: int32(2),
					Mincount: int32(1),
					Type: "minecraft:dolphin",
					Weight: int32(2),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:nautilus",
					Weight: int32(5),
				},
			},
		},
		Temperature: float32(0.5),
	},
	"lush_caves": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:audio/background_music": map[string]any{
				"default": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.overworld.lush_caves",
				},
			},
			"minecraft:visual/sky_color": "#7ba4ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.5),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:ore_clay",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:patch_tall_grass_2",
				"minecraft:lush_caves_ceiling_vegetation",
				"minecraft:cave_vines",
				"minecraft:lush_caves_clay",
				"minecraft:lush_caves_vegetation",
				"minecraft:rooted_azalea_tree",
				"minecraft:spore_blossom",
				"minecraft:classic_vines_cave_feature",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:axolotl",
					Weight: int32(10),
				},
			},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:tropical_fish",
					Weight: int32(25),
				},
			},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(0.5),
	},
	"mangrove_swamp": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:audio/background_music": map[string]any{
				"default": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.overworld.swamp",
				},
			},
			"minecraft:gameplay/increased_fire_burnout": true,
			"minecraft:visual/fog_color": "#c0d8ff",
			"minecraft:visual/sky_color": "#78a7ff",
			"minecraft:visual/water_fog_color": "#4d7a60",
			"minecraft:visual/water_fog_end_distance": map[string]any{
				"argument": float32(0.85),
				"modifier": "multiply",
			},
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.9),
		Effects: map[string]any{
			"dry_foliage_color": "#7b5334",
			"foliage_color": "#8db127",
			"grass_color_modifier": "swamp",
			"water_color": "#3a7a6a",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:fossil_upper",
				"minecraft:fossil_lower",
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_grass",
				"minecraft:disk_clay",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:trees_mangrove",
				"minecraft:patch_grass_normal",
				"minecraft:patch_dead_bush",
				"minecraft:patch_waterlily",
				"minecraft:seagrass_swamp",
				"minecraft:patch_firefly_bush_near_water",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(5),
					Mincount: int32(2),
					Type: "minecraft:frog",
					Weight: int32(10),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(70),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:slime",
					Weight: int32(1),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:bogged",
					Weight: int32(30),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:tropical_fish",
					Weight: int32(25),
				},
			},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(0.8),
	},
	"meadow": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:audio/background_music": map[string]any{
				"default": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.overworld.meadow",
				},
			},
			"minecraft:visual/sky_color": "#7ba4ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.8),
		Effects: map[string]any{
			"water_color": "#0e4ecf",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
				"minecraft:ore_emerald",
			},
			[]any{
				"minecraft:ore_infested",
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:patch_tall_grass_2",
				"minecraft:patch_grass_meadow",
				"minecraft:flower_meadow",
				"minecraft:trees_meadow",
				"minecraft:wildflowers_meadow",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(2),
					Mincount: int32(1),
					Type: "minecraft:donkey",
					Weight: int32(1),
				},
				{
					Maxcount: int32(6),
					Mincount: int32(2),
					Type: "minecraft:rabbit",
					Weight: int32(2),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(2),
					Type: "minecraft:sheep",
					Weight: int32(2),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(0.5),
	},
	"mushroom_fields": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:gameplay/can_pillager_patrol_spawn": false,
			"minecraft:gameplay/increased_fire_burnout": true,
			"minecraft:visual/sky_color": "#77a8ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(1.0),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:mushroom_island_vegetation",
				"minecraft:brown_mushroom_taiga",
				"minecraft:red_mushroom_taiga",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(4),
					Type: "minecraft:mooshroom",
					Weight: int32(8),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(0.9),
	},
	"nether_wastes": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:audio/ambient_sounds": map[string]any{
				"additions": map[string]any{
					"sound": "minecraft:ambient.nether_wastes.additions",
					"tick_chance": float32(0.0111),
				},
				"loop": "minecraft:ambient.nether_wastes.loop",
				"mood": map[string]any{
					"block_search_extent": int32(8),
					"offset": float32(2.0),
					"sound": "minecraft:ambient.nether_wastes.mood",
					"tick_delay": int32(6000),
				},
			},
			"minecraft:audio/background_music": map[string]any{
				"default": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.nether.nether_wastes",
				},
			},
			"minecraft:visual/fog_color": "#330808",
		},
		Carvers: "minecraft:nether_cave",
		Downfall: float32(0.0),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
			},
			[]any{
			},
			[]any{
			},
			[]any{
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:spring_open",
				"minecraft:patch_fire",
				"minecraft:patch_soul_fire",
				"minecraft:glowstone_extra",
				"minecraft:glowstone",
				"minecraft:brown_mushroom_nether",
				"minecraft:red_mushroom_nether",
				"minecraft:ore_magma",
				"minecraft:spring_closed",
				"minecraft:ore_gravel_nether",
				"minecraft:ore_blackstone",
				"minecraft:ore_gold_nether",
				"minecraft:ore_quartz_nether",
				"minecraft:ore_ancient_debris_large",
				"minecraft:ore_debris_small",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_lava",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
			},
		},
		HasPrecipitation: false,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(2),
					Mincount: int32(1),
					Type: "minecraft:strider",
					Weight: int32(60),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:ghast",
					Weight: int32(50),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombified_piglin",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:magma_cube",
					Weight: int32(2),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:enderman",
					Weight: int32(1),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:piglin",
					Weight: int32(15),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(2.0),
	},
	"ocean": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
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
				"underwater": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.under_water",
				},
			},
			"minecraft:visual/sky_color": "#7ba4ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.5),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:trees_water",
				"minecraft:flower_default",
				"minecraft:patch_grass_badlands",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
				"minecraft:seagrass_normal",
				"minecraft:kelp_cold",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:drowned",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(3),
					Type: "minecraft:cod",
					Weight: int32(10),
				},
			},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:squid",
					Weight: int32(1),
				},
				{
					Maxcount: int32(2),
					Mincount: int32(1),
					Type: "minecraft:dolphin",
					Weight: int32(1),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:nautilus",
					Weight: int32(5),
				},
			},
		},
		Temperature: float32(0.5),
	},
	"old_growth_birch_forest": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:audio/background_music": map[string]any{
				"default": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.overworld.forest",
				},
			},
			"minecraft:visual/sky_color": "#7aa5ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.6),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:forest_flowers",
				"minecraft:wildflowers_birch_forest",
				"minecraft:birch_tall",
				"minecraft:patch_bush",
				"minecraft:flower_default",
				"minecraft:patch_grass_forest",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:sheep",
					Weight: int32(12),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:pig",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:chicken",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:cow",
					Weight: int32(8),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(0.6),
	},
	"old_growth_pine_taiga": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:audio/background_music": map[string]any{
				"default": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.overworld.old_growth_taiga",
				},
			},
			"minecraft:visual/sky_color": "#7ca3ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.8),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
				"minecraft:forest_rock",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:patch_large_fern",
				"minecraft:trees_old_growth_pine_taiga",
				"minecraft:flower_default",
				"minecraft:patch_grass_taiga",
				"minecraft:patch_dead_bush",
				"minecraft:brown_mushroom_old_growth",
				"minecraft:red_mushroom_old_growth",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
				"minecraft:patch_berry_common",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:sheep",
					Weight: int32(12),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:pig",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:chicken",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:cow",
					Weight: int32(8),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:wolf",
					Weight: int32(8),
				},
				{
					Maxcount: int32(3),
					Mincount: int32(2),
					Type: "minecraft:rabbit",
					Weight: int32(4),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(2),
					Type: "minecraft:fox",
					Weight: int32(8),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(100),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(25),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(0.3),
	},
	"old_growth_spruce_taiga": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:audio/background_music": map[string]any{
				"default": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.overworld.old_growth_taiga",
				},
			},
			"minecraft:visual/sky_color": "#7da3ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.8),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
				"minecraft:forest_rock",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:patch_large_fern",
				"minecraft:trees_old_growth_spruce_taiga",
				"minecraft:flower_default",
				"minecraft:patch_grass_taiga",
				"minecraft:patch_dead_bush",
				"minecraft:brown_mushroom_old_growth",
				"minecraft:red_mushroom_old_growth",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
				"minecraft:patch_berry_common",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:sheep",
					Weight: int32(12),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:pig",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:chicken",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:cow",
					Weight: int32(8),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:wolf",
					Weight: int32(8),
				},
				{
					Maxcount: int32(3),
					Mincount: int32(2),
					Type: "minecraft:rabbit",
					Weight: int32(4),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(2),
					Type: "minecraft:fox",
					Weight: int32(8),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(0.25),
	},
	"pale_garden": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:audio/background_music": map[string]any{
			},
			"minecraft:audio/music_volume": float32(0.0),
			"minecraft:visual/fog_color": "#817770",
			"minecraft:visual/sky_color": "#b9b9b9",
			"minecraft:visual/water_fog_color": "#556980",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.8),
		Effects: map[string]any{
			"dry_foliage_color": "#a0a69c",
			"foliage_color": "#878d76",
			"grass_color": "#778272",
			"water_color": "#76889d",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:pale_garden_vegetation",
				"minecraft:pale_moss_patch",
				"minecraft:pale_garden_flowers",
				"minecraft:flower_pale_garden",
				"minecraft:patch_grass_forest",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(0.7),
	},
	"plains": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:visual/sky_color": "#78a7ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.4),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:patch_tall_grass_2",
				"minecraft:patch_bush",
				"minecraft:trees_plains",
				"minecraft:flower_plains",
				"minecraft:patch_grass_plain",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:sheep",
					Weight: int32(12),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:pig",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:chicken",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:cow",
					Weight: int32(8),
				},
				{
					Maxcount: int32(6),
					Mincount: int32(2),
					Type: "minecraft:horse",
					Weight: int32(5),
				},
				{
					Maxcount: int32(3),
					Mincount: int32(1),
					Type: "minecraft:donkey",
					Weight: int32(1),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(90),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_horse",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(0.8),
	},
	"river": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
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
				"underwater": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.under_water",
				},
			},
			"minecraft:visual/sky_color": "#7ba4ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.5),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:trees_water",
				"minecraft:patch_bush",
				"minecraft:flower_default",
				"minecraft:patch_grass_badlands",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
				"minecraft:seagrass_river",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:drowned",
					Weight: int32(100),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(5),
					Mincount: int32(1),
					Type: "minecraft:salmon",
					Weight: int32(5),
				},
			},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:squid",
					Weight: int32(2),
				},
			},
		},
		Temperature: float32(0.5),
	},
	"savanna": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:gameplay/snow_golem_melts": true,
			"minecraft:visual/sky_color": "#6eb1ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.0),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:patch_tall_grass",
				"minecraft:trees_savanna",
				"minecraft:flower_warm",
				"minecraft:patch_grass_savanna",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: false,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:sheep",
					Weight: int32(12),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:pig",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:chicken",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:cow",
					Weight: int32(8),
				},
				{
					Maxcount: int32(6),
					Mincount: int32(2),
					Type: "minecraft:horse",
					Weight: int32(1),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:donkey",
					Weight: int32(1),
				},
				{
					Maxcount: int32(3),
					Mincount: int32(2),
					Type: "minecraft:armadillo",
					Weight: int32(10),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(90),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_horse",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(2.0),
	},
	"savanna_plateau": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:gameplay/snow_golem_melts": true,
			"minecraft:visual/sky_color": "#6eb1ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.0),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:patch_tall_grass",
				"minecraft:trees_savanna",
				"minecraft:flower_warm",
				"minecraft:patch_grass_savanna",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: false,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:sheep",
					Weight: int32(12),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:pig",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:chicken",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:cow",
					Weight: int32(8),
				},
				{
					Maxcount: int32(6),
					Mincount: int32(2),
					Type: "minecraft:horse",
					Weight: int32(1),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:donkey",
					Weight: int32(1),
				},
				{
					Maxcount: int32(3),
					Mincount: int32(2),
					Type: "minecraft:armadillo",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:llama",
					Weight: int32(8),
				},
				{
					Maxcount: int32(8),
					Mincount: int32(4),
					Type: "minecraft:wolf",
					Weight: int32(8),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(90),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_horse",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(2.0),
	},
	"small_end_islands": struct {
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Carvers: []any{
		},
		Downfall: float32(0.5),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
				"minecraft:end_island_decorated",
			},
		},
		HasPrecipitation: false,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(0.5),
	},
	"snowy_beach": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:visual/sky_color": "#7fa1ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.3),
		Effects: map[string]any{
			"water_color": "#3d57d6",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:flower_default",
				"minecraft:patch_grass_badlands",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(0.05),
	},
	"snowy_plains": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		CreatureSpawnProbability float32 `nbt:"creature_spawn_probability"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:visual/sky_color": "#7fa1ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		CreatureSpawnProbability: float32(0.07),
		Downfall: float32(0.5),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:trees_snowy",
				"minecraft:flower_default",
				"minecraft:patch_grass_badlands",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(3),
					Mincount: int32(2),
					Type: "minecraft:rabbit",
					Weight: int32(10),
				},
				{
					Maxcount: int32(2),
					Mincount: int32(1),
					Type: "minecraft:polar_bear",
					Weight: int32(1),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(90),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_horse",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(20),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:stray",
					Weight: int32(80),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(0.0),
	},
	"snowy_slopes": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:audio/background_music": map[string]any{
				"default": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.overworld.snowy_slopes",
				},
			},
			"minecraft:gameplay/increased_fire_burnout": true,
			"minecraft:visual/sky_color": "#829fff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.9),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
				"minecraft:ore_emerald",
			},
			[]any{
				"minecraft:ore_infested",
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
				"minecraft:spring_lava_frozen",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:patch_pumpkin",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(3),
					Mincount: int32(2),
					Type: "minecraft:rabbit",
					Weight: int32(4),
				},
				{
					Maxcount: int32(3),
					Mincount: int32(1),
					Type: "minecraft:goat",
					Weight: int32(5),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(-0.3),
	},
	"snowy_taiga": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:visual/sky_color": "#839eff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.4),
		Effects: map[string]any{
			"water_color": "#3d57d6",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:patch_large_fern",
				"minecraft:trees_taiga",
				"minecraft:flower_default",
				"minecraft:patch_grass_taiga_2",
				"minecraft:brown_mushroom_taiga",
				"minecraft:red_mushroom_taiga",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
				"minecraft:patch_berry_rare",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:sheep",
					Weight: int32(12),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:pig",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:chicken",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:cow",
					Weight: int32(8),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:wolf",
					Weight: int32(8),
				},
				{
					Maxcount: int32(3),
					Mincount: int32(2),
					Type: "minecraft:rabbit",
					Weight: int32(4),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(2),
					Type: "minecraft:fox",
					Weight: int32(8),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(-0.5),
	},
	"soul_sand_valley": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:audio/ambient_sounds": map[string]any{
				"additions": map[string]any{
					"sound": "minecraft:ambient.soul_sand_valley.additions",
					"tick_chance": float32(0.0111),
				},
				"loop": "minecraft:ambient.soul_sand_valley.loop",
				"mood": map[string]any{
					"block_search_extent": int32(8),
					"offset": float32(2.0),
					"sound": "minecraft:ambient.soul_sand_valley.mood",
					"tick_delay": int32(6000),
				},
			},
			"minecraft:audio/background_music": map[string]any{
				"default": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.nether.soul_sand_valley",
				},
			},
			"minecraft:visual/ambient_particles": []any{
				map[string]any{
					"particle": map[string]any{
						"type": "minecraft:ash",
					},
					"probability": float32(0.00625),
				},
			},
			"minecraft:visual/fog_color": "#1b4745",
		},
		Carvers: "minecraft:nether_cave",
		Downfall: float32(0.0),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:basalt_pillar",
			},
			[]any{
			},
			[]any{
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:spring_open",
				"minecraft:patch_fire",
				"minecraft:patch_soul_fire",
				"minecraft:glowstone_extra",
				"minecraft:glowstone",
				"minecraft:patch_crimson_roots",
				"minecraft:ore_magma",
				"minecraft:spring_closed",
				"minecraft:ore_soul_sand",
				"minecraft:ore_gravel_nether",
				"minecraft:ore_blackstone",
				"minecraft:ore_gold_nether",
				"minecraft:ore_quartz_nether",
				"minecraft:ore_ancient_debris_large",
				"minecraft:ore_debris_small",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_lava",
			},
		},
		HasPrecipitation: false,
		SpawnCosts: map[string]any{
			"minecraft:enderman": map[string]any{
				"charge": float32(0.7),
				"energy_budget": float32(0.15),
			},
			"minecraft:ghast": map[string]any{
				"charge": float32(0.7),
				"energy_budget": float32(0.15),
			},
			"minecraft:skeleton": map[string]any{
				"charge": float32(0.7),
				"energy_budget": float32(0.15),
			},
			"minecraft:strider": map[string]any{
				"charge": float32(0.7),
				"energy_budget": float32(0.15),
			},
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(2),
					Mincount: int32(1),
					Type: "minecraft:strider",
					Weight: int32(60),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(5),
					Mincount: int32(5),
					Type: "minecraft:skeleton",
					Weight: int32(20),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:ghast",
					Weight: int32(50),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:enderman",
					Weight: int32(1),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(2.0),
	},
	"sparse_jungle": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:audio/background_music": map[string]any{
				"default": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.overworld.sparse_jungle",
				},
			},
			"minecraft:visual/sky_color": "#77a8ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.8),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:trees_sparse_jungle",
				"minecraft:flower_warm",
				"minecraft:patch_grass_jungle",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
				"minecraft:vines",
				"minecraft:patch_melon_sparse",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:sheep",
					Weight: int32(12),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:pig",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:chicken",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:cow",
					Weight: int32(8),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:chicken",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(2),
					Type: "minecraft:wolf",
					Weight: int32(8),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(0.95),
	},
	"stony_peaks": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:audio/background_music": map[string]any{
				"default": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.overworld.stony_peaks",
				},
			},
			"minecraft:visual/sky_color": "#76a8ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.3),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
				"minecraft:ore_emerald",
			},
			[]any{
				"minecraft:ore_infested",
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(1.0),
	},
	"stony_shore": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:visual/sky_color": "#7da2ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.3),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:flower_default",
				"minecraft:patch_grass_badlands",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(0.2),
	},
	"sunflower_plains": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:visual/sky_color": "#78a7ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.4),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:patch_tall_grass_2",
				"minecraft:patch_sunflower",
				"minecraft:trees_plains",
				"minecraft:flower_plains",
				"minecraft:patch_grass_plain",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:sheep",
					Weight: int32(12),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:pig",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:chicken",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:cow",
					Weight: int32(8),
				},
				{
					Maxcount: int32(6),
					Mincount: int32(2),
					Type: "minecraft:horse",
					Weight: int32(5),
				},
				{
					Maxcount: int32(3),
					Mincount: int32(1),
					Type: "minecraft:donkey",
					Weight: int32(1),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(90),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_horse",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(0.8),
	},
	"swamp": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:audio/background_music": map[string]any{
				"default": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.overworld.swamp",
				},
			},
			"minecraft:gameplay/increased_fire_burnout": true,
			"minecraft:visual/sky_color": "#78a7ff",
			"minecraft:visual/water_fog_color": "#232317",
			"minecraft:visual/water_fog_end_distance": map[string]any{
				"argument": float32(0.85),
				"modifier": "multiply",
			},
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.9),
		Effects: map[string]any{
			"dry_foliage_color": "#7b5334",
			"foliage_color": "#6a7039",
			"grass_color_modifier": "swamp",
			"water_color": "#617b64",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:fossil_upper",
				"minecraft:fossil_lower",
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_clay",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:trees_swamp",
				"minecraft:flower_swamp",
				"minecraft:patch_grass_normal",
				"minecraft:patch_dead_bush",
				"minecraft:patch_waterlily",
				"minecraft:brown_mushroom_swamp",
				"minecraft:red_mushroom_swamp",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_sugar_cane_swamp",
				"minecraft:patch_pumpkin",
				"minecraft:patch_firefly_bush_swamp",
				"minecraft:patch_firefly_bush_near_water_swamp",
				"minecraft:seagrass_swamp",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:sheep",
					Weight: int32(12),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:pig",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:chicken",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:cow",
					Weight: int32(8),
				},
				{
					Maxcount: int32(5),
					Mincount: int32(2),
					Type: "minecraft:frog",
					Weight: int32(10),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(70),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:slime",
					Weight: int32(1),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:bogged",
					Weight: int32(30),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(0.8),
	},
	"taiga": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:visual/sky_color": "#7da3ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.8),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:patch_large_fern",
				"minecraft:trees_taiga",
				"minecraft:flower_default",
				"minecraft:patch_grass_taiga_2",
				"minecraft:brown_mushroom_taiga",
				"minecraft:red_mushroom_taiga",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
				"minecraft:patch_berry_common",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:sheep",
					Weight: int32(12),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:pig",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:chicken",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:cow",
					Weight: int32(8),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:wolf",
					Weight: int32(8),
				},
				{
					Maxcount: int32(3),
					Mincount: int32(2),
					Type: "minecraft:rabbit",
					Weight: int32(4),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(2),
					Type: "minecraft:fox",
					Weight: int32(8),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(0.25),
	},
	"the_end": struct {
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Carvers: []any{
		},
		Downfall: float32(0.5),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:end_spike",
			},
			[]any{
			},
			[]any{
			},
			[]any{
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:end_platform",
			},
		},
		HasPrecipitation: false,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(0.5),
	},
	"the_void": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:visual/sky_color": "#7ba4ff",
		},
		Carvers: []any{
		},
		Downfall: float32(0.5),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
			},
			[]any{
			},
			[]any{
			},
			[]any{
			},
			[]any{
			},
			[]any{
			},
			[]any{
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:void_start_platform",
			},
		},
		HasPrecipitation: false,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(0.5),
	},
	"warm_ocean": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
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
				"underwater": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.under_water",
				},
			},
			"minecraft:visual/sky_color": "#7ba4ff",
			"minecraft:visual/water_fog_color": "#041f33",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.5),
		Effects: map[string]any{
			"water_color": "#43d5ee",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:trees_water",
				"minecraft:flower_default",
				"minecraft:patch_grass_badlands",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
				"minecraft:warm_ocean_vegetation",
				"minecraft:seagrass_warm",
				"minecraft:sea_pickle",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:drowned",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(3),
					Mincount: int32(1),
					Type: "minecraft:pufferfish",
					Weight: int32(15),
				},
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:tropical_fish",
					Weight: int32(25),
				},
			},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:nautilus",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:squid",
					Weight: int32(10),
				},
				{
					Maxcount: int32(2),
					Mincount: int32(1),
					Type: "minecraft:dolphin",
					Weight: int32(2),
				},
			},
		},
		Temperature: float32(0.5),
	},
	"warped_forest": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:audio/ambient_sounds": map[string]any{
				"additions": map[string]any{
					"sound": "minecraft:ambient.warped_forest.additions",
					"tick_chance": float32(0.0111),
				},
				"loop": "minecraft:ambient.warped_forest.loop",
				"mood": map[string]any{
					"block_search_extent": int32(8),
					"offset": float32(2.0),
					"sound": "minecraft:ambient.warped_forest.mood",
					"tick_delay": int32(6000),
				},
			},
			"minecraft:audio/background_music": map[string]any{
				"default": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.nether.warped_forest",
				},
			},
			"minecraft:visual/ambient_particles": []any{
				map[string]any{
					"particle": map[string]any{
						"type": "minecraft:warped_spore",
					},
					"probability": float32(0.01428),
				},
			},
			"minecraft:visual/fog_color": "#1a051a",
		},
		Carvers: "minecraft:nether_cave",
		Downfall: float32(0.0),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
			},
			[]any{
			},
			[]any{
			},
			[]any{
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:spring_open",
				"minecraft:patch_fire",
				"minecraft:patch_soul_fire",
				"minecraft:glowstone_extra",
				"minecraft:glowstone",
				"minecraft:ore_magma",
				"minecraft:spring_closed",
				"minecraft:ore_gravel_nether",
				"minecraft:ore_blackstone",
				"minecraft:ore_gold_nether",
				"minecraft:ore_quartz_nether",
				"minecraft:ore_ancient_debris_large",
				"minecraft:ore_debris_small",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_lava",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:warped_fungi",
				"minecraft:warped_forest_vegetation",
				"minecraft:nether_sprouts",
				"minecraft:twisting_vines",
			},
		},
		HasPrecipitation: false,
		SpawnCosts: map[string]any{
			"minecraft:enderman": map[string]any{
				"charge": float32(1.0),
				"energy_budget": float32(0.12),
			},
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(2),
					Mincount: int32(1),
					Type: "minecraft:strider",
					Weight: int32(60),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:enderman",
					Weight: int32(1),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(2.0),
	},
	"windswept_forest": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:visual/sky_color": "#7da2ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.3),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
				"minecraft:ore_emerald",
			},
			[]any{
				"minecraft:ore_infested",
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:trees_windswept_forest",
				"minecraft:patch_bush",
				"minecraft:flower_default",
				"minecraft:patch_grass_badlands",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:sheep",
					Weight: int32(12),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:pig",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:chicken",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:cow",
					Weight: int32(8),
				},
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:llama",
					Weight: int32(5),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(0.2),
	},
	"windswept_gravelly_hills": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:visual/sky_color": "#7da2ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.3),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
				"minecraft:ore_emerald",
			},
			[]any{
				"minecraft:ore_infested",
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:trees_windswept_hills",
				"minecraft:patch_bush",
				"minecraft:flower_default",
				"minecraft:patch_grass_badlands",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:sheep",
					Weight: int32(12),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:pig",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:chicken",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:cow",
					Weight: int32(8),
				},
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:llama",
					Weight: int32(5),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(0.2),
	},
	"windswept_hills": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:visual/sky_color": "#7da2ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.3),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
				"minecraft:ore_emerald",
			},
			[]any{
				"minecraft:ore_infested",
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:trees_windswept_hills",
				"minecraft:patch_bush",
				"minecraft:flower_default",
				"minecraft:patch_grass_badlands",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: true,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:sheep",
					Weight: int32(12),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:pig",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:chicken",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:cow",
					Weight: int32(8),
				},
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:llama",
					Weight: int32(5),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(0.2),
	},
	"windswept_savanna": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:gameplay/snow_golem_melts": true,
			"minecraft:visual/sky_color": "#6eb1ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		Downfall: float32(0.0),
		Effects: map[string]any{
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:trees_windswept_savanna",
				"minecraft:flower_default",
				"minecraft:patch_grass_normal",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_pumpkin",
				"minecraft:patch_sugar_cane",
				"minecraft:patch_firefly_bush_near_water",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: false,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:sheep",
					Weight: int32(12),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:pig",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:chicken",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:cow",
					Weight: int32(8),
				},
				{
					Maxcount: int32(6),
					Mincount: int32(2),
					Type: "minecraft:horse",
					Weight: int32(1),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:donkey",
					Weight: int32(1),
				},
				{
					Maxcount: int32(3),
					Mincount: int32(2),
					Type: "minecraft:armadillo",
					Weight: int32(10),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(90),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_horse",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(2.0),
	},
	"wooded_badlands": struct {
		Attributes any `nbt:"attributes"`
		Carvers any `nbt:"carvers"`
		CreatureSpawnProbability float32 `nbt:"creature_spawn_probability"`
		Downfall float32 `nbt:"downfall"`
		Effects any `nbt:"effects"`
		Features [][]any `nbt:"features"`
		HasPrecipitation bool `nbt:"has_precipitation"`
		SpawnCosts any `nbt:"spawn_costs"`
		Spawners struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		} `nbt:"spawners"`
		Temperature float32 `nbt:"temperature"`
	}{
		Attributes: map[string]any{
			"minecraft:audio/background_music": map[string]any{
				"default": map[string]any{
					"max_delay": int32(24000),
					"min_delay": int32(12000),
					"sound": "minecraft:music.overworld.badlands",
				},
			},
			"minecraft:gameplay/snow_golem_melts": true,
			"minecraft:visual/sky_color": "#6eb1ff",
		},
		Carvers: []any{
			"minecraft:cave",
			"minecraft:cave_extra_underground",
			"minecraft:canyon",
		},
		CreatureSpawnProbability: float32(0.04),
		Downfall: float32(0.0),
		Effects: map[string]any{
			"foliage_color": "#9e814d",
			"grass_color": "#90814d",
			"water_color": "#3f76e4",
		},
		Features: [][]any{
			[]any{
			},
			[]any{
				"minecraft:lake_lava_underground",
				"minecraft:lake_lava_surface",
			},
			[]any{
				"minecraft:amethyst_geode",
			},
			[]any{
				"minecraft:monster_room",
				"minecraft:monster_room_deep",
			},
			[]any{
			},
			[]any{
			},
			[]any{
				"minecraft:ore_dirt",
				"minecraft:ore_gravel",
				"minecraft:ore_granite_upper",
				"minecraft:ore_granite_lower",
				"minecraft:ore_diorite_upper",
				"minecraft:ore_diorite_lower",
				"minecraft:ore_andesite_upper",
				"minecraft:ore_andesite_lower",
				"minecraft:ore_tuff",
				"minecraft:ore_coal_upper",
				"minecraft:ore_coal_lower",
				"minecraft:ore_iron_upper",
				"minecraft:ore_iron_middle",
				"minecraft:ore_iron_small",
				"minecraft:ore_gold",
				"minecraft:ore_gold_lower",
				"minecraft:ore_redstone",
				"minecraft:ore_redstone_lower",
				"minecraft:ore_diamond",
				"minecraft:ore_diamond_medium",
				"minecraft:ore_diamond_large",
				"minecraft:ore_diamond_buried",
				"minecraft:ore_lapis",
				"minecraft:ore_lapis_buried",
				"minecraft:ore_copper",
				"minecraft:underwater_magma",
				"minecraft:ore_gold_extra",
				"minecraft:disk_sand",
				"minecraft:disk_clay",
				"minecraft:disk_gravel",
			},
			[]any{
			},
			[]any{
				"minecraft:spring_water",
				"minecraft:spring_lava",
			},
			[]any{
				"minecraft:glow_lichen",
				"minecraft:trees_badlands",
				"minecraft:patch_grass_badlands",
				"minecraft:patch_dry_grass_badlands",
				"minecraft:patch_dead_bush_badlands",
				"minecraft:brown_mushroom_normal",
				"minecraft:red_mushroom_normal",
				"minecraft:patch_sugar_cane_badlands",
				"minecraft:patch_pumpkin",
				"minecraft:patch_cactus_decorated",
				"minecraft:patch_firefly_bush_near_water",
			},
			[]any{
				"minecraft:freeze_top_layer",
			},
		},
		HasPrecipitation: false,
		SpawnCosts: map[string]any{
		},
		Spawners: struct {
			Ambient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"ambient"`
			Axolotls []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"axolotls"`
			Creature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"creature"`
			Misc []any `nbt:"misc"`
			Monster []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"monster"`
			UndergroundWaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"underground_water_creature"`
			WaterAmbient []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_ambient"`
			WaterCreature []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			} `nbt:"water_creature"`
		}{
			Ambient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(8),
					Mincount: int32(8),
					Type: "minecraft:bat",
					Weight: int32(10),
				},
			},
			Axolotls: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			Creature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:sheep",
					Weight: int32(12),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:pig",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:chicken",
					Weight: int32(10),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:cow",
					Weight: int32(8),
				},
				{
					Maxcount: int32(2),
					Mincount: int32(1),
					Type: "minecraft:armadillo",
					Weight: int32(6),
				},
				{
					Maxcount: int32(8),
					Mincount: int32(4),
					Type: "minecraft:wolf",
					Weight: int32(2),
				},
			},
			Misc: []any{},
			Monster: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:spider",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:zombie",
					Weight: int32(95),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:zombie_villager",
					Weight: int32(5),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:skeleton",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:creeper",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(4),
					Type: "minecraft:slime",
					Weight: int32(100),
				},
				{
					Maxcount: int32(4),
					Mincount: int32(1),
					Type: "minecraft:enderman",
					Weight: int32(10),
				},
				{
					Maxcount: int32(1),
					Mincount: int32(1),
					Type: "minecraft:witch",
					Weight: int32(5),
				},
			},
			UndergroundWaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{
				{
					Maxcount: int32(6),
					Mincount: int32(4),
					Type: "minecraft:glow_squid",
					Weight: int32(10),
				},
			},
			WaterAmbient: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
			WaterCreature: []struct {
				Maxcount int32 `nbt:"maxCount"`
				Mincount int32 `nbt:"minCount"`
				Type string `nbt:"type"`
				Weight int32 `nbt:"weight"`
			}{},
		},
		Temperature: float32(2.0),
	},
}
