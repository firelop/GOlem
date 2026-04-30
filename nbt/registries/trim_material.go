package registries

var TrimMaterial = map[string]any{
	"amethyst": struct {
		AssetName string `nbt:"asset_name"`
		Description struct {
			Color string `nbt:"color"`
			Translate string `nbt:"translate"`
		} `nbt:"description"`
	}{
		AssetName: "amethyst",
		Description: struct {
			Color string `nbt:"color"`
			Translate string `nbt:"translate"`
		}{
			Color: "#9A5CC6",
			Translate: "trim_material.minecraft.amethyst",
		},
	},
	"copper": struct {
		AssetName string `nbt:"asset_name"`
		Description struct {
			Color string `nbt:"color"`
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		OverrideArmorAssets any `nbt:"override_armor_assets"`
	}{
		AssetName: "copper",
		Description: struct {
			Color string `nbt:"color"`
			Translate string `nbt:"translate"`
		}{
			Color: "#B4684D",
			Translate: "trim_material.minecraft.copper",
		},
		OverrideArmorAssets: map[string]any{
			"minecraft:copper": "copper_darker",
		},
	},
	"diamond": struct {
		AssetName string `nbt:"asset_name"`
		Description struct {
			Color string `nbt:"color"`
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		OverrideArmorAssets any `nbt:"override_armor_assets"`
	}{
		AssetName: "diamond",
		Description: struct {
			Color string `nbt:"color"`
			Translate string `nbt:"translate"`
		}{
			Color: "#6EECD2",
			Translate: "trim_material.minecraft.diamond",
		},
		OverrideArmorAssets: map[string]any{
			"minecraft:diamond": "diamond_darker",
		},
	},
	"emerald": struct {
		AssetName string `nbt:"asset_name"`
		Description struct {
			Color string `nbt:"color"`
			Translate string `nbt:"translate"`
		} `nbt:"description"`
	}{
		AssetName: "emerald",
		Description: struct {
			Color string `nbt:"color"`
			Translate string `nbt:"translate"`
		}{
			Color: "#11A036",
			Translate: "trim_material.minecraft.emerald",
		},
	},
	"gold": struct {
		AssetName string `nbt:"asset_name"`
		Description struct {
			Color string `nbt:"color"`
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		OverrideArmorAssets any `nbt:"override_armor_assets"`
	}{
		AssetName: "gold",
		Description: struct {
			Color string `nbt:"color"`
			Translate string `nbt:"translate"`
		}{
			Color: "#DEB12D",
			Translate: "trim_material.minecraft.gold",
		},
		OverrideArmorAssets: map[string]any{
			"minecraft:gold": "gold_darker",
		},
	},
	"iron": struct {
		AssetName string `nbt:"asset_name"`
		Description struct {
			Color string `nbt:"color"`
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		OverrideArmorAssets any `nbt:"override_armor_assets"`
	}{
		AssetName: "iron",
		Description: struct {
			Color string `nbt:"color"`
			Translate string `nbt:"translate"`
		}{
			Color: "#ECECEC",
			Translate: "trim_material.minecraft.iron",
		},
		OverrideArmorAssets: map[string]any{
			"minecraft:iron": "iron_darker",
		},
	},
	"lapis": struct {
		AssetName string `nbt:"asset_name"`
		Description struct {
			Color string `nbt:"color"`
			Translate string `nbt:"translate"`
		} `nbt:"description"`
	}{
		AssetName: "lapis",
		Description: struct {
			Color string `nbt:"color"`
			Translate string `nbt:"translate"`
		}{
			Color: "#416E97",
			Translate: "trim_material.minecraft.lapis",
		},
	},
	"netherite": struct {
		AssetName string `nbt:"asset_name"`
		Description struct {
			Color string `nbt:"color"`
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		OverrideArmorAssets any `nbt:"override_armor_assets"`
	}{
		AssetName: "netherite",
		Description: struct {
			Color string `nbt:"color"`
			Translate string `nbt:"translate"`
		}{
			Color: "#625859",
			Translate: "trim_material.minecraft.netherite",
		},
		OverrideArmorAssets: map[string]any{
			"minecraft:netherite": "netherite_darker",
		},
	},
	"quartz": struct {
		AssetName string `nbt:"asset_name"`
		Description struct {
			Color string `nbt:"color"`
			Translate string `nbt:"translate"`
		} `nbt:"description"`
	}{
		AssetName: "quartz",
		Description: struct {
			Color string `nbt:"color"`
			Translate string `nbt:"translate"`
		}{
			Color: "#E3D4C4",
			Translate: "trim_material.minecraft.quartz",
		},
	},
	"redstone": struct {
		AssetName string `nbt:"asset_name"`
		Description struct {
			Color string `nbt:"color"`
			Translate string `nbt:"translate"`
		} `nbt:"description"`
	}{
		AssetName: "redstone",
		Description: struct {
			Color string `nbt:"color"`
			Translate string `nbt:"translate"`
		}{
			Color: "#971607",
			Translate: "trim_material.minecraft.redstone",
		},
	},
	"resin": struct {
		AssetName string `nbt:"asset_name"`
		Description struct {
			Color string `nbt:"color"`
			Translate string `nbt:"translate"`
		} `nbt:"description"`
	}{
		AssetName: "resin",
		Description: struct {
			Color string `nbt:"color"`
			Translate string `nbt:"translate"`
		}{
			Color: "#FC7812",
			Translate: "trim_material.minecraft.resin",
		},
	},
}
