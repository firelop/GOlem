package registries

var TestEnvironment = map[string]any{
	"default": struct {
		Definitions []any `nbt:"definitions"`
		Type string `nbt:"type"`
	}{
		Definitions: []any{},
		Type: "minecraft:all_of",
	},
}
