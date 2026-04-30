package registries

var TestInstance = map[string]any{
	"always_pass": struct {
		Environment string `nbt:"environment"`
		Function string `nbt:"function"`
		MaxTicks int32 `nbt:"max_ticks"`
		Required bool `nbt:"required"`
		SetupTicks int32 `nbt:"setup_ticks"`
		Structure string `nbt:"structure"`
		Type string `nbt:"type"`
	}{
		Environment: "minecraft:default",
		Function: "minecraft:always_pass",
		MaxTicks: int32(1),
		Required: false,
		SetupTicks: int32(1),
		Structure: "minecraft:empty",
		Type: "minecraft:function",
	},
}
