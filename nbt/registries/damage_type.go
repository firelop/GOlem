package registries

var DamageType = map[string]any{
	"arrow": struct {
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Exhaustion: float32(0.1),
		MessageId: "arrow",
		Scaling: "when_caused_by_living_non_player",
	},
	"bad_respawn_point": struct {
		DeathMessageType string `nbt:"death_message_type"`
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		DeathMessageType: "intentional_game_design",
		Exhaustion: float32(0.1),
		MessageId: "badRespawnPoint",
		Scaling: "always",
	},
	"cactus": struct {
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Exhaustion: float32(0.1),
		MessageId: "cactus",
		Scaling: "when_caused_by_living_non_player",
	},
	"campfire": struct {
		Effects string `nbt:"effects"`
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Effects: "burning",
		Exhaustion: float32(0.1),
		MessageId: "inFire",
		Scaling: "when_caused_by_living_non_player",
	},
	"cramming": struct {
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Exhaustion: float32(0.0),
		MessageId: "cramming",
		Scaling: "when_caused_by_living_non_player",
	},
	"dragon_breath": struct {
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Exhaustion: float32(0.0),
		MessageId: "dragonBreath",
		Scaling: "when_caused_by_living_non_player",
	},
	"drown": struct {
		Effects string `nbt:"effects"`
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Effects: "drowning",
		Exhaustion: float32(0.0),
		MessageId: "drown",
		Scaling: "when_caused_by_living_non_player",
	},
	"dry_out": struct {
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Exhaustion: float32(0.1),
		MessageId: "dryout",
		Scaling: "when_caused_by_living_non_player",
	},
	"ender_pearl": struct {
		DeathMessageType string `nbt:"death_message_type"`
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		DeathMessageType: "fall_variants",
		Exhaustion: float32(0.0),
		MessageId: "fall",
		Scaling: "when_caused_by_living_non_player",
	},
	"explosion": struct {
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Exhaustion: float32(0.1),
		MessageId: "explosion",
		Scaling: "always",
	},
	"fall": struct {
		DeathMessageType string `nbt:"death_message_type"`
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		DeathMessageType: "fall_variants",
		Exhaustion: float32(0.0),
		MessageId: "fall",
		Scaling: "when_caused_by_living_non_player",
	},
	"falling_anvil": struct {
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Exhaustion: float32(0.1),
		MessageId: "anvil",
		Scaling: "when_caused_by_living_non_player",
	},
	"falling_block": struct {
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Exhaustion: float32(0.1),
		MessageId: "fallingBlock",
		Scaling: "when_caused_by_living_non_player",
	},
	"falling_stalactite": struct {
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Exhaustion: float32(0.1),
		MessageId: "fallingStalactite",
		Scaling: "when_caused_by_living_non_player",
	},
	"fireball": struct {
		Effects string `nbt:"effects"`
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Effects: "burning",
		Exhaustion: float32(0.1),
		MessageId: "fireball",
		Scaling: "when_caused_by_living_non_player",
	},
	"fireworks": struct {
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Exhaustion: float32(0.1),
		MessageId: "fireworks",
		Scaling: "when_caused_by_living_non_player",
	},
	"fly_into_wall": struct {
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Exhaustion: float32(0.0),
		MessageId: "flyIntoWall",
		Scaling: "when_caused_by_living_non_player",
	},
	"freeze": struct {
		Effects string `nbt:"effects"`
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Effects: "freezing",
		Exhaustion: float32(0.0),
		MessageId: "freeze",
		Scaling: "when_caused_by_living_non_player",
	},
	"generic": struct {
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Exhaustion: float32(0.0),
		MessageId: "generic",
		Scaling: "when_caused_by_living_non_player",
	},
	"generic_kill": struct {
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Exhaustion: float32(0.0),
		MessageId: "genericKill",
		Scaling: "when_caused_by_living_non_player",
	},
	"hot_floor": struct {
		Effects string `nbt:"effects"`
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Effects: "burning",
		Exhaustion: float32(0.1),
		MessageId: "hotFloor",
		Scaling: "when_caused_by_living_non_player",
	},
	"in_fire": struct {
		Effects string `nbt:"effects"`
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Effects: "burning",
		Exhaustion: float32(0.1),
		MessageId: "inFire",
		Scaling: "when_caused_by_living_non_player",
	},
	"in_wall": struct {
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Exhaustion: float32(0.0),
		MessageId: "inWall",
		Scaling: "when_caused_by_living_non_player",
	},
	"indirect_magic": struct {
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Exhaustion: float32(0.0),
		MessageId: "indirectMagic",
		Scaling: "when_caused_by_living_non_player",
	},
	"lava": struct {
		Effects string `nbt:"effects"`
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Effects: "burning",
		Exhaustion: float32(0.1),
		MessageId: "lava",
		Scaling: "when_caused_by_living_non_player",
	},
	"lightning_bolt": struct {
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Exhaustion: float32(0.1),
		MessageId: "lightningBolt",
		Scaling: "when_caused_by_living_non_player",
	},
	"mace_smash": struct {
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Exhaustion: float32(0.1),
		MessageId: "mace_smash",
		Scaling: "when_caused_by_living_non_player",
	},
	"magic": struct {
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Exhaustion: float32(0.0),
		MessageId: "magic",
		Scaling: "when_caused_by_living_non_player",
	},
	"mob_attack": struct {
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Exhaustion: float32(0.1),
		MessageId: "mob",
		Scaling: "when_caused_by_living_non_player",
	},
	"mob_attack_no_aggro": struct {
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Exhaustion: float32(0.1),
		MessageId: "mob",
		Scaling: "when_caused_by_living_non_player",
	},
	"mob_projectile": struct {
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Exhaustion: float32(0.1),
		MessageId: "mob",
		Scaling: "when_caused_by_living_non_player",
	},
	"on_fire": struct {
		Effects string `nbt:"effects"`
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Effects: "burning",
		Exhaustion: float32(0.0),
		MessageId: "onFire",
		Scaling: "when_caused_by_living_non_player",
	},
	"out_of_world": struct {
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Exhaustion: float32(0.0),
		MessageId: "outOfWorld",
		Scaling: "when_caused_by_living_non_player",
	},
	"outside_border": struct {
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Exhaustion: float32(0.0),
		MessageId: "outsideBorder",
		Scaling: "when_caused_by_living_non_player",
	},
	"player_attack": struct {
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Exhaustion: float32(0.1),
		MessageId: "player",
		Scaling: "when_caused_by_living_non_player",
	},
	"player_explosion": struct {
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Exhaustion: float32(0.1),
		MessageId: "explosion.player",
		Scaling: "always",
	},
	"sonic_boom": struct {
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Exhaustion: float32(0.0),
		MessageId: "sonic_boom",
		Scaling: "always",
	},
	"spear": struct {
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Exhaustion: float32(0.1),
		MessageId: "spear",
		Scaling: "when_caused_by_living_non_player",
	},
	"spit": struct {
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Exhaustion: float32(0.1),
		MessageId: "mob",
		Scaling: "when_caused_by_living_non_player",
	},
	"stalagmite": struct {
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Exhaustion: float32(0.0),
		MessageId: "stalagmite",
		Scaling: "when_caused_by_living_non_player",
	},
	"starve": struct {
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Exhaustion: float32(0.0),
		MessageId: "starve",
		Scaling: "when_caused_by_living_non_player",
	},
	"sting": struct {
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Exhaustion: float32(0.1),
		MessageId: "sting",
		Scaling: "when_caused_by_living_non_player",
	},
	"sweet_berry_bush": struct {
		Effects string `nbt:"effects"`
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Effects: "poking",
		Exhaustion: float32(0.1),
		MessageId: "sweetBerryBush",
		Scaling: "when_caused_by_living_non_player",
	},
	"thorns": struct {
		Effects string `nbt:"effects"`
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Effects: "thorns",
		Exhaustion: float32(0.1),
		MessageId: "thorns",
		Scaling: "when_caused_by_living_non_player",
	},
	"thrown": struct {
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Exhaustion: float32(0.1),
		MessageId: "thrown",
		Scaling: "when_caused_by_living_non_player",
	},
	"trident": struct {
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Exhaustion: float32(0.1),
		MessageId: "trident",
		Scaling: "when_caused_by_living_non_player",
	},
	"unattributed_fireball": struct {
		Effects string `nbt:"effects"`
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Effects: "burning",
		Exhaustion: float32(0.1),
		MessageId: "onFire",
		Scaling: "when_caused_by_living_non_player",
	},
	"wind_charge": struct {
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Exhaustion: float32(0.1),
		MessageId: "mob",
		Scaling: "when_caused_by_living_non_player",
	},
	"wither": struct {
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Exhaustion: float32(0.0),
		MessageId: "wither",
		Scaling: "when_caused_by_living_non_player",
	},
	"wither_skull": struct {
		Exhaustion float32 `nbt:"exhaustion"`
		MessageId string `nbt:"message_id"`
		Scaling string `nbt:"scaling"`
	}{
		Exhaustion: float32(0.1),
		MessageId: "witherSkull",
		Scaling: "when_caused_by_living_non_player",
	},
}
