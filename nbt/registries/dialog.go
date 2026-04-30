package registries

var Dialog = map[string]any{
	"custom_options": struct {
		ButtonWidth int32 `nbt:"button_width"`
		Columns int32 `nbt:"columns"`
		Dialogs string `nbt:"dialogs"`
		ExitAction struct {
			Label struct {
				Translate string `nbt:"translate"`
			} `nbt:"label"`
			Width int32 `nbt:"width"`
		} `nbt:"exit_action"`
		ExternalTitle struct {
			Translate string `nbt:"translate"`
		} `nbt:"external_title"`
		Title struct {
			Translate string `nbt:"translate"`
		} `nbt:"title"`
		Type string `nbt:"type"`
	}{
		ButtonWidth: int32(310),
		Columns: int32(1),
		Dialogs: "#minecraft:pause_screen_additions",
		ExitAction: struct {
			Label struct {
				Translate string `nbt:"translate"`
			} `nbt:"label"`
			Width int32 `nbt:"width"`
		}{
			Label: struct {
				Translate string `nbt:"translate"`
			}{
				Translate: "gui.back",
			},
			Width: int32(200),
		},
		ExternalTitle: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "menu.custom_options",
		},
		Title: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "menu.custom_options.title",
		},
		Type: "minecraft:dialog_list",
	},
	"quick_actions": struct {
		ButtonWidth int32 `nbt:"button_width"`
		Columns int32 `nbt:"columns"`
		Dialogs string `nbt:"dialogs"`
		ExitAction struct {
			Label struct {
				Translate string `nbt:"translate"`
			} `nbt:"label"`
			Width int32 `nbt:"width"`
		} `nbt:"exit_action"`
		ExternalTitle struct {
			Translate string `nbt:"translate"`
		} `nbt:"external_title"`
		Title struct {
			Translate string `nbt:"translate"`
		} `nbt:"title"`
		Type string `nbt:"type"`
	}{
		ButtonWidth: int32(310),
		Columns: int32(1),
		Dialogs: "#minecraft:quick_actions",
		ExitAction: struct {
			Label struct {
				Translate string `nbt:"translate"`
			} `nbt:"label"`
			Width int32 `nbt:"width"`
		}{
			Label: struct {
				Translate string `nbt:"translate"`
			}{
				Translate: "gui.back",
			},
			Width: int32(200),
		},
		ExternalTitle: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "menu.quick_actions",
		},
		Title: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "menu.quick_actions.title",
		},
		Type: "minecraft:dialog_list",
	},
	"server_links": struct {
		ButtonWidth int32 `nbt:"button_width"`
		Columns int32 `nbt:"columns"`
		ExitAction struct {
			Label struct {
				Translate string `nbt:"translate"`
			} `nbt:"label"`
			Width int32 `nbt:"width"`
		} `nbt:"exit_action"`
		ExternalTitle struct {
			Translate string `nbt:"translate"`
		} `nbt:"external_title"`
		Title struct {
			Translate string `nbt:"translate"`
		} `nbt:"title"`
		Type string `nbt:"type"`
	}{
		ButtonWidth: int32(310),
		Columns: int32(1),
		ExitAction: struct {
			Label struct {
				Translate string `nbt:"translate"`
			} `nbt:"label"`
			Width int32 `nbt:"width"`
		}{
			Label: struct {
				Translate string `nbt:"translate"`
			}{
				Translate: "gui.back",
			},
			Width: int32(200),
		},
		ExternalTitle: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "menu.server_links",
		},
		Title: struct {
			Translate string `nbt:"translate"`
		}{
			Translate: "menu.server_links.title",
		},
		Type: "minecraft:server_links",
	},
}
