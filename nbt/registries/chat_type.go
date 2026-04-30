package registries

var ChatType = map[string]any{
	"chat": struct {
		Chat any `nbt:"chat"`
		Narration struct {
			Parameters []string `nbt:"parameters"`
			TranslationKey string `nbt:"translation_key"`
		} `nbt:"narration"`
	}{
		Chat: map[string]any{
			"parameters": []any{
				"sender",
				"content",
			},
			"translation_key": "chat.type.text",
		},
		Narration: struct {
			Parameters []string `nbt:"parameters"`
			TranslationKey string `nbt:"translation_key"`
		}{
			Parameters: []string{
				"sender",
				"content",
			},
			TranslationKey: "chat.type.text.narrate",
		},
	},
	"emote_command": struct {
		Chat any `nbt:"chat"`
		Narration struct {
			Parameters []string `nbt:"parameters"`
			TranslationKey string `nbt:"translation_key"`
		} `nbt:"narration"`
	}{
		Chat: map[string]any{
			"parameters": []any{
				"sender",
				"content",
			},
			"translation_key": "chat.type.emote",
		},
		Narration: struct {
			Parameters []string `nbt:"parameters"`
			TranslationKey string `nbt:"translation_key"`
		}{
			Parameters: []string{
				"sender",
				"content",
			},
			TranslationKey: "chat.type.emote",
		},
	},
	"msg_command_incoming": struct {
		Chat any `nbt:"chat"`
		Narration struct {
			Parameters []string `nbt:"parameters"`
			TranslationKey string `nbt:"translation_key"`
		} `nbt:"narration"`
	}{
		Chat: map[string]any{
			"parameters": []any{
				"sender",
				"content",
			},
			"style": map[string]any{
				"color": "gray",
				"italic": true,
			},
			"translation_key": "commands.message.display.incoming",
		},
		Narration: struct {
			Parameters []string `nbt:"parameters"`
			TranslationKey string `nbt:"translation_key"`
		}{
			Parameters: []string{
				"sender",
				"content",
			},
			TranslationKey: "chat.type.text.narrate",
		},
	},
	"msg_command_outgoing": struct {
		Chat any `nbt:"chat"`
		Narration struct {
			Parameters []string `nbt:"parameters"`
			TranslationKey string `nbt:"translation_key"`
		} `nbt:"narration"`
	}{
		Chat: map[string]any{
			"parameters": []any{
				"target",
				"content",
			},
			"style": map[string]any{
				"color": "gray",
				"italic": true,
			},
			"translation_key": "commands.message.display.outgoing",
		},
		Narration: struct {
			Parameters []string `nbt:"parameters"`
			TranslationKey string `nbt:"translation_key"`
		}{
			Parameters: []string{
				"sender",
				"content",
			},
			TranslationKey: "chat.type.text.narrate",
		},
	},
	"say_command": struct {
		Chat any `nbt:"chat"`
		Narration struct {
			Parameters []string `nbt:"parameters"`
			TranslationKey string `nbt:"translation_key"`
		} `nbt:"narration"`
	}{
		Chat: map[string]any{
			"parameters": []any{
				"sender",
				"content",
			},
			"translation_key": "chat.type.announcement",
		},
		Narration: struct {
			Parameters []string `nbt:"parameters"`
			TranslationKey string `nbt:"translation_key"`
		}{
			Parameters: []string{
				"sender",
				"content",
			},
			TranslationKey: "chat.type.text.narrate",
		},
	},
	"team_msg_command_incoming": struct {
		Chat any `nbt:"chat"`
		Narration struct {
			Parameters []string `nbt:"parameters"`
			TranslationKey string `nbt:"translation_key"`
		} `nbt:"narration"`
	}{
		Chat: map[string]any{
			"parameters": []any{
				"target",
				"sender",
				"content",
			},
			"translation_key": "chat.type.team.text",
		},
		Narration: struct {
			Parameters []string `nbt:"parameters"`
			TranslationKey string `nbt:"translation_key"`
		}{
			Parameters: []string{
				"sender",
				"content",
			},
			TranslationKey: "chat.type.text.narrate",
		},
	},
	"team_msg_command_outgoing": struct {
		Chat any `nbt:"chat"`
		Narration struct {
			Parameters []string `nbt:"parameters"`
			TranslationKey string `nbt:"translation_key"`
		} `nbt:"narration"`
	}{
		Chat: map[string]any{
			"parameters": []any{
				"target",
				"sender",
				"content",
			},
			"translation_key": "chat.type.team.sent",
		},
		Narration: struct {
			Parameters []string `nbt:"parameters"`
			TranslationKey string `nbt:"translation_key"`
		}{
			Parameters: []string{
				"sender",
				"content",
			},
			TranslationKey: "chat.type.text.narrate",
		},
	},
}
