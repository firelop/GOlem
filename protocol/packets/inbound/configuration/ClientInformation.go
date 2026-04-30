package configuration

import (
	"github.com/firelop/GOlem/nbt/registries"
	"github.com/firelop/GOlem/protocol"
	"github.com/firelop/GOlem/protocol/packets"
	"github.com/firelop/GOlem/protocol/packets/outbound/configuration"
	"github.com/firelop/GOlem/server"
)

type ClientInformation struct {
	Locale               string
	ViewDistance         byte
	ChatMode             int32
	ChatColors           bool
	SkinParts            byte
	MainHand             int32
	TextFilteringEnabled bool
	AllowServerListing   bool
	ParticleStatus       int32
}

func (c *ClientInformation) New() packets.InboundPacket {
	return &ClientInformation{}
}

func (c *ClientInformation) ID() int32 {
	return 0x00
}

func (c *ClientInformation) Read(pb *protocol.PacketBuffer) {
	c.Locale, _ = pb.ReadString()
	c.ViewDistance, _ = pb.ReadByte()
	c.ChatMode, _ = pb.ReadVarInt()
	c.ChatColors, _ = pb.ReadBool()
	c.SkinParts, _ = pb.ReadByte()
	c.MainHand, _ = pb.ReadVarInt()
	c.TextFilteringEnabled, _ = pb.ReadBool()
	c.AllowServerListing, _ = pb.ReadBool()
	c.ParticleStatus, _ = pb.ReadVarInt()
}

func (c *ClientInformation) Handle(session *server.Session) {
	// TODO: Save those settings in the session

	packets.SendPacket(session, &configuration.KnownPacks{
		Entries: []configuration.KnownPacksEntry{
			{
				Namespace: "minecraft",
				ID:        "core",
				Version:   "1.21.11",
			},
		},
	})

	packets.SendPacket(session, &configuration.UpdateTags{ // TODO: Send actual tags instead of empty ones. I don't really know why this packet is necessary
		Registries: []configuration.RegistryTagsEntry{
			{
				Identifier: "minecraft:timeline",
				Tags: []configuration.TagEntry{
					{
						Identifier: "minecraft:in_end",
						Tags:       []int32{},
					},
					{
						Identifier: "minecraft:in_nether",
						Tags:       []int32{},
					},
					{
						Identifier: "minecraft:in_overworld",
						Tags:       []int32{},
					},
				},
			},
		},
	})

	// TODO: Don't parse every registry every time
	//sendRegistry(session, "minecraft:timeline", registries.Timeline)
	sendRegistry(session, "minecraft:dimension_type", registries.DimensionType)
	sendRegistry(session, "minecraft:worldgen/biome", registries.Biome)
	sendRegistry(session, "minecraft:damage_type", registries.DamageType)
	sendRegistry(session, "minecraft:wolf_variant", registries.WolfVariant)
	sendRegistry(session, "minecraft:wolf_sound_variant", registries.WolfSoundVariant)
	sendRegistry(session, "minecraft:painting_variant", registries.PaintingVariant)
	sendRegistry(session, "minecraft:cat_variant", registries.CatVariant)
	sendRegistry(session, "minecraft:chicken_variant", registries.ChickenVariant)
	sendRegistry(session, "minecraft:cow_variant", registries.CowVariant)
	sendRegistry(session, "minecraft:frog_variant", registries.FrogVariant)
	sendRegistry(session, "minecraft:pig_variant", registries.PigVariant)
	sendRegistry(session, "minecraft:zombie_nautilus_variant", registries.ZombieNautilusVariant)

	packets.SendPacket(session, &configuration.FinishConfiguration{})
}

func sendRegistry(session *server.Session, id string, data map[string]any) {
	regData := registries.RegistryStructToRegistryData(id, data)
	packets.SendPacket(session, &regData)
}
