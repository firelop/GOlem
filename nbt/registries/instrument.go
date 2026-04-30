package registries

var Instrument = map[string]any{
	"admire_goat_horn": struct {
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Range float32 `nbt:"range"`
		SoundEvent string `nbt:"sound_event"`
		UseDuration float32 `nbt:"use_duration"`
	}{
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "instrument.minecraft.admire_goat_horn",
		},
		Range: float32(256.0),
		SoundEvent: "minecraft:item.goat_horn.sound.4",
		UseDuration: float32(7.0),
	},
	"call_goat_horn": struct {
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Range float32 `nbt:"range"`
		SoundEvent string `nbt:"sound_event"`
		UseDuration float32 `nbt:"use_duration"`
	}{
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "instrument.minecraft.call_goat_horn",
		},
		Range: float32(256.0),
		SoundEvent: "minecraft:item.goat_horn.sound.5",
		UseDuration: float32(7.0),
	},
	"dream_goat_horn": struct {
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Range float32 `nbt:"range"`
		SoundEvent string `nbt:"sound_event"`
		UseDuration float32 `nbt:"use_duration"`
	}{
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "instrument.minecraft.dream_goat_horn",
		},
		Range: float32(256.0),
		SoundEvent: "minecraft:item.goat_horn.sound.7",
		UseDuration: float32(7.0),
	},
	"feel_goat_horn": struct {
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Range float32 `nbt:"range"`
		SoundEvent string `nbt:"sound_event"`
		UseDuration float32 `nbt:"use_duration"`
	}{
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "instrument.minecraft.feel_goat_horn",
		},
		Range: float32(256.0),
		SoundEvent: "minecraft:item.goat_horn.sound.3",
		UseDuration: float32(7.0),
	},
	"ponder_goat_horn": struct {
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Range float32 `nbt:"range"`
		SoundEvent string `nbt:"sound_event"`
		UseDuration float32 `nbt:"use_duration"`
	}{
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "instrument.minecraft.ponder_goat_horn",
		},
		Range: float32(256.0),
		SoundEvent: "minecraft:item.goat_horn.sound.0",
		UseDuration: float32(7.0),
	},
	"seek_goat_horn": struct {
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Range float32 `nbt:"range"`
		SoundEvent string `nbt:"sound_event"`
		UseDuration float32 `nbt:"use_duration"`
	}{
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "instrument.minecraft.seek_goat_horn",
		},
		Range: float32(256.0),
		SoundEvent: "minecraft:item.goat_horn.sound.2",
		UseDuration: float32(7.0),
	},
	"sing_goat_horn": struct {
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Range float32 `nbt:"range"`
		SoundEvent string `nbt:"sound_event"`
		UseDuration float32 `nbt:"use_duration"`
	}{
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "instrument.minecraft.sing_goat_horn",
		},
		Range: float32(256.0),
		SoundEvent: "minecraft:item.goat_horn.sound.1",
		UseDuration: float32(7.0),
	},
	"yearn_goat_horn": struct {
		Description struct {
			Translate string `nbt:"translate"`
		} `nbt:"description"`
		Range float32 `nbt:"range"`
		SoundEvent string `nbt:"sound_event"`
		UseDuration float32 `nbt:"use_duration"`
	}{
		Description: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "instrument.minecraft.yearn_goat_horn",
		},
		Range: float32(256.0),
		SoundEvent: "minecraft:item.goat_horn.sound.6",
		UseDuration: float32(7.0),
	},
}
