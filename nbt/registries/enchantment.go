package registries

var Enchantment = map[string]any{
	"aqua_affinity": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(4),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.aqua_affinity",
		},
		Effects: map[string]any{
			"minecraft:attributes": []any{
				map[string]any{
					"amount": map[string]any{
						"base": float32(4.0),
						"per_level_above_first": float32(4.0),
						"type": "minecraft:linear",
					},
					"attribute": "minecraft:submerged_mining_speed",
					"id": "minecraft:enchantment.aqua_affinity",
					"operation": "add_multiplied_total",
				},
			},
		},
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(41),
			PerLevelAboveFirst: int32(0),
		},
		MaxLevel: int32(1),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(1),
			PerLevelAboveFirst: int32(0),
		},
		Slots: []string{
			"head",
		},
		SupportedItems: "#minecraft:enchantable/head_armor",
		Weight: int32(2),
	},
	"bane_of_arthropods": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		ExclusiveSet string `nbt:"exclusive_set"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		PrimaryItems string `nbt:"primary_items"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(2),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.bane_of_arthropods",
		},
		Effects: map[string]any{
			"minecraft:damage": []any{
				map[string]any{
					"effect": map[string]any{
						"type": "minecraft:add",
						"value": map[string]any{
							"base": float32(2.5),
							"per_level_above_first": float32(2.5),
							"type": "minecraft:linear",
						},
					},
					"requirements": map[string]any{
						"condition": "minecraft:entity_properties",
						"entity": "this",
						"predicate": map[string]any{
							"type": "#minecraft:sensitive_to_bane_of_arthropods",
						},
					},
				},
			},
			"minecraft:post_attack": []any{
				map[string]any{
					"affected": "victim",
					"effect": map[string]any{
						"max_amplifier": float32(3.0),
						"max_duration": map[string]any{
							"base": float32(1.5),
							"per_level_above_first": float32(0.5),
							"type": "minecraft:linear",
						},
						"min_amplifier": float32(3.0),
						"min_duration": float32(1.5),
						"to_apply": "minecraft:slowness",
						"type": "minecraft:apply_mob_effect",
					},
					"enchanted": "attacker",
					"requirements": map[string]any{
						"condition": "minecraft:all_of",
						"terms": []any{
							map[string]any{
								"condition": "minecraft:entity_properties",
								"entity": "this",
								"predicate": map[string]any{
									"type": "#minecraft:sensitive_to_bane_of_arthropods",
								},
							},
							map[string]any{
								"condition": "minecraft:damage_source_properties",
								"predicate": map[string]any{
									"is_direct": true,
								},
							},
						},
					},
				},
			},
		},
		ExclusiveSet: "#minecraft:exclusive_set/damage",
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(25),
			PerLevelAboveFirst: int32(8),
		},
		MaxLevel: int32(5),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(5),
			PerLevelAboveFirst: int32(8),
		},
		PrimaryItems: "#minecraft:enchantable/melee_weapon",
		Slots: []string{
			"mainhand",
		},
		SupportedItems: "#minecraft:enchantable/weapon",
		Weight: int32(5),
	},
	"binding_curse": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(8),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.binding_curse",
		},
		Effects: map[string]any{
			"minecraft:prevent_armor_change": map[string]any{
			},
		},
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(50),
			PerLevelAboveFirst: int32(0),
		},
		MaxLevel: int32(1),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(25),
			PerLevelAboveFirst: int32(0),
		},
		Slots: []string{
			"armor",
		},
		SupportedItems: "#minecraft:enchantable/equippable",
		Weight: int32(1),
	},
	"blast_protection": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		ExclusiveSet string `nbt:"exclusive_set"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(4),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.blast_protection",
		},
		Effects: map[string]any{
			"minecraft:attributes": []any{
				map[string]any{
					"amount": map[string]any{
						"base": float32(0.15),
						"per_level_above_first": float32(0.15),
						"type": "minecraft:linear",
					},
					"attribute": "minecraft:explosion_knockback_resistance",
					"id": "minecraft:enchantment.blast_protection",
					"operation": "add_value",
				},
			},
			"minecraft:damage_protection": []any{
				map[string]any{
					"effect": map[string]any{
						"type": "minecraft:add",
						"value": map[string]any{
							"base": float32(2.0),
							"per_level_above_first": float32(2.0),
							"type": "minecraft:linear",
						},
					},
					"requirements": map[string]any{
						"condition": "minecraft:damage_source_properties",
						"predicate": map[string]any{
							"tags": []any{
								map[string]any{
									"expected": true,
									"id": "minecraft:is_explosion",
								},
								map[string]any{
									"expected": false,
									"id": "minecraft:bypasses_invulnerability",
								},
							},
						},
					},
				},
			},
		},
		ExclusiveSet: "#minecraft:exclusive_set/armor",
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(13),
			PerLevelAboveFirst: int32(8),
		},
		MaxLevel: int32(4),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(5),
			PerLevelAboveFirst: int32(8),
		},
		Slots: []string{
			"armor",
		},
		SupportedItems: "#minecraft:enchantable/armor",
		Weight: int32(2),
	},
	"breach": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		ExclusiveSet string `nbt:"exclusive_set"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(4),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.breach",
		},
		Effects: map[string]any{
			"minecraft:armor_effectiveness": []any{
				map[string]any{
					"effect": map[string]any{
						"type": "minecraft:add",
						"value": map[string]any{
							"base": float32(-0.15),
							"per_level_above_first": float32(-0.15),
							"type": "minecraft:linear",
						},
					},
				},
			},
		},
		ExclusiveSet: "#minecraft:exclusive_set/damage",
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(65),
			PerLevelAboveFirst: int32(9),
		},
		MaxLevel: int32(4),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(15),
			PerLevelAboveFirst: int32(9),
		},
		Slots: []string{
			"mainhand",
		},
		SupportedItems: "#minecraft:enchantable/mace",
		Weight: int32(2),
	},
	"channeling": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(8),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.channeling",
		},
		Effects: map[string]any{
			"minecraft:hit_block": []any{
				map[string]any{
					"effect": map[string]any{
						"effects": []any{
							map[string]any{
								"entity": "minecraft:lightning_bolt",
								"type": "minecraft:summon_entity",
							},
							map[string]any{
								"pitch": float32(1.0),
								"sound": "minecraft:item.trident.thunder",
								"type": "minecraft:play_sound",
								"volume": float32(5.0),
							},
						},
						"type": "minecraft:all_of",
					},
					"requirements": map[string]any{
						"condition": "minecraft:all_of",
						"terms": []any{
							map[string]any{
								"condition": "minecraft:weather_check",
								"thundering": true,
							},
							map[string]any{
								"condition": "minecraft:entity_properties",
								"entity": "this",
								"predicate": map[string]any{
									"type": "minecraft:trident",
								},
							},
							map[string]any{
								"condition": "minecraft:location_check",
								"predicate": map[string]any{
									"block": map[string]any{
										"blocks": "#minecraft:lightning_rods",
									},
									"can_see_sky": true,
								},
							},
						},
					},
				},
			},
			"minecraft:post_attack": []any{
				map[string]any{
					"affected": "victim",
					"effect": map[string]any{
						"effects": []any{
							map[string]any{
								"entity": "minecraft:lightning_bolt",
								"type": "minecraft:summon_entity",
							},
							map[string]any{
								"pitch": float32(1.0),
								"sound": "minecraft:item.trident.thunder",
								"type": "minecraft:play_sound",
								"volume": float32(5.0),
							},
						},
						"type": "minecraft:all_of",
					},
					"enchanted": "attacker",
					"requirements": map[string]any{
						"condition": "minecraft:all_of",
						"terms": []any{
							map[string]any{
								"condition": "minecraft:weather_check",
								"thundering": true,
							},
							map[string]any{
								"condition": "minecraft:entity_properties",
								"entity": "this",
								"predicate": map[string]any{
									"location": map[string]any{
										"can_see_sky": true,
									},
								},
							},
							map[string]any{
								"condition": "minecraft:entity_properties",
								"entity": "direct_attacker",
								"predicate": map[string]any{
									"type": "minecraft:trident",
								},
							},
						},
					},
				},
			},
		},
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(50),
			PerLevelAboveFirst: int32(0),
		},
		MaxLevel: int32(1),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(25),
			PerLevelAboveFirst: int32(0),
		},
		Slots: []string{
			"mainhand",
		},
		SupportedItems: "#minecraft:enchantable/trident",
		Weight: int32(1),
	},
	"density": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		ExclusiveSet string `nbt:"exclusive_set"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(2),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.density",
		},
		Effects: map[string]any{
			"minecraft:smash_damage_per_fallen_block": []any{
				map[string]any{
					"effect": map[string]any{
						"type": "minecraft:add",
						"value": map[string]any{
							"base": float32(0.5),
							"per_level_above_first": float32(0.5),
							"type": "minecraft:linear",
						},
					},
				},
			},
		},
		ExclusiveSet: "#minecraft:exclusive_set/damage",
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(25),
			PerLevelAboveFirst: int32(8),
		},
		MaxLevel: int32(5),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(5),
			PerLevelAboveFirst: int32(8),
		},
		Slots: []string{
			"mainhand",
		},
		SupportedItems: "#minecraft:enchantable/mace",
		Weight: int32(5),
	},
	"depth_strider": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		ExclusiveSet string `nbt:"exclusive_set"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(4),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.depth_strider",
		},
		Effects: map[string]any{
			"minecraft:attributes": []any{
				map[string]any{
					"amount": map[string]any{
						"base": float32(0.33333334),
						"per_level_above_first": float32(0.33333334),
						"type": "minecraft:linear",
					},
					"attribute": "minecraft:water_movement_efficiency",
					"id": "minecraft:enchantment.depth_strider",
					"operation": "add_value",
				},
			},
		},
		ExclusiveSet: "#minecraft:exclusive_set/boots",
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(25),
			PerLevelAboveFirst: int32(10),
		},
		MaxLevel: int32(3),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(10),
			PerLevelAboveFirst: int32(10),
		},
		Slots: []string{
			"feet",
		},
		SupportedItems: "#minecraft:enchantable/foot_armor",
		Weight: int32(2),
	},
	"efficiency": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(1),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.efficiency",
		},
		Effects: map[string]any{
			"minecraft:attributes": []any{
				map[string]any{
					"amount": map[string]any{
						"added": float32(1.0),
						"type": "minecraft:levels_squared",
					},
					"attribute": "minecraft:mining_efficiency",
					"id": "minecraft:enchantment.efficiency",
					"operation": "add_value",
				},
			},
		},
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(51),
			PerLevelAboveFirst: int32(10),
		},
		MaxLevel: int32(5),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(1),
			PerLevelAboveFirst: int32(10),
		},
		Slots: []string{
			"mainhand",
		},
		SupportedItems: "#minecraft:enchantable/mining",
		Weight: int32(10),
	},
	"feather_falling": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(2),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.feather_falling",
		},
		Effects: map[string]any{
			"minecraft:damage_protection": []any{
				map[string]any{
					"effect": map[string]any{
						"type": "minecraft:add",
						"value": map[string]any{
							"base": float32(3.0),
							"per_level_above_first": float32(3.0),
							"type": "minecraft:linear",
						},
					},
					"requirements": map[string]any{
						"condition": "minecraft:damage_source_properties",
						"predicate": map[string]any{
							"tags": []any{
								map[string]any{
									"expected": true,
									"id": "minecraft:is_fall",
								},
								map[string]any{
									"expected": false,
									"id": "minecraft:bypasses_invulnerability",
								},
							},
						},
					},
				},
			},
		},
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(11),
			PerLevelAboveFirst: int32(6),
		},
		MaxLevel: int32(4),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(5),
			PerLevelAboveFirst: int32(6),
		},
		Slots: []string{
			"armor",
		},
		SupportedItems: "#minecraft:enchantable/foot_armor",
		Weight: int32(5),
	},
	"fire_aspect": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		PrimaryItems string `nbt:"primary_items"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(4),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.fire_aspect",
		},
		Effects: map[string]any{
			"minecraft:post_attack": []any{
				map[string]any{
					"affected": "victim",
					"effect": map[string]any{
						"duration": map[string]any{
							"base": float32(4.0),
							"per_level_above_first": float32(4.0),
							"type": "minecraft:linear",
						},
						"type": "minecraft:ignite",
					},
					"enchanted": "attacker",
					"requirements": map[string]any{
						"condition": "minecraft:damage_source_properties",
						"predicate": map[string]any{
							"is_direct": true,
						},
					},
				},
			},
		},
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(60),
			PerLevelAboveFirst: int32(20),
		},
		MaxLevel: int32(2),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(10),
			PerLevelAboveFirst: int32(20),
		},
		PrimaryItems: "#minecraft:enchantable/melee_weapon",
		Slots: []string{
			"mainhand",
		},
		SupportedItems: "#minecraft:enchantable/fire_aspect",
		Weight: int32(2),
	},
	"fire_protection": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		ExclusiveSet string `nbt:"exclusive_set"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(2),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.fire_protection",
		},
		Effects: map[string]any{
			"minecraft:attributes": []any{
				map[string]any{
					"amount": map[string]any{
						"base": float32(-0.15),
						"per_level_above_first": float32(-0.15),
						"type": "minecraft:linear",
					},
					"attribute": "minecraft:burning_time",
					"id": "minecraft:enchantment.fire_protection",
					"operation": "add_multiplied_base",
				},
			},
			"minecraft:damage_protection": []any{
				map[string]any{
					"effect": map[string]any{
						"type": "minecraft:add",
						"value": map[string]any{
							"base": float32(2.0),
							"per_level_above_first": float32(2.0),
							"type": "minecraft:linear",
						},
					},
					"requirements": map[string]any{
						"condition": "minecraft:all_of",
						"terms": []any{
							map[string]any{
								"condition": "minecraft:damage_source_properties",
								"predicate": map[string]any{
									"tags": []any{
										map[string]any{
											"expected": true,
											"id": "minecraft:is_fire",
										},
										map[string]any{
											"expected": false,
											"id": "minecraft:bypasses_invulnerability",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		ExclusiveSet: "#minecraft:exclusive_set/armor",
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(18),
			PerLevelAboveFirst: int32(8),
		},
		MaxLevel: int32(4),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(10),
			PerLevelAboveFirst: int32(8),
		},
		Slots: []string{
			"armor",
		},
		SupportedItems: "#minecraft:enchantable/armor",
		Weight: int32(5),
	},
	"flame": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(4),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.flame",
		},
		Effects: map[string]any{
			"minecraft:projectile_spawned": []any{
				map[string]any{
					"effect": map[string]any{
						"duration": float32(100.0),
						"type": "minecraft:ignite",
					},
				},
			},
		},
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(50),
			PerLevelAboveFirst: int32(0),
		},
		MaxLevel: int32(1),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(20),
			PerLevelAboveFirst: int32(0),
		},
		Slots: []string{
			"mainhand",
		},
		SupportedItems: "#minecraft:enchantable/bow",
		Weight: int32(2),
	},
	"fortune": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		ExclusiveSet string `nbt:"exclusive_set"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(4),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.fortune",
		},
		ExclusiveSet: "#minecraft:exclusive_set/mining",
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(65),
			PerLevelAboveFirst: int32(9),
		},
		MaxLevel: int32(3),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(15),
			PerLevelAboveFirst: int32(9),
		},
		Slots: []string{
			"mainhand",
		},
		SupportedItems: "#minecraft:enchantable/mining_loot",
		Weight: int32(2),
	},
	"frost_walker": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		ExclusiveSet string `nbt:"exclusive_set"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(4),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.frost_walker",
		},
		Effects: map[string]any{
			"minecraft:damage_immunity": []any{
				map[string]any{
					"effect": map[string]any{
					},
					"requirements": map[string]any{
						"condition": "minecraft:damage_source_properties",
						"predicate": map[string]any{
							"tags": []any{
								map[string]any{
									"expected": true,
									"id": "minecraft:burn_from_stepping",
								},
								map[string]any{
									"expected": false,
									"id": "minecraft:bypasses_invulnerability",
								},
							},
						},
					},
				},
			},
			"minecraft:location_changed": []any{
				map[string]any{
					"effect": map[string]any{
						"block_state": map[string]any{
							"state": map[string]any{
								"Name": "minecraft:frosted_ice",
								"Properties": map[string]any{
									"age": "0",
								},
							},
							"type": "minecraft:simple_state_provider",
						},
						"height": float32(1.0),
						"offset": []any{
							int32(0),
							int32(-1),
							int32(0),
						},
						"predicate": map[string]any{
							"predicates": []any{
								map[string]any{
									"offset": []any{
										int32(0),
										int32(1),
										int32(0),
									},
									"tag": "minecraft:air",
									"type": "minecraft:matching_block_tag",
								},
								map[string]any{
									"blocks": "minecraft:water",
									"type": "minecraft:matching_blocks",
								},
								map[string]any{
									"fluids": "minecraft:water",
									"type": "minecraft:matching_fluids",
								},
								map[string]any{
									"type": "minecraft:unobstructed",
								},
							},
							"type": "minecraft:all_of",
						},
						"radius": map[string]any{
							"max": float32(16.0),
							"min": float32(0.0),
							"type": "minecraft:clamped",
							"value": map[string]any{
								"base": float32(3.0),
								"per_level_above_first": float32(1.0),
								"type": "minecraft:linear",
							},
						},
						"trigger_game_event": "minecraft:block_place",
						"type": "minecraft:replace_disk",
					},
					"requirements": map[string]any{
						"condition": "minecraft:all_of",
						"terms": []any{
							map[string]any{
								"condition": "minecraft:entity_properties",
								"entity": "this",
								"predicate": map[string]any{
									"flags": map[string]any{
										"is_on_ground": true,
									},
								},
							},
							map[string]any{
								"condition": "minecraft:inverted",
								"term": map[string]any{
									"condition": "minecraft:entity_properties",
									"entity": "this",
									"predicate": map[string]any{
										"vehicle": map[string]any{
										},
									},
								},
							},
						},
					},
				},
			},
		},
		ExclusiveSet: "#minecraft:exclusive_set/boots",
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(25),
			PerLevelAboveFirst: int32(10),
		},
		MaxLevel: int32(2),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(10),
			PerLevelAboveFirst: int32(10),
		},
		Slots: []string{
			"feet",
		},
		SupportedItems: "#minecraft:enchantable/foot_armor",
		Weight: int32(2),
	},
	"impaling": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		ExclusiveSet string `nbt:"exclusive_set"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(4),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.impaling",
		},
		Effects: map[string]any{
			"minecraft:damage": []any{
				map[string]any{
					"effect": map[string]any{
						"type": "minecraft:add",
						"value": map[string]any{
							"base": float32(2.5),
							"per_level_above_first": float32(2.5),
							"type": "minecraft:linear",
						},
					},
					"requirements": map[string]any{
						"condition": "minecraft:entity_properties",
						"entity": "this",
						"predicate": map[string]any{
							"type": "#minecraft:sensitive_to_impaling",
						},
					},
				},
			},
		},
		ExclusiveSet: "#minecraft:exclusive_set/damage",
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(21),
			PerLevelAboveFirst: int32(8),
		},
		MaxLevel: int32(5),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(1),
			PerLevelAboveFirst: int32(8),
		},
		Slots: []string{
			"mainhand",
		},
		SupportedItems: "#minecraft:enchantable/trident",
		Weight: int32(2),
	},
	"infinity": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		ExclusiveSet string `nbt:"exclusive_set"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(8),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.infinity",
		},
		Effects: map[string]any{
			"minecraft:ammo_use": []any{
				map[string]any{
					"effect": map[string]any{
						"type": "minecraft:set",
						"value": float32(0.0),
					},
					"requirements": map[string]any{
						"condition": "minecraft:match_tool",
						"predicate": map[string]any{
							"items": "minecraft:arrow",
						},
					},
				},
			},
		},
		ExclusiveSet: "#minecraft:exclusive_set/bow",
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(50),
			PerLevelAboveFirst: int32(0),
		},
		MaxLevel: int32(1),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(20),
			PerLevelAboveFirst: int32(0),
		},
		Slots: []string{
			"mainhand",
		},
		SupportedItems: "#minecraft:enchantable/bow",
		Weight: int32(1),
	},
	"knockback": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(2),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.knockback",
		},
		Effects: map[string]any{
			"minecraft:knockback": []any{
				map[string]any{
					"effect": map[string]any{
						"type": "minecraft:add",
						"value": map[string]any{
							"base": float32(1.0),
							"per_level_above_first": float32(1.0),
							"type": "minecraft:linear",
						},
					},
				},
			},
		},
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(55),
			PerLevelAboveFirst: int32(20),
		},
		MaxLevel: int32(2),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(5),
			PerLevelAboveFirst: int32(20),
		},
		Slots: []string{
			"mainhand",
		},
		SupportedItems: "#minecraft:enchantable/melee_weapon",
		Weight: int32(5),
	},
	"looting": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(4),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.looting",
		},
		Effects: map[string]any{
			"minecraft:equipment_drops": []any{
				map[string]any{
					"effect": map[string]any{
						"type": "minecraft:add",
						"value": map[string]any{
							"base": float32(0.01),
							"per_level_above_first": float32(0.01),
							"type": "minecraft:linear",
						},
					},
					"enchanted": "attacker",
					"requirements": map[string]any{
						"condition": "minecraft:entity_properties",
						"entity": "attacker",
						"predicate": map[string]any{
							"type": "minecraft:player",
						},
					},
				},
			},
		},
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(65),
			PerLevelAboveFirst: int32(9),
		},
		MaxLevel: int32(3),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(15),
			PerLevelAboveFirst: int32(9),
		},
		Slots: []string{
			"mainhand",
		},
		SupportedItems: "#minecraft:enchantable/melee_weapon",
		Weight: int32(2),
	},
	"loyalty": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(2),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.loyalty",
		},
		Effects: map[string]any{
			"minecraft:trident_return_acceleration": []any{
				map[string]any{
					"effect": map[string]any{
						"type": "minecraft:add",
						"value": map[string]any{
							"base": float32(1.0),
							"per_level_above_first": float32(1.0),
							"type": "minecraft:linear",
						},
					},
				},
			},
		},
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(50),
			PerLevelAboveFirst: int32(0),
		},
		MaxLevel: int32(3),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(12),
			PerLevelAboveFirst: int32(7),
		},
		Slots: []string{
			"mainhand",
		},
		SupportedItems: "#minecraft:enchantable/trident",
		Weight: int32(5),
	},
	"luck_of_the_sea": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(4),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.luck_of_the_sea",
		},
		Effects: map[string]any{
			"minecraft:fishing_luck_bonus": []any{
				map[string]any{
					"effect": map[string]any{
						"type": "minecraft:add",
						"value": map[string]any{
							"base": float32(1.0),
							"per_level_above_first": float32(1.0),
							"type": "minecraft:linear",
						},
					},
				},
			},
		},
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(65),
			PerLevelAboveFirst: int32(9),
		},
		MaxLevel: int32(3),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(15),
			PerLevelAboveFirst: int32(9),
		},
		Slots: []string{
			"mainhand",
		},
		SupportedItems: "#minecraft:enchantable/fishing",
		Weight: int32(2),
	},
	"lunge": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(2),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.lunge",
		},
		Effects: map[string]any{
			"minecraft:post_piercing_attack": []any{
				map[string]any{
					"effect": map[string]any{
						"effects": []any{
							map[string]any{
								"amount": float32(1.0),
								"type": "minecraft:change_item_damage",
							},
							map[string]any{
								"amount": map[string]any{
									"base": float32(4.0),
									"per_level_above_first": float32(4.0),
									"type": "minecraft:linear",
								},
								"type": "minecraft:apply_exhaustion",
							},
							map[string]any{
								"coordinate_scale": []any{
									float32(1.0),
									float32(0.0),
									float32(1.0),
								},
								"direction": []any{
									float32(0.0),
									float32(0.0),
									float32(1.0),
								},
								"magnitude": map[string]any{
									"base": float32(0.458),
									"per_level_above_first": float32(0.458),
									"type": "minecraft:linear",
								},
								"type": "minecraft:apply_impulse",
							},
							map[string]any{
								"pitch": float32(1.0),
								"sound": []any{
									"minecraft:item.spear.lunge_1",
									"minecraft:item.spear.lunge_2",
									"minecraft:item.spear.lunge_3",
								},
								"type": "minecraft:play_sound",
								"volume": float32(1.0),
							},
						},
						"type": "minecraft:all_of",
					},
					"requirements": map[string]any{
						"condition": "minecraft:all_of",
						"terms": []any{
							map[string]any{
								"condition": "minecraft:inverted",
								"term": map[string]any{
									"condition": "minecraft:entity_properties",
									"entity": "this",
									"predicate": map[string]any{
										"vehicle": map[string]any{
										},
									},
								},
							},
							map[string]any{
								"condition": "minecraft:entity_properties",
								"entity": "this",
								"predicate": map[string]any{
									"flags": map[string]any{
										"is_fall_flying": false,
									},
								},
							},
							map[string]any{
								"condition": "minecraft:entity_properties",
								"entity": "this",
								"predicate": map[string]any{
									"flags": map[string]any{
										"is_in_water": false,
									},
								},
							},
						},
					},
				},
			},
		},
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(25),
			PerLevelAboveFirst: int32(8),
		},
		MaxLevel: int32(3),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(5),
			PerLevelAboveFirst: int32(8),
		},
		Slots: []string{
			"hand",
		},
		SupportedItems: "#minecraft:enchantable/lunge",
		Weight: int32(5),
	},
	"lure": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(4),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.lure",
		},
		Effects: map[string]any{
			"minecraft:fishing_time_reduction": []any{
				map[string]any{
					"effect": map[string]any{
						"type": "minecraft:add",
						"value": map[string]any{
							"base": float32(5.0),
							"per_level_above_first": float32(5.0),
							"type": "minecraft:linear",
						},
					},
				},
			},
		},
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(65),
			PerLevelAboveFirst: int32(9),
		},
		MaxLevel: int32(3),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(15),
			PerLevelAboveFirst: int32(9),
		},
		Slots: []string{
			"mainhand",
		},
		SupportedItems: "#minecraft:enchantable/fishing",
		Weight: int32(2),
	},
	"mending": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(4),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.mending",
		},
		Effects: map[string]any{
			"minecraft:repair_with_xp": []any{
				map[string]any{
					"effect": map[string]any{
						"factor": float32(2.0),
						"type": "minecraft:multiply",
					},
				},
			},
		},
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(75),
			PerLevelAboveFirst: int32(25),
		},
		MaxLevel: int32(1),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(25),
			PerLevelAboveFirst: int32(25),
		},
		Slots: []string{
			"any",
		},
		SupportedItems: "#minecraft:enchantable/durability",
		Weight: int32(2),
	},
	"multishot": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		ExclusiveSet string `nbt:"exclusive_set"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(4),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.multishot",
		},
		Effects: map[string]any{
			"minecraft:projectile_count": []any{
				map[string]any{
					"effect": map[string]any{
						"type": "minecraft:add",
						"value": map[string]any{
							"base": float32(2.0),
							"per_level_above_first": float32(2.0),
							"type": "minecraft:linear",
						},
					},
				},
			},
			"minecraft:projectile_spread": []any{
				map[string]any{
					"effect": map[string]any{
						"type": "minecraft:add",
						"value": map[string]any{
							"base": float32(10.0),
							"per_level_above_first": float32(10.0),
							"type": "minecraft:linear",
						},
					},
				},
			},
		},
		ExclusiveSet: "#minecraft:exclusive_set/crossbow",
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(50),
			PerLevelAboveFirst: int32(0),
		},
		MaxLevel: int32(1),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(20),
			PerLevelAboveFirst: int32(0),
		},
		Slots: []string{
			"mainhand",
		},
		SupportedItems: "#minecraft:enchantable/crossbow",
		Weight: int32(2),
	},
	"piercing": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		ExclusiveSet string `nbt:"exclusive_set"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(1),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.piercing",
		},
		Effects: map[string]any{
			"minecraft:projectile_piercing": []any{
				map[string]any{
					"effect": map[string]any{
						"type": "minecraft:add",
						"value": map[string]any{
							"base": float32(1.0),
							"per_level_above_first": float32(1.0),
							"type": "minecraft:linear",
						},
					},
				},
			},
		},
		ExclusiveSet: "#minecraft:exclusive_set/crossbow",
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(50),
			PerLevelAboveFirst: int32(0),
		},
		MaxLevel: int32(4),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(1),
			PerLevelAboveFirst: int32(10),
		},
		Slots: []string{
			"mainhand",
		},
		SupportedItems: "#minecraft:enchantable/crossbow",
		Weight: int32(10),
	},
	"power": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(1),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.power",
		},
		Effects: map[string]any{
			"minecraft:damage": []any{
				map[string]any{
					"effect": map[string]any{
						"type": "minecraft:add",
						"value": map[string]any{
							"base": float32(1.0),
							"per_level_above_first": float32(0.5),
							"type": "minecraft:linear",
						},
					},
					"requirements": map[string]any{
						"condition": "minecraft:entity_properties",
						"entity": "direct_attacker",
						"predicate": map[string]any{
							"type": "#minecraft:arrows",
						},
					},
				},
			},
		},
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(16),
			PerLevelAboveFirst: int32(10),
		},
		MaxLevel: int32(5),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(1),
			PerLevelAboveFirst: int32(10),
		},
		Slots: []string{
			"mainhand",
		},
		SupportedItems: "#minecraft:enchantable/bow",
		Weight: int32(10),
	},
	"projectile_protection": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		ExclusiveSet string `nbt:"exclusive_set"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(2),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.projectile_protection",
		},
		Effects: map[string]any{
			"minecraft:damage_protection": []any{
				map[string]any{
					"effect": map[string]any{
						"type": "minecraft:add",
						"value": map[string]any{
							"base": float32(2.0),
							"per_level_above_first": float32(2.0),
							"type": "minecraft:linear",
						},
					},
					"requirements": map[string]any{
						"condition": "minecraft:damage_source_properties",
						"predicate": map[string]any{
							"tags": []any{
								map[string]any{
									"expected": true,
									"id": "minecraft:is_projectile",
								},
								map[string]any{
									"expected": false,
									"id": "minecraft:bypasses_invulnerability",
								},
							},
						},
					},
				},
			},
		},
		ExclusiveSet: "#minecraft:exclusive_set/armor",
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(9),
			PerLevelAboveFirst: int32(6),
		},
		MaxLevel: int32(4),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(3),
			PerLevelAboveFirst: int32(6),
		},
		Slots: []string{
			"armor",
		},
		SupportedItems: "#minecraft:enchantable/armor",
		Weight: int32(5),
	},
	"protection": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		ExclusiveSet string `nbt:"exclusive_set"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(1),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.protection",
		},
		Effects: map[string]any{
			"minecraft:damage_protection": []any{
				map[string]any{
					"effect": map[string]any{
						"type": "minecraft:add",
						"value": map[string]any{
							"base": float32(1.0),
							"per_level_above_first": float32(1.0),
							"type": "minecraft:linear",
						},
					},
					"requirements": map[string]any{
						"condition": "minecraft:damage_source_properties",
						"predicate": map[string]any{
							"tags": []any{
								map[string]any{
									"expected": false,
									"id": "minecraft:bypasses_invulnerability",
								},
							},
						},
					},
				},
			},
		},
		ExclusiveSet: "#minecraft:exclusive_set/armor",
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(12),
			PerLevelAboveFirst: int32(11),
		},
		MaxLevel: int32(4),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(1),
			PerLevelAboveFirst: int32(11),
		},
		Slots: []string{
			"armor",
		},
		SupportedItems: "#minecraft:enchantable/armor",
		Weight: int32(10),
	},
	"punch": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(4),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.punch",
		},
		Effects: map[string]any{
			"minecraft:knockback": []any{
				map[string]any{
					"effect": map[string]any{
						"type": "minecraft:add",
						"value": map[string]any{
							"base": float32(1.0),
							"per_level_above_first": float32(1.0),
							"type": "minecraft:linear",
						},
					},
					"requirements": map[string]any{
						"condition": "minecraft:entity_properties",
						"entity": "direct_attacker",
						"predicate": map[string]any{
							"type": "#minecraft:arrows",
						},
					},
				},
			},
		},
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(37),
			PerLevelAboveFirst: int32(20),
		},
		MaxLevel: int32(2),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(12),
			PerLevelAboveFirst: int32(20),
		},
		Slots: []string{
			"mainhand",
		},
		SupportedItems: "#minecraft:enchantable/bow",
		Weight: int32(2),
	},
	"quick_charge": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(2),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.quick_charge",
		},
		Effects: map[string]any{
			"minecraft:crossbow_charge_time": map[string]any{
				"type": "minecraft:add",
				"value": map[string]any{
					"base": float32(-0.25),
					"per_level_above_first": float32(-0.25),
					"type": "minecraft:linear",
				},
			},
			"minecraft:crossbow_charging_sounds": []any{
				map[string]any{
					"end": "minecraft:item.crossbow.loading_end",
					"start": "minecraft:item.crossbow.quick_charge_1",
				},
				map[string]any{
					"end": "minecraft:item.crossbow.loading_end",
					"start": "minecraft:item.crossbow.quick_charge_2",
				},
				map[string]any{
					"end": "minecraft:item.crossbow.loading_end",
					"start": "minecraft:item.crossbow.quick_charge_3",
				},
			},
		},
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(50),
			PerLevelAboveFirst: int32(0),
		},
		MaxLevel: int32(3),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(12),
			PerLevelAboveFirst: int32(20),
		},
		Slots: []string{
			"mainhand",
			"offhand",
		},
		SupportedItems: "#minecraft:enchantable/crossbow",
		Weight: int32(5),
	},
	"respiration": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(4),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.respiration",
		},
		Effects: map[string]any{
			"minecraft:attributes": []any{
				map[string]any{
					"amount": map[string]any{
						"base": float32(1.0),
						"per_level_above_first": float32(1.0),
						"type": "minecraft:linear",
					},
					"attribute": "minecraft:oxygen_bonus",
					"id": "minecraft:enchantment.respiration",
					"operation": "add_value",
				},
			},
		},
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(40),
			PerLevelAboveFirst: int32(10),
		},
		MaxLevel: int32(3),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(10),
			PerLevelAboveFirst: int32(10),
		},
		Slots: []string{
			"head",
		},
		SupportedItems: "#minecraft:enchantable/head_armor",
		Weight: int32(2),
	},
	"riptide": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		ExclusiveSet string `nbt:"exclusive_set"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(4),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.riptide",
		},
		Effects: map[string]any{
			"minecraft:trident_sound": []any{
				"minecraft:item.trident.riptide_1",
				"minecraft:item.trident.riptide_2",
				"minecraft:item.trident.riptide_3",
			},
			"minecraft:trident_spin_attack_strength": map[string]any{
				"type": "minecraft:add",
				"value": map[string]any{
					"base": float32(1.5),
					"per_level_above_first": float32(0.75),
					"type": "minecraft:linear",
				},
			},
		},
		ExclusiveSet: "#minecraft:exclusive_set/riptide",
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(50),
			PerLevelAboveFirst: int32(0),
		},
		MaxLevel: int32(3),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(17),
			PerLevelAboveFirst: int32(7),
		},
		Slots: []string{
			"hand",
		},
		SupportedItems: "#minecraft:enchantable/trident",
		Weight: int32(2),
	},
	"sharpness": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		ExclusiveSet string `nbt:"exclusive_set"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		PrimaryItems string `nbt:"primary_items"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(1),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.sharpness",
		},
		Effects: map[string]any{
			"minecraft:damage": []any{
				map[string]any{
					"effect": map[string]any{
						"type": "minecraft:add",
						"value": map[string]any{
							"base": float32(1.0),
							"per_level_above_first": float32(0.5),
							"type": "minecraft:linear",
						},
					},
				},
			},
		},
		ExclusiveSet: "#minecraft:exclusive_set/damage",
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(21),
			PerLevelAboveFirst: int32(11),
		},
		MaxLevel: int32(5),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(1),
			PerLevelAboveFirst: int32(11),
		},
		PrimaryItems: "#minecraft:enchantable/melee_weapon",
		Slots: []string{
			"mainhand",
		},
		SupportedItems: "#minecraft:enchantable/sharp_weapon",
		Weight: int32(10),
	},
	"silk_touch": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		ExclusiveSet string `nbt:"exclusive_set"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(8),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.silk_touch",
		},
		Effects: map[string]any{
			"minecraft:block_experience": []any{
				map[string]any{
					"effect": map[string]any{
						"type": "minecraft:set",
						"value": float32(0.0),
					},
				},
			},
		},
		ExclusiveSet: "#minecraft:exclusive_set/mining",
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(65),
			PerLevelAboveFirst: int32(0),
		},
		MaxLevel: int32(1),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(15),
			PerLevelAboveFirst: int32(0),
		},
		Slots: []string{
			"mainhand",
		},
		SupportedItems: "#minecraft:enchantable/mining_loot",
		Weight: int32(1),
	},
	"smite": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		ExclusiveSet string `nbt:"exclusive_set"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		PrimaryItems string `nbt:"primary_items"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(2),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.smite",
		},
		Effects: map[string]any{
			"minecraft:damage": []any{
				map[string]any{
					"effect": map[string]any{
						"type": "minecraft:add",
						"value": map[string]any{
							"base": float32(2.5),
							"per_level_above_first": float32(2.5),
							"type": "minecraft:linear",
						},
					},
					"requirements": map[string]any{
						"condition": "minecraft:entity_properties",
						"entity": "this",
						"predicate": map[string]any{
							"type": "#minecraft:sensitive_to_smite",
						},
					},
				},
			},
		},
		ExclusiveSet: "#minecraft:exclusive_set/damage",
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(25),
			PerLevelAboveFirst: int32(8),
		},
		MaxLevel: int32(5),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(5),
			PerLevelAboveFirst: int32(8),
		},
		PrimaryItems: "#minecraft:enchantable/melee_weapon",
		Slots: []string{
			"mainhand",
		},
		SupportedItems: "#minecraft:enchantable/weapon",
		Weight: int32(5),
	},
	"soul_speed": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(8),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.soul_speed",
		},
		Effects: map[string]any{
			"minecraft:location_changed": []any{
				map[string]any{
					"effect": map[string]any{
						"effects": []any{
							map[string]any{
								"amount": map[string]any{
									"base": float32(0.0405),
									"per_level_above_first": float32(0.0105),
									"type": "minecraft:linear",
								},
								"attribute": "minecraft:movement_speed",
								"id": "minecraft:enchantment.soul_speed",
								"operation": "add_value",
								"type": "minecraft:attribute",
							},
							map[string]any{
								"amount": float32(1.0),
								"attribute": "minecraft:movement_efficiency",
								"id": "minecraft:enchantment.soul_speed",
								"operation": "add_value",
								"type": "minecraft:attribute",
							},
						},
						"type": "minecraft:all_of",
					},
					"requirements": map[string]any{
						"condition": "minecraft:all_of",
						"terms": []any{
							map[string]any{
								"condition": "minecraft:inverted",
								"term": map[string]any{
									"condition": "minecraft:entity_properties",
									"entity": "this",
									"predicate": map[string]any{
										"vehicle": map[string]any{
										},
									},
								},
							},
							map[string]any{
								"condition": "minecraft:any_of",
								"terms": []any{
									map[string]any{
										"condition": "minecraft:all_of",
										"terms": []any{
											map[string]any{
												"active": true,
												"condition": "minecraft:enchantment_active_check",
											},
											map[string]any{
												"condition": "minecraft:entity_properties",
												"entity": "this",
												"predicate": map[string]any{
													"flags": map[string]any{
														"is_flying": false,
													},
												},
											},
											map[string]any{
												"condition": "minecraft:any_of",
												"terms": []any{
													map[string]any{
														"condition": "minecraft:entity_properties",
														"entity": "this",
														"predicate": map[string]any{
															"movement_affected_by": map[string]any{
																"block": map[string]any{
																	"blocks": "#minecraft:soul_speed_blocks",
																},
															},
														},
													},
													map[string]any{
														"condition": "minecraft:entity_properties",
														"entity": "this",
														"predicate": map[string]any{
															"flags": map[string]any{
																"is_on_ground": false,
															},
														},
													},
												},
											},
										},
									},
									map[string]any{
										"condition": "minecraft:all_of",
										"terms": []any{
											map[string]any{
												"active": false,
												"condition": "minecraft:enchantment_active_check",
											},
											map[string]any{
												"condition": "minecraft:entity_properties",
												"entity": "this",
												"predicate": map[string]any{
													"flags": map[string]any{
														"is_flying": false,
													},
													"movement_affected_by": map[string]any{
														"block": map[string]any{
															"blocks": "#minecraft:soul_speed_blocks",
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
				map[string]any{
					"effect": map[string]any{
						"amount": float32(1.0),
						"type": "minecraft:change_item_damage",
					},
					"requirements": map[string]any{
						"condition": "minecraft:all_of",
						"terms": []any{
							map[string]any{
								"chance": map[string]any{
									"amount": float32(0.04),
									"type": "minecraft:enchantment_level",
								},
								"condition": "minecraft:random_chance",
							},
							map[string]any{
								"condition": "minecraft:entity_properties",
								"entity": "this",
								"predicate": map[string]any{
									"flags": map[string]any{
										"is_on_ground": true,
									},
									"movement_affected_by": map[string]any{
										"block": map[string]any{
											"blocks": "#minecraft:soul_speed_blocks",
										},
									},
								},
							},
						},
					},
				},
			},
			"minecraft:tick": []any{
				map[string]any{
					"effect": map[string]any{
						"horizontal_position": map[string]any{
							"type": "in_bounding_box",
						},
						"horizontal_velocity": map[string]any{
							"movement_scale": float32(-0.2),
						},
						"particle": map[string]any{
							"type": "minecraft:soul",
						},
						"speed": float32(1.0),
						"type": "minecraft:spawn_particles",
						"vertical_position": map[string]any{
							"offset": float32(0.1),
							"type": "entity_position",
						},
						"vertical_velocity": map[string]any{
							"base": float32(0.1),
						},
					},
					"requirements": map[string]any{
						"condition": "minecraft:entity_properties",
						"entity": "this",
						"predicate": map[string]any{
							"flags": map[string]any{
								"is_flying": false,
								"is_on_ground": true,
							},
							"movement": map[string]any{
								"horizontal_speed": map[string]any{
									"min": float32(9.999999747378752e-06),
								},
							},
							"movement_affected_by": map[string]any{
								"block": map[string]any{
									"blocks": "#minecraft:soul_speed_blocks",
								},
							},
							"periodic_tick": int32(5),
						},
					},
				},
				map[string]any{
					"effect": map[string]any{
						"pitch": map[string]any{
							"max_exclusive": float32(1.0),
							"min_inclusive": float32(0.6),
							"type": "minecraft:uniform",
						},
						"sound": "minecraft:particle.soul_escape",
						"type": "minecraft:play_sound",
						"volume": float32(0.6),
					},
					"requirements": map[string]any{
						"condition": "minecraft:all_of",
						"terms": []any{
							map[string]any{
								"chance": float32(0.35),
								"condition": "minecraft:random_chance",
							},
							map[string]any{
								"condition": "minecraft:entity_properties",
								"entity": "this",
								"predicate": map[string]any{
									"flags": map[string]any{
										"is_flying": false,
										"is_on_ground": true,
									},
									"movement": map[string]any{
										"horizontal_speed": map[string]any{
											"min": float32(9.999999747378752e-06),
										},
									},
									"movement_affected_by": map[string]any{
										"block": map[string]any{
											"blocks": "#minecraft:soul_speed_blocks",
										},
									},
									"periodic_tick": int32(5),
								},
							},
						},
					},
				},
			},
		},
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(25),
			PerLevelAboveFirst: int32(10),
		},
		MaxLevel: int32(3),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(10),
			PerLevelAboveFirst: int32(10),
		},
		Slots: []string{
			"feet",
		},
		SupportedItems: "#minecraft:enchantable/foot_armor",
		Weight: int32(1),
	},
	"sweeping_edge": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(4),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.sweeping_edge",
		},
		Effects: map[string]any{
			"minecraft:attributes": []any{
				map[string]any{
					"amount": map[string]any{
						"denominator": map[string]any{
							"base": float32(2.0),
							"per_level_above_first": float32(1.0),
							"type": "minecraft:linear",
						},
						"numerator": map[string]any{
							"base": float32(1.0),
							"per_level_above_first": float32(1.0),
							"type": "minecraft:linear",
						},
						"type": "minecraft:fraction",
					},
					"attribute": "minecraft:sweeping_damage_ratio",
					"id": "minecraft:enchantment.sweeping_edge",
					"operation": "add_value",
				},
			},
		},
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(20),
			PerLevelAboveFirst: int32(9),
		},
		MaxLevel: int32(3),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(5),
			PerLevelAboveFirst: int32(9),
		},
		Slots: []string{
			"mainhand",
		},
		SupportedItems: "#minecraft:enchantable/sweeping",
		Weight: int32(2),
	},
	"swift_sneak": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(8),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.swift_sneak",
		},
		Effects: map[string]any{
			"minecraft:attributes": []any{
				map[string]any{
					"amount": map[string]any{
						"base": float32(0.15),
						"per_level_above_first": float32(0.15),
						"type": "minecraft:linear",
					},
					"attribute": "minecraft:sneaking_speed",
					"id": "minecraft:enchantment.swift_sneak",
					"operation": "add_value",
				},
			},
		},
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(75),
			PerLevelAboveFirst: int32(25),
		},
		MaxLevel: int32(3),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(25),
			PerLevelAboveFirst: int32(25),
		},
		Slots: []string{
			"legs",
		},
		SupportedItems: "#minecraft:enchantable/leg_armor",
		Weight: int32(1),
	},
	"thorns": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		PrimaryItems string `nbt:"primary_items"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(8),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.thorns",
		},
		Effects: map[string]any{
			"minecraft:post_attack": []any{
				map[string]any{
					"affected": "attacker",
					"effect": map[string]any{
						"effects": []any{
							map[string]any{
								"damage_type": "minecraft:thorns",
								"max_damage": float32(5.0),
								"min_damage": float32(1.0),
								"type": "minecraft:damage_entity",
							},
							map[string]any{
								"amount": float32(2.0),
								"type": "minecraft:change_item_damage",
							},
						},
						"type": "minecraft:all_of",
					},
					"enchanted": "victim",
					"requirements": map[string]any{
						"chance": map[string]any{
							"amount": map[string]any{
								"base": float32(0.15),
								"per_level_above_first": float32(0.15),
								"type": "minecraft:linear",
							},
							"type": "minecraft:enchantment_level",
						},
						"condition": "minecraft:random_chance",
					},
				},
			},
		},
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(60),
			PerLevelAboveFirst: int32(20),
		},
		MaxLevel: int32(3),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(10),
			PerLevelAboveFirst: int32(20),
		},
		PrimaryItems: "#minecraft:enchantable/chest_armor",
		Slots: []string{
			"any",
		},
		SupportedItems: "#minecraft:enchantable/armor",
		Weight: int32(1),
	},
	"unbreaking": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(2),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.unbreaking",
		},
		Effects: map[string]any{
			"minecraft:item_damage": []any{
				map[string]any{
					"effect": map[string]any{
						"chance": map[string]any{
							"denominator": map[string]any{
								"base": float32(10.0),
								"per_level_above_first": float32(5.0),
								"type": "minecraft:linear",
							},
							"numerator": map[string]any{
								"base": float32(2.0),
								"per_level_above_first": float32(2.0),
								"type": "minecraft:linear",
							},
							"type": "minecraft:fraction",
						},
						"type": "minecraft:remove_binomial",
					},
					"requirements": map[string]any{
						"condition": "minecraft:match_tool",
						"predicate": map[string]any{
							"items": "#minecraft:enchantable/armor",
						},
					},
				},
				map[string]any{
					"effect": map[string]any{
						"chance": map[string]any{
							"denominator": map[string]any{
								"base": float32(2.0),
								"per_level_above_first": float32(1.0),
								"type": "minecraft:linear",
							},
							"numerator": map[string]any{
								"base": float32(1.0),
								"per_level_above_first": float32(1.0),
								"type": "minecraft:linear",
							},
							"type": "minecraft:fraction",
						},
						"type": "minecraft:remove_binomial",
					},
					"requirements": map[string]any{
						"condition": "minecraft:inverted",
						"term": map[string]any{
							"condition": "minecraft:match_tool",
							"predicate": map[string]any{
								"items": "#minecraft:enchantable/armor",
							},
						},
					},
				},
			},
		},
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(55),
			PerLevelAboveFirst: int32(8),
		},
		MaxLevel: int32(3),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(5),
			PerLevelAboveFirst: int32(8),
		},
		Slots: []string{
			"any",
		},
		SupportedItems: "#minecraft:enchantable/durability",
		Weight: int32(5),
	},
	"vanishing_curse": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(8),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.vanishing_curse",
		},
		Effects: map[string]any{
			"minecraft:prevent_equipment_drop": map[string]any{
			},
		},
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(50),
			PerLevelAboveFirst: int32(0),
		},
		MaxLevel: int32(1),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(25),
			PerLevelAboveFirst: int32(0),
		},
		Slots: []string{
			"any",
		},
		SupportedItems: "#minecraft:enchantable/vanishing",
		Weight: int32(1),
	},
	"wind_burst": struct {
		AnvilCost int32 `nbt:"anvil_cost"`
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Effects any `nbt:"effects"`
		MaxCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"max_cost"`
		MaxLevel int32 `nbt:"max_level"`
		MinCost struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		} `nbt:"min_cost"`
		Slots []string `nbt:"slots"`
		SupportedItems string `nbt:"supported_items"`
		Weight int32 `nbt:"weight"`
	}{
		AnvilCost: int32(4),
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "enchantment.minecraft.wind_burst",
		},
		Effects: map[string]any{
			"minecraft:post_attack": []any{
				map[string]any{
					"affected": "attacker",
					"effect": map[string]any{
						"block_interaction": "trigger",
						"immune_blocks": "#minecraft:blocks_wind_charge_explosions",
						"knockback_multiplier": map[string]any{
							"fallback": map[string]any{
								"base": float32(1.5),
								"per_level_above_first": float32(0.35),
								"type": "minecraft:linear",
							},
							"type": "minecraft:lookup",
							"values": []any{
								float32(1.2),
								float32(1.75),
								float32(2.2),
							},
						},
						"large_particle": map[string]any{
							"type": "minecraft:gust_emitter_large",
						},
						"radius": float32(3.5),
						"small_particle": map[string]any{
							"type": "minecraft:gust_emitter_small",
						},
						"sound": "minecraft:entity.wind_charge.wind_burst",
						"type": "minecraft:explode",
					},
					"enchanted": "attacker",
					"requirements": map[string]any{
						"condition": "minecraft:entity_properties",
						"entity": "direct_attacker",
						"predicate": map[string]any{
							"flags": map[string]any{
								"is_flying": false,
							},
							"movement": map[string]any{
								"fall_distance": map[string]any{
									"min": float32(1.5),
								},
							},
						},
					},
				},
			},
		},
		MaxCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(65),
			PerLevelAboveFirst: int32(9),
		},
		MaxLevel: int32(3),
		MinCost: struct {
			Base int32 `nbt:"base"`
			PerLevelAboveFirst int32 `nbt:"per_level_above_first"`
		}{
			Base: int32(15),
			PerLevelAboveFirst: int32(9),
		},
		Slots: []string{
			"mainhand",
		},
		SupportedItems: "#minecraft:enchantable/mace",
		Weight: int32(2),
	},
}
